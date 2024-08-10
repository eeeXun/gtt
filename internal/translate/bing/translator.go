package bing

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/eeeXun/gtt/internal/translate/core"
)

const (
	setUpURL = "https://www.bing.com/translator"
	textURL  = "https://www.bing.com/ttranslatev3?IG=%s&IID=%s"
	posURL   = "https://www.bing.com/tlookupv3?IG=%s&IID=%s"
	ttsURL   = "https://www.bing.com/tfettts?IG=%s&IID=%s"
	ttsSSML  = "<speak version='1.0' xml:lang='%[1]s'><voice xml:lang='%[1]s' xml:gender='Female' name='%s'><prosody rate='-20.00%%'>%s</prosody></voice></speak>"
)

type Translator struct {
	*core.Server
	*core.Language
	*core.TTS
	core.EngineName
}

type setUpData struct {
	ig    string
	iid   string
	key   string
	token string
}

func NewTranslator() *Translator {
	return &Translator{
		Server:     new(core.Server),
		Language:   new(core.Language),
		TTS:        core.NewTTS(),
		EngineName: core.NewEngineName("Bing"),
	}
}

func (t *Translator) GetAllLang() []string {
	return lang
}

func (t *Translator) setUp() (*setUpData, error) {
	data := new(setUpData)

	res, err := http.Get(setUpURL)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	bodyStr := string(body)
	igData := regexp.MustCompile(`IG:"([^"]+)"`).FindStringSubmatch(bodyStr)
	if len(igData) < 2 {
		return nil, errors.New(t.GetEngineName() + " IG not found")
	}
	data.ig = igData[1]
	iidData := regexp.MustCompile(`data-iid="([^"]+)`).FindStringSubmatch(bodyStr)
	if len(iidData) < 2 {
		return nil, errors.New(t.GetEngineName() + " IID not found")
	}
	data.iid = iidData[1]
	params := regexp.MustCompile(`params_AbusePreventionHelper = ([^;]+);`).FindStringSubmatch(bodyStr)
	if len(params) < 2 {
		return nil, errors.New(t.GetEngineName() + " Key and Token not found")
	}
	paramsStr := strings.Split(params[1][1:len(params[1])-1], ",")
	data.key = paramsStr[0]
	data.token = paramsStr[1][1 : len(paramsStr[1])-1]

	return data, nil
}

func (t *Translator) Translate(message string) (translation *core.Translation, err error) {
	translation = new(core.Translation)
	var data []interface{}

	initData, err := t.setUp()
	if err != nil {
		return nil, err
	}
	userData := url.Values{
		"fromLang": {langCode[t.GetSrcLang()]},
		"to":       {langCode[t.GetDstLang()]},
		"text":     {message},
		"key":      {initData.key},
		"token":    {initData.token},
	}
	req, _ := http.NewRequest(http.MethodPost,
		fmt.Sprintf(textURL, initData.ig, initData.iid),
		strings.NewReader(userData.Encode()),
	)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", core.UserAgent)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	if len(data) <= 0 {
		return nil, errors.New("Translation not found")
	}

	// translation
	translation.TEXT =
		data[0].(map[string]interface{})["translations"].([]interface{})[0].(map[string]interface{})["text"].(string)

	// request part of speech
	userData.Del("fromLang")
	userData.Add("from", langCode[t.GetSrcLang()])
	req, _ = http.NewRequest(http.MethodPost,
		fmt.Sprintf(posURL, initData.ig, initData.iid),
		strings.NewReader(userData.Encode()),
	)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", core.UserAgent)
	res, err = http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	// Bing will return the request with list when success.
	// Otherwises, it would return map. Then the following err would not be nil.
	if err = json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	if len(data) <= 0 {
		return nil, errors.New("Translation not found")
	}

	poses := make(posSet)
	for _, pos := range data[0].(map[string]interface{})["translations"].([]interface{}) {
		pos := pos.(map[string]interface{})
		var words posWords

		words.target = pos["displayTarget"].(string)
		for _, backTranslation := range pos["backTranslations"].([]interface{}) {
			backTranslation := backTranslation.(map[string]interface{})
			words.add(backTranslation["displayText"].(string))
		}
		poses.add(pos["posTag"].(string), words)
	}
	translation.POS = poses.format()

	return translation, nil
}

func (t *Translator) PlayTTS(lang, message string) error {
	defer t.ReleaseLock()

	name, ok := voiceName[lang]
	if !ok {
		return errors.New(t.GetEngineName() + " does not support text to speech of " + lang)
	}
	initData, err := t.setUp()
	if err != nil {
		return err
	}
	userData := url.Values{
		// lang='%s' in ssml should be xx-XX, e.g. en-US
		// But xx also works, e.g. en
		// So don't do extra work to get xx-XX
		"ssml":  {fmt.Sprintf(ttsSSML, langCode[lang], name, message)},
		"key":   {initData.key},
		"token": {initData.token},
	}
	req, _ := http.NewRequest(http.MethodPost,
		fmt.Sprintf(ttsURL, initData.ig, initData.iid),
		strings.NewReader(userData.Encode()),
	)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", core.UserAgent)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	return t.Play(res.Body)
}

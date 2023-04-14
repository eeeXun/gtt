package bingtranslate

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/eeeXun/gtt/internal/translate/core"
)

const (
	setUpURL = "https://www.bing.com/translator"
	textURL  = "https://www.bing.com/ttranslatev3?IG=%s&IID=%s"
)

type BingTranslate struct {
	*core.Language
	*core.TTSLock
	core.EngineName
}

type setUpData struct {
	ig    string
	iid   string
	key   string
	token string
}

func NewBingTranslate() *BingTranslate {
	return &BingTranslate{
		Language:   core.NewLanguage(),
		TTSLock:    core.NewTTSLock(),
		EngineName: core.NewEngineName("BingTranslate"),
	}
}

func (t *BingTranslate) GetAllLang() []string {
	return lang
}

func (t *BingTranslate) setUp() (*setUpData, error) {
	var data setUpData

	res, err := http.Get(setUpURL)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
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

	return &data, nil
}

func (t *BingTranslate) Translate(message string) (translation, definition, partOfSpeech string, err error) {
	var data []interface{}

	initData, err := t.setUp()
	if err != nil {
		return "", "", "", err
	}
	userData := url.Values{
		"fromLang": {langCode[t.GetSrcLang()]},
		"to":       {langCode[t.GetDstLang()]},
		"text":     {message},
		"key":      {initData.key},
		"token":    {initData.token},
	}
	req, err := http.NewRequest("POST",
		fmt.Sprintf(textURL, initData.ig, initData.iid),
		strings.NewReader(userData.Encode()),
	)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", core.UserAgent)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", "", "", err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", "", "", err
	}
	if err = json.Unmarshal(body, &data); err != nil {
		return "", "", "", err
	}

	if len(data) <= 0 {
		return "", "", "", errors.New("Translation not found")
	}

	// translation
	translation = fmt.Sprintf("%v",
		data[0].(map[string]interface{})["translations"].([]interface{})[0].(map[string]interface{})["text"])

	return translation, definition, partOfSpeech, nil
}

func (t *BingTranslate) PlayTTS(lang, message string) error {
	defer t.ReleaseLock()

	return errors.New(t.GetEngineName() + " does not support text to speech")
}

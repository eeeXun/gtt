package lingvatranslate

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	// "time"

	"github.com/eeeXun/gtt/internal/translate/core"
	// "github.com/hajimehoshi/go-mp3"
	// "github.com/hajimehoshi/oto/v2"
)

const (
	textURL = "https://lingva.ml/_next/data/3qnDcUVykFKnSC3cdRX2t/%s/%s/%s.json"
)

type LingvaTranslate struct {
	*core.Language
	*core.TTSLock
	core.EngineName
}

func NewLingvaTranslate() *LingvaTranslate {
	return &LingvaTranslate{
		Language:   core.NewLanguage(),
		TTSLock:    core.NewTTSLock(),
		EngineName: core.NewEngineName("LingvaTranslate"),
	}
}

func (t *LingvaTranslate) GetAllLang() []string {
	return lang
}

func (t *LingvaTranslate) Translate(message string) (translation, definition, partOfSpeech string, err error) {
	var data map[string]interface{}

	urlStr := fmt.Sprintf(
		textURL,
		langCode[t.GetSrcLang()],
		langCode[t.GetDstLang()],
		url.PathEscape(message),
	)
	res, err := http.Get(urlStr)
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
	translation = fmt.Sprintf("%v", data["pageProps"].(map[string]interface{})["translation"])

	return translation, definition, partOfSpeech, nil
}

func (t *LingvaTranslate) PlayTTS(lang, message string) error {
	defer t.ReleaseLock()

	return nil
}

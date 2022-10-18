package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	api_url = "https://translate.googleapis.com/translate_a/single?client=gtx&sl=%s&tl=%s&dt=t&q=%s"
	sound_url = "https://translate.google.com.vn/translate_tts?ie=UTF-8&q=%s&tl=%s&client=tw-ob"
)

type Translator struct {
	src_lang  string
	dest_lang string
}

func (t Translator) Translate(message string) (string, error) {
	var data []interface{}
	var translated string

	url_str := fmt.Sprintf(
		api_url,
		Lang_Code[t.src_lang],
		Lang_Code[t.dest_lang],
		url.QueryEscape(message),
	)
	res, err := http.Get(url_str)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}

	if len(data) > 0 {
		result := data[0]
		for _, lines := range result.([]interface{}) {
			translated_line := lines.([]interface{})[0]
			translated += fmt.Sprintf("%v", translated_line)
		}
		return translated, nil
	}
	return "", errors.New("Translation not found")
}

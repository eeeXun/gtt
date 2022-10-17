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
	API_URL = "https://translate.googleapis.com/translate_a/single?client=gtx&sl=%s&tl=%s&dt=t&q=%s"
)

type Translator struct {
	src_lang  string
	dest_lang string
}

func (t Translator) Translate(message string) (string, error) {
	var data []interface{}
	var translated string

	url_str := fmt.Sprintf(
		API_URL,
		Lang[t.src_lang],
		Lang[t.dest_lang],
		url.QueryEscape(message),
	)
	res, err := http.Get(url_str)
	if err != nil {
		return err.Error(), err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err.Error(), err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return err.Error(), err
	}

	if len(data) > 0 {
		result := data[0]
		for _, lines := range result.([]interface{}) {
			translated_line := lines.([]interface{})[0]
			translated += fmt.Sprintf("%v", translated_line)
		}
		return translated, nil
	}
	return "Translation not found", errors.New("Translation not found")
}

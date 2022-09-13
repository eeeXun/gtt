package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Translate(message string, source_lang string, target_lang string) string {
	url_str := fmt.Sprintf(
		"https://translate.googleapis.com/translate_a/single?client=gtx&sl=%s&tl=%s&dt=t&q=%s",
		Lang[source_lang],
		Lang[target_lang],
		url.QueryEscape(message),
	)
	res, err := http.Get(url_str)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	return string(body[:])
}

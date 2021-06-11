package apis

import (
	"io/ioutil"
	"net/http"
)

var (
	url = "https://api.dictionaryapi.dev/api/v2/entries/"
)

func DictionarySearch(word string, langCode string) []byte {
	urlAfterAppend := url + langCode + "/" + word
	ret, err := http.Get(urlAfterAppend)
	if err != nil {
		return nil
	}
	body, err := ioutil.ReadAll(ret.Body)
	if err != nil {
		return nil
	}

	return body

}

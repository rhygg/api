package apis

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func GetSentiment(sentence string) ([]byte, error) {
	postBody, _ := json.Marshal(map[string]string{
		"text": sentence,
	})
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post("https://sentim-api.herokuapp.com/api/v1/", "application/json", responseBody)
	if err != nil {
		return nil, errors.New("Could Not Retrive A proper Sentiment, please try again")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Could Not Retrive A proper Sentiment, please try again")
	}
	return body, nil
}

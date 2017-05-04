package textprocessing

import (
	"net/http"
	"net/url"
	"encoding/json"
	"io/ioutil"
	"errors"
	"fmt"
)

const PhraseApiUrl = "http://text-processing.com/api/phrases/"

type PhraseExtractionResult map[string][]string

func ExtractPhrase(text string) (PhraseExtractionResult, error) {
	return ExtractPhraseByLanguage(text, "english")
}

func ExtractPhraseByLanguage(text string, language string) (PhraseExtractionResult, error) {
	postData := url.Values{}
	postData.Add("text", text)
	postData.Add("language", language)
	res, err := http.PostForm(PhraseApiUrl, postData)
	if err != nil {
		return nil, err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Expected status 200, but got %v. Response body: \n------\n%s", res.StatusCode, resBody))
	}

	result := PhraseExtractionResult{}
	err = json.Unmarshal(resBody, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
// Package textprocessing is a golang wrapper / client for the text-processing.com APIs.
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

// PhraseExtractionResult represents a result from the Phrase Extraction API.
type PhraseExtractionResult map[string][]string

// ExtractPhrase sends the given text to the Phrase Extraction and Named Entity Recognition API and results gives the result back. Use this for English-only analysis.
func ExtractPhrase(text string) (PhraseExtractionResult, error) {
	return ExtractPhraseByLanguage(text, "english")
}

// ExtractPhrase sends the given text to the Phrase Extraction and Named Entity Recognition API with the provided language as context and results gives the result back.
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
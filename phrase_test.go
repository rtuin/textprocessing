package textprocessing

import (
	"testing"
	"reflect"
)

func TestSimplePhraseRequest(t *testing.T) {
	result, err := ExtractPhrase("John hit the ball.")
	if err != nil {
		t.Errorf("Did not expect an error from the API, got: %v", err)
	}

	if !reflect.DeepEqual(result["PERSON"], []string{"John"}) {
		t.Errorf("Result was expected to contain []string{\"John\"}, but was: %v", result["PERSON"])
	}
}

func TestBadPhraseRequest(t *testing.T) {
	_, err := ExtractPhrase("")
	if err == nil {
		t.Errorf("Expected an error from the API but didn't get one.", err)
	}
}
package i18n

import (
	"encoding/json"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var bundle *i18n.Bundle

// supported languages
const (
	EN = "en"
	TR = "tr"
)

func InitBundle(languagePath string) {
	bundle = i18n.NewBundle(language.Turkish)

	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	var languages = []string{
		languagePath + "/en.json",
		languagePath + "/tr.json",
	}

	for _, language := range languages {
		bundle.MustLoadMessageFile(language)
	}
}

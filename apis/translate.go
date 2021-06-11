package apis

import (
	"github.com/bregydoc/gtranslate"
)

func Translate(text string, from string, to string) string {
	translated, err := gtranslate.TranslateWithParams(
		text,
		gtranslate.TranslationParams{
			From: from,
			To:   to,
		},
	)
	if err != nil {
		return "error while accquring a valid response."
	}
	return translated
}

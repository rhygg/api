package apis

import (
	"github.com/detectlanguage/detectlanguage-go"
)

func DetectLanguage(text string) string {
	client := detectlanguage.New("")

	detections, err := client.Detect(text)
	if err != nil {
		return "error while detecting the language"
	}
	return detections[0].Language
}

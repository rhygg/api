package apis

import (
	"github.com/detectlanguage/detectlanguage-go"
	"os"
)

func DetectLanguage(text string) string {
	client := detectlanguage.New(os.Getenv("KEY"))

	detections, err := client.Detect(text)
	if err != nil {
		return "error while detecting the language"
	}
	return detections[0].Language
}

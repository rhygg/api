package apis

import (
	"github.com/detectlanguage/detectlanguage-go"
)

func DetectLanguage(text string) string {
	client := detectlanguage.New("c2d939351dfbed486754612b9016b754")

	detections, err := client.Detect(text)
	if err != nil {
		return "error while detecting the language"
	}
	return detections[0].Language
}

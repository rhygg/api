package main

import (
	"rhydderchc/api/apis"

	"encoding/json"

	"rhydderchc/api/structures"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {

		return ctx.SendString("Hey this is my api! Its based on text analysis, you should go over to https://docs.rhydderchc.rocks and take a look at the documentation and numerous endpoints this text analysis api provides. If you like the api please do star the repository at https://github.com/rhydderchc/api !")
	})
	api := app.Group("/api/v1/")
	api.Get("/sentiment", func(ctx *fiber.Ctx) error {
		resp, err := apis.GetSentiment(ctx.Query("sentence"))
		if err != nil {
			return ctx.JSON(fiber.Map{
				"error": "Could not initiate the json request.",
			})
		}
		getResp := structures.Sentiment{}
		err = json.Unmarshal(resp, &getResp)
		if err != nil {
			return ctx.JSON(fiber.Map{
				"error": "Could not initiate the json request.",
			})
		}

		result := getResp.Result

		return ctx.JSON(fiber.Map{
			"polarity": result.Polarity,
			"type":     result.Type,
		})
	})
	api.Get("/detect_language", func(ctx *fiber.Ctx) error {
		lang := apis.DetectLanguage(ctx.Query("text"))
		return ctx.JSON(fiber.Map{
			"language": lang,
			"credits":  "Rhydderchc's Mocha API v1",
			"text":     ctx.Query("text"),
		})
	})
	api.Get("/translate", func(ctx *fiber.Ctx) error {
		text := ctx.Query("text")
		to := ctx.Query("to")
		from := ctx.Query("from")
		resp := apis.Translate(text, from, to)
		return ctx.JSON(fiber.Map{
			"translation": resp,
			"from":        from,
			"to":          to,
			"credits":     "Rhydderchc's Mocha API v1",
		})
	})
	api.Get("/dictionary", func(ctx *fiber.Ctx) error {
		word := ctx.Query("word")
		language_code := ctx.Query("lc")
		answers := apis.DictionarySearch(word, language_code)

		var reply []structures.DictionarySearch
		_ = json.Unmarshal(answers, &reply)
		return ctx.JSON(reply)
	})
	app.Listen(":3000")
}

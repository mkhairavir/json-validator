package main

// import (
// 	"fmt"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/gofiber/fiber/v2/log"
// )

// // field
// // type
// // value
// // required

// type Rule struct {
// 	Types    string
// 	Required bool
// }

// var (
// 	Rules = map[string]Rule{
// 		"a": {
// 			Types:    "string",
// 			Required: false,
// 		},
// 		"b": {
// 			Types:    "number",
// 			Required: false,
// 		},
// 		"c": {
// 			Types:    "bool",
// 			Required: true,
// 		},
// 		"d": {
// 			Types:    "string",
// 			Required: true,
// 		},
// 		"e": {
// 			Types:    "float",
// 			Required: false,
// 		},
// 	}
// )

// func validateJson(c *fiber.Ctx) error {
// 	bodyReq := map[string]interface{}{}

// 	_ = c.BodyParser(&bodyReq)

// 	for i, k := range Rules {
// 		input := bodyReq[i]
// 		if k.Required && input == nil {
// 			log.Error("kok nil sihhhhhh")
// 			return c.SendString(fmt.Sprintf("field %s ngga ada nih! padahal perlu", i))
// 		}

// 		switch input.(type) {
// 		case string:
// 			if k.Types != "string" {
// 				log.Error("kok bukan string sihhhhhh")
// 				return c.SendString(fmt.Sprintf("field %s ngga ada nih! padahal perlu", i))
// 			}
// 		case int, float32, float64:
// 			if k.Types != "number" {
// 				log.Error("kok bukan float sihhhhhh")
// 				return c.SendString(fmt.Sprintf("field %s ngga ada nih! padahal perlu", i))
// 			}
// 		case bool:
// 			if k.Types != "bool" {
// 				log.Error("kok bukan bool sihhhhhh")
// 				return c.SendString(fmt.Sprintf("field %s ngga ada nih! padahal perlu", i))
// 			}
// 		default:
// 			log.Error("kok nil sihhhhhh")
// 			return c.SendString(fmt.Sprintf("field %s ngga ada nih! padahal perlu", i))
// 		}

// 		if input != nil {

// 		}
// 	}
// 	// dapetin json
// 	// example: {"a": "aku", "b": "adalah", "c": "suneo", "d": "bukan"}
// 	// append field kalo valid based existing rules
// 	// kalo ga sama, ignore aja
// 	// hasil akhir: {"a": "aku"}

// 	return c.Send([]byte("haiya"))
// }

// func main() {
// 	app := fiber.New()

// 	// app.Use(validateJson)

// 	app.Get("/", func(c *fiber.Ctx) error {
// 		return c.SendString("yahooo!")
// 	})

// 	app.Get("/validate", validateJson)

// 	app.Listen(":4001")
// }

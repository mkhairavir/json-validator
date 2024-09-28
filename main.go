package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

type ValidationRule struct {
	Field    string `json:"field"`
	Type     string `json:"type"`
	Required bool   `json:"required"`
}

func ValidateJSON(input map[string]interface{}, rules []ValidationRule) (map[string]interface{}, error) {
	validated := make(map[string]interface{})
	for _, rule := range rules {
		value, exists := input[rule.Field]

		if rule.Required {
			if !exists {
				return nil, fmt.Errorf("required field '%s' is missing", rule.Field)
			}

			if !isValidType(value, rule.Type) {
				return nil, fmt.Errorf("required field '%s' value is not valid", rule.Field)
			}
		} else {
			if !exists || !isValidType(value, rule.Type) {
				continue
			}
		}
		validated[rule.Field] = value
	}
	return validated, nil
}

func isValidType(value interface{}, expectedType string) bool {
	actualType := reflect.TypeOf(value).String()

	switch expectedType {
	case "string":
		return actualType == "string"
	case "number":
		return actualType == "float64" || actualType == "int"
	case "bool":
		return actualType == "bool"
	default:
		return false
	}
}

func main() {
	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		var input map[string]interface{}
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid JSON format",
			})
		}

		rules := []ValidationRule{
			{"a", "string", true},
			{"b", "number", true},
			{"c", "bool", true},
			{"d", "string", false},
		}

		validatedJSON, err := ValidateJSON(input, rules)
		if err != nil {
			fmt.Println("Validation error:", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		} else {
			validatedOutput, _ := json.Marshal(validatedJSON)
			fmt.Println("Validated JSON:", string(validatedOutput))
		}
		return c.JSON(validatedJSON)
	})

	app.Listen(":8881")
}

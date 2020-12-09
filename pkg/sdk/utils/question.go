package utils

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

func ShowQuestion(message string) bool {
	allowedValues := [...]string{"y", "yes", "no", "n"}

	validate := func(input string) error {
		for _, value := range allowedValues {
			if strings.ToLower(input) == value {
				return nil
			}
		}
		return fmt.Errorf("Number should be one of the values %v", allowedValues)
	}

	prompt := promptui.Prompt{
		Label:    message,
		Validate: validate,
		Default:  "y",
	}

	result, err := prompt.Run()
	if err != nil {
		return ShowQuestion(message)
	}

	result = strings.ToLower(result)

	return result == "y" || result == "yes"
}

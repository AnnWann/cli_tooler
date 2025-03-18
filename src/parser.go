package cli_tooler

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func (c *__CLI_TOOLER__) parse(str string) error {
	var option string
	var modifiers = make(map[string]string)
	var arguments []string

	lexemes, err := getLexemes(str, c.Language)
	if err != nil {
		return err
	}

	if len(lexemes) < 2 {
		return nil
	}

	hasModifier, err := checkSyntax(lexemes, c.Language)
	if err != nil {
		return err
	}

	if hasModifier {
		option, modifiers, arguments = divideLexemesWithModifiers(lexemes)
	} else {
		option, arguments = divideLexemesWithoutModifiers(lexemes)
		modifiers = nil
	}

	c.Option = option
	c.Modifiers = modifiers
	c.Arguments = arguments
	return nil
}

func getLexemes(str string, language string) ([]string, error) {
	re := regexp.MustCompile(`"[^"]*"|\S+`)
	splitStr := re.FindAllString(str, -1)

	if len(splitStr) == 0 {
		errorMessage := dictionary[language]["noLexemesFound"]
		return nil, errors.New(errorMessage)
	}

	return splitStr, nil
}

func checkSyntax(lexemes []string, language string) (bool, error) {

	// Check if the first lexeme is a variable assignment
	if strings.HasPrefix(lexemes[0], `$`) {
		if len(lexemes) < 2 || len(lexemes) > 2 {
			errorMessage := dictionary[language]["notEnoughLexemes"]
			return false, fmt.Errorf(errorMessage, 3, len(lexemes))
		}

		return false, nil
	}

	// Check if the first lexeme is an option
	if !expectOption(lexemes[0]) {
		errorMessage := dictionary[language]["syntaxError"]
		return false, fmt.Errorf(errorMessage, lexemes[0], lexemes[1], lexemes[0], lexemes[1])
	}

	// Check if the second lexeme is an argument, if so, we have a command without modifiers so we just check for arguments
	if expectArgument(lexemes[1]) {
		lexemes = lexemes[1:]

		for i := 0; i < len(lexemes); i++ {
			if !expectArgument(lexemes[i]) {
				errorMessage := dictionary[language]["syntaxError"]
				return false, fmt.Errorf(errorMessage, lexemes[i-1], lexemes[i], name(expectArgument, language), name(expectModifier, language))
			}
		}
		return false, nil
	}

	lexemes = lexemes[2:]

	if len(lexemes) == 0 {
		return true, nil
	}

	// Check if the second lexeme is an argument, if so, we have a command with a modifier and a chain of arguments
	if expectArgument(lexemes[0]) {
		for i := 0; i < len(lexemes); i++ {
			if !expectArgument(lexemes[i]) {
				errorMessage := dictionary[language]["syntaxError"]
				return false, fmt.Errorf(errorMessage, lexemes[i-1], lexemes[i], name(expectArgument, language), name(expectModifier, language))
			}
		}
		return true, nil
	}

	var expect expectFn
	expect = expectModifier
	// Run through the lexemes expecting a modifier and then an argument
	for i := 0; i < len(lexemes); i++ {
		if expect(lexemes[i]) {
			expect = flipExpect(expect)
			continue
		}
		errorMessage := dictionary[language]["syntaxError"]
		return true, fmt.Errorf(errorMessage, lexemes[i-1], lexemes[i], name(expect, language), name(flipExpect(expect), language))
	}
	return true, nil
}

type expectFn func(string) bool

func expectOption(lexemes string) bool {
	return !strings.HasPrefix(lexemes, "--")
}

func expectModifier(lexemes string) bool {
	return strings.HasPrefix(lexemes, "--")
}

func expectArgument(lexemes string) bool {
	return !strings.HasPrefix(lexemes, `--`)
}

func name(fn func(string) bool, language string) string {
	if getFunctionName(fn) == getFunctionName(expectOption) {
		return dictionary[language]["option"]
	}
	if getFunctionName(fn) == getFunctionName(expectModifier) {
		return dictionary[language]["modifier"]
	}
	if getFunctionName(fn) == getFunctionName(expectArgument) {
		return dictionary[language]["argument"]
	}
	return dictionary[language]["unknown"]
}

func flipExpect(expect expectFn) expectFn {
	if getFunctionName(expect) == getFunctionName(expectModifier) {
		return expectArgument
	}
	return expectModifier
}

func getFunctionName(i interface{}) string {
	return strings.TrimPrefix(strings.TrimSuffix(strings.TrimPrefix(strings.TrimSuffix(fmt.Sprintf("%v", i), "func(string) bool"), "func("), ")"), "main.")
}

func divideLexemesWithModifiers(lexemes []string) (string, map[string]string, []string) {
	var option string
	var modifiers = make(map[string]string)
	var arguments []string

	option = lexemes[0]
	firstModifier := lexemes[1]
	modifiers[firstModifier] = ""
	lexemesRest := lexemes[2:]

	for i := 0; i < len(lexemesRest); i++ {
		if strings.HasPrefix(lexemesRest[i], "--") {
			if i+1 < len(lexemesRest) && !strings.HasPrefix(lexemesRest[i+1], "--") {
				if strings.HasPrefix(lexemesRest[i+1], `$`) {

					modifiers[lexemesRest[i]] = CLI_TOOLER.getVariable(strings.Trim(lexemesRest[i+1], `$`))
				} else {
					modifiers[lexemesRest[i]] = strings.Trim(lexemesRest[i+1], `"`)
					i++
				}
			} else {
				modifiers[lexemesRest[i]] = ""
			}
		} else {
			if strings.HasPrefix(lexemesRest[i], `$`) {
				arguments = append(arguments, CLI_TOOLER.getVariable(strings.Trim(lexemesRest[i], `$`)))
			} else {
				arguments = append(arguments, strings.Trim(lexemesRest[i], `"`))
			}
		}
	}

	return option, modifiers, arguments
}

func divideLexemesWithoutModifiers(lexemes []string) (string, []string) {
	var option string
	var arguments []string

	if strings.HasPrefix(lexemes[0], `$`) {
		option = "$"
		arguments = append(arguments, CLI_TOOLER.getVariable(strings.Trim(lexemes[0], `$`)))
		arguments = append(arguments, strings.Trim(lexemes[2], `"`))
		return option, arguments
	}

	option = lexemes[0]
	lexemesRest := lexemes[1:]

	for i := 0; i < len(lexemesRest); i++ {
		if strings.HasPrefix(lexemesRest[i], `$`) {
			arguments = append(arguments, CLI_TOOLER.getVariable(strings.Trim(lexemesRest[i], `$`)))
		} else {
			arguments = append(arguments, strings.Trim(lexemesRest[i], `"`))
		}
	}

	return option, arguments
}

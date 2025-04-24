package parser

import (
	"cli_tooler/internal/dictionary"
	errorhandler "cli_tooler/internal/error_handler"
	"errors"
	"regexp"
)

type LexicalErrorOutput struct {
	tokens           []token
	problematicToken token
}

func lexical(str string) ([]token, error) {
	lexemes, err := getLexemes(str)
	if err != nil {
		return nil, err
	}

	var tokens []token
	for i, lexeme := range lexemes {
		switch {
		case matchEntireString(assignmentREG, lexeme):
			tokens = append(tokens, token{Value: lexeme, Kind: "assignment", Column: i})
		case matchEntireString(logicalREG, lexeme):
			tokens = append(tokens, token{Value: lexeme, Kind: "logical", Column: i})
		case matchEntireString(mathREG, lexeme):
			tokens = append(tokens, token{Value: lexeme, Kind: "math", Column: i})
		case matchEntireString(parenthesis_OREG, lexeme):
			tokens = append(tokens, token{Value: lexeme, Kind: "parenthesis_open", Column: i})
		case matchEntireString(parenthesis_CREG, lexeme):
			tokens = append(tokens, token{Value: lexeme, Kind: "parenthesis_close", Column: i})
		case matchEntireString(callREG, lexeme):
			tokens = append(tokens, token{Value: lexeme, Kind: "call", Column: i})
		case matchEntireString(stringREG, lexeme):
			lexemeWithoutQuotes := lexeme[1 : len(lexeme)-1]
			tokens = append(tokens, token{Value: lexemeWithoutQuotes, Kind: "string", Column: i})
		case matchEntireString(nameREG, lexeme):
			tokens = append(tokens, token{Value: lexeme, Kind: "name", Column: i})
		case matchEntireString(floatREG, lexeme):
			tokens = append(tokens, token{Value: lexeme, Kind: "float", Column: i})
		case matchEntireString(integerREG, lexeme):
			tokens = append(tokens, token{Value: lexeme, Kind: "integer", Column: i})
		default:
			return nil, errorhandler.HandleError("lexicalError", LexicalErrorOutput{
				tokens:           tokens,
				problematicToken: token{Value: lexeme, Kind: "unknown", Column: i},
			})
		}
	}

	return tokens, nil
}

func getLexemes(str string) ([]string, error) {
	var lexemes = make([]string, 0)
	s := ""
	for c := 0; c < len(str); c++ {
		if str[c] == ' ' {
			if len(s) > 0 {
				lexemes = append(lexemes, s)
				s = ""
			}

			continue
		}
		if str[c] == '(' || str[c] == ')' {
			if len(s) > 0 {
				lexemes = append(lexemes, s)
				s = ""
			}
			lexemes = append(lexemes, string(str[c]))
			continue
		}
		if str[c] == '$' {
			if len(s) > 0 {
				lexemes = append(lexemes, s)
				s = ""
			}
			lexemes = append(lexemes, "$")
			continue
		}
		if str[c] == ':' {
			if len(s) > 0 {
				lexemes = append(lexemes, s)
				s = ""
			}
			s += ":"
			if str[c+1] == '=' {
				s += "="
				c++
			}
			lexemes = append(lexemes, s)
			s = ""
			continue
		}
		if str[c] == '=' {
			if len(s) > 0 {
				lexemes = append(lexemes, s)
				s = ""
			}
			s += "="
			if str[c+1] == '=' {
				s += "="
				c++
			}
			lexemes = append(lexemes, s)
			s = ""
			continue
		}
		if str[c] == '>' {
			if len(s) > 0 {
				lexemes = append(lexemes, s)
				s = ""
			}
			s += ">"
			if str[c+1] == '=' {
				s += "="
				c++
			}
			lexemes = append(lexemes, s)
			s = ""
			continue
		}
		if str[c] == '<' {
			if len(s) > 0 {
				lexemes = append(lexemes, s)
				s = ""
			}
			s += "<"
			if str[c+1] == '=' {
				s += "="
				c++
			}
			lexemes = append(lexemes, s)
			s = ""
			continue
		}
		if str[c] == '!' {
			if len(s) > 0 {
				lexemes = append(lexemes, s)
				s = ""
			}
			s += "!"
			if str[c+1] == '=' {
				s += "="
				c++
			}
			lexemes = append(lexemes, s)
			s = ""
			continue
		}
		if str[c] == '&' {
			if len(s) > 0 {
				lexemes = append(lexemes, s)
				s = ""
			}
			s += "&"
			if str[c+1] == '&' {
				s += "&"
				c++
			}
			lexemes = append(lexemes, s)
			s = ""
			continue
		}
		if str[c] == '|' {
			if len(s) > 0 {
				lexemes = append(lexemes, s)
				s = ""
			}
			s += "|"
			if str[c+1] == '|' {
				s += "|"
				c++
			}
			lexemes = append(lexemes, s)
			s = ""
			continue
		}
		if str[c] == '+' || str[c] == '-' || str[c] == '/' || str[c] == '*' || str[c] == '^' {
			if len(s) > 0 {
				lexemes = append(lexemes, s)
				s = ""
			}
			s += string(str[c])
			lexemes = append(lexemes, s)
			s = ""
			continue
		}

		if str[c] == '"' {
			if len(s) > 0 {
				lexemes = append(lexemes, s)
				s = ""
			}
			s += "\""
			i := c + 1
			for ; i < len(str) && str[i] != '"'; i++ {
				s += string(str[i])
			}
			s += "\""
			lexemes = append(lexemes, s)
			c = i
			s = ""
			continue
		}

		s += string(str[c])
	}

	if len(s) > 0 {
		lexemes = append(lexemes, s)
	}

	if len(lexemes) == 0 {
		errorMessage := dictionary.GetString("noLexemesFound")
		return nil, errors.New(errorMessage)
	}

	return lexemes, nil
}

const (
	mathREG          = `(\+|-|/|\*|\^)`
	logicalREG       = `(==|!=|<=|>=|&&|\|\||>|<)`
	assignmentREG    = `(:=|=)`
	parenthesis_OREG = `[\(]`
	parenthesis_CREG = `[\)]`
	nameREG          = `[a-zA-Z_][a-zA-Z0-9_]*`
	callREG          = `[$]`
	stringREG        = `"[^"]*"`
	integerREG       = `[0-9]+`
	floatREG         = `[0-9]+\.[0-9]+`
)

func matchEntireString(regExp string, str string) bool {
	match := regexp.MustCompile(regExp).FindSubmatch([]byte(str))
	if match == nil {
		return false
	}
	return string(match[0]) == str
}

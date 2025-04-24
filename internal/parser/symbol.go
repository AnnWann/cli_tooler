package parser

import (
	"cli_tooler/internal/cli_tooler"
	"cli_tooler/internal/dictionary"
	"math"
	"regexp"
	"strconv"
	"strings"
)

const maxPrecedence = math.MaxInt
const nonValueMaxPrecedence = maxPrecedence - 1

type Symbol struct {
	token        token
	precedence   int
	apply        func(string, string) string
	nextIsValid  func(string) bool
	errorMessage string
	isUnary      bool
}

func createSymbol(t token) Symbol {
	switch t.Value {
	case "+":
		return plus(t)
	case "-":
		return minus(t)
	case "*":
		return multiply(t)
	case "/":
		return divide(t)
	case "!":
		return negation(t)
	case "$":
		return call(t)
	case ":=":
		return assignment(t)
	case ">":
		return greater(t)
	case "<":
		return less(t)
	case ">=":
		return greaterOrEqual(t)
	case "<=":
		return lessOrEqual(t)
	case "=":
		return equal(t)
	case "!=":
		return notEqual(t)
	case "&&":
		return and(t)
	case "||":
		return or(t)
	default:
		for word, data := range cli_tooler.CLI.GetReservedWords() {
			if t.Value == word {
				return Symbol{
					token:        t,
					precedence:   data.Precedence,
					apply:        data.Fn,
					nextIsValid:  data.NextIsValid,
					errorMessage: data.ErrorMessage,
					isUnary:      data.IsUnary,
				}
			}
		}
		var precedence int
		if t.Value == "(" || t.Value == ")" {
			precedence = -1
		} else {
			precedence = maxPrecedence
		}

		return Symbol{
			token:       t,
			precedence:  precedence,
			nextIsValid: func(string) bool { return true },
		}
	}
}

func (s Symbol) LessOrEqualInPrecedence(sp Symbol) bool {
	return s.precedence <= sp.precedence
}

func nextIsValidMathOperators(symbol string) bool {
	return regexp.MustCompile(`[$()]|[a-zA-Z_][a-zA-Z0-9_]*|[0-9]+\.[0-9]+|[0-9]+`).MatchString(symbol)
}

func nextIsValidLogicalOperators(symbol string) bool {
	return regexp.MustCompile(`[$()!]|[a-zA-Z_][a-zA-Z0-9_]*|"[^"]*"|[0-9]+\.[0-9]+|[0-9]+`).MatchString(symbol)
}

func plus(t token) Symbol {
	return Symbol{
		t,
		nonValueMaxPrecedence - 1,
		func(a, b string) string {
			integer, err := strconv.Atoi(a)
			if err == nil {
				integer2, err := strconv.Atoi(b)
				if err == nil {
					return strconv.Itoa(integer + integer2)
				}
			}
			float, err := strconv.ParseFloat(a, 64)
			if err == nil {
				float2, err := strconv.ParseFloat(b, 64)
				if err == nil {
					return strconv.FormatFloat(float+float2, 'f', -1, 64)
				}
			}
			return a + b
		},
		nextIsValidMathOperators,
		dictionary.GetString("operatorError"),
		false,
	}
}

func minus(t token) Symbol {
	return Symbol{
		t,
		nonValueMaxPrecedence - 1,
		func(a, b string) string {
			integer, err := strconv.Atoi(a)
			if err == nil {
				integer2, err := strconv.Atoi(b)
				if err == nil {
					return strconv.Itoa(integer - integer2)
				}
			}
			float, err := strconv.ParseFloat(a, 64)
			if err == nil {
				float2, err := strconv.ParseFloat(b, 64)
				if err == nil {
					return strconv.FormatFloat(float-float2, 'f', -1, 64)
				}
			}
			return strings.Trim(a, b)
		},
		nextIsValidMathOperators,
		dictionary.GetString("operatorError"),
		false,
	}
}

func multiply(t token) Symbol {
	return Symbol{
		t,
		nonValueMaxPrecedence,
		func(a, b string) string {
			integer, err := strconv.Atoi(a)
			if err == nil {
				integer2, err := strconv.Atoi(b)
				if err == nil {
					return strconv.Itoa(integer * integer2)
				}
			}
			float, err := strconv.ParseFloat(a, 64)
			if err == nil {
				float2, err := strconv.ParseFloat(b, 64)
				if err == nil {
					return strconv.FormatFloat(float*float2, 'f', -1, 64)
				}
			}
			return ""
		},
		nextIsValidMathOperators,
		dictionary.GetString("operatorError"),
		false,
	}
}

func divide(t token) Symbol {
	return Symbol{
		t,
		nonValueMaxPrecedence,
		func(a, b string) string {
			integer, err := strconv.Atoi(a)
			if err == nil {
				integer2, err := strconv.Atoi(b)
				if err == nil {
					return strconv.Itoa(integer / integer2)
				}
			}
			float, err := strconv.ParseFloat(a, 64)
			if err == nil {
				float2, err := strconv.ParseFloat(b, 64)
				if err == nil {
					return strconv.FormatFloat(float/float2, 'f', -1, 64)
				}
			}
			return ""
		},
		nextIsValidMathOperators,
		dictionary.GetString("operatorError"),
		false,
	}
}

func equal(t token) Symbol {
	return Symbol{
		t,
		nonValueMaxPrecedence - 2,
		func(a, b string) string {
			if a == b {
				return "true"
			}
			return "false"
		},
		nextIsValidLogicalOperators,
		dictionary.GetString("operatorError"),
		false,
	}
}

func greater(t token) Symbol {
	return Symbol{
		t,
		nonValueMaxPrecedence - 2,
		func(a, b string) string {
			integer, err := strconv.Atoi(a)
			if err == nil {
				integer2, err := strconv.Atoi(b)
				if err == nil {
					if integer > integer2 {
						return "true"
					}
					return "false"
				}
			}
			float, err := strconv.ParseFloat(a, 64)
			if err == nil {
				float2, err := strconv.ParseFloat(b, 64)
				if err == nil {
					if float > float2 {
						return "true"
					}
					return "false"
				}
			}
			return "false"
		},
		nextIsValidLogicalOperators,
		dictionary.GetString("operatorError"),
		false,
	}
}

func less(t token) Symbol {
	return Symbol{
		t,
		nonValueMaxPrecedence - 2,
		func(a, b string) string {
			integer, err := strconv.Atoi(a)
			if err == nil {
				integer2, err := strconv.Atoi(b)
				if err == nil {
					if integer < integer2 {
						return "true"
					}
					return "false"
				}
			}
			float, err := strconv.ParseFloat(a, 64)
			if err == nil {
				float2, err := strconv.ParseFloat(b, 64)
				if err == nil {
					if float < float2 {
						return "true"
					}
					return "false"
				}
			}
			return "false"
		},
		nextIsValidLogicalOperators,
		dictionary.GetString("operatorError"),
		false,
	}
}

func greaterOrEqual(t token) Symbol {
	return Symbol{
		t,
		nonValueMaxPrecedence - 2,
		func(a, b string) string {
			integer, err := strconv.Atoi(a)
			if err == nil {
				integer2, err := strconv.Atoi(b)
				if err == nil {
					if integer >= integer2 {
						return "true"
					}
					return "false"
				}
			}
			float, err := strconv.ParseFloat(a, 64)
			if err == nil {
				float2, err := strconv.ParseFloat(b, 64)
				if err == nil {
					if float >= float2 {
						return "true"
					}
					return "false"
				}
			}
			return "false"
		},
		nextIsValidLogicalOperators,
		dictionary.GetString("operatorError"),
		false,
	}
}

func lessOrEqual(t token) Symbol {
	return Symbol{
		t,
		nonValueMaxPrecedence - 2,
		func(a, b string) string {
			integer, err := strconv.Atoi(a)
			if err == nil {
				integer2, err := strconv.Atoi(b)
				if err == nil {
					if integer < integer2 {
						return "true"
					}
					return "false"
				}
			}
			float, err := strconv.ParseFloat(a, 64)
			if err == nil {
				float2, err := strconv.ParseFloat(b, 64)
				if err == nil {
					if float < float2 {
						return "true"
					}
					return "false"
				}
			}
			return "false"
		},
		nextIsValidLogicalOperators,
		dictionary.GetString("operatorError"),
		false,
	}
}

func notEqual(t token) Symbol {
	return Symbol{
		t,
		nonValueMaxPrecedence - 2,
		func(a, b string) string {
			if a != b {
				return "true"
			}
			return "false"
		},
		nextIsValidLogicalOperators,
		dictionary.GetString("operatorError"),
		false,
	}
}

func negation(t token) Symbol {
	return Symbol{
		t,
		nonValueMaxPrecedence - 3,
		func(a, b string) string {
			if a == "true" {
				return "false"
			}
			return "true"
		},
		nextIsValidLogicalOperators,
		dictionary.GetString("operatorError"),
		true,
	}
}

func and(t token) Symbol {
	return Symbol{
		t,
		nonValueMaxPrecedence - 4,
		func(a, b string) string {
			aBool := a == "true"
			bBool := b == "true"
			return strconv.FormatBool(aBool && bBool)
		},
		nextIsValidLogicalOperators,
		dictionary.GetString("operatorError"),
		false,
	}
}

func or(t token) Symbol {
	return Symbol{
		t,
		nonValueMaxPrecedence - 5,
		func(a, b string) string {
			aBool := a == "true"
			bBool := b == "true"
			return strconv.FormatBool(aBool && bBool)
		},
		nextIsValidLogicalOperators,
		dictionary.GetString("operatorError"),
		false,
	}
}

func assignment(t token) Symbol {
	return Symbol{
		t,
		nonValueMaxPrecedence - 6,
		func(a, b string) string {
			cli_tooler.CLI.SetVariable(a, b)
			return b
		},
		nextIsValidLogicalOperators,
		dictionary.GetString("operatorError"),
		false,
	}
}

func call(t token) Symbol {
	return Symbol{
		t,
		nonValueMaxPrecedence - 6,
		func(a, b string) string {
			return cli_tooler.CLI.GetVariable(a)
		},
		func(token string) bool {
			return regexp.MustCompile(`[a-zA-Z_][a-zA-Z0-9_]*`).MatchString(token)
		},
		dictionary.GetString("operatorError"),
		true,
	}
}

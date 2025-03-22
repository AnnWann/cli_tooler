package parser

import (
	"cli_tooler/internal/cli_tooler"
	"cli_tooler/internal/dictionary"
	"cli_tooler/internal/variable_table"
	"math"
	"regexp"
)

type Symbol struct {
	sign         string
	precedence   int
	apply        interface{}
	nextIsValid  func(string) bool
	errorMessage string
}

func createSymbol(sign string) Symbol {
	switch sign {
	case "+":
		return plus()
	case "-":
		return minus()
	case "*":
		return multiply()
	case "/":
		return divide()
	case "!":
		return negation()
	case "$":
		return call()
	case ":=":
		return assignment()
	case ">":
		return greater()
	case "<":
		return less()
	case ">=":
		return greaterOrEqual()
	case "<=":
		return lessOrEqual()
	case "=":
		return equal()
	case "!=":
		return notEqual()
	case "&&":
		return and()
	case "||":
		return or()
	default:
		for word, data := range cli_tooler.ReservedWords {
			if sign == word {
				return Symbol{
					sign:         sign,
					precedence:   data.Precedence,
					apply:        data.Fn,
					nextIsValid:  data.NextIsValid,
					errorMessage: data.ErrorMessage,
				}
			}
		}
		return Symbol{
			sign:        sign,
			precedence:  0,
			apply:       nil,
			nextIsValid: func(string) bool { return true },
		}
	}
}

func (s Symbol) lessOrEqualInPrecedence(sp Symbol) bool {
	return s.precedence <= sp.precedence
}

func nextIsValidMathOperators(symbol string) bool {
	return regexp.MustCompile(`[$]|[a-zA-Z_][a-zA-Z0-9_]*|[0-9]+\.[0-9]+|[0-9]+`).MatchString(symbol)
}

func nextIsValidLogicalOperators(symbol string) bool {
	return regexp.MustCompile(`[$]|[a-zA-Z_][a-zA-Z0-9_]*|"[^"]*"|[0-9]+\.[0-9]+|[0-9]+`).MatchString(symbol)
}

func plus() Symbol {
	return Symbol{
		"+",
		math.MaxInt - 1,
		struct {
			_int   func(int, int) int
			_float func(float64, float64) float64
		}{
			_int:   func(a, b int) int { return a + b },
			_float: func(a, b float64) float64 { return a + b },
		},
		nextIsValidMathOperators,
		dictionary.GetString("operatorError"),
	}
}

func minus() Symbol {
	return Symbol{
		"-",
		math.MaxInt - 1,
		struct {
			_int   func(int, int) int
			_float func(float64, float64) float64
		}{
			_int:   func(a, b int) int { return a - b },
			_float: func(a, b float64) float64 { return a - b },
		},
		nextIsValidMathOperators,
		dictionary.GetString("operatorError"),
	}
}

func multiply() Symbol {
	return Symbol{
		"*",
		math.MaxInt,
		struct {
			_int   func(int, int) int
			_float func(float64, float64) float64
		}{
			_int:   func(a, b int) int { return a * b },
			_float: func(a, b float64) float64 { return a * b },
		},
		nextIsValidMathOperators,
		dictionary.GetString("operatorError"),
	}
}

func divide() Symbol {
	return Symbol{
		"/",
		math.MaxInt,
		struct {
			_int   func(int, int) int
			_float func(float64, float64) float64
		}{
			_int:   func(a, b int) int { return a / b },
			_float: func(a, b float64) float64 { return a / b },
		},
		nextIsValidMathOperators,
		dictionary.GetString("operatorError"),
	}
}

func equal() Symbol {
	return Symbol{
		"=",
		math.MaxInt - 2,
		func(a, b int) bool {
			return a == b
		},
		nextIsValidLogicalOperators,
		dictionary.GetString("operatorError"),
	}
}

func greater() Symbol {
	return Symbol{
		">",
		math.MaxInt - 2,
		func(a, b int) bool {
			return a > b
		},
		nextIsValidLogicalOperators,
		dictionary.GetString("operatorError"),
	}
}

func less() Symbol {
	return Symbol{
		"<",
		math.MaxInt - 2,
		func(a, b int) bool {
			return a < b
		},
		nextIsValidLogicalOperators,
		dictionary.GetString("operatorError")}
}

func greaterOrEqual() Symbol {
	return Symbol{">=",
		math.MaxInt - 2,
		func(a, b int) bool {
			return a >= b
		},
		nextIsValidLogicalOperators,
		dictionary.GetString("operatorError")}
}

func lessOrEqual() Symbol {
	return Symbol{
		"<=",
		math.MaxInt - 2,
		func(a, b int) bool {
			return a <= b
		},
		nextIsValidLogicalOperators,
		dictionary.GetString("operatorError")}
}

func notEqual() Symbol {
	return Symbol{
		"!=",
		math.MaxInt - 2,
		func(a, b int) bool {
			return a != b
		},
		nextIsValidLogicalOperators,
		dictionary.GetString("operatorError"),
	}
}

func negation() Symbol {
	return Symbol{
		"!",
		math.MaxInt - 3,
		func(a bool) bool { return !a },
		nextIsValidLogicalOperators,
		dictionary.GetString("operatorError")}
}

func and() Symbol {
	return Symbol{
		"&&",
		math.MaxInt - 4,
		func(a, b bool) bool {
			return a && b
		},
		nextIsValidLogicalOperators,
		dictionary.GetString("operatorError")}
}

func or() Symbol {
	return Symbol{
		"||",
		math.MaxInt - 5,
		func(a, b bool) bool {
			return a || b
		},
		nextIsValidLogicalOperators,
		dictionary.GetString("operatorError")}
}

func assignment() Symbol {
	return Symbol{
		":=",
		math.MaxInt - 6,
		func(a, b string) {
			variable_table.SetVariable(a, b)
		},
		nextIsValidLogicalOperators,
		dictionary.GetString("operatorError")}
}

func call() Symbol {
	return Symbol{
		"$", math.MaxInt - 6,
		func(a string) string {
			return variable_table.GetVariable(a)
		},
		func(token string) bool {
			return regexp.MustCompile(`[a-zA-Z_][a-zA-Z0-9_]*`).MatchString(token)
		},
		dictionary.GetString("operatorError")}
}

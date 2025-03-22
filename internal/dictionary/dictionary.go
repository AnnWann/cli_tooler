package dictionary

func SetLanguage(lang string) {
	language = lang
}

func GetString(key string) string {
	return dictionary[language][key]
}

var language string

var dictionary = map[string]map[string]string{
	"pt-br": {
		"noLexemesFound":            "nenhum léxico encontrado",
		"lexicalError":              "erro léxico encontrado em '%s %s', léxico não reconhecido",
		"syntaxError":               "erro de sintaxe encontrado em '%s %s'",
		"invalidNumberOfLexemes":    "não há léxicos suficientes, esperavasse pelo menos %d, mas recebeu %d",
		"operatorError":             "após um operador se espera um valor, não se pode usar outro operador ou palavra reservada",
		"missingClosingParenthesis": "falta fechar um parêntese",

		"option":   "opção",
		"modifier": "modificador",
		"argument": "argumento",
		"value":    "valor",
		"operator": "operador",
		"unknown":  "desconhecido",
	},
	"en-us": {
		"noLexemesFound":            "no lexemes found",
		"lexicalError":              "lexical error found in '%s %s', lexeme not recognized",
		"syntaxError":               "syntax error found in '%s %s'",
		"invalidNumberOfLexemes":    "not enough lexemes, expected at least %d, but received %d",
		"operatorError":             "after an operator a value is expected, another operator or reserved word cannot be used",
		"missingClosingParenthesis": "missing closing parenthesis",

		"option":   "option",
		"modifier": "modifier",
		"argument": "argument",
		"value":    "value",
		"operator": "operator",
		"unknown":  "unknown",
	},
}

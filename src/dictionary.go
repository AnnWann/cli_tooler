package cli_tooler

var dictionary = map[string]map[string]string{
	"pt-br": {
		"noLexemesFound":         "nenhum léxico encontrado",
		"syntaxError":            "erro de sintaxe encontrado em '%s %s', recebemos um %s mas esperavamos um %s",
		"invalidNumberOfLexemes": "não há léxicos suficientes, esperavasse pelo menos %d, mas recebeu %d",

		"option":   "opção",
		"modifier": "modificador",
		"argument": "argumento",
		"unknown":  "desconhecido",
	},
	"en-us": {
		"noLexemesFound":         "no lexemes found",
		"syntaxError":            "syntax error found in '%s %s', received a %s but expected a %s",
		"invalidNumberOfLexemes": "not enough lexemes, expected at least %d, but received %d",

		"option":   "option",
		"modifier": "modifier",
		"argument": "argument",
		"unknown":  "unknown",
	},
}

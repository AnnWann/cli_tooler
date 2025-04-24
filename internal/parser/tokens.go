package parser

type token struct {
	Value  string
	Kind   string
	Row    int
	Column int
}

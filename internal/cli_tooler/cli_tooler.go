package cli_tooler

import "cli_tooler/internal/variable_table"

type Word struct {
	Fn           func(string, string) string
	Precedence   int
	NextIsValid  func(string) bool
	ErrorMessage string
	IsUnary      bool
}

var reservedWords map[string]Word
var variables variable_table.VariableTable
var CLI CLI_TOOLER

type CLI_TOOLER int

type Word_Builder struct {
	name string
	Word Word
}

func BuildWord(name string) *Word_Builder {
	return &Word_Builder{name: name}
}

func (wb *Word_Builder) SetFn(fn func(string, string) string) *Word_Builder {
	wb.Word.Fn = fn
	return wb
}

func (wb *Word_Builder) SetPrecedence(precedence int) *Word_Builder {
	wb.Word.Precedence = precedence
	return wb
}

func (wb *Word_Builder) SetNextIsValid(nextIsValid func(string) bool) *Word_Builder {
	wb.Word.NextIsValid = nextIsValid
	return wb
}

func (wb *Word_Builder) SetErrorMessage(errorMessage string) *Word_Builder {
	wb.Word.ErrorMessage = errorMessage
	return wb
}

func (wb *Word_Builder) SetIsUnary(isUnary bool) *Word_Builder {
	wb.Word.IsUnary = isUnary
	return wb
}

func (wb *Word_Builder) Build() {
	CLI.AddReservedWord(wb.name, wb.Word)
}

func (ct *CLI_TOOLER) AddReservedWord(name string, word Word) {
	reservedWords[name] = word
}

func (ct *CLI_TOOLER) GetReservedWord(name string) Word {
	return reservedWords[name]
}

func (ct *CLI_TOOLER) GetReservedWords() map[string]Word {
	return reservedWords
}

func (ct *CLI_TOOLER) SetVariable(name string, value string) {
	variables.SetVariable(name, value)
}

func (ct *CLI_TOOLER) GetVariable(name string) string {
	return variables.GetVariable(name)
}

func (ct *CLI_TOOLER) GetVariables() variable_table.VariableTable {
	return variables
}

func (ct *CLI_TOOLER) ClearVariableTable() {
	variables.ClearVariableTable()
}

func (ct *CLI_TOOLER) InitVariableTable(file string) error {
	return variables.InitVariableTable(file)
}

//var CLI_TOOLER CLI_TOOLER = CLI_TOOLER{}

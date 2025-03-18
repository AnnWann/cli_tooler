package cli_tooler

var CLI_TOOLER __CLI_TOOLER__

type __CLI_TOOLER__ struct {
	Language      string
	Lexemes       []string
	Option        string
	Modifiers     map[string]string
	Arguments     []string
	VariableTable map[string]string
}

type __CLI_TOOLER_INTERFACE__ interface {
	init(language string) __CLI_TOOLER_INTERFACE__

	parse(str string) error

	setVariable(args []string)
	getVariable(name string) string
	deleteVariable(name string)
	clearVariableTable()
	initVariableTable() error
}

func (c *__CLI_TOOLER__) init(language string) __CLI_TOOLER_INTERFACE__ {
	c.Language = language
	c.initVariableTable()
	return c
}

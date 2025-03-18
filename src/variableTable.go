package cli_tooler

import (
	"errors"
	"os"
	"strings"
)

func (c *__CLI_TOOLER__) setVariable(args []string) {
	if len(args) < 2 {
		return
	}

	if c.VariableTable == nil {
		c.VariableTable = make(map[string]string)
	}

	c.VariableTable[args[0]] = args[1]
}

func (c *__CLI_TOOLER__) getVariable(name string) string {
	return c.VariableTable[name]
}

func (c *__CLI_TOOLER__) deleteVariable(name string) {
	delete(c.VariableTable, name)
}

func (c *__CLI_TOOLER__) clearVariableTable() {
	c.VariableTable = nil
	c.initVariableTable()
}

func (c *__CLI_TOOLER__) initVariableTable() error {
	c.VariableTable = make(map[string]string)
	variables, err := os.ReadFile("variables")
	if err != nil {
		return errors.New("could not read variables file")
	}
	variablesStr := string(variables)
	variablesArr := strings.Split(variablesStr, "\n")
	for _, variable := range variablesArr {
		if variable == "" {
			continue
		}
		variableParts := strings.Split(variable, "=")
		c.VariableTable[variableParts[0]] = variableParts[1]
	}
	return nil
}

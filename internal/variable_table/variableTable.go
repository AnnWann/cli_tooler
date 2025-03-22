package variable_table

import (
	"errors"
	"os"
	"strings"
)

var __VARIABLE_TABLE__ = map[string]string{}

func SetVariable(a string, b string) {

	if __VARIABLE_TABLE__ == nil {
		__VARIABLE_TABLE__ = make(map[string]string)
	}

	__VARIABLE_TABLE__[a] = b
}

func GetVariable(name string) string {
	return __VARIABLE_TABLE__[name]
}

func DeleteVariable(name string) {
	delete(__VARIABLE_TABLE__, name)
}

func ClearVariableTable() {
	__VARIABLE_TABLE__ = nil
	InitVariableTable()
}

func InitVariableTable() error {
	__VARIABLE_TABLE__ = make(map[string]string)

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
		__VARIABLE_TABLE__[variableParts[0]] = variableParts[1]
	}
	return nil
}

package variable_table

import (
	"errors"
	"os"
	"strings"
)

type VariableTable map[string]string

func (v VariableTable) SetVariable(a string, b string) {

	if v == nil {
		v = make(map[string]string)
	}

	v[a] = b
}

func (v VariableTable) GetVariable(name string) string {
	return v[name]
}

func (v VariableTable) DeleteVariable(name string) {
	delete(v, name)
}

func (v *VariableTable) ClearVariableTable() {
	*v = make(map[string]string)
}

func (v VariableTable) InitVariableTable(file string) error {
	if v == nil {
		v = make(map[string]string)
	}

	if file == "" {
		return nil
	}

	variables, err := os.ReadFile(file)
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
		v[variableParts[0]] = variableParts[1]
	}
	return nil
}

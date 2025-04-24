package variable_table

import (
	"os"
	"testing"
)

func Test_VariableT(t *testing.T) {
	t.Parallel()

	t.Run("Test_Init With No Variables", func(t *testing.T) {

		v := VariableTable{}

		err := v.InitVariableTable("")
		if err != nil {
			t.Errorf("Expected err to be nil, but got %v", err)
		}
		if len(v) != 0 {
			t.Errorf("Expected variableTable to be empty, but got %v", v)
		}

		v.SetVariable("a", "1")
		v.SetVariable("b", "2")
		v.SetVariable("c", "3")

		if v.GetVariable("a") != "1" {
			t.Errorf("Expected variableTable[a] to be 1, but got %s", v["a"])
		}

		if v.GetVariable("b") != "2" {
			t.Errorf("Expected variableTable[b] to be 2, but got %s", v["b"])
		}

		if v.GetVariable("c") != "3" {
			t.Errorf("Expected variableTable[c] to be 3, but got %s", v["c"])
		}

		v.DeleteVariable("a")

		if va := v.GetVariable("a"); va != "" {
			t.Errorf("Expected variableTable[a] to be empty, but got %s", v["a"])
		}

		v.ClearVariableTable()
		if len(v) != 0 {
			t.Errorf("Expected variableTable to be empty, but got %v", v)
		}
	})

	t.Run("Test_Init With Variables", func(t *testing.T) {

		v := VariableTable{}

		variables := "a=1\nb=2\nc=3"
		file_name := "variables"
		// input variables in file
		err := os.WriteFile(file_name, []byte(variables), 0644)
		if err != nil {
			t.Errorf("Expected err to be nil, but got %v", err)
		}

		defer os.Remove(file_name)

		//expected output
		expected := map[string]string{"a": "1", "b": "2", "c": "3"}

		err = v.InitVariableTable(file_name)
		if err != nil {
			t.Errorf("Expected err to be nil, but got %v", err)
		}
		if len(v) != 0 {
			t.Errorf("Expected variableTable to be empty, but got %v", v)
		}

		for key, value := range expected {
			if v[key] != value {
				t.Errorf("Expected variableTable[%s] to be %s, but got %s", key, value, v[key])
			}
		}
	})

}

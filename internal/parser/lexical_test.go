package parser

import (
	"testing"
)

func Test_Lexical(t *testing.T) {

	t.Run("Test \"a sentence of only words\"", func(t *testing.T) {
		// Test "a sentence of only words"
		// Test case for the lexical function.
		// The input is a string "a sentence of only words".
		// The output is a list of tokens.

		expected := []token{
			{value: "a", kind: "name"},
			{value: "sentence", kind: "name"},
			{value: "of", kind: "name"},
			{value: "only", kind: "name"},
			{value: "words", kind: "name"},
		}

		succedingTest(t, "a sentence of only words", expected)
	})

	t.Run("Test \"90 40 30 21 59 3280 33230 64000 39291\"", func(t *testing.T) {
		// Test "90 40 30 21 59 3280 33230 64000 39291"
		// Test case for the lexical function.
		// The input is a string "90 40 30 21 59 3280 33230 64000 39291".
		// The output is a list of tokens.

		expected := []token{
			{value: "90", kind: "integer"},
			{value: "40", kind: "integer"},
			{value: "30", kind: "integer"},
			{value: "21", kind: "integer"},
			{value: "59", kind: "integer"},
			{value: "3280", kind: "integer"},
			{value: "33230", kind: "integer"},
			{value: "64000", kind: "integer"},
			{value: "39291", kind: "integer"},
		}

		succedingTest(t, "90 40 30 21 59 3280 33230 64000 39291", expected)
	})

	t.Run("Test \"30.5 1.5 0.5 0.1 0.0001 0.0", func(t *testing.T) {
		// Test "30.5 1.5 0.5 0.1 0.0001 0.0"
		// Test case for the lexical function.
		// The input is a string "30.5 1.5 0.5 0.1 0.0001 0.0".
		// The output is a list of tokens.

		expected := []token{
			{value: "30.5", kind: "float"},
			{value: "1.5", kind: "float"},
			{value: "0.5", kind: "float"},
			{value: "0.1", kind: "float"},
			{value: "0.0001", kind: "float"},
			{value: "0.0", kind: "float"},
		}

		succedingTest(t, "30.5 1.5 0.5 0.1 0.0001 0.0", expected)
	})

	t.Run("Test \"( ) ( ( ) )\"", func(t *testing.T) {
		// Test "( ) ( ( ) )"
		// Test case for the lexical function.
		// The input is a string "( ) ( ( ) )".
		// The output is a list of tokens.7

		expected := []token{
			{value: "(", kind: "parenthesis_open"},
			{value: ")", kind: "parenthesis_close"},
			{value: "(", kind: "parenthesis_open"},
			{value: "(", kind: "parenthesis_open"},
			{value: ")", kind: "parenthesis_close"},
			{value: ")", kind: "parenthesis_close"},
		}

		succedingTest(t, "( ) ( ( ) )", expected)
	})

	t.Run("Test \"==:=<=>=< > + -/ * ^\"", func(t *testing.T) {
		// Test "== := <= >= < > + - / * ^"
		// Test case for the lexical function.
		// The input is a string "== := <= >= < > + - / * ^".
		// The output is a list of tokens.

		expected := []token{
			{value: "==", kind: "logical"},
			{value: ":=", kind: "assignment"},
			{value: "<=", kind: "logical"},
			{value: ">=", kind: "logical"},
			{value: "<", kind: "logical"},
			{value: ">", kind: "logical"},
			{value: "+", kind: "math"},
			{value: "-", kind: "math"},
			{value: "/", kind: "math"},
			{value: "*", kind: "math"},
			{value: "^", kind: "math"},
		}

		succedingTest(t, "==:=<=>=< > + -/ * ^", expected)
	})

	t.Run("Test \"a := 30.5 + 1.5 * 0.5\"", func(t *testing.T) {
		// Test "a := 30.5 + 1.5 * 0.5"
		// Test case for the lexical function.
		// The input is a string "a := 30.5 + 1.5 * 0.5".
		// The output is a list of tokens.

		expected := []token{
			{value: "a", kind: "name"},
			{value: ":=", kind: "assignment"},
			{value: "30.5", kind: "float"},
			{value: "+", kind: "math"},
			{value: "1.5", kind: "float"},
			{value: "*", kind: "math"},
			{value: "0.5", kind: "float"},
		}

		succedingTest(t, "a := 30.5 + 1.5 * 0.5", expected)
	})

	t.Run("Test \"reservedWord reservedWord2 reservedWord3\"", func(t *testing.T) {
		// Test "reservedWord reservedWord2 reservedWord3"
		// Test case for the lexical function.
		// The input is a string "reservedWord reservedWord2 reservedWord3".
		// The output is a list of tokens.

		expected := []token{
			{value: "reservedWord", kind: "name"},
			{value: "reservedWord2", kind: "name"},
			{value: "reservedWord3", kind: "name"},
		}

		succedingTest(t, "reservedWord reservedWord2 reservedWord3", expected)
	})

	t.Run("Test \"\"value\" \"value2\" \"value3\"\"", func(t *testing.T) {
		// Test "\"value\" \"value2\" \"value3\""
		// Test case for the lexical function.
		// The input is a string "\"value" "value2" "value3"".
		// The output is a list of tokens

		expected := []token{
			{value: "value", kind: "string"},
			{value: "value2", kind: "string"},
			{value: "value3", kind: "string"},
		}

		succedingTest(t, "\"value\" \"value2\" \"value3\"", expected)
	})

	t.Run("Test \"banana:=$batata+salada\"", func(t *testing.T) {
		// Test "banana:=$batata+salada"
		// Test case for the lexical function.
		// The input is a string "banana:=$batata+salada".
		// The output is a list of tokens.

		expected := []token{
			{value: "banana", kind: "name"},
			{value: ":=", kind: "assignment"},
			{value: "$", kind: "call"},
			{value: "batata", kind: "name"},
			{value: "+", kind: "math"},
			{value: "salada", kind: "name"},
		}

		succedingTest(t, "banana:=$batata+salada", expected)
	})

}

func succedingTest(t *testing.T, input string, expectedOutput []token) {
	//t.Helper()

	// Test "input"
	// Test case for the lexical function.
	// The input is a string "input".
	// The output is a list of tokens.

	// Call the lexical function with the input.
	// Check if the output is equal to the expected output.

	got, err := lexical(input)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(got) != len(expectedOutput) {
		t.Errorf("expected %v, got %v", expectedOutput, got)
	}

	for i := range got {
		if got[i] != expectedOutput[i] {
			t.Errorf("expected %v, got %v", expectedOutput, got)
		}
	}
}

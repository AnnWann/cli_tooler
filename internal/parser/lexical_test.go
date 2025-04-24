package parser

import (
	"testing"
)

func Test_Lexical(t *testing.T) {
	t.Parallel()

	t.Run("Test \"a sentence of only words\"", func(t *testing.T) {
		// Test "a sentence of only words"
		// Test case for the lexical function.
		// The input is a string "a sentence of only words".
		// The output is a list of tokens.

		expected := []token{
			{Value: "a", Kind: "name"},
			{Value: "sentence", Kind: "name"},
			{Value: "of", Kind: "name"},
			{Value: "only", Kind: "name"},
			{Value: "words", Kind: "name"},
		}

		succedingLexicalTest(t, "a sentence of only words", expected)
	})

	t.Run("Test \"90 40 30 21 59 3280 33230 64000 39291\"", func(t *testing.T) {
		// Test "90 40 30 21 59 3280 33230 64000 39291"
		// Test case for the lexical function.
		// The input is a string "90 40 30 21 59 3280 33230 64000 39291".
		// The output is a list of tokens.

		expected := []token{
			{Value: "90", Kind: "integer"},
			{Value: "40", Kind: "integer"},
			{Value: "30", Kind: "integer"},
			{Value: "21", Kind: "integer"},
			{Value: "59", Kind: "integer"},
			{Value: "3280", Kind: "integer"},
			{Value: "33230", Kind: "integer"},
			{Value: "64000", Kind: "integer"},
			{Value: "39291", Kind: "integer"},
		}

		succedingLexicalTest(t, "90 40 30 21 59 3280 33230 64000 39291", expected)
	})

	t.Run("Test \"30.5 1.5 0.5 0.1 0.0001 0.0", func(t *testing.T) {
		// Test "30.5 1.5 0.5 0.1 0.0001 0.0"
		// Test case for the lexical function.
		// The input is a string "30.5 1.5 0.5 0.1 0.0001 0.0".
		// The output is a list of tokens.

		expected := []token{
			{Value: "30.5", Kind: "float"},
			{Value: "1.5", Kind: "float"},
			{Value: "0.5", Kind: "float"},
			{Value: "0.1", Kind: "float"},
			{Value: "0.0001", Kind: "float"},
			{Value: "0.0", Kind: "float"},
		}

		succedingLexicalTest(t, "30.5 1.5 0.5 0.1 0.0001 0.0", expected)
	})

	t.Run("Test \"( ) ( ( ) )\"", func(t *testing.T) {
		// Test "( ) ( ( ) )"
		// Test case for the lexical function.
		// The input is a string "( ) ( ( ) )".
		// The output is a list of tokens.7

		expected := []token{
			{Value: "(", Kind: "parenthesis_open"},
			{Value: ")", Kind: "parenthesis_close"},
			{Value: "(", Kind: "parenthesis_open"},
			{Value: "(", Kind: "parenthesis_open"},
			{Value: ")", Kind: "parenthesis_close"},
			{Value: ")", Kind: "parenthesis_close"},
		}

		succedingLexicalTest(t, "( ) ( ( ) )", expected)
	})

	t.Run("Test \"==:=<=>=< > + -/ * ^\"", func(t *testing.T) {
		// Test "== := <= >= < > + - / * ^"
		// Test case for the lexical function.
		// The input is a string "== := <= >= < > + - / * ^".
		// The output is a list of tokens.

		expected := []token{
			{Value: "==", Kind: "logical"},
			{Value: ":=", Kind: "assignment"},
			{Value: "<=", Kind: "logical"},
			{Value: ">=", Kind: "logical"},
			{Value: "<", Kind: "logical"},
			{Value: ">", Kind: "logical"},
			{Value: "+", Kind: "math"},
			{Value: "-", Kind: "math"},
			{Value: "/", Kind: "math"},
			{Value: "*", Kind: "math"},
			{Value: "^", Kind: "math"},
		}

		succedingLexicalTest(t, "==:=<=>=< > + -/ * ^", expected)
	})

	t.Run("Test \"a := 30.5 + 1.5 * 0.5\"", func(t *testing.T) {
		// Test "a := 30.5 + 1.5 * 0.5"
		// Test case for the lexical function.
		// The input is a string "a := 30.5 + 1.5 * 0.5".
		// The output is a list of tokens.

		expected := []token{
			{Value: "a", Kind: "name"},
			{Value: ":=", Kind: "assignment"},
			{Value: "30.5", Kind: "float"},
			{Value: "+", Kind: "math"},
			{Value: "1.5", Kind: "float"},
			{Value: "*", Kind: "math"},
			{Value: "0.5", Kind: "float"},
		}

		succedingLexicalTest(t, "a := 30.5 + 1.5 * 0.5", expected)
	})

	t.Run("Test \"reservedWord reservedWord2 reservedWord3\"", func(t *testing.T) {
		// Test "reservedWord reservedWord2 reservedWord3"
		// Test case for the lexical function.
		// The input is a string "reservedWord reservedWord2 reservedWord3".
		// The output is a list of tokens.

		expected := []token{
			{Value: "reservedWord", Kind: "name"},
			{Value: "reservedWord2", Kind: "name"},
			{Value: "reservedWord3", Kind: "name"},
		}

		succedingLexicalTest(t, "reservedWord reservedWord2 reservedWord3", expected)
	})

	t.Run("Test \"\"value\" \"value2\" \"value3\"\"", func(t *testing.T) {
		// Test "\"value\" \"value2\" \"value3\""
		// Test case for the lexical function.
		// The input is a string "\"value" "value2" "value3"".
		// The output is a list of tokens

		expected := []token{
			{Value: "value", Kind: "string"},
			{Value: "value2", Kind: "string"},
			{Value: "value3", Kind: "string"},
		}

		succedingLexicalTest(t, "\"value\" \"value2\" \"value3\"", expected)
	})

	t.Run("Test \"banana:=$batata+salada\"", func(t *testing.T) {
		// Test "banana:=$batata+salada"
		// Test case for the lexical function.
		// The input is a string "banana:=$batata+salada".
		// The output is a list of tokens.

		expected := []token{
			{Value: "banana", Kind: "name"},
			{Value: ":=", Kind: "assignment"},
			{Value: "$", Kind: "call"},
			{Value: "batata", Kind: "name"},
			{Value: "+", Kind: "math"},
			{Value: "salada", Kind: "name"},
		}

		succedingLexicalTest(t, "banana:=$batata+salada", expected)
	})

	t.Run("Test multi-line input", func(t *testing.T) {
	})

	t.Run("Error cases", func(t *testing.T) {

	})
}

func succedingLexicalTest(t *testing.T, input string, expectedOutput []token) {
	t.Helper()

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

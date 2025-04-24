package parser

import (
	"cli_tooler/internal/cli_tooler"
	"cli_tooler/internal/utility"
	"fmt"
	"testing"
)

func Test_Syntax(t *testing.T) {
	t.Parallel()
	t.Run("Mathematical operations", func(t *testing.T) {
		// Test mathematical operations
		t.Run("Test \"5+2\"", func(t *testing.T) {
			// Test "5+2"
			// Test case for the syntof tokens made from the string "5+2".
			// The output is a binary tree.

			expected := &utility.BinaryTreeNode[Symbol]{
				Value: Symbol{token: token{Value: "+", Kind: "math", Column: 1, Row: 0}},
				Left:  &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: token{Value: "5", Kind: "integer", Column: 0, Row: 0}}},
				Right: &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: token{Value: "2", Kind: "integer", Column: 2, Row: 0}}},
			}

			input := []token{
				{Value: "5", Kind: "integer", Column: 0, Row: 0},
				{Value: "+", Kind: "math", Column: 1, Row: 0},
				{Value: "2", Kind: "integer", Column: 2, Row: 0},
			}

			succedingSyntaticTest(t, input, expected)
		})

		t.Run("Test \"5-2\"", func(t *testing.T) {
			// Test "5-2"
			// Test case for the syntax function.
			// The input is a list of tokens made from the string "5-2".
			// The output is a binary tree.

			expected := &utility.BinaryTreeNode[Symbol]{
				Value: Symbol{token: token{Value: "-", Kind: "math", Column: 1, Row: 0}},
				Left:  &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: token{Value: "5", Kind: "integer", Column: 0, Row: 0}}},
				Right: &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: token{Value: "2", Kind: "integer", Column: 2, Row: 0}}},
			}

			input := []token{
				{Value: "5", Kind: "integer", Column: 0, Row: 0},
				{Value: "-", Kind: "math", Column: 1, Row: 0},
				{Value: "2", Kind: "integer", Column: 2, Row: 0},
			}

			succedingSyntaticTest(t, input, expected)
		})

		t.Run("Test \"5*3\"", func(t *testing.T) {
			// Test "5*3"
			// Test case for the syntax function.
			// The input is a list of tokens made from the string "5*3".
			// The output is a binary tree.

			expected := &utility.BinaryTreeNode[Symbol]{
				Value: Symbol{token: token{Value: "*", Kind: "math", Column: 1, Row: 0}},
				Left:  &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: token{Value: "5", Kind: "integer", Column: 0, Row: 0}}},
				Right: &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: token{Value: "3", Kind: "integer", Column: 2, Row: 0}}},
			}

			input := []token{
				{Value: "5", Kind: "integer", Column: 0, Row: 0},
				{Value: "*", Kind: "math", Column: 1, Row: 0},
				{Value: "3", Kind: "integer", Column: 2, Row: 0},
			}

			succedingSyntaticTest(t, input, expected)
		})

		t.Run("Test \"5/3\"", func(t *testing.T) {
			// Test "5/3"
			// Test case for the syntax function.
			// The input is a list of tokens made from the string "5/3".
			// The output is a binary tree.

			expected := &utility.BinaryTreeNode[Symbol]{
				Value: Symbol{token: token{Value: "/", Kind: "math", Column: 1, Row: 0}},
				Left:  &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: token{Value: "5", Kind: "integer", Column: 0, Row: 0}}},
				Right: &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: token{Value: "3", Kind: "integer", Column: 2, Row: 0}}},
			}

			input := []token{
				{Value: "5", Kind: "integer", Column: 0, Row: 0},
				{Value: "/", Kind: "math", Column: 1, Row: 0},
				{Value: "3", Kind: "integer", Column: 2, Row: 0},
			}

			succedingSyntaticTest(t, input, expected)
		})

		t.Run("Test \"5^2\"", func(t *testing.T) {
			// Test "5^2"
			// Test case for the syntax function.
			// The input is a list of tokens made from the string "5^2".
			// The output is a binary tree.

			expected := &utility.BinaryTreeNode[Symbol]{
				Value: Symbol{token: token{Value: "^", Kind: "math", Column: 1, Row: 0}},
				Left:  &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: token{Value: "5", Kind: "integer", Column: 0, Row: 0}}},
				Right: &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: token{Value: "2", Kind: "integer", Column: 2, Row: 0}}},
			}

			input := []token{
				{Value: "5", Kind: "integer", Column: 0, Row: 0},
				{Value: "^", Kind: "math", Column: 1, Row: 0},
				{Value: "2", Kind: "integer", Column: 2, Row: 0},
			}

			succedingSyntaticTest(t, input, expected)
		})

		t.Run("Test \"5+2*3\"", func(t *testing.T) {
			// Test "5+2*3"
			// Test case for the syntax function.
			// The input is a list of tokens made from the string "5+2*3".
			// The output is a binary tree.

			expected := &utility.BinaryTreeNode[Symbol]{
				Value: Symbol{token: token{Value: "+", Kind: "math", Column: 1, Row: 0}},
				Left:  &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: token{Value: "5", Kind: "integer", Column: 0, Row: 0}}},
				Right: &utility.BinaryTreeNode[Symbol]{
					Value: Symbol{token: token{Value: "*", Kind: "math", Column: 3, Row: 0}},
					Left:  &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: token{Value: "2", Kind: "integer", Column: 2, Row: 0}}},
					Right: &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: token{Value: "3", Kind: "integer", Column: 4, Row: 0}}},
				},
			}

			input := []token{
				{Value: "5", Kind: "integer", Column: 0, Row: 0},
				{Value: "+", Kind: "math", Column: 1, Row: 0},
				{Value: "2", Kind: "integer", Column: 2, Row: 0},
				{Value: "*", Kind: "math", Column: 3, Row: 0},
				{Value: "3", Kind: "integer", Column: 4, Row: 0},
			}

			succedingSyntaticTest(t, input, expected)
		})

		t.Run("Test \"5/2+2*5\"", func(t *testing.T) {
			// Test "5/2+2*5"
			// Test case for the syntax function.
			// The input is a list of tokens made from the string "5/2+2*5".
			// The output is a binary tree.

			expected := &utility.BinaryTreeNode[Symbol]{
				Value: Symbol{token: token{Value: "+", Kind: "math", Column: 3, Row: 0}},
				Left: &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: token{Value: "/", Kind: "math", Column: 1, Row: 0}},
					Left:  &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: token{Value: "5", Kind: "integer", Column: 0, Row: 0}}},
					Right: &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: token{Value: "2", Kind: "integer", Column: 2, Row: 0}}},
				},
				Right: &utility.BinaryTreeNode[Symbol]{
					Value: Symbol{token: token{Value: "*", Kind: "math", Column: 5, Row: 0}},
					Left:  &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: token{Value: "2", Kind: "integer", Column: 4, Row: 0}}},
					Right: &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: token{Value: "5", Kind: "integer", Column: 6, Row: 0}}},
				},
			}

			input := []token{
				{Value: "5", Kind: "integer", Column: 0, Row: 0},
				{Value: "/", Kind: "math", Column: 1, Row: 0},
				{Value: "2", Kind: "integer", Column: 2, Row: 0},
				{Value: "+", Kind: "math", Column: 3, Row: 0},
				{Value: "2", Kind: "integer", Column: 4, Row: 0},
				{Value: "*", Kind: "math", Column: 5, Row: 0},
				{Value: "5", Kind: "integer", Column: 6, Row: 0},
			}

			succedingSyntaticTest(t, input, expected)
		})

		t.Run("Test \"5+2*3+1\"", func(t *testing.T) {
			// Test "5+2*3+1"
			// Test case for the syntax function.
			// The input is a list of tokens made from the string "5+2*3+1".
			// The output is a binary tree.

			expected := &utility.BinaryTreeNode[Symbol]{
				Value: Symbol{token: token{Value: "+", Kind: "math", Column: 5, Row: 0}},
				Left: &utility.BinaryTreeNode[Symbol]{
					Value: Symbol{token: token{Value: "+", Kind: "math", Column: 1, Row: 0}},
					Left:  &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: token{Value: "5", Kind: "integer", Column: 0, Row: 0}}},
					Right: &utility.BinaryTreeNode[Symbol]{
						Value: Symbol{token: token{Value: "*", Kind: "math", Column: 3, Row: 0}},
						Left:  &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: token{Value: "2", Kind: "integer", Column: 2, Row: 0}}},
						Right: &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: token{Value: "3", Kind: "integer", Column: 4, Row: 0}}},
					},
				},
				Right: &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: token{Value: "1", Kind: "integer", Column: 6, Row: 0}}},
			}

			input := []token{
				{Value: "5", Kind: "integer", Column: 0, Row: 0},
				{Value: "+", Kind: "math", Column: 1, Row: 0},
				{Value: "2", Kind: "integer", Column: 2, Row: 0},
				{Value: "*", Kind: "math", Column: 3, Row: 0},
				{Value: "3", Kind: "integer", Column: 4, Row: 0},
				{Value: "+", Kind: "math", Column: 5, Row: 0},
				{Value: "1", Kind: "integer", Column: 6, Row: 0},
			}

			succedingSyntaticTest(t, input, expected)
		})

		t.Run("Test \"(5+2)*3\"", func(t *testing.T) {
			// Test "(5+2)*3"
			// Test case for the syntax function.
			// The input is a list of tokens made from the string "(5+2)*3".
			// The output is a binary tree.

			expected := &utility.BinaryTreeNode[Symbol]{
				Value: Symbol{token: token{Value: "*", Kind: "math", Column: 6, Row: 0}},
				Left: &utility.BinaryTreeNode[Symbol]{
					Value: Symbol{token: token{Value: "+", Kind: "math", Column: 1, Row: 0}},
					Left:  &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: token{Value: "5", Kind: "integer", Column: 1, Row: 0}}},
					Right: &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: token{Value: "2", Kind: "integer", Column: 3, Row: 0}}},
				},
				Right: &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: token{Value: "3", Kind: "integer", Column: 5, Row: 0}}},
			}

			input := []token{
				{Value: "(", Kind: "parenthesis_open", Column: 0, Row: 0},
				{Value: "5", Kind: "integer", Column: 1, Row: 0},
				{Value: "+", Kind: "math", Column: 2, Row: 0},
				{Value: "2", Kind: "integer", Column: 3, Row: 0},
				{Value: ")", Kind: "parenthesis_close", Column: 4, Row: 0},
				{Value: "*", Kind: "math", Column: 5, Row: 0},
				{Value: "3", Kind: "integer", Column: 6, Row: 0},
			}

			succedingSyntaticTest(t, input, expected)
		})

		t.Run("Test \"3*(5+2)\"", func(t *testing.T) {
			// Test "3*(5+2)"
			// Test case for the syntax function.
			// The input is a list of tokens made from the string "3*(5+2)".
			// The output is a binary tree.

			expected := &utility.BinaryTreeNode[Symbol]{
				Value: Symbol{token: token{Value: "*", Kind: "math", Column: 1, Row: 0}},
				Left:  &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: token{Value: "3", Kind: "integer", Column: 0, Row: 0}}},
				Right: &utility.BinaryTreeNode[Symbol]{
					Value: Symbol{token: token{Value: "+", Kind: "math", Column: 4, Row: 0}},
					Left:  &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: token{Value: "5", Kind: "integer", Column: 2, Row: 0}}},
					Right: &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: token{Value: "2", Kind: "integer", Column: 3, Row: 0}}},
				},
			}
			input := []token{
				{Value: "3", Kind: "integer", Column: 0, Row: 0},
				{Value: "*", Kind: "math", Column: 1, Row: 0},
				{Value: "(", Kind: "parenthesis_open", Column: 2, Row: 0},
				{Value: "5", Kind: "integer", Column: 2, Row: 0},
				{Value: "+", Kind: "math", Column: 3, Row: 0},
				{Value: "2", Kind: "integer", Column: 4, Row: 0},
				{Value: ")", Kind: "parenthesis_close", Column: 5, Row: 0},
			}

			succedingSyntaticTest(t, input, expected)
		})
	})

	t.Run("Logical operations", func(t *testing.T) {

		t.Run("Test \"true == true\"", func(t *testing.T) {
			// Test "true == true"
			// Test case for the syntax function.
			// The input is a list of tokens made from the string "true == true".
			// The output is a binary tree.

			expected := &utility.BinaryTreeNode[Symbol]{
				Value: Symbol{token: "=="},
				Left:  &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: "true"}},
				Right: &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: "true"}},
			}

			input := []token{
				{Value: "true", Kind: "string"},
				{Value: "==", Kind: "logical"},
				{Value: "true", Kind: "string"},
			}

			succedingSyntaticTest(t, input, expected)

		})

		t.Run("Test \"true != false\"", func(t *testing.T) {
			// Test "true != false"
			// Test case for the syntax function.
			// The input is a list of tokens made from the string "true != false".
			// The output is a binary tree.

			expected := &utility.BinaryTreeNode[Symbol]{
				Value: Symbol{token: "!="},
				Left:  &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: "true"}},
				Right: &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: "false"}},
			}

			input := []token{
				{Value: "true", Kind: "string"},
				{Value: "!=", Kind: "logical"},
				{Value: "false", Kind: "string"},
			}

			succedingSyntaticTest(t, input, expected)

		})

		t.Run("Test \"true && false\"", func(t *testing.T) {
			// Test "true && false"
			// Test case for the syntax function.
			// The input is a list of tokens made from the string "true && false".
			// The output is a binary tree.

			expected := &utility.BinaryTreeNode[Symbol]{
				Value: Symbol{token: "&&"},
				Left:  &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: "true"}},
				Right: &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: "false"}},
			}

			input := []token{
				{Value: "true", Kind: "string"},
				{Value: "&&", Kind: "logical"},
				{Value: "false", Kind: "string"},
			}

			succedingSyntaticTest(t, input, expected)

		})

		t.Run("Test \"true || false\"", func(t *testing.T) {
			// Test "true || false"
			// Test case for the syntax function.
			// The input is a list of tokens made from the string "true || false".
			// The output is a binary tree.

			expected := &utility.BinaryTreeNode[Symbol]{
				Value: Symbol{token: "||"},
				Left:  &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: "true"}},
				Right: &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: "false"}},
			}

			input := []token{
				{Value: "true", Kind: "string"},
				{Value: "||", Kind: "logical"},
				{Value: "false", Kind: "string"},
			}

			succedingSyntaticTest(t, input, expected)

		})

		t.Run("Test \"!true\"", func(t *testing.T) {
			// Test "!true"
			// Test case for the syntax function.
			// The input is a list of tokens made from the string "!true".
			// The output is a binary tree.

			expected := &utility.BinaryTreeNode[Symbol]{
				Value: Symbol{token: "!"},
				Right: &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: "true"}},
			}

			input := []token{
				{Value: "!", Kind: "logical"},
				{Value: "true", Kind: "string"},
			}

			succedingSyntaticTest(t, input, expected)

		})

		t.Run("Test \"!(true && false)\"", func(t *testing.T) {
			// Test "!(true && false)"
			// Test case for the syntax function.
			// The input is a list of tokens made from the string "!(true && false)".
			// The output is a binary tree.

			expected := &utility.BinaryTreeNode[Symbol]{
				Value: Symbol{token: "!"},
				Right: &utility.BinaryTreeNode[Symbol]{
					Value: Symbol{token: "&&"},
					Left:  &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: "true"}},
					Right: &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: "false"}},
				},
			}

			input := []token{
				{Value: "!", Kind: "logical"},
				{Value: "(", Kind: "parenthesis_open"},
				{Value: "true", Kind: "string"},
				{Value: "&&", Kind: "logical"},
				{Value: "false", Kind: "string"},
				{Value: ")", Kind: "parenthesis_close"},
			}

			succedingSyntaticTest(t, input, expected)

		})

		t.Run("Test \"true && false || true\"", func(t *testing.T) {
			// Test "true && false || true"
			// Test case for the syntax function.
			// The input is a list of tokens made from the string "true && false || true".
			// The output is a binary tree.

			expected := &utility.BinaryTreeNode[Symbol]{
				Value: Symbol{token: "||"},
				Left: &utility.BinaryTreeNode[Symbol]{
					Value: Symbol{token: "&&"},
					Left:  &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: "true"}},
					Right: &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: "false"}},
				},
				Right: &utility.BinaryTreeNode[Symbol]{
					Value: Symbol{token: "true"},
				},
			}

			input := []token{
				{Value: "true", Kind: "string"},
				{Value: "&&", Kind: "logical"},
				{Value: "false", Kind: "string"},
				{Value: "||", Kind: "logical"},
				{Value: "true", Kind: "string"},
			}

			succedingSyntaticTest(t, input, expected)

		})

		t.Run("Test \"true && (false || true)\"", func(t *testing.T) {
			// Test "true && (false || true)"
			// Test case for the syntax function.
			// The input is a list of tokens made from the string "true && (false || true)".
			// The output is a binary tree.

			expected := &utility.BinaryTreeNode[Symbol]{
				Value: Symbol{token: "&&"},
				Left:  &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: "true"}},
				Right: &utility.BinaryTreeNode[Symbol]{
					Value: Symbol{token: "||"},
					Left:  &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: "false"}},
					Right: &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: "true"}},
				},
			}

			input := []token{
				{Value: "true", Kind: "string"},
				{Value: "&&", Kind: "logical"},
				{Value: "(", Kind: "parenthesis_open"},
				{Value: "false", Kind: "string"},
				{Value: "||", Kind: "logical"},
				{Value: "true", Kind: "string"},
				{Value: ")", Kind: "parenthesis_close"},
			}

			succedingSyntaticTest(t, input, expected)

		})

		t.Run("Test \"true && !(false || true)\"", func(t *testing.T) {
			// Test "true && !(false || true)"
			// Test case for the syntax function.
			// The input is a list of tokens made from the string "true && !(false || true)".
			// The output is a binary tree.

			expected := &utility.BinaryTreeNode[Symbol]{
				Value: Symbol{token: "&&"},
				Left:  &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: "true"}},
				Right: &utility.BinaryTreeNode[Symbol]{
					Value: Symbol{token: "!"},
					Right: &utility.BinaryTreeNode[Symbol]{
						Value: Symbol{token: "||"},
						Left:  &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: "false"}},
						Right: &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: "true"}},
					},
				},
			}

			input := []token{
				{Value: "true", Kind: "string"},
				{Value: "&&", Kind: "logical"},
				{Value: "!", Kind: "logical"},
				{Value: "(", Kind: "parenthesis_open"},
				{Value: "false", Kind: "string"},
				{Value: "||", Kind: "logical"},
				{Value: "true", Kind: "string"},
				{Value: ")", Kind: "parenthesis_close"},
			}

			succedingSyntaticTest(t, input, expected)

		})

		t.Run("Test \"!(false || true) && false\"", func(t *testing.T) {
			// Test "!(false || true) && false"
			// Test case for the syntax function.
			// The input is a list of tokens made from the string "!(false || true) && false".
			// The output is a binary tree.

			expected := &utility.BinaryTreeNode[Symbol]{
				Value: Symbol{token: "&&"},
				Left: &utility.BinaryTreeNode[Symbol]{
					Value: Symbol{token: "!"},
					Right: &utility.BinaryTreeNode[Symbol]{
						Value: Symbol{token: "||"},
						Left:  &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: "false"}},
						Right: &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: "true"}},
					},
				},
				Right: &utility.BinaryTreeNode[Symbol]{
					Value: Symbol{token: "false"},
				},
			}

			input := []token{
				{Value: "!", Kind: "logical"},
				{Value: "(", Kind: "parenthesis_open"},
				{Value: "false", Kind: "string"},
				{Value: "||", Kind: "logical"},
				{Value: "true", Kind: "string"},
				{Value: ")", Kind: "parenthesis_close"},
				{Value: "&&", Kind: "logical"},
				{Value: "false", Kind: "string"},
			}

			succedingSyntaticTest(t, input, expected)

		})

		t.Run("Test \"!(false || !(true || false)) && false\"", func(t *testing.T) {
			// Test "!(false || true) && false"
			// Test case for the syntax function.
			// The input is a list of tokens made from the string "!(false || true) && false".
			// The output is a binary tree.

			expected := &utility.BinaryTreeNode[Symbol]{
				Value: Symbol{token: "&&"},
				Left: &utility.BinaryTreeNode[Symbol]{
					Value: Symbol{token: "!"},
					Right: &utility.BinaryTreeNode[Symbol]{
						Value: Symbol{token: "||"},
						Left:  &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: "false"}},
						Right: &utility.BinaryTreeNode[Symbol]{
							Value: Symbol{token: "!"},
							Right: &utility.BinaryTreeNode[Symbol]{
								Value: Symbol{token: "||"},
								Left:  &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: "true"}},
								Right: &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: "false"}},
							},
						},
					},
				},
				Right: &utility.BinaryTreeNode[Symbol]{
					Value: Symbol{token: "false"},
				},
			}

			input := []token{
				{Value: "!", Kind: "logical"},
				{Value: "(", Kind: "parenthesis_open"},
				{Value: "false", Kind: "string"},
				{Value: "||", Kind: "logical"},
				{Value: "!", Kind: "logical"},
				{Value: "(", Kind: "parenthesis_open"},
				{Value: "true", Kind: "string"},
				{Value: "||", Kind: "logical"},
				{Value: "false", Kind: "string"},
				{Value: ")", Kind: "parenthesis_close"},
				{Value: ")", Kind: "parenthesis_close"},
				{Value: "&&", Kind: "logical"},
				{Value: "false", Kind: "string"},
			}

			succedingSyntaticTest(t, input, expected)

		})

		t.Run("Test numeric comparison", func(t *testing.T) {

			t.Run("Test \"5 > 2\"", func(t *testing.T) {
				// Test "5 > 2"
				// Test case for the syntax function.
				// The input is a list of tokens made from the string "5 > 2".
				// The output is a binary tree.

				expected := &utility.BinaryTreeNode[Symbol]{
					Value: Symbol{token: ">"},
					Left:  &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: "5"}},
					Right: &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: "2"}},
				}

				input := []token{
					{Value: "5", Kind: "integer"},
					{Value: ">", Kind: "logical"},
					{Value: "2", Kind: "integer"},
				}

				succedingSyntaticTest(t, input, expected)
			})

			t.Run("Test \"5 < 2\"", func(t *testing.T) {
				// Test "5 < 2"
				// Test case for the syntax function.
				// The input is a list of tokens made from the string "5 < 2".
				// The output is a binary tree.

				expected := &utility.BinaryTreeNode[Symbol]{
					Value: Symbol{token: "<"},
					Left:  &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: "5"}},
					Right: &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: "2"}},
				}

				input := []token{
					{Value: "5", Kind: "integer"},
					{Value: "<", Kind: "logical"},
					{Value: "2", Kind: "integer"},
				}

				succedingSyntaticTest(t, input, expected)
			})

			t.Run("Test \"5 >= 2\"", func(t *testing.T) {
				// Test "5 >= 2"
				// Test case for the syntax function.
				// The input is a list of tokens made from the string "5 >= 2".
				// The output is a binary tree.

				expected := &utility.BinaryTreeNode[Symbol]{
					Value: Symbol{token: ">="},
					Left:  &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: "5"}},
					Right: &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: "2"}},
				}

				input := []token{
					{Value: "5", Kind: "integer"},
					{Value: ">=", Kind: "logical"},
					{Value: "2", Kind: "integer"},
				}

				succedingSyntaticTest(t, input, expected)
			})

			t.Run("Test \"5 <= 2\"", func(t *testing.T) {
				// Test "5 <= 2"
				// Test case for the syntax function.
				// The input is a list of tokens made from the string "5 <= 2".
				// The output is a binary tree.

				expected := &utility.BinaryTreeNode[Symbol]{
					Value: Symbol{token: "<="},
					Left:  &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: "5"}},
					Right: &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: "2"}},
				}

				input := []token{
					{Value: "5", Kind: "integer"},
					{Value: "<=", Kind: "logical"},
					{Value: "2", Kind: "integer"},
				}

				succedingSyntaticTest(t, input, expected)
			})

			t.Run("Test \"5 == 2\"", func(t *testing.T) {
				// Test "5 == 2"
				// Test case for the syntax function.
				// The input is a list of tokens made from the string "5 == 2".
				// The output is a binary tree.

				expected := &utility.BinaryTreeNode[Symbol]{
					Value: Symbol{token: "=="},
					Left:  &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: "5"}},
					Right: &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: "2"}},
				}

				input := []token{
					{Value: "5", Kind: "integer"},
					{Value: "==", Kind: "logical"},
					{Value: "2", Kind: "integer"},
				}

				succedingSyntaticTest(t, input, expected)
			})

			t.Run("Test \"5 != 2\"", func(t *testing.T) {
				// Test "5 != 2"
				// Test case for the syntax function.
				// The input is a list of tokens made from the string "5 != 2".
				// The output is a binary tree.

				expected := &utility.BinaryTreeNode[Symbol]{
					Value: Symbol{token: "!="},
					Left:  &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: "5"}},
					Right: &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: "2"}},
				}

				input := []token{
					{Value: "5", Kind: "integer"},
					{Value: "!=", Kind: "logical"},
					{Value: "2", Kind: "integer"},
				}

				succedingSyntaticTest(t, input, expected)
			})
		})
	})

	t.Run("Assignment and Call", func(t *testing.T) {

		t.Run("Test \"a := 5\"", func(t *testing.T) {
			// Test "a = 5"
			// Test case for the syntax function.
			// The input is a list of tokens made from the string "a = 5".
			// The output is a binary tree.

			expected := &utility.BinaryTreeNode[Symbol]{
				Value: Symbol{token: "="},
				Left:  &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: "a"}},
				Right: &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: "5"}},
			}

			input := []token{
				{Value: "a", Kind: "identifier"},
				{Value: ":=", Kind: "assignment"},
				{Value: "5", Kind: "integer"},
			}

			succedingSyntaticTest(t, input, expected)
		})

		t.Run("Test \"$a := $b\"", func(t *testing.T) {
			// Test "$a = $b"
			// Test case for the syntax function.
			// The input is a list of tokens made from the string "$a = $b".
			// The output is a binary tree.

			expected := &utility.BinaryTreeNode[Symbol]{
				Value: Symbol{token: "="},
				Left: &utility.BinaryTreeNode[Symbol]{
					Value: Symbol{token: "$"},
					Right: &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: "a"}}},
				Right: &utility.BinaryTreeNode[Symbol]{
					Value: Symbol{token: "$"},
					Right: &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: "b"}}},
			}

			input := []token{
				{Value: "$", Kind: "call"},
				{Value: "a", Kind: "string"},
				{Value: ":=", Kind: "assignment"},
				{Value: "$", Kind: "call"},
				{Value: "b", Kind: "string"},
			}

			succedingSyntaticTest(t, input, expected)

		})
	})

	t.Run("User defined functions", func(t *testing.T) {

		var build_general = func(name string) { cli_tooler.BuildWord(name).SetPrecedence(1).Build() }
		words := []string{"word1", "word2", "word3"}

		for _, word := range words {
			build_general(word)
		}

		t.Run("Test \"word1\"", func(t *testing.T) {

			// Test "word1"
			// Test case for the syntax function.
			// The input is a list of tokens made from the string "word1".
			// The output is a binary tree.

			expected := &utility.BinaryTreeNode[Symbol]{
				Value: Symbol{token: "word1"},
			}

			input := []token{
				{Value: "word1", Kind: "name"},
			}

			succedingSyntaticTest(t, input, expected)
		})

		t.Run("Test \"word1 batata cenoura\"", func(t *testing.T) {
			// Test "word1 batata cenoura"
			// Test case for the syntax function.
			// The input is a list of tokens made from the string "word1 batata cenoura".
			// The output is a binary tree.

			expected := &utility.BinaryTreeNode[Symbol]{
				Value: Symbol{token: "word1"},
				Left:  &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: "batata"}},
				Right: &utility.BinaryTreeNode[Symbol]{Value: Symbol{token: "cenoura"}},
			}

			input := []token{
				{Value: "word1", Kind: "name"},
				{Value: "batata", Kind: "value"},
				{Value: "cenoura", Kind: "value"},
			}

			succedingSyntaticTest(t, input, expected)
		})
	})

	t.Run("Test Multi-line inputs", func(t *testing.T) {

	})
	t.Run("Error cases", func(t *testing.T) {

	})

}

func succedingSyntaticTest(t *testing.T, input []token, expected *utility.BinaryTreeNode[Symbol]) {
	t.Helper()

	// Test "input"
	// Test case for the syntax function.

	tree, err := syntax(input)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if !isIdentical(tree, expected) {
		t.Errorf("Expected: %s\nGot: %s", formatTree(expected, 0), formatTree(tree, 0))
	}

}

func formatTree(tree *utility.BinaryTreeNode[Symbol], tabs int) string {
	if tree == nil {
		return ""
	}
	tabsStr := ""
	for i := 0; i < tabs; i++ {
		tabsStr += "\t"
	}

	right, left := formatTree(tree.Right, tabs+1), formatTree(tree.Left, tabs+1)

	if left != "" {
		return fmt.Sprintf("%svalue: %v\n%sleft: %v\n%sright: %v", tabsStr, tree.Value, tabsStr, left, tabsStr, right)
	}
	if right != "" {
		return fmt.Sprintf("%svalue: %v\n%sright: %v", tabsStr, tree.Value, tabsStr, right)
	}
	if right == "" && left == "" {
		return fmt.Sprintf("%svalue: %v", tabsStr, tree.Value)
	}

	return fmt.Sprintf("%svalue: %v\n%sleft: %v\n%sright: %v", tabsStr, tree.Value, tabsStr, formatTree(tree.Left, tabs+1), tabsStr, formatTree(tree.Right, tabs+1))
}

func isIdentical(tree1, tree2 *utility.BinaryTreeNode[Symbol]) bool {

	if tree1 == nil && tree2 == nil {
		return true
	}

	if tree1 == nil || tree2 == nil {
		return false
	}

	return tree1.Value.token == tree2.Value.token && isIdentical(tree1.Left, tree2.Left) && isIdentical(tree1.Right, tree2.Right)
}

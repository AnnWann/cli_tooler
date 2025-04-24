package parser

import (
	errorhandler "cli_tooler/internal/error_handler"
	"cli_tooler/internal/utility"
)

type SyntaxErrorOutput struct {
	Head              *utility.BinaryTreeNode[Symbol]
	CommandStack      *utility.Stack[utility.BinaryTreeNode[Symbol]]
	OutputStack       *utility.Stack[utility.BinaryTreeNode[Symbol]]
	Tokens            []token
	ProblematicSymbol Symbol
}

func syntax(tokens []token) (*utility.BinaryTreeNode[Symbol], error) {
	head, commandStack, outputStack := &utility.BinaryTreeNode[Symbol]{}, utility.Stack[utility.BinaryTreeNode[Symbol]]{}, utility.Stack[utility.BinaryTreeNode[Symbol]]{}

	lastSymbol := Symbol{nextIsValid: func(string) bool { return true }}
	needsClosingParenthesis := 0
	for _, token := range tokens {
		symbol := createSymbol(token)
		if !lastSymbol.nextIsValid(symbol.token.Value) {
			return nil, errorhandler.HandleError("syntaxError", SyntaxErrorOutput{
				Head:              head,
				CommandStack:      &commandStack,
				OutputStack:       &outputStack,
				Tokens:            tokens,
				ProblematicSymbol: symbol,
			})
		}

		if token.Kind == "parenthesis_open" {
			commandStack.Push(utility.BinaryTreeNode[Symbol]{Value: createSymbol(token)})
			needsClosingParenthesis++
		} else if token.Kind == "parenthesis_close" {
			if !commandStack.IsEmpty() {
				subtree := createSubTree(commandStack.Pop().Value, nil, &commandStack, &outputStack)
				outputStack.Push(*subtree)
			}
			needsClosingParenthesis--

		} else if token.Kind != "string" && token.Kind != "integer" && token.Kind != "float" {
			if !commandStack.IsEmpty() && symbol.LessOrEqualInPrecedence(commandStack.Peek().Value) {
				subtree := createSubTree(commandStack.Pop().Value, nil, &commandStack, &outputStack)
				outputStack.Push(*subtree)
			}
			commandStack.Push(utility.BinaryTreeNode[Symbol]{Value: symbol})
		} else {
			outputStack.Push(utility.BinaryTreeNode[Symbol]{Value: createSymbol(token)})
		}
		lastSymbol = symbol
	}

	if needsClosingParenthesis > 0 {
		return nil, errorhandler.HandleError("parenthesis", SyntaxErrorOutput{
			Head:         head,
			CommandStack: &commandStack,
			OutputStack:  &outputStack,
			Tokens:       tokens,
		})
	}

	for !commandStack.IsEmpty() {
		head = updateTree(head, &commandStack, &outputStack)
	}

	if head.Value.token.Value == "" {
		temp := outputStack.Pop()
		head = &temp
	}

	return head, nil
}

func createSubTree(
	symbol Symbol,
	tree *utility.BinaryTreeNode[Symbol],
	commandStack *utility.Stack[utility.BinaryTreeNode[Symbol]],
	outputStack *utility.Stack[utility.BinaryTreeNode[Symbol]]) *utility.BinaryTreeNode[Symbol] {
	if symbol.token.Value == "(" || symbol.token.Value == ")" {
		return tree
	}
	if tree == nil || tree.Value.token.Value == "" {
		tree = &utility.BinaryTreeNode[Symbol]{Value: symbol}
		if symbol.isUnary {
			right := outputStack.Pop()
			rightPtr := &right
			if right.Value.token.Value == "" {
				rightPtr = nil
			}
			tree.Right = rightPtr
		} else {
			right, left := outputStack.Pop(), outputStack.Pop()
			rightPtr, leftPtr := &right, &left
			if left.Value.token.Value == "" {
				leftPtr = nil
			}
			if right.Value.token.Value == "" {
				rightPtr = nil
			}
			tree.Right, tree.Left = rightPtr, leftPtr
		}
	} else {
		if symbol.isUnary {
			tree = &utility.BinaryTreeNode[Symbol]{Value: symbol, Right: tree}
		} else {
			subtree, left := tree, outputStack.Pop()
			leftPtr := &left
			if left.Value.token.Value == "" {
				leftPtr = nil
			}
			tree = &utility.BinaryTreeNode[Symbol]{Value: symbol, Left: leftPtr, Right: subtree}
		}
	}

	if !commandStack.IsEmpty() && commandStack.Peek().Value.LessOrEqualInPrecedence(symbol) {
		return createSubTree(commandStack.Pop().Value, tree, commandStack, outputStack)
	}

	return tree
}

func updateTree(head *utility.BinaryTreeNode[Symbol], commandStack *utility.Stack[utility.BinaryTreeNode[Symbol]], outputStack *utility.Stack[utility.BinaryTreeNode[Symbol]]) *utility.BinaryTreeNode[Symbol] {
	symbol := commandStack.Pop()

	if head == nil || head.Value.token.Value == "" {
		output := outputStack.Pop()
		if symbol.Value.isUnary {
			head = &utility.BinaryTreeNode[Symbol]{Value: symbol.Value, Right: &output}
		} else {
			left := outputStack.Pop()
			ptr := &left
			if left.Value.token.Value == "" {
				ptr = nil
			}
			head = &utility.BinaryTreeNode[Symbol]{Value: symbol.Value, Left: ptr, Right: &output}
		}
	} else {
		if symbol.Value.isUnary {
			head = &utility.BinaryTreeNode[Symbol]{Value: symbol.Value, Right: head}
		} else {
			output := outputStack.Pop()
			subtree := head
			head = &utility.BinaryTreeNode[Symbol]{Value: symbol.Value, Left: &output, Right: subtree}
		}
	}

	return head
}

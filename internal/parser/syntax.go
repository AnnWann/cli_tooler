package parser

import (
	"cli_tooler/internal/dictionary"
	"fmt"
)

func syntax(tokens []token) (*BinaryTreeNode, error) {
	head, symbolStack, outputStack := &BinaryTreeNode{}, stack{}, stack{}

	lastSymbol := Symbol{nextIsValid: func(string) bool { return true }}
	needsClosingParenthesis := 0
	for i, token := range tokens {
		symbol := createSymbol(token.value)
		if !lastSymbol.nextIsValid(symbol.sign) {
			errorMessage := dictionary.GetString("syntaxError") + "\n" + symbol.errorMessage
			return nil, fmt.Errorf(errorMessage, tokens[i-1].value, tokens[i].value, symbol.sign)
		}

		if token.kind == "parenthesis_open" {
			symbolStack.push(BinaryTreeNode{Value: createSymbol(token.value)})
			needsClosingParenthesis++
		} else if token.kind == "parenthesis_close" {
			for !symbolStack.isEmpty() && symbolStack.peek().Value.(Symbol).sign != "(" {
				subtree := createSubTree(symbolStack.pop().Value.(Symbol), nil, &symbolStack, &outputStack)
				outputStack.push(*subtree)
			}
			symbolStack.pop()
			needsClosingParenthesis--

		} else if token.kind != "value" {
			symbol := createSymbol(token.value)
			if !symbolStack.isEmpty() && symbolStack.peek().Value.lessOrEqualInPrecedence(symbol) {
				subtree := createSubTree(symbolStack.pop().Value.(Symbol), nil, &symbolStack, &outputStack)
				outputStack.push(*subtree)
			}
			symbolStack.push(BinaryTreeNode{Value: symbol})
		} else {
			outputStack.push(BinaryTreeNode{Value: createSymbol(token.value)})
		}
	}

	if needsClosingParenthesis > 0 {
		return nil, fmt.Errorf(dictionary.GetString("syntaxError") + "\n" + dictionary.GetString("missingClosingParenthesis"))
	}

	for !symbolStack.isEmpty() {
		updateTree(head, &symbolStack, &outputStack)
	}

	return head, nil
}

func createSubTree(symbol Symbol, tree *BinaryTreeNode, symbolStack *stack, outputStack *stack) *BinaryTreeNode {
	if tree == nil {
		tree = &BinaryTreeNode{}
		right, left := outputStack.pop(), outputStack.pop()
		tree.Right, tree.Left = &right, &left
	} else {
		subtree, left := tree, outputStack.pop()
		tree = &BinaryTreeNode{Value: symbol, Left: &left, Right: subtree}
	}

	if !symbolStack.isEmpty() && symbolStack.peek().Value.lessOrEqualInPrecedence(symbol) {
		return createSubTree(symbolStack.pop().Value.(Symbol), tree, symbolStack, outputStack)
	}

	return tree
}

func updateTree(head *BinaryTreeNode, symbolStack *stack, outputStack *stack) {
	symbol, output := symbolStack.pop(), outputStack.pop()

	if head == nil {
		left := outputStack.pop()
		head = &BinaryTreeNode{Value: symbol.Value.(Symbol), Left: &left, Right: &output}
	} else {
		subtree := head
		head = &BinaryTreeNode{Value: symbol.Value.(Symbol), Left: &output, Right: subtree}
	}
}

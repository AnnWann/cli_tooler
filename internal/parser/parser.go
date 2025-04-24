package parser

import "cli_tooler/internal/utility"

func Parse(str string) (*utility.BinaryTreeNode[Symbol], error) {

	tokens, err := lexical(str)
	if err != nil {
		return nil, err
	}

	binaryTree, err := syntax(tokens)
	if err != nil {
		return nil, err
	}

	return binaryTree, nil
}

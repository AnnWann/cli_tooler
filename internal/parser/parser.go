package parser

func Parse(str string) (*BinaryTreeNode, error) {

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

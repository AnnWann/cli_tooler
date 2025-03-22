package cli_tooler

type Word struct {
	Fn           any
	Precedence   int
	NextIsValid  func(string) bool
	ErrorMessage string
}

var ReservedWords map[string]Word

//var CLI_TOOLER CLI_TOOLER = CLI_TOOLER{}

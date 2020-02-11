package main

import (
	"os"
	"strings"

	"./lexer"
	"./zerror"
)

func main() {
	lex := lexer.Lexer{}
	if len(os.Args) == 1 {
		zerror.Fatal("UnknownFile", "the file you're passing doesn't exist", "InterpreterRuntime")
	} else if os.Args[1] == "-v" {
		lex.Parser(os.Args[2], true)
		lex.Analyze()
	} else if strings.Contains(os.Args[1], ".zl") != true {
		zerror.Fatal("InvalidFileType", "the file you're trying to compile isn't supported by the compiler", "InterpreterStartup")
	} else {
		lex.Parser(os.Args[1], false)
		lex.Analyze()
	}

}

package main

import (
	"flag"

	"./lexer"
)

var (
	verbose  bool
	filename string
)

func init() {
	flag.BoolVar(&verbose, "v", false, "Wether or not to run the interpreter in verbose mode")
	flag.BoolVar(&verbose, "verbose", false, "Wether or not to run the interpreter in verbose mode")
	flag.StringVar(&filename, "f", "main.shh", "The file that you want the interpreter to run")
	flag.StringVar(&filename, "file", "main.shh", "The file that you want the interpreter to run")
	flag.Parse()
}

func isVerbose(v bool) (reponse bool) {
	if v {
		return true
	} else {
		return false
	}
}

func main() {
	lex := lexer.Lexer{}
	lex.Parser(filename, isVerbose(verbose))
	lex.Analyze()
}

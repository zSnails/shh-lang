package lexer

import (
	"fmt"
	"io/ioutil"
	"strings"

	shherror "../shh-error"
)

//Token the tokens we'll use
type Token struct {
	Type  string
	Value string
}

var (
	//ttplus tt_plus
	ttplus = "PLUS"
	//ttminus tt_plus
	ttminus = "MINUS"
	//ttmrthan tt_plus
	ttmrthan = "MRTHAN"
	//ttlsrthan tt_plus
	ttlsrthan = "LSRTHAN"
	output    []string
	result    string
)

//Lexer lexer type def
type Lexer struct {
	Text        string
	Pos         uint16
	CurrentChar []int
	Verboose    bool
	Word        string
}

//Analyze analyze the code and interpret it
func (l *Lexer) Analyze() {
	var verbosedOut = arr("START")
	l.Pos = 0
	l.CurrentChar = make([]int, 30000)
	lines := strings.SplitN(l.Text, "\n", -1)
	for k := 0; k < len(lines); k++ {
		textPerLine := strings.SplitN(lines[k], ".>", -1)
		currentLine := strings.SplitN(strings.Join(textPerLine, ""), "", -1)
		l.Word += string(len(currentLine))
		verbosedOut = arr("END")
	}
	if l.Verboose {
		fmt.Println("Verbose: ", verbosedOut)
	}
	println(l.Word)
}

//Parser parse the source code
func (l *Lexer) Parser(filename string, verboose bool) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		shherror.Fatal("FileNotFound", "The file you tried to compile does not exist", "CompileTime")
	}
	l.Text = string(data)
	l.Verboose = verboose
}

//Arr create an arr and then output it to the console, only if the current mode is set to verbose
func arr(token string) []string {
	output = append(output, token)
	return output
}

func checkImports(i []string) {
	println(len(i))
	fmt.Printf("%v", i)
	for is := range i {
		println(i[is])
	}
}

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
	//TT_PLUS tt_plus
	TT_PLUS = "PLUS"
	//TT_MINUS tt_plus
	TT_MINUS = "MINUS"
	//TT_MRTA+HAN tt_plus
	TT_MRTHAN = "MRTHAN"
	//TT_LSRTHAN tt_plus
	TT_LSRTHAN = "LSRTHAN"
	output     []string
	result     string
)

//Lexer lexer type def
type Lexer struct {
	Text        string
	Pos         uint16
	CurrentChar []int
	Verboose    bool
}

//Analyze analyze the code and interpret it
func (l *Lexer) Analyze() {
	var verbosedOut = arr("START")
	l.Pos = 0
	l.CurrentChar = make([]int, 30000)
	for i := 0; i < len(l.Text); i++ {
		switch l.Text[i] {
		case 'i':
			iSection := strings.SplitN(l.Text, "i ", -1)
			valuesText := []string{}
			for i := range iSection {
				text := string(iSection[i])
				valuesText = append(valuesText, text)
			}
			result += strings.Join(valuesText, " ")
			result = strings.Replace(result, "?", "", -1)
			iStatements := strings.SplitN(result, " | ", -1)
			checkImports(iStatements)
			break
		case '+':
			verbosedOut = arr(TT_PLUS)
			l.CurrentChar[l.Pos]++
			break
		case '-':
			verbosedOut = arr(TT_MINUS)
			l.CurrentChar[l.Pos]--
			break
		case '>':
			verbosedOut = arr(TT_MRTHAN)
			l.CurrentChar[l.Pos] = 0
			l.Pos++
			break
		case '<':
			verbosedOut = arr(TT_LSRTHAN)
			l.CurrentChar[l.Pos] = 0
			l.Pos--
			break
		case '.':
			verbosedOut = arr("PRINT_STATEMENT")
			valuesText := []string{}
			for i := range l.CurrentChar {
				text := string(l.CurrentChar[i])
				valuesText = append(valuesText, text)
			}
			result += strings.Join(valuesText, "")
			break
		case '%':
			verbosedOut = arr("RUNNABLE_MODULE")
			print(result)
			break
		}
	}
	if l.Verboose {
		fmt.Println("Verbose: ", verbosedOut)
	}
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

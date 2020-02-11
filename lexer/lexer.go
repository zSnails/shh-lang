package lexer

import (
	"io/ioutil"

	"../zerror"
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
)

//Lexer lexer type def
type Lexer struct {
	Text        string
	Pos         uint16
	CurrentChar []uint16
	Verboose    bool
}

//Analyze analyze the code and interpret it
func (l *Lexer) Analyze() {
	var verbosedOut = Arr("START")
	l.Pos = 0
	l.CurrentChar = make([]uint16, 3000)
	//already handled file reading and turned it into a string
	for i := 0; i < len(l.Text); i++ {
		if l.Text[i] == '+' {
			verbosedOut = Arr(TT_PLUS)
			l.CurrentChar[l.Pos]++
		} else if l.Text[i] == '-' {
			verbosedOut = Arr(TT_MINUS)
			l.CurrentChar[l.Pos]--
		} else if l.Text[i] == '>' {
			verbosedOut = Arr(TT_LSRTHAN)
			l.CurrentChar[l.Pos] = 0
			l.Pos++
		} else if l.Text[i] == '<' {
			verbosedOut = Arr(TT_MRTHAN)
			l.CurrentChar[l.Pos] = 0
			l.Pos--
		} else if l.Text[i] == '.' {
			verbosedOut = Arr("PRINT_STATEMENT")
			for c := 0; c < len(l.CurrentChar); c++ {
				txt := string(l.CurrentChar[c])
				print(txt)
			}
		}
	}
	if l.Verboose {
		println(verbosedOut)
	}
}

//Parser parse the source code
func (l *Lexer) Parser(filename string, verboose bool) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		zerror.Fatal("FileNotFound", "The file you tried to compile does not exist", "CompileTime")
	}
	l.Text = string(data)
	l.Verboose = verboose
}

func Arr(token string) []string {
	var output []string
	output = append(output, token)
	return output
}

package shherror

import (
	"fmt"
	"os"
)

//Error error type def
type Error struct {
	Name, Content, At string
}

//Fatal fatal error
func Fatal(name, content, at string) {
	var err Error
	err.Name = name
	err.Content = content
	err.At = at

	fmt.Println(err.Name + "\n\n\t - " + err.Content + "\n\nAt: " + err.At)
	os.Exit(1)
}

//Warning log a warning into the console
func Warning(name, content string) {
	var err Error
	err.Name = name
	err.Content = content

	fmt.Println(err.Name + "\n\n\t - " + err.Content)
}

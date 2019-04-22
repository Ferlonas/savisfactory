package main

import (
	"fmt"

	"github.com/ferlonas/savisfactory/parser"
)

func main() {
	p := parser.NewParser("save1.sav")
	f, err := p.ParseFile()
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
	f.Dump()
}

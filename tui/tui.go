package tui

import (
	"fmt"
	"github.com/nathantebbs/gpac/calculator"
	"github.com/nathantebbs/gpac/student"
)

func init() {
	fmt.Println("Hello, Tui!")
	calculator.init()
	student.init()
}

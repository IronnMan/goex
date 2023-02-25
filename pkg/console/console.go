package console

import (
	"fmt"
	"github.com/mgutz/ansi"
	"os"
)

func Success(msg string) {
	colorOut(msg, "green")
}

func Error(msg string) {
	colorOut(msg, "red")
}

func Warning(msg string) {
	colorOut(msg, "yellow")
}

func ExitIf(err error) {
	if err != nil {
		Exit(err.Error())
	}
}

func Exit(msg string) {
	Error(msg)
	os.Exit(1)
}

func colorOut(msg string, color string) {
	fmt.Fprintln(os.Stdout, ansi.Color(msg, color))
}

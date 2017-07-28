package cliwriter

import "github.com/fatih/color"
import "os"
import "fmt"

var errWriter *color.Color
var tipWriter *color.Color
var successWriter *color.Color

func init() {
	errWriter = color.New(color.FgRed)
	tipWriter = color.New(color.FgBlue)
	successWriter = color.New(color.FgGreen)
}

func Errorf(format string, args ...interface{}) {
	errWriter.Fprintf(os.Stderr, format, args)
}
func Error(content string) {
	errWriter.Fprint(os.Stderr, content)
}
func Errorln(content string) {
	errWriter.Fprintln(os.Stderr, content)
}
func Infof(format string, args ...interface{}) {
	fmt.Printf(format, args)
}
func Info(content string) {
	fmt.Print(content)
}
func Infoln(content string) {
	fmt.Println(content)
}

func Tipf(format string, args ...interface{}) {
	tipWriter.Fprintf(os.Stdout, format)
}
func Tip(content string) {
	tipWriter.Fprint(os.Stdout, content)
}
func Tipln(content string) {
	tipWriter.Fprintln(os.Stdout, content)
}

func Successf(format string, args ...interface{}) {
	successWriter.Fprintf(os.Stdout, format, args)
}

func Success(content string) {
	successWriter.Fprint(os.Stdout, content)
}

func Successln(content string) {
	successWriter.Fprintln(os.Stdout, content)
}

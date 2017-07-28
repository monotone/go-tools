package cliwriter

import "github.com/fatih/color"
import "os"
import "fmt"

var (
	errWriter   *color.Color
	ErrorColor  = color.FgRed
	ErrorPrefix = ""

	tipWriter *color.Color
	TipColor  = color.FgBlue
	TipPrefix = "==> "

	successWriter *color.Color
	SuccessPrefix = ""
	SuccessColor  = color.FgGreen
)

func init() {
	errWriter = color.New(color.FgRed)
	tipWriter = color.New(color.FgBlue)
	successWriter = color.New(color.FgGreen)
}

func ResetColor() {
	errWriter.Add(ErrorColor)

	tipWriter.Add(TipColor)

	successWriter.Add(SuccessColor)
}

func Errorf(format string, args ...interface{}) {
	errWriter.Fprintf(os.Stderr, ErrorPrefix+format, args)
}
func Error(content string) {
	errWriter.Fprint(os.Stderr, ErrorPrefix+content)
}
func Errorln(content string) {
	errWriter.Fprintln(os.Stderr, ErrorPrefix+content)
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
	tipWriter.Fprintf(os.Stdout, TipPrefix+format, args)
}
func Tip(content string) {
	tipWriter.Fprint(os.Stdout, TipPrefix+content)
}
func Tipln(content string) {
	tipWriter.Fprintln(os.Stdout, TipPrefix+content)
}

func Successf(format string, args ...interface{}) {
	successWriter.Fprintf(os.Stdout, SuccessPrefix+format, args)
}

func Success(content string) {
	successWriter.Fprint(os.Stdout, SuccessPrefix+content)
}

func Successln(content string) {
	successWriter.Fprintln(os.Stdout, SuccessPrefix+content)
}

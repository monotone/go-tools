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
	errWriter.Fprintf(os.Stderr, ErrorPrefix+format, args...)
}
func Error(args ...interface{}) {
	errWriter.Fprint(os.Stderr, ErrorPrefix)
	errWriter.Fprint(os.Stderr, args...)
}
func Errorln(args ...interface{}) {
	errWriter.Fprint(os.Stderr, ErrorPrefix)
	errWriter.Fprintln(os.Stderr, args...)
}

func Infof(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
func Info(args ...interface{}) {
	fmt.Print(args...)
}
func Infoln(args ...interface{}) {
	fmt.Println(args...)
}

func Tipf(format string, args ...interface{}) {
	tipWriter.Fprintf(os.Stdout, TipPrefix+format, args...)
}
func Tip(args ...interface{}) {
	tipWriter.Fprint(os.Stderr, TipPrefix)
	tipWriter.Fprint(os.Stderr, args...)
}
func Tipln(args ...interface{}) {
	tipWriter.Fprint(os.Stderr, TipPrefix)
	tipWriter.Fprintln(os.Stderr, args...)
}

func Successf(format string, args ...interface{}) {
	successWriter.Fprintf(os.Stdout, SuccessPrefix+format, args...)
}

func Success(args ...interface{}) {
	successWriter.Fprint(os.Stderr, SuccessPrefix)
	successWriter.Fprint(os.Stderr, args...)
}

func Successln(args ...interface{}) {
	successWriter.Fprint(os.Stderr, SuccessPrefix)
	successWriter.Fprintln(os.Stderr, args...)
}

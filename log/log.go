package log

import "fmt"

var DebugOn bool

func Debug(msg string) {
	if !DebugOn  {
		return
	}
	fmt.Println("[debug]", msg)
}

func Error(msg string, err error) {
	fmt.Println("[error]", msg, err)
}

func AppError(msg string) {
	fmt.Println("[error]", msg)
}

func Info(msg string) {
	fmt.Println("[info]", msg)
}

func Warn(msg string) {
	fmt.Println("[warn]", msg)
}
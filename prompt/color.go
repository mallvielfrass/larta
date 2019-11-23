package prompt

import (
	"fmt"
	"strconv"
)

type color struct {
	red, green, blue int
}

func colorized(str string, code color) string {
	red := strconv.Itoa(code.red)
	green :=strconv.Itoa(code.green)
	blue :=strconv.Itoa(code.blue)
	s := fmt.Sprintf("\x1b[38;2;%s;%s;%sm%s\x1b[0m", red, green, blue, str)
	return s
}

func magenta(str string) string {
	code := color{231, 0, 247}
	return colorized(str, code) 
}

func green(str string) string {
	code := color{0, 250, 0}
	return colorized(str, code)
}

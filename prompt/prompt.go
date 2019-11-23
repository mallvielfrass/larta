package prompt

import (
	"fmt"
	"os"
	"os/user"
	"strings"
)

func PrintPrompt() {
	usr, err := user.Current(); loged(err)
	home := usr.HomeDir

	dir, err := os.Getwd(); loged(err)

	dir = strings.Replace(dir, home, "~", 1)

	python := green(GetPython())
	pwd := magenta(dir)

	promptStr := fmt.Sprintf("on %s with %s", pwd, python)
	fmt.Println(promptStr)
	fmt.Println(magenta(" ~>"))
}

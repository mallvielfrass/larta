package main

import (
	"fmt"
	"os"
	"os/user"
	"strings"
)

func main() {
	usr, err := user.Current(); loged(err)
	home := usr.HomeDir
	dir, err := os.Getwd(); loged(err)
	dir = strings.Replace(dir, home, "~", 1)
	pwd := magenta(dir)
	fmt.Println(pwd)
}

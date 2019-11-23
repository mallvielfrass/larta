package prompt

import (
	"os/exec"
	"strings"
)

func GetPython() string {
	out, err := exec.Command("python", "-V").Output(); loged(err)
	return strings.TrimSpace(string(out))
}

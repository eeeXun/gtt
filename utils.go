package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func IndexOf(candidate string, arr []string) int {
	for index, element := range arr {
		if element == candidate {
			return index
		}
	}
	return -1
}

func SetTermTitle(title string) {
	print("\033]0;", title, "\007")
}

func CopyToClipboard(text string) {
	switch runtime.GOOS {
	case "linux":
		exec.Command("sh", "-c",
			fmt.Sprintf("echo -n '%s' | xclip -selection clipboard", text)).
			Start()
	case "darwin":
		exec.Command("sh", "-c",
			fmt.Sprintf("echo -n '%s' | pbcopy", text)).
			Start()
	case "windows":
		exec.Command("cmd", "/c",
			fmt.Sprintf("echo %s | clip", text)).
			Start()
	}
}

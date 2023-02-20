package main

import (
	"fmt"
	"os"
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

func SetTermTitle(name string) {
	fmt.Printf("\033]0;gtt - %s\007", name)
}

func CopyToClipboard(text string) {
	switch runtime.GOOS {
	case "linux":
		switch os.Getenv("XDG_SESSION_TYPE") {
		case "x11":
			exec.Command("sh", "-c",
				fmt.Sprintf("echo -n '%s' | xclip -selection clipboard", text)).
				Start()
		case "wayland":
			exec.Command("sh", "-c",
				fmt.Sprintf("echo -n '%s' | wl-copy", text)).
				Start()
		}
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

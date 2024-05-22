package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
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
	if uiStyle.OSC52 {
		fmt.Printf("\033]52;c;%s\a", base64.StdEncoding.EncodeToString([]byte(text)))
		return
	}

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		switch os.Getenv("XDG_SESSION_TYPE") {
		case "x11":
			cmd = exec.Command("xclip", "-selection", "clipboard")
		case "wayland":
			cmd = exec.Command("wl-copy")
		}
	case "darwin":
		cmd = exec.Command("pbcopy")
	case "windows":
		cmd = exec.Command("clip")
	}

	cmd.Stdin = strings.NewReader(text)
	cmd.Start()
}

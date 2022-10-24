package main

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

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

type Lock struct {
	stop        bool
	threadCount int8
}

func NewLock() *Lock {
	return &Lock{
		stop:        true,
		threadCount: 0,
	}
}

func (l *Lock) Aquire() {
	l.stop = false
	l.threadCount++
}

func (l *Lock) Release() {
	l.stop = true
	l.threadCount--
}

package translate

type Lock struct {
	Stop        bool
	threadCount int8
}

func NewLock() *Lock {
	return &Lock{
		Stop:        true,
		threadCount: 0,
	}
}

func (l *Lock) Available() bool {
	return l.Stop && l.threadCount == 0
}

func (l *Lock) Acquire() {
	l.Stop = false
	l.threadCount++
}

func (l *Lock) Release() {
	l.Stop = true
	l.threadCount--
}

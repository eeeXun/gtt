package core

type TTSLock struct {
	stop        bool
	threadCount int8
}

func NewTTSLock() *TTSLock {
	return &TTSLock{
		stop:        true,
		threadCount: 0,
	}
}

func (l *TTSLock) LockAvailable() bool {
	return l.stop && l.threadCount == 0
}

func (l *TTSLock) AcquireLock() {
	l.stop = false
	l.threadCount++
}

func (l *TTSLock) IsStopped() bool {
	return l.stop
}

func (l *TTSLock) StopTTS() {
	l.stop = true
}

func (l *TTSLock) ReleaseLock() {
	l.stop = true
	l.threadCount--
}

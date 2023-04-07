package core

type TTSLock struct {
	using bool
	stop  bool
}

func NewTTSLock() *TTSLock {
	return &TTSLock{
		stop:  true,
		using: false,
	}
}

func (l *TTSLock) LockAvailable() bool {
	return l.stop && !l.using
}

func (l *TTSLock) AcquireLock() {
	l.stop = false
	l.using = true
}

func (l *TTSLock) IsStopped() bool {
	return l.stop
}

func (l *TTSLock) StopTTS() {
	l.stop = true
}

func (l *TTSLock) ReleaseLock() {
	l.stop = true
	l.using = false
}

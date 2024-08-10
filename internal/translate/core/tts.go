package core

import (
	"io"
	"time"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
)

type TTS struct {
	stop  bool
	using bool
}

func NewTTS() *TTS {
	return &TTS{
		stop:  true,
		using: false,
	}
}

func (s *TTS) LockAvailable() bool {
	return s.stop && !s.using
}

func (s *TTS) AcquireLock() {
	s.stop = false
	s.using = true
}

func (s *TTS) IsStopped() bool {
	return s.stop
}

func (s *TTS) StopTTS() {
	s.stop = true
}

func (s *TTS) ReleaseLock() {
	s.stop = true
	s.using = false
}

func (s *TTS) Play(body io.Reader) error {
	decoder, err := mp3.NewDecoder(body)
	if err != nil {
		return err
	}
	otoCtx, readyChan, err := oto.NewContext(decoder.SampleRate(), 2, 2)
	if err != nil {
		return err
	}
	<-readyChan
	player := otoCtx.NewPlayer(decoder)
	player.Play()
	for player.IsPlaying() {
		if s.IsStopped() {
			return player.Close()
		}
		time.Sleep(time.Millisecond)
	}
	return player.Close()
}

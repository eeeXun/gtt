package core

import (
	"io"
	"time"

	"github.com/gen2brain/malgo"
	"github.com/hajimehoshi/go-mp3"
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
	ctx, err := malgo.InitContext(nil, malgo.ContextConfig{}, nil)
	if err != nil {
		return err
	}
	defer func() {
		ctx.Uninit()
		ctx.Free()
	}()
	deviceConfig := malgo.DefaultDeviceConfig(malgo.Playback)
	deviceConfig.Playback.Format = malgo.FormatS16
	deviceConfig.Playback.Channels = 2
	deviceConfig.SampleRate = uint32(decoder.SampleRate())
	isPlaying := true
	deviceCallbacks := malgo.DeviceCallbacks{
		Data: func(outputSamples, inputSamples []byte, frameCount uint32) {
			read, _ := io.ReadFull(decoder, outputSamples)
			if read <= 0 {
				isPlaying = false
				return
			}
		}}
	device, err := malgo.InitDevice(ctx.Context, deviceConfig, deviceCallbacks)
	if err != nil {
		return err
	}
	defer device.Uninit()
	err = device.Start()
	if err != nil {
		return err
	}
	for isPlaying {
		if s.IsStopped() {
			return nil
		}
		time.Sleep(time.Millisecond)
	}
	return nil
}

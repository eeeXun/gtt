// echo 'nice' | base64
// curl -A "Mozilla/4.0" \
// 	"https://voice.reverso.net/RestPronunciation.svc/v1/output=json/GetVoiceStream/voiceName=Heather22k?voiceSpeed=80&inputText=bmljZQ==" \
// 	--output a.mp
package reversotranslate

const (
	ttsURL = "https://translate.google.com.vn/translate_tts?ie=UTF-8&q=%s&tl=%s&client=tw-ob"
)

func (t *ReversoTranslate) LockAvailable() bool {
	return t.SoundLock.Available()
}

func (t *ReversoTranslate) LockAcquire() {
	t.SoundLock.Acquire()
}

func (t *ReversoTranslate) StopTTS() {
	t.SoundLock.Stop = true
}

func (t *ReversoTranslate) PlayTTS(lang, message string) error {
	// urlStr := fmt.Sprintf(
	// 	ttsURL,
	// 	url.QueryEscape(message),
	// 	langCode[lang],
	// )
	// res, err := http.Get(urlStr)
	// if err != nil {
	// 	t.SoundLock.Release()
	// 	return err
	// }
	// decoder, err := mp3.NewDecoder(res.Body)
	// if err != nil {
	// 	t.SoundLock.Release()
	// 	return err
	// }
	// otoCtx, readyChan, err := oto.NewContext(decoder.SampleRate(), 2, 2)
	// if err != nil {
	// 	t.SoundLock.Release()
	// 	return err
	// }
	// <-readyChan
	// player := otoCtx.NewPlayer(decoder)
	// player.Play()
	// for player.IsPlaying() {
	// 	if t.SoundLock.Stop {
	// 		t.SoundLock.Release()
	// 		return nil
	// 	} else {
	// 		time.Sleep(time.Millisecond)
	// 	}
	// }
	// if err = player.Close(); err != nil {
	// 	t.SoundLock.Release()
	// 	return err
	// }
	//
	t.SoundLock.Release()
	return nil
}

package core

type Language struct {
	srcLang string
	dstLang string
}

func NewLanguage() *Language {
	return &Language{}
}

func (l *Language) GetSrcLang() string {
	return l.srcLang
}

func (l *Language) GetDstLang() string {
	return l.dstLang
}

func (l *Language) SetSrcLang(lang string) {
	l.srcLang = lang
}

func (l *Language) SetDstLang(lang string) {
	l.dstLang = lang
}

func (l *Language) SwapLang() {
	l.srcLang, l.dstLang = l.dstLang, l.srcLang
}

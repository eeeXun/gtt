package core

const (
	UserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36"
)

type Translation struct {
	// translation text
	TEXT string

	// translation definition or example
	DEF string

	// translation part of speech
	POS string
}

package lingvatranslate

type Decoder struct {
	sampleRate    int
	length        int64
	pos           int64
	buf           []byte
	bytesPerFrame int64
}

func NewDecoder(buf []byte)*Decoder {
	return &Decoder{
	}
}
func (d *Decoder) Read(buf []byte) (int, error) {
	n := copy(buf, d.buf)
	d.buf = d.buf[n:]
	d.pos += int64(n)
	return n, nil
}

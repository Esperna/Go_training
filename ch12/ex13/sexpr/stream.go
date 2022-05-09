package sexpr

import (
	"fmt"
	"io"
)

type Decoder struct {
	r io.Reader
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: r}
}

func (dec *Decoder) Decode(v any) error {
	b, err := io.ReadAll(dec.r)
	if err != nil {
		return fmt.Errorf("read failed:%s", err)
	}
	Unmarshal(b, v)
	return nil
}

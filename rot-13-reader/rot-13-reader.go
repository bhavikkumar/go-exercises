package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(p []byte) (int, error) {
	s, err := rot.r.Read(p)

	for i := range p {
		if (p[i] >= 'A' && p[i] < 'N') || (p[i] >='a' && p[i] < 'n') {
		  p[i] += 13
		} else if (p[i] > 'M' && p[i] <= 'Z') || (p[i] > 'm' && p[i] <= 'z'){
		  p[i] -= 13
		}
	}
	return s, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

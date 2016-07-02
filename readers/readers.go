package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (myReader MyReader) Read(bytesToRead []byte) (int, error) {
	for i := range bytesToRead {
		bytesToRead[i] = byte('A')
	}
	return len(bytesToRead), nil	
}

func main() {
	reader.Validate(MyReader{})
}

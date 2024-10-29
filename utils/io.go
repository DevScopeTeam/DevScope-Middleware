package utils

import (
	"bytes"
	"encoding/base64"
	"io"
)

func FileToBase64(file io.Reader) (string, error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(file)
	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

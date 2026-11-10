package byte

import (
	"bytes"
	"encoding/hex"
	"encoding/gob"
)

func Hash(a interface{}) []byte {
	value, _ := HashWithError(a)
	return value
}

func HashWithError(a interface{}) ([]byte, error) {
	var b bytes.Buffer
	if err := gob.NewEncoder(&b).Encode(a); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func HashString(a interface{}) (string, error) {
	value, err := HashWithError(a)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(value), nil
}

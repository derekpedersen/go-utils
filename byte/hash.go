package byte

import (
	"bytes"
	"encoding/gob"
)

func Hash(a interface{}) []byte {
	var b bytes.Buffer
	gob.NewEncoder(&b).Encode(a)
	return b.Bytes()
}

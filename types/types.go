package types

import (
	"bytes"
	"encoding/gob"
)

type EntityType int

type Entity struct {
	ID   string
	Type EntityType
}

func GobMarshal(v interface{}) ([]byte, error) {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	err := enc.Encode(v)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func GobUnmarshal(b []byte, v interface{}) error {
	r := bytes.NewReader(b)
	dec := gob.NewDecoder(r)
	return dec.Decode(v)
}

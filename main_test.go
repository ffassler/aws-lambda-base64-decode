package main

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	source := "In go we trust"
	s := base64.StdEncoding.EncodeToString([]byte(source))
	dest, err := Decode64(s)

	assert.IsType(t, nil, err)
	assert.Equal(t, source, dest)
}

func TestDecodeErr(t *testing.T) {
	_, err := Decode64(string([]byte{0, 1, 2, 3}))

	assert.NotNil(t, err)
}

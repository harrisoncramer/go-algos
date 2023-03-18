package decodestring

import (
	b64 "encoding/base64"
	"strings"
)

func EncodeString(s string) (string, error) {

	shifted := strings.Map(func(r rune) rune {
		return r + 1
	}, string(s))

	encoded := b64.StdEncoding.EncodeToString([]byte(shifted))

	return encoded, nil
}

func DecodeString(s string) (string, error) {
	res, err := b64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}

	result := strings.Map(func(r rune) rune {
		return r - 1
	}, string(res))

	return result, nil
}

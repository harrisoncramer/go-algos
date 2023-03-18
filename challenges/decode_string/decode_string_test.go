package decodestring

import "testing"

/*
Write a function that decodes a string that's been translated into base64 encoding
and shifted one character up
*/

func TestDecodeString(t *testing.T) {
	t.Run("Should decode the string", func(t *testing.T) {
		input := "VEZEU0ZVVFVTSk9I"
		got, err := DecodeString(input)

		if err != nil {
			t.Errorf("Got error during decoding: %e", err)
		}

		want := "SECRETSTRING"

		if got != want {
			t.Errorf("Got %s but wanted %s", got, want)
		}
	})

	t.Run("Should encode and decode a string", func(t *testing.T) {
		input := "Hello world"
		encoded, err := EncodeString(input)

		if err != nil {
			t.Errorf("Got error during encoding: %s", err)
		}

		got, err := DecodeString(encoded)

		if err != nil {
			t.Errorf("Got error during decoding: %s", err)
		}

		if got != input {
			t.Errorf("Got %s but wanted %s", got, input)
		}
	})

	t.Run("Should handle empty string", func(t *testing.T) {
		input := ""
		want := ""

		got, err := DecodeString(input)
		if err != nil {
			t.Errorf("Got error during decoding: %s", err)
		}

		if got != want {
			t.Errorf("Got %s but wanted %s", got, want)
		}
	})

	t.Run("Should handle non-ASCII characters", func(t *testing.T) {
		input := "Hello 🦾 world 👄"
		encoded, err := EncodeString(input)

		if err != nil {
			t.Errorf("Got error during encoding: %s", err)
		}

		got, err := DecodeString(encoded)

		if err != nil {
			t.Errorf("Got error during decoding: %s", err)
		}

		if got != input {
			t.Errorf("Got %s but wanted %s", got, input)
		}
	})
}

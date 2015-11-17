package jsonbox

import (
	"encoding/json"
	"errors"

	"golang.org/x/crypto/nacl/secretbox"
)

// length for nonce
const length = 24

// Key is private key for Salsa20
var Key [32]byte

// Secret is custom string type that implements Marshal- and UnmarshalJSON
type Secret string

// MarshalJSON implements json.Marshaler interface.
func (s Secret) MarshalJSON() ([]byte, error) {
	nonce := [length]byte{}
	out := []byte{}
	out = secretbox.Seal(out, []byte(s), &nonce, &Key)
	res := append(nonce[:], out...)
	return json.Marshal(res)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (s *Secret) UnmarshalJSON(data []byte) error {
	var box []byte
	if err := json.Unmarshal(data, &box); err != nil {
		return err
	}
	nonce := [length]byte{}
	copy(nonce[:], box[:length])
	out := []byte{}
	ok := false
	out, ok = secretbox.Open(out, box[length:], &nonce, &Key)
	if !ok {
		return errors.New("Failed to open box")
	}
	*s = Secret(string(out))
	return nil
}

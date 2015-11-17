package jsonbox

import (
	"encoding/hex"
	"encoding/json"
	"testing"
)

// Person uses Secret type for Code field
type Person struct {
	Name string `json:"name"`
	Code Secret `json:"code"`
}

func init() {
	// use random key of length 32
	// rand.Reader.Read(Key[:])
	// log.Printf("%x", Key)
	s := "108743504038048198717d7ff329464397ec2e199bca88e22423717f86ecaae7"
	k, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	copy(Key[:], k[:])
}

func Test(t *testing.T) {
	// create new person instance
	john := Person{
		Name: "john",
		Code: "open sesame",
	}
	res, err := json.Marshal(john)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(res))
	// unmarshal json into new struct
	steve := Person{}
	err = json.Unmarshal(res, &steve)
	if err != nil {
		t.Error(err)
	}
	if steve.Code != "open sesame" {
		t.Error("invalid value")
	}
	t.Logf("%+v", steve)
}

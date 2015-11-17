
# jsonbox

[![Build Status](https://travis-ci.org/zemirco/jsonbox.svg)](https://travis-ci.org/zemirco/jsonbox)
[![Coverage Status](https://coveralls.io/repos/zemirco/jsonbox/badge.svg?branch=master&service=github)](https://coveralls.io/github/zemirco/jsonbox?branch=master)
[![GoDoc](https://godoc.org/github.com/zemirco/jsonbox?status.svg)](https://godoc.org/github.com/zemirco/jsonbox)

Encrypt and decrypt JSON with [secretbox](https://godoc.org/golang.org/x/crypto/nacl/secretbox).

## Example

```go
package main

import (
  "encoding/json"

  "github.com/zemirco/jsonbox"
)

type Person struct {
	Name string `json:"name"`
  // use jsonbox.Secret as type (string is underlying type)
	Code jsonbox.Secret `json:"code"`
}

func main() {
  // use your own key of length 32 in production
  rand.Reader.Read(jsonbox.Key[:])
  // create new instance
  john := Person{
    Name: "john",
    Code: "open sesame",
  }
  res, err := json.Marshal(john)
  if err != nil {
    panic(err)
  }
  // make sure code field is encrypted
  log.Println(string(res))

  // unmarshal json into new struct
  steve := Person{}
  err = json.Unmarshal(res, &steve)
  if err != nil {
    panic(err)
  }
  log.Printf("%+v", steve)
}
```

## Test

```go
go test
```

## License

MIT

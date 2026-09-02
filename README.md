# validator-go

[Katydid](https://katydid.org.za) is a validation language. `validator-go` is a validator for Katydid in Go.

[![GoDoc](https://godoc.org/katydid.org.za/go/validator-go?status.svg)](https://godoc.org/katydid.org.za/go/validator-go)
[![Build Status](https://git.katydid.org.za/validator-go/actions/workflows/build.yml/badge.svg)](https://git.katydid.org.za/validator-go/actions)

![Katydid Logo](https://katydid.org.za/logo.png)

Katydid consists of:

  * A validator: a regular expression type language for serialized data that matches up to 1000000s of records per second,
  * A collection of parsers (protobuf, json, xml, reflected go structures, yaml) which are easily extendable.

## Usage Example

```go
package main

import (
  "katydid.org.za/go/validator-go/validator"
  "katydid.org.za/go/parser-go-json/json"
)

func main() {
  data := json.NewParser()
	data.Init([]byte(`
    {
      "WhatsUp": "E",
      "DragonsExist": false,
      "MonkeysSmart": true
    }`))
  ...
  ast, err := validator.Parse(`.WhatsUp == "E"`)
  ...
  // creates memoizing validator that increases in speed the more it is used.
  mem, err := validator.Prepare(ast)
  ...
  valid, err := validator.Validate(mem, data)
  ...
  if valid {
    fmt.Printf("the serailized json contains a field WhatsUp with a value E\n")
  }
}
```
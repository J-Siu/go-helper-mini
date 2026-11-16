# basestruct

Provides a simple struct with 5 common fields to be embedded by other structs.

## Installation

```sh
go get github.com/J-Siu/go-helper-mini
```

## Usage

```go
import "github.com/J-Siu/go-helper-mini/basestruct"
```

## TYPES

### Embedded

A simple struct to be embedded by other struct
```go
type Embedded struct {
  Debug       bool   `json:"Debug"` // simple log level control
  Err         error  `json:"Err,omitempty"`
  Initialized bool   `json:"Initialized,omitempty"`
  LogLevel    uint8  `json:"LogLevel,omitempty"` // finer log control
  MyType      string `json:"MyType"`             // Store typename. Cheaper way than reflector for logging.
}
```

### IsInitAndNoErr

```go
func (b *Embedded) IsInitAndNoErr(prefix string) (pass bool)
```

To be put at the beginning of Check error and initialization state in following order:

1. If `Err` not nil -> return `false`
2. If `Initialized` is `false` -> set `Err` -> return `false`
3. Else -> return `true`

### License

The MIT License (MIT)

Copyright © 2026 John, Sing Dao, Siu <john.sd.siu@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

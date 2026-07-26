/*
Copyright © 2026 John, Sing Dao, Siu <john.sd.siu@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

// # basestruct
//
// Provides a simple struct with common fields to be embedded by other structs.
package basestruct

import (
	"errors"
)

// A simple struct to be embedded by other struct
type Embedded struct {
	Debug       bool   `json:"Debug"` // simple log level control
	Err         error  `json:"Err,omitempty"`
	Initialized bool   `json:"Initialized,omitempty"`
	LogLevel    uint8  `json:"LogLevel,omitempty"` // finer log control
	MyType      string `json:"MyType"`             // Store typename. Cheaper way than reflector for logging.
}

// Check error and initialization state in following order:
//  1. If `Err` not nil -> return `false`
//  2. If `Initialized` is `false` -> set `Err` -> return `false`
//  3. Else, check passed -> return `true`
func (t *Embedded) IsInitAndNoErr(prefix string) (pass bool) {
	pass = true
	// check error first
	if t.Err != nil {
		pass = false
	} else if !t.Initialized {
		errMsg := "not initialized"
		if prefix != "" {
			errMsg = prefix + ": " + errMsg
		} else if t.MyType != "" {
			errMsg = t.MyType + ": " + errMsg
		}
		t.Err = errors.New(errMsg)
		pass = false
	}
	return pass
}

// for interface
func (t *Embedded) GetErr() error { return t.Err }

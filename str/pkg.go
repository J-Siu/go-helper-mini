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

// Non-struct, function only
package str

import "encoding/json"

func Struct(prefix string, in any, out *string) (err error) {
	var b []byte
	if prefix != "" {
		*out = prefix + ":" + NewLine
	}
	if b, err = json.MarshalIndent(in, "", "  "); err == nil {
		*out += string(b)
	}
	return
}

// if json(sonic).MarshalIndent err, return empty
func Struct2String(prefix string, in any) (s string) {
	Struct(prefix, in, &s)
	return
}

// Return "OK"/"Fail"
func Ok(b bool) string {
	if b {
		return "OK"
	}
	return "Fail"
}

// Return "Success"/"Fail"
func Success(b bool) string {
	if b {
		return "Success"
	}
	return "Fail"
}

// Return "Yes"/"No"
func YesNo(b bool) string {
	if b {
		return "Yes"
	}
	return "No"
}

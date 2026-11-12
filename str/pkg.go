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

import (
	"encoding/json"
	"strings"
)

// Check if string array contain a string.
//
// Return false if arrP is nil
func ArrayContains(arrP *[]string, str string, caseSensitive bool) bool {
	if arrP != nil {
		for _, s := range *arrP {
			if caseSensitive && s == str || strings.EqualFold(s, str) {
				return true
			}
		}
	}
	return false
}

// Check if string array contains a substring
//
// Return false if arrP is nil
func ArrayContainsSubString(arrP *[]string, subStr string, caseSensitive bool) bool {
	if arrP != nil {
		for _, s := range *arrP {
			if caseSensitive && strings.Contains(s, subStr) || strings.Contains(strings.ToLower(s), strings.ToLower(subStr)) {
				return true
			}
		}
	}
	return false
}

// Check if string contains any substring of an array
//   - result (bool)
//   - if result is true, `resultVal` == matching substring
//   - if result is false, `resultVal` == ""
func ContainsAnySubStrings(str string, subStrArrP *[]string, caseSensitive bool) (result bool, resultVal string) {
	// prefix := "matchList"
	if subStrArrP != nil {
		for _, subStr := range *subStrArrP {
			if caseSensitive && strings.Contains(str, subStr) || strings.Contains(strings.ToLower(str), strings.ToLower(subStr)) {
				result = true
				resultVal = subStr
				break
			}
		}
	}
	return result, resultVal
}

// Return only bool from [ContainsAnySubStrings]
func ContainsAnySubStringsBool(str string, subStrArrP *[]string, caseSensitive bool) (result bool) {
	result, _ = ContainsAnySubStrings(str, subStrArrP, caseSensitive)
	return result
}

func Struct(prefix string, in any, indent bool) (out string, err error) {
	var b []byte
	if indent {
		b, err = json.MarshalIndent(in, "", "  ")
	} else {
		b, err = json.Marshal(in)
	}
	if err == nil {
		if prefix != "" {
			out = prefix + ":" + NewLine + string(b)
		} else {
			out = string(b)
		}
	}
	return
}

// if json(sonic).MarshalIndent err, return empty
func Struct2String(prefix string, in any, indent bool) (s string) {
	s, _ = Struct(prefix, in, indent)
	return
}

// cspell:words charlievieth

// Return "OK"/"Fail"
func Ok(b bool) string {
	if b {
		return "OK"
	}
	return "Fail"
}

// Return "Success"/"Fail"
func SuccessFail(b bool) string {
	if b {
		return "Success"
	}
	return "Fail"
}

// Return "Pass"/"Fail"
func PassFail(b bool) string {
	if b {
		return "PAss"
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

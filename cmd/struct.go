/*
The MIT License

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

package cmd

import (
	"bytes"
	"fmt"
	"log/slog"
	"os/exec"
	"sync"

	"fj.vms.local/jsds/go-helper-mini/basestruct"
	"fj.vms.local/jsds/go-helper-mini/file"
	"fj.vms.local/jsds/go-helper-mini/str"
)

type Cmd struct {
	basestruct.Base
	Args     []string      `json:"ArgsP,omitempty"`    // In : Command args
	CmdLn    string        `json:"CmdLn,omitempty"`    // Out: Command line
	CmdName  string        `json:"CmdName,omitempty"`  // In : Command name
	Dir      string        `json:"Dir,omitempty"`      // In : Command working dir
	ExitCode int           `json:"ExitCode,omitempty"` // Out: Exit Code
	Ran      bool          `json:"Ran,omitempty"`      // Out: Set to true by Run()
	Stderr   *bytes.Buffer `json:"Stderr,omitempty"`   // Out: Stderr
	Stdout   *bytes.Buffer `json:"Stdout,omitempty"`   // Out: Stdout
}

// Setup and return cmd pointer.
//   - If <workPathP> is empty/nil, current directory is used.
func (t *Cmd) New(cmdName string, argsP *[]string, workPathP *string) *Cmd {
	t.MyType = "Cmd"
	t.Args = *argsP
	t.CmdName = cmdName
	t.CmdLn = ""
	t.Err = nil
	t.ExitCode = 0
	t.Ran = false
	t.Dir = file.FullPath(*workPathP)
	t.Stderr = new(bytes.Buffer)
	t.Stdout = new(bytes.Buffer)
	return t
}

// A exec.Cmd.Run() wrapper.
func (t *Cmd) Run() *Cmd {
	execCmd := exec.Command(t.CmdName, t.Args...)
	execCmd.Dir = t.Dir
	execCmd.Stderr = t.Stderr
	execCmd.Stdout = t.Stdout
	t.CmdLn = execCmd.String()
	t.Err = execCmd.Run()
	t.Ran = true

	if exitErr, ok := t.Err.(*exec.ExitError); ok {
		t.ExitCode = exitErr.ExitCode()
	}

	prefix := "cmd.Run"
	if t.Debug {
		slog.Debug(prefix, "struct", str.Struct2String("", t, false))
	}
	return t
}

// run func wrapper with sync.WaitGroup support.
//   - If <workPathP> is empty/nil, current directory is used.
//   - Print if <output> is true
func (t *Cmd) RunWg(title *string, wgP *sync.WaitGroup, output bool) *Cmd {
	if wgP != nil {
		defer wgP.Done()
	}
	t.Run()
	if output {
		t.Log(*title)
	}
	return t
}

func (t *Cmd) Error() error { return t.Err }

func (t *Cmd) Log(title string) *Cmd {
	fmt.Printf("%s: %s: %s%s", title, "Stdout", t.Stdout.String(), str.NewLine)
	fmt.Printf("%s: %s: %s%s", title, "Stderr", t.Stderr.String(), str.NewLine)
	return t
}

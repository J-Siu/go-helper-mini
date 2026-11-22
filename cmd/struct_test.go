package cmd

import (
	"strings"
	"testing"

	"github.com/J-Siu/go-helper-mini/str"
)

func TestCmdStructSerialization(t *testing.T) {
	cmd := (&Cmd{}).New("echo", []string{"hello"}, "")
	cmd.Run()

	out, err := str.Struct("Cmd", cmd, true)
	if err != nil {
		t.Fatalf("serialize Cmd: %v", err)
	}
	if !strings.Contains(out, `"CmdLn"`) {
		t.Fatalf("serialized Cmd does not contain command line: %s", out)
	}
	if !strings.Contains(out, `"Stdout": "hello`) {
		t.Fatalf("serialized Cmd does not contain stdout: %s", out)
	}
}

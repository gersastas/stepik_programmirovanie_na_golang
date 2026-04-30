package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestMain_LikePrintTree(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"\n", "I like Go!\nI like Go!\nI like Go!\n"},
	}
	for _, tt := range tests {
		oldStdin := os.Stdin
		r, w, _ := os.Pipe()
		os.Stdin = r
		w.WriteString(tt.input)
		w.Close()

		oldStdout := os.Stdout
		ro, wo, _ := os.Pipe()
		os.Stdout = wo

		main()

		wo.Close()
		os.Stdout = oldStdout
		os.Stdin = oldStdin

		var buf bytes.Buffer
		io.Copy(&buf, ro)

		if buf.String() != tt.want {
			t.Errorf("input %q: got %q, want %q", tt.input, buf.String(), tt.want)
		}
	}
}

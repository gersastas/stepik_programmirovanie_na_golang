package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestMain_NumSquareInt(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"2\n", "4\n"},
		{"3\n", "9\n"},
		{"5\n", "25\n"},
		{"4\n", "16\n"},
		{"7\n", "49\n"},
	}

	for _, tt := range tests {
		// Подменяем stdin
		oldStdin := os.Stdin
		r, w, _ := os.Pipe()
		os.Stdin = r
		w.WriteString(tt.input)
		w.Close()

		// Перехватываем stdout
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

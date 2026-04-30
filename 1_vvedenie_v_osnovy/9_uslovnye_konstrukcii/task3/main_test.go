package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestMain_FirstDigitCheck(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"1234\n", "1\n"},
		{"756\n", "7\n"},
		{"48\n", "4\n"},
		{"8\n", "8\n"},
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

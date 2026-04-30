package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestMain_ElemEvenIndexOfArray(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"5\n1 2 3 -1 -4", "3\n"},
		{"5\n-1 -2 -3 -4 -5", "0\n"},
		{"5\n1 2 3 4 5", "5\n"},
		{"1\n0", "0\n"},
		{"1\n1", "1\n"},
		{"4\n0 0 0 0", "0\n"},
		{"6\n-3 0 5 -1 0 7", "2\n"},
		{"3\n-100 0 100", "1\n"},
		{"10\n1 -1 2 -2 3 -3 4 -4 5 -5", "5\n"},
		{"4\n0 1 2 3", "3\n"},
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

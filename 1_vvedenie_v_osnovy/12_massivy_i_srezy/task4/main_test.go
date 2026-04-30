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
		{"6\n4 5 3 4 2 3", "4 3 2 "},
		{"1\n7", "7 "},
		{"2\n10 20", "10 "},
		{"5\n1 2 3 4 5", "1 3 5 "},
		{"4\n9 8 7 6", "9 7 "},
		{"7\n10 20 30 40 50 60 70", "10 30 50 70 "},
		{"3\n-1 -2 -3", "-1 -3 "},
		{"10\n1 2 3 4 5 6 7 8 9 10", "1 3 5 7 9 "},
		{"6\n0 0 0 0 0 0", "0 0 0 "},
		{"5\n-5 10 -3 8 -1", "-5 -3 -1 "},
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

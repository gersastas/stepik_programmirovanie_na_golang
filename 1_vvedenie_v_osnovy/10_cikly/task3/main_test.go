package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestMain_SumNumXX(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"5\n38 24 800 8 16\n", "40\n"},
		{"7\n6 16 24 32 48 96 100\n", "216\n"},
		{"4\n3 8 9 10\n", "0\n"},
		{"5\n4 80 88 96 104\n", "264\n"},
		{"2\n1 56\n", "56\n"},
		{"4\n3 0 8 800\n", "0\n"},
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

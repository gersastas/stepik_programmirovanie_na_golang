package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestMain_CalcArithmeticMean(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"3 5\n", "4\n"},
		{"277 154\n", "215.5\n"},
		{"1 1\n", "1\n"},
		{"1 2\n", "1.5\n"},
		{"10 20\n", "15\n"},
		{"7 8\n", "7.5\n"},
		{"100 200\n", "150\n"},
		{"1 99\n", "50\n"},
		{"3 4\n", "3.5\n"},
		{"1000 1000\n", "1000\n"},
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

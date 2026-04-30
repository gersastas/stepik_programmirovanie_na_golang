package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestMain_CheckTringle(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"4 5 6\n", "Существует\n"},
		{"3 4 5\n", "Существует\n"},
		{"1 1 1\n", "Существует\n"},
		{"1 2 3\n", "Не существует\n"},
		{"1 1 2\n", "Не существует\n"},
		{"1 10 1\n", "Не существует\n"},
		{"5 5 5\n", "Существует\n"},
		{"7 3 5\n", "Существует\n"},
		{"100 1 1\n", "Не существует\n"},
		{"10 10 1\n", "Существует\n"},
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

package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestMain_MaxNumOfArray(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"2 3 56 45 21\n", "56\n"},
		{"-5 -3 -10 -1 -8", "-1\n"},
		{"500 400 300 200 100", "500\n"},
		{"0 0 0 0 1\n", "1\n"},
		{"-1 0 1 2 3\n", "3\n"},
		{"42 42 42 42 43\n", "43\n"},
		{"7 3 9 1 5\n", "9\n"},
		{"0 0 0 0 0\n", "0\n"},
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

package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestMain_CalcHourMinutes(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"13257\n", "It is 3 hours 40 minutes."},
		{"3600\n", "It is 1 hours 0 minutes."},
		{"7200\n", "It is 2 hours 0 minutes."},
		{"60\n", "It is 0 hours 1 minutes."},
		{"59\n", "It is 0 hours 0 minutes."},
		{"43200\n", "It is 12 hours 0 minutes."},
		{"1\n", "It is 0 hours 0 minutes."},
		{"36000\n", "It is 10 hours 0 minutes."},
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

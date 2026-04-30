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
		{"5\n41 -231 24 49 6", "49\n"},
		{"7\n6 16 24 32 48 96 100\n", "32\n"},
		{"5\n4 80 88 96 104\n", "96\n"},
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

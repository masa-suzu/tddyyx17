package fizzbuzz

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestItoa(t *testing.T) {
	tests := []struct {
		in   int
		want string
	}{
		// 数字から文字列に変える
		{1, "1"},
		{2, "2"},
		{4, "4"},
		// 3の倍数はFizz
		{3, "Fizz"},
		{6, "Fizz"},
		// 5の倍数ならBuzz
		{5, "Buzz"},
		{10, "Buzz"},
		// 15の倍数ならFizzBuzz
		{15, "FizzBuzz"},
		{45, "FizzBuzz"},
	}

	for _, tt := range tests {
		got := Itoa(tt.in)

		if got != tt.want {
			t.Errorf("answer(%v) == %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestPrint(t *testing.T) {

	out := captureStdout(func() {
		print(1)
	})
	if out != "1" {
		t.Errorf("out == %v, want '1'", out)
	}
	out = captureStdout(func() {
		print(2)
	})
	if out != "2" {
		t.Errorf("out == %v, want '2'", out)
	}
}

func TestPrintby(t *testing.T) {

	out := captureStdout(func() {
		PrintBy(20)
	})
	want := `1
2
Fizz
4
Buzz
Fizz
7
8
Fizz
Buzz
11
Fizz
13
14
FizzBuzz
16
17
Fizz
19
Buzz
`
	if out != want {
		t.Errorf("out == %#v, want %#v", out, want)
	}
}

func captureStdout(f func()) string {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	stdout := os.Stdout
	os.Stdout = w

	f()

	os.Stdout = stdout
	w.Close()

	var buf bytes.Buffer
	io.Copy(&buf, r)

	return buf.String()
}

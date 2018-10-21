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

func TestPrintBy(t *testing.T) {

	out := captureStdout(func() {
		PrintBy(100)
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
Fizz
22
23
Fizz
Buzz
26
Fizz
28
29
FizzBuzz
31
32
Fizz
34
Buzz
Fizz
37
38
Fizz
Buzz
41
Fizz
43
44
FizzBuzz
46
47
Fizz
49
Buzz
Fizz
52
53
Fizz
Buzz
56
Fizz
58
59
FizzBuzz
61
62
Fizz
64
Buzz
Fizz
67
68
Fizz
Buzz
71
Fizz
73
74
FizzBuzz
76
77
Fizz
79
Buzz
Fizz
82
83
Fizz
Buzz
86
Fizz
88
89
FizzBuzz
91
92
Fizz
94
Buzz
Fizz
97
98
Fizz
Buzz
`

	if out != want {
		t.Errorf("\ngot  %#v, \nwant %#v", out, want)
	}
}

func captureStdout(f func()) string {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	stdout := os.Stdout
	os.Stdout = w

	restoreStdout := func() {
		os.Stdout = stdout
	}

	defer restoreStdout()

	f()

	os.Stdout = stdout
	w.Close()

	var buf bytes.Buffer
	io.Copy(&buf, r)

	return buf.String()
}

package wordwrapper

import "testing"

func Test_Wrap_no(t *testing.T) {
	length := 5
	text := "Hello"
	expected := "Hello"
	got := Wrap(text, length)

	if got != expected {
		t.Errorf("\nGot:\n%s\nExpected:\n%s\n", got, expected)
	}
}

func Test_Wrap_Ex1(t *testing.T) {
	length := 10
	text := "This is a very long sentence that will be wrapped"
	expected := "This is a\nvery long\nsentence\nthat will\nbe wrapped"
	got := Wrap(text, length)

	if got != expected {
		t.Errorf("\nGot:\n%s\nExpected:\n%s\n", got, expected)
	}
}

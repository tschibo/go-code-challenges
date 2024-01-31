package conway

import "testing"

func Test_Blinker1(t *testing.T) {
	input := ".....\n" +
		".....\n" +
		".***.\n" +
		".....\n" +
		".....\n"

	expected := ".....\n" +
		"..*..\n" +
		"..*..\n" +
		"..*..\n" +
		".....\n"

	got := ConwayConv(input)

	if got != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, got)
	}
}

func Test_Blinker2(t *testing.T) {
	input := ".....\n" +
		"..*..\n" +
		"..*..\n" +
		"..*..\n" +
		".....\n"

	expected := ".....\n" +
		".....\n" +
		".***.\n" +
		".....\n" +
		".....\n"

	got := ConwayConv(input)

	if got != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, got)
	}
}

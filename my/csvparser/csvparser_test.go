package csvparser

import "testing"

func Test_Parse(t *testing.T) {
	winner := Parse("turing.csv")

	if winner[1966] != "Alan Perlis" {
		t.Errorf("Did not get 'Alan Perlis'")
	}

	if winner[1968] != "Richard Hamming" {
		t.Errorf("Did not get 'Richard Hamming'")
	}

	if winner[2001] != "Ole-Johan Dahl,Kristen Nygaard" {
		t.Errorf("Did not get 'Ole-Johan Dahl,Kristen Nygaard'")
	}
}

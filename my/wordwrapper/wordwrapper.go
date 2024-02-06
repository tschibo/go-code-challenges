package wordwrapper

import (
	"strings"
)

func Wrap(text string, length int) string {
	// Split text apart
	words := strings.Split(text, " ")

	// Reconstruct wrapped text
	pos := 0
	wrapped := ""
	for _, word := range words {
		if pos+len(word) <= length {
			// Line will be withing length
			wrapped += word + " "
			pos += len(word) + 1
		} else {
			// Line will exceed length: new line (restart pos from 0)
			// To get a new line, the tailing " " has to be replaced by "\n"
			wrapped = wrapped[:len(wrapped)-1]
			wrapped += "\n" + word + " "
			pos = len(word) + 1
		}
	}

	// remove tailing whitespace
	wrapped = wrapped[:len(wrapped)-1]

	return wrapped
}

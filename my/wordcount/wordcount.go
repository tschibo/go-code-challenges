package wordcount

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type FileContentInfo struct {
	Words                     int
	Characters                int
	CharactersExcludingSpaces int
}

func ReadFileContent(filename string) FileContentInfo {
	// Create empty result struct
	var fi FileContentInfo

	// Open file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// Scan through words
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		fi.Words++
		fi.CharactersExcludingSpaces += len(word)
		fi.Characters += len(word) + 1
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return fi
}

func ReadFileContent2(filename string) FileContentInfo {
	// Create empty result struct
	var fi FileContentInfo

	// Open file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// Scan through words
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		processLine(scanner.Text(), &fi)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return fi
}

func processLine(line string, fi *FileContentInfo) {
	fi.Characters += len(line)
	fi.CharactersExcludingSpaces += len(strings.Replace(line, " ", "", -1))
	fi.Words += len(strings.Fields(line))
}

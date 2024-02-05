package csvparser

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Parse(fileName string) map[int]string {
	result := make(map[int]string)

	file, err := os.Open("/workspaces/go-code-challenges/" + fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		segs := strings.Split(line, ",")
		year, err := strconv.Atoi(segs[0])
		if err != nil {
			panic(err)
		}
		if entry, ok := result[year]; ok {
			result[year] = entry + "," + segs[1]
		} else {
			result[year] = segs[1]
		}
	}

	return result
}

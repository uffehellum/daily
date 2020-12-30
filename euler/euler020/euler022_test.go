package euler020

import (
	"bufio"
	"io"
	"log"
	"os"
	"sort"
	"testing"
)

const pathToNameFile = "./p022_names.txt"

func TestScanner(t *testing.T) {
	names := readNames(pathToNameFile)
	AssertEqual(t, len(names), 5163)
	last := names[len(names)-1]
	if last != "ALONSO" {
		t.Fatal(last)
	}
}

func TestAlphaValue(t *testing.T) {
	AssertEqual(t, AlphaValue("COLIN"), 53)
}

func TestAlphaScoreSum(t *testing.T) {
	names := readNames(pathToNameFile)
	sort.Strings(names)
	AssertEqual(t, AlphaScoreSum(names), 871198282)
}

func AlphaScoreSum(names []string) int {
	r := 0
	for i, name := range names {
		r += (i + 1) * AlphaValue(name)
	}
	return r
}

func AlphaValue(name string) int {
	s := 0
	for _, c := range name {
		s += int(c-'A') + 1
	}
	return s
}

func readNames(filename string) []string {
	// https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// scanner := bufio.NewScanner(file)
	reader := bufio.NewReader(file)
	r := []string{}
	for {
		line, err := reader.ReadString(',')
		if err != nil && err != io.EOF {
			break
		}
		name := trimName(line)
		r = append(r, name)
		if err != nil {
			break
		}
	}
	return r
}

func trimName(s string) string {
	// https://stackoverflow.com/questions/44222554/how-to-remove-quotes-from-around-a-string-in-golang
	for len(s) > 0 && s[0] == '"' {
		s = s[1:]

	}
	for len(s) > 0 && (s[len(s)-1] == '"' || s[len(s)-1] == ',') {
		s = s[:len(s)-1]
	}
	return s
}

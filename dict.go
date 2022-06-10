package main

import (
	"bufio"
	"os"
	"strings"
	"unicode/utf8"
)

const WordleLen = 5

func dictRead(file string) ([]string, error) {
	words := []string{}

	f, err := os.Open(file)
	if err != nil {
		return words, err
	}

	s := bufio.NewScanner(f)
	for s.Scan() {
		word := strings.ToLower(s.Text())
		if utf8.RuneCountInString(word) == WordleLen {
			words = append(words, word)
		}
	}
	err = s.Err()
	if err != nil {
		return words, err
	}

	return words, nil
}

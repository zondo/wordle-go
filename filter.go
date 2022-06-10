package main

import (
	"strconv"
)

type Grade int

const (
	NoMatch Grade = iota
	OtherPos
	Match
)

func filter(words []string, guess, graded string) []string {
	res := []string{}

	for _, word := range words {
		if graded == grades(guess, word) {
			res = append(res, word)
		}
	}

	return res
}

func grades(guess, want string) string {
	hints := make([]Grade, len(guess))
	for i, _ := range hints {
		hints[i] = NoMatch
	}

	wantMap := map[byte]int{}

	// wanted letter counts
	for i := 0; i < len(want); i++ {
		wantMap[want[i]] += 1
	}

	// full matches
	for i := 0; i < len(guess); i++ {
		guessRune := guess[i]
		if guessRune == want[i] {
			hints[i] = Match
			wantMap[guessRune] -= 1
		}
	}

	for i := 0; i < len(guess); i++ {
		guessRune := guess[i]
		if hints[i] == Match {
			continue
		}
		if wantMap[guessRune] > 0 {
			hints[i] = OtherPos
			wantMap[guessRune] -= 1
		}
	}

	res := ""
	for _, hint := range hints {
		res += strconv.Itoa(int(hint))
	}

	return res
}

func bestBang(words []string) string {
	best := ""
	count := 0

	for _, word := range words {
		runes := map[rune]bool{}

		for _, rune := range word {
			runes[rune] = true
		}

		if count == 0 || len(runes) > count {
			count = len(runes)
			best = word
		}
	}

	return best
}

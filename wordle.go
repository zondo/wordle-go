package main

import (
	"fmt"
)

const dictFile = "dict.txt"
const startWord = "oater"

func main() {
	dict, err := dictRead(dictFile)
	if err != nil {
		panic(err)
	}

	newWord := startWord

	for round := 0; ; round++ {
		list(dict, false)

		if round != 0 {
			newWord = bestBang(dict)
		}

		fmt.Printf("Try next: '%s'\n", newWord)
		fmt.Printf("hints(%d)> ", len(dict))

		var word, score string
		fmt.Scanf("%s %s", &word, &score)

		if len(score) == 0 {
			if word == "l" {
				list(dict, true)
			} else {
				fmt.Printf("Invalid input\n")
			}
			continue
		}

		dict = filter(dict, word, score)
	}
}

func list(matches []string, full bool) {
	if len(matches) > 30 && !full {
		fmt.Printf("%d matches ('l' to list).\n",
			len(matches))
	} else {
		for _, w := range matches {
			fmt.Printf("%s\n", w)
		}
	}
}

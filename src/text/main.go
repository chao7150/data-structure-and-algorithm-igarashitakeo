package main

import "fmt"

func KMP(text string, pattern string, slideMap []int) int {
	textRunes := []rune(text)
	textLength := len(textRunes)
	patternRunes := []rune(pattern)
	patternLength := len(patternRunes)
	if textLength < patternLength {
		return 0
	}
	if len(slideMap) != patternLength {
		return 0
	}
	t := 0
	skip := 0
	for {
		for p := range patternRunes {
			fmt.Printf("t: %d, text: %c(%U), p: %d, pattern: %c(%U)\n", t, textRunes[t+p], textRunes[t+p], p, patternRunes[p], patternRunes[p])
			if textRunes[t+p] != patternRunes[p] {
				skip = slideMap[p]
				break
			}
			if p == patternLength-1 {
				return t
			}
		}
		if textLength-patternLength == t {
			// これ以上ループを回すとout of rangeになる
			return 0
		}
		t = t + skip
	}
}

func main() {
	const text = "aあ^~"
	for i := 0; i < len(text); i++ {
		fmt.Printf("%X ", text[i]) // 61 E3 81 82 5E 7E
	}
	for i, v := range text {
		fmt.Printf("%#U starts at byte position %d\n", v, i)
		// U+0061 'a' starts at byte position 0
		// U+3042 'あ' starts at byte position 1
		// U+005E '^' starts at byte position 4
		// U+007E '~' starts at byte position 5
	}
	runes := []rune(text)
	for _, v := range runes {
		fmt.Println(v)
		// 97
		// 12354
		// 94
		// 126
	}
}

// Based on https://gist.github.com/iwinux/4018005
//
// Enhanced to handle unicode char's
//
package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"unicode"
	"unicode/utf8"
)

func main() {
	input := "என்னால் என் கண்ணகளை நம்ப முடியவில்லை. நான் படிக்கும் இந்த வார்த்தைகளை என்னால் புரிந்து கொள்ள முடிகிறது. மனித மூளையின் தனித்துவ சக்தி: கேம்பிரிட்ஜ் பல்கலைக்கழகத்தில் நடத்தியா ஆராய்ச்சியில் தெரியவந்தது என்னவன்றல், வார்த்தையில் உள்ள எழுத்துகளடைய வரிசை முக்கியமானது இல்லை. முதல் மற்றும் கடைசி எழுத்துகள் மட்டும் சரியன இடத்தில் இருந்தால் போதும், மற்ற எழுத்துகள் கலைந்து இருந்தலூம் நம்மால் படித்து பூரிந்து கொள்ள முடியும். இது ஏனெனில் மனித மூளை ஒவ்வரு எழுத்துக்களாக படிபதில்லை, முழு வார்த்தைகளாக படிக்கிறது. இதற்கு பெயர் typoglycemia."

	inputArray := []byte(input)

	fmt.Printf("\n\n%v\n", string(shuffleWords(inputArray)))
}

func shuffleWords(input []byte) []byte {
	words := bytes.Fields(input)

	if len(words) == 0 {
		return input
	}

	for _, word := range words {
		lenOfWord := len(word)

		if lenOfWord <= 3 {
			continue
		}

		lastAlphaIndex := findLastAlpha(word, lenOfWord-1)
		if lastAlphaIndex == -1 || lastAlphaIndex <= 2 {
			continue
		}

		shuffleWord(word, 1, lastAlphaIndex-1)
	}

	return bytes.Join(words, []byte(" "))
}

func shuffleWord(word []byte, start, end int) {
	var posArray []int
	var lenArray []int
	var unicodeCharLength int = 0
	i, j, k, w := 0, 0, 0, 0

	posArray = make([]int, len(word))
	lenArray = make([]int, len(word))
	inputStr := string(word[:])
	inputArray := []byte(inputStr)

	posArray[0] = 0

	j = 0
	for i < len(word) {
		runeValue, width := utf8.DecodeRuneInString(inputStr[i:])
		if !unicode.IsMark(runeValue) && i != 0 && i != len(word)-1 {
			lenArray[k] = i - unicodeCharLength
			unicodeCharLength = i
			k += 1
			posArray[k] = i
		}

		j += 1
		i += width
	}

	lenArray[k] = i - unicodeCharLength - w
	unicodeCharLength = i
	k += 1
	posArray[k] = i

	if k < 3 {
		return
	}

	destPosArray := make([]int, k-1)
	destLenArray := make([]int, k-1)

	perm := rand.Perm(k - 1)
	for j, k := range perm {
		destPosArray[k] = posArray[j]
		destLenArray[k] = lenArray[j]
	}

	i = 0
	offset := 0
	for i < lenArray[0] {
		word[i+offset] = inputArray[i]
		i += 1
	}
	offset = offset + i

	j = 0
	i = 0
	for i < len(destPosArray) {
		if destPosArray[i] == 0 {
			i += 1
			continue
		}
		j = 0
		for j < destLenArray[i] {
			word[j+offset] = inputArray[destPosArray[i]+j]
			j += 1
		}
		offset += j
		i += 1
	}

	i = 0
	for i < lenArray[k-1] {
		word[i+offset] = inputArray[i+posArray[k-1]]
		i += 1
	}
	offset += i
}

func findLastAlpha(str []byte, lastIndex int) int {
	for i := lastIndex; i > 1; i-- {
		if unicode.IsLetter(rune(str[i])) {
			return i
		}
	}
	return -1
}

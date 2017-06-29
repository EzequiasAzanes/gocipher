package gocipher

import "strings"

/*
 * Caesar cipher
 * Keyed Caesar cipher
 */

// CaesarEncipher enciphers string using Caesar cipher according to key.
func CaesarEncipher(text string, key int) string {
	if key == 0 {
		return text
	}
	return mapAlpha(text, func(i, char int) int {
		return char + key
	})
}

// CaesarDecipher deciphers string using Caesar cipher according to key.
func CaesarDecipher(text string, key int) string {
	return CaesarEncipher(text, -key)
}

// CaesarKeyedEncipher enciphers string using keyed Caesar cipher according to key.
func CaesarKeyedEncipher(text string, shift int, key string) string {
	alphabet := KeyedAlphabetRange(strings.ToUpper(key), 'A', 'Z')
	alpha := []rune(alphabet)
	s := rune(shift)
	runes := []rune(text)
	for i, char := range runes {
		if char >= 'A' && char <= 'Z' {
			runes[i] = alpha[modRune(char+s-'A', 26)]
		} else if char >= 'a' && char <= 'z' {
			runes[i] = alpha[modRune(char+s-'a', 26)] - 'A' + 'a'
		}
	}
	return string(runes)
}

// CaesarKeyedDecipher deciphers string using keyed Caesar cipher according to key.
func CaesarKeyedDecipher(text string, shift int, key string) string {
	return CaesarKeyedEncipher(text, -shift, key)
}

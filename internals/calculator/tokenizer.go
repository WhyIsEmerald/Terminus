package calculator

import (
	"errors"
	"unicode"
)

func Tokenize(input string) ([]Token, error) {
	tokens := []Token{}
	currentToken := ""
	i := 0
	for i < len(input) {
		char := rune(input[i])
		if unicode.IsSpace(char) {
			i++
			continue
		}

		if unicode.IsNumber(char) || char == '.' {
			currentToken, i = number(i, input)
			tokens = append(tokens, generateToken(currentToken, "number"))

			continue
		}

		if unicode.IsLetter(char) {
			currentToken, i = letter(i, input)

			if currentToken == "err:unknown" {
				return nil, errors.New("unknown token")
			}
			tokens = append(tokens, generateToken(currentToken, "operator"))
			continue
		}
		if _, ok := operators[string(char)]; ok {
			currentToken = string(char)
			tokens = append(tokens, generateToken(currentToken, "operator"))
			i++
			continue
		}
		if char == '(' {
			currentToken = string(char)
			tokens = append(tokens, generateToken(currentToken, "leftParenthesis"))
			i++

			continue
		}
		if char == ')' {
			currentToken = string(char)
			tokens = append(tokens, generateToken(currentToken, "rightParenthesis"))
			i++
			continue
		}
	}
	return tokens, nil
}

func number(i int, input string) (string, int) {
	currentToken := ""
	char := rune(input[i])
	for unicode.IsNumber(rune(char)) || char == '.' {
		currentToken += string(char)
		i++
		if i < len(input) {
			char = rune(input[i])
		} else {
			break
		}
	}
	return currentToken, i
}
func letter(i int, input string) (string, int) {
	currentToken := ""
	char := rune(input[i])
	for unicode.IsLetter(rune(char)) {
		currentToken += string(char)
		i++
		if i < len(input) {
			char = rune(input[i])
		} else {
			break
		}
	}
	if _, ok := operators[currentToken]; ok {
		return currentToken, i
	} else {
		return "err:unknown", i + 1
	}
}

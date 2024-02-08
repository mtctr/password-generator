package password_generator

import "strings"

var LOWERCASE_LETTERS = []rune("abcdefghijklmnopqrstuvwxyz")
var UPPERCASE_LETTERS = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

type PasswordRules struct {
	hasUppercaseLetters bool
	hasLowercaseLetters bool
	length              int
}

func generatePassword(generateRandom func(int) int, rules PasswordRules) string {
	var passwordBuilder strings.Builder

	for i := 0; i < rules.length; i++ {
		var letter rune
		if rules.hasLowercaseLetters && !rules.hasUppercaseLetters {
			letter = LOWERCASE_LETTERS[generateRandom(len(LOWERCASE_LETTERS))]
		} else if !rules.hasLowercaseLetters && rules.hasUppercaseLetters {
			letter = UPPERCASE_LETTERS[generateRandom(len(UPPERCASE_LETTERS))]
		} else {
			if generateRandom(2) == 1 {
				letter = LOWERCASE_LETTERS[generateRandom(len(LOWERCASE_LETTERS))]

			} else {
				letter = UPPERCASE_LETTERS[generateRandom(len(UPPERCASE_LETTERS))]
			}

		}
		passwordBuilder.WriteRune(letter)
	}

	return passwordBuilder.String()
}

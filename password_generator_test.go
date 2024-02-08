package password_generator

import (
	"math/rand"
	"testing"
	"unicode"
)

func TestGeneratePassword(t *testing.T) {
	t.Run("should generate a string of random lowercase letters with a given length", func(t *testing.T) {
		rules := PasswordRules{hasLowercaseLetters: true, hasUppercaseLetters: false, length: rand.Intn(32)}

		generatedPassword := generatePassword(rand.Intn, rules)

		assertLength(t, rules, generatedPassword)

		hasUpper := false
		hasLower := false

		for _, r := range generatedPassword {
			if unicode.IsUpper(r) {
				hasUpper = true
			}
			if unicode.IsLower(r) {
				hasLower = true
			}
		}

		assertDoesNotHaveUppercase(t, hasUpper, rules, generatedPassword)
		assertDoesHaveLowercase(t, hasLower, rules, generatedPassword)
	})

	t.Run("should generate a string of random uppercase letters with a given length", func(t *testing.T) {
		rules := PasswordRules{hasLowercaseLetters: false, hasUppercaseLetters: true, length: rand.Intn(32)}

		generatedPassword := generatePassword(rand.Intn, rules)

		assertLength(t, rules, generatedPassword)

		hasUpper := false
		hasLower := false

		for _, r := range generatedPassword {
			if unicode.IsUpper(r) {
				hasUpper = true
			}
			if unicode.IsLower(r) {
				hasLower = true
			}
		}

		assertDoesHaveUppercase(t, hasUpper, rules, generatedPassword)
		assertDoesNotHaveLowercase(t, hasLower, rules, generatedPassword)
	})

	t.Run("should generate a string of random lowercase and uppercase letters with a given length", func(t *testing.T) {
		rules := PasswordRules{hasLowercaseLetters: true, hasUppercaseLetters: true, length: rand.Intn(32)}

		generatedPassword := generatePassword(rand.Intn, rules)

		assertLength(t, rules, generatedPassword)

		hasUpper := false
		hasLower := false

		for _, r := range generatedPassword {
			if unicode.IsUpper(r) {
				hasUpper = true
			}
			if unicode.IsLower(r) {
				hasLower = true
			}
		}

		assertDoesHaveUppercase(t, hasUpper, rules, generatedPassword)
		assertDoesHaveLowercase(t, hasLower, rules, generatedPassword)

	})
}

func assertLength(t *testing.T, rules PasswordRules, generatedPassword string) {
	length := len(generatedPassword)

	if length != rules.length {
		t.Fatalf("string should have a length of %d, but has a length of %d", rules.length, length)
	}
}

func assertDoesHaveUppercase(t *testing.T, hasUpper bool, rules PasswordRules, generatedPassword string) {
	if !hasUpper && rules.hasUppercaseLetters {
		t.Errorf("'%s' should have uppercase letters", generatedPassword)
	}
}

func assertDoesNotHaveUppercase(t *testing.T, hasUpper bool, rules PasswordRules, generatedPassword string) {
	if hasUpper && !rules.hasUppercaseLetters {
		t.Errorf("'%s' should not have uppercase letters", generatedPassword)
	}
}

func assertDoesHaveLowercase(t *testing.T, hasLower bool, rules PasswordRules, generatedPassword string) {
	if !hasLower && rules.hasLowercaseLetters {
		t.Errorf("'%s' should have lowercase letters", generatedPassword)
	}
}

func assertDoesNotHaveLowercase(t *testing.T, hasLower bool, rules PasswordRules, generatedPassword string) {
	if hasLower && !rules.hasLowercaseLetters {
		t.Errorf("'%s' should not have lowercase letters", generatedPassword)
	}
}

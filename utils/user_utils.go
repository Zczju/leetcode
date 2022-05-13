package utils

import (
	"unicode/utf16"
)

func UserNameValid(username string) bool {
	if len(utf16.Encode([]rune(username))) > 32 {
		return false
	}
	if len(utf16.Encode([]rune(username))) == 0 {
		return false
	}
	return true
}

func PasswordValid(password string) bool {
	// TODO some testify on length and char
	return len(utf16.Encode([]rune(password))) <= 32
}

func UserIdValid(userId int64) bool {
	// TODO some testify on length and char
	return userId != 0
}

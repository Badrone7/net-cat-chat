package NetCat

/*
	This file contains additional utility functions for the NetCat chat server,
	including message validation, string-to-integer conversion, and name length checking.
	These functions are used throughout the server to ensure that messages are valid,
	to convert string inputs to integers safely, and to enforce name length restrictions.
*/

// IsAllowedMessage checks if the message contains only allowed characters.
func IsAllowedMessage(s string) bool {
	for _, r := range s {
		allowed := false
		for _, a := range AllowedChars {
			if r == a {
				allowed = true
				break
			}
		}
		if !allowed {
			return false
		}
	}
	return true
}

// Personal Atoi Function
func PersoAtoi(s string) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}
	n := 0
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		n = n*10 + int(s[i]-'0')
	}
	return n, true
}

func LenChecker(name string) bool {
	return len(name) <= 20
}

package NetCat

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

package validators

import "regexp"

// Valida que un email sea válido
func IsValidEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}

// Valida que un texto no esté vacío y no exceda cierta longitud
func IsValidText(text string, maxLength int) bool {
	return len(text) > 0 && len(text) <= maxLength
}

package letters

// IsUppercase returns true if the byte is a capital letter.
// Capital letters are intentionally placed letters, while lowercase letters are randomly placed and
// can be overwritten with intentionally placed letters. Everything in the grid should either be a
// lowercase or uppercase letter.
func IsUppercase(b byte) bool {
	return b >= 'A' && b <= 'Z'
}

// ToLowercase turns an byte into lowercase
func ToLowercase(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + 32
	}
	return b
}

// ToUppercase turns a byte into uppercase
func ToUppercase(b byte) byte {
	if b >= 'a' && b <= 'z' {
		return b - 32
	}
	return b
}

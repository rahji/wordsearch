package letters

// IsUppercase returns true if the byte is a capital letter.
func IsUppercase(b byte) bool {
	return b >= 'A' && b <= 'Z'
}

// IsLowercase returns true if the byte is a lowercase letter.
func IsLowercase(b byte) bool {
	return b >= 'a' && b <= 'z'
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

package utils

const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Base62Encode(num int64) string {
	if num == 0 {
		return "0"
	}

	encoded := ""

	for num > 0 {
		rem := num%62
        encoded = string(base62Chars[rem]) + encoded
		num/=62
	}

	return encoded
}
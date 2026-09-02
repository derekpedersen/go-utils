package math

import cryptorand "crypto/rand"

func RandomString(length int) (string, error) {
	if length < 0 {
		return "", ErrInvalidLength
	}
	values := make([]byte, length)
	if _, err := cryptorand.Read(values); err != nil {
		return "", err
	}
	for index, value := range values {
		values[index] = byte(letters[int(value)%len(letters)])
	}
	return string(values), nil
}

func RandomBytes(length int) ([]byte, error) {
	if length < 0 {
		return nil, ErrInvalidLength
	}
	values := make([]byte, length)
	_, err := cryptorand.Read(values)
	return values, err
}

package helper

import (
	"crypto/rand"
	"strconv"
)

func RandomNumbers(length int) (int, error) {

	const numbers = "1234567890"

	tmp := make([]byte, length)

	_, err := rand.Read(tmp)

	if err != nil {
		return 0, err
	}

	numLength := len(numbers)

	for i := 0; i < length; i++ {
		tmp[i] = numbers[int(tmp[i])%numLength]
	}

	return strconv.Atoi(string(tmp))

}

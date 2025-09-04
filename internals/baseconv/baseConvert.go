package baseconv

import (
	"errors"
	"strconv"
	"strings"
)

func BaseConvert(number string, fromBase, toBase int) (string, error) {
	if fromBase < 2 || fromBase > 36 || toBase < 2 || toBase > 36 {
		return "", errors.New("invalid base")
	}

	value, err := strconv.ParseInt(number, fromBase, 64)
	if err != nil {
		return "", err
	}

	if value == 0 {
		return "0", nil
	}

	const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var result strings.Builder

	for value > 0 {
		digit := value % int64(toBase)
		result.WriteByte(charset[digit])
		value /= int64(toBase)
	}

	
	runes := []rune(result.String())
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes), nil
}
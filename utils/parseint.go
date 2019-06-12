package utils

import (
	"fmt"
	"errors"
)

//ParseLen , parse byte to int ,in order to get data size
func ParseLen(p []byte) (int, error) {
	if len(p) == 0 {
		return -1, errors.New("malformed length")
	}
	if p[0] == '-' && len(p) == 2 && p[1] == '1' {
		// handle $-1 and $-1 null replies.
		return -1, nil
	}
	var n int

	for _, b := range p {
		n *= 10

		if b < '0' || b > '9' {
			return -1, errors.New("illegal bytes here in length")
		}
		n += int(b - '0')
	}
	return n, nil
}


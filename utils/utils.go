package utils


import (
	"bufio"
	"errors"
)

//ParseLen get params to caculate length
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


//ReadLen parse byte to int32
func ReadLen(br *bufio.Reader) (int, error) {
	//prefix :=byte('$')
	ls, err := br.ReadBytes('\n')

	ls = ls[:len(ls)-2]  // delete \n chracaters

	if err != nil {
		return 0, err
	}
	if len(ls) < 2 {

		return 0, errors.New("illegal bytes ddd in length")
	}
	if ls[0] != '$' { // start flag

		return 0, errors.New("illegal bytes bbb  in length")
	}
	return ParseLen(ls[1:])
}
package diamond

import (
	"errors"
	"fmt"
	"strings"
)

const (
	A     = 'A'
	Z     = 'Z'
	Space = " "
)

// Index starts from '1'
func indexOf(c byte) (m int) {
	switch {
	case c >= A && c <= Z:
		m = int(c - A + 1)
	default:
		return -1
	}

	return
}

// Makes the diamond
func makeDiamond(s, middleRow, a byte) (string, error) {

	// Index of the element till where the Diamond has to be drawn
	m := indexOf(middleRow)

	if m < 0 {
		return Space, fmt.Errorf("%s", "invalid char")
	}

	// Only the 'A' or 'a'
	if 1 == middleRow {
		return string(s) + "\n", nil
	}

	// Total number of rows (index 1 based so total should be (- 1))
	cols := 2*m - 1

	// Will contain the Diamond representation
	sb := make([]string, cols+1)

	// Representation of a single row in the diamond
	var str []byte

	pivot := m - 1
	for i := 0; i < m; i++ {

		// Fill in the string
		str = []byte(strings.Repeat(Space, cols))

		if i == 0 {

			str[pivot] = s
		} else {

			str[pivot-i] = s
			str[pivot+i] = s
		}

		// Add the row to the array
		sb[i] = string(str)

		// if it is the 0th row
		if i == 0 {
			sb[cols-1] = sb[i]
		} else if i != pivot {
			sb[cols-1-i] = sb[i]
		}

		s++
	}
	return strings.Join(sb, "\n"), nil
}

func Gen(c byte) (s string, e error) {
	switch {
	case c >= A && c <= Z:
		s, e = makeDiamond(A, c, A)
	default:
		return s, errors.New("not a valid character")
	}
	return s, e
}

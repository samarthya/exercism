// Implement variable length quantity encoding and decoding.
package variablelengthquantity

import (
	"errors"
	"fmt"
	"log"
)

const (
	mask = 0x7F
	move = 0x07
	mark = 0x01 << move
)

// encodes a single uint32 into a byte stream
func unpack(num uint32) []byte {
	var (
		out []byte // byte stream for a given uint32
		lsb uint32 // least significant byte
	)

	log.Printf("(Type) %T (Hex) %X (Binary) %b\n", num, num, num)

	//if number is greater than 127 0x7F
	for num > mask {
		// As per the requirement, only the first 7 bits of each byte is significant, which can be retrieved by masking it with 0x7F
		lsb = num & mask // retrieve last 7 bits from num
		log.Printf("MASK'ng: LSB (Hex) %x (Binary) %b\n", lsb, lsb)
		// Next 7 Bits of the number can be obtained by Shifting 7 bits out
		num = num >> move // right shift num by 7 bits
		log.Printf("Next bits  (>> 7)  (Hex) %X (Binary) %b\n", num, num)
		out = append(out, byte(lsb))
	}

	//Append the last byte
	out = append(out, byte(num))

	//Now encode the 28 bit obtained.
	return encode(out)
}

// To indicate which is the last byte of the series,
// you leave bit #7 clear. In all of the preceding bytes, you set bit #7.
func encode(inp []byte) []byte {
	var (
		out  = make([]byte, len(inp))
		i, j = 0, len(inp) - 1
	)

	fmt.Printf(" Byte chunks (7 Bits) %d %#v\n", len(inp), inp)

	// J is initialized as the last element, if it is 0 no need to loop in.
	for j > 0 {
		out[i] = inp[j] | mark // mark the most significant bit to 1
		fmt.Printf(" Enable the significant 7th %b %X\n", out[i], out[i])

		i, j = i+1, j-1
		if i == len(inp) {
			break
		}
	}

	out[i] = inp[j]
	return out
}

// EncodeVarint encodes the uint
func EncodeVarint(input []uint32) (s []byte) {
	for _, num := range input {
		s = append(s, unpack(num)...)
	}
	return
}

// decode the VLQ
func decode(lsb byte, num uint32) (res uint32, more bool) {
	log.Printf(" Decode %X %b\n", lsb, lsb)

	if lsb&mark == mark {
		lsb &= mask
		more = true
	}

	res = num | uint32(lsb)
	log.Printf(" Num: %X %b \n", res, res)
	return
}

// DecodeVarint decodes the VLQ
func DecodeVarint(input []byte) (out []uint32, err error) {
	var (
		num  uint32 // decoded num
		more bool   // flag to check if the sequence terminated
	)

	fmt.Printf(" Input %#v\n", input)

	for _, lsb := range input {
		switch num, more = decode(lsb, num); more {
		case true:
			num = num << move // left shift num by 7 bits to store next lsb
		case false:
			out = append(out, num)
			num = 0x00
		}
	}

	// Handle error
	if more {
		err = errors.New("incomplete sequence causes error")
	}
	return
}

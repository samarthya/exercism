// Module encode
package encode

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func RunLengthEncode(input string) string {
	var runes = []rune(input)
	var sm strings.Builder
	var cnt int
	var i = 0

	// Empty string, do nothing.
	if len(input) <= 0 {
		return input
	}

	for {
		v := runes[i]

		if (i + 1) < len(runes) {
			i++
		} else {
			fmt.Fprintf(&sm, "%c", v)
			break
		}

		if v == runes[i] {
			for cnt = 1; cnt < len(runes); {
				if runes[i] == v {
					cnt++
					if (i + 1) < len(runes) {
						i++
					} else {
						fmt.Fprintf(&sm, "%d%c", cnt, v)
						return sm.String()
					}
					continue
				} else {
					if cnt > 1 {
						fmt.Fprintf(&sm, "%d%c", cnt, v)
					} else {
						fmt.Fprintf(&sm, "%c", v)
					}
					break
				}
			}
		} else {
			// the next one is not the match
			fmt.Fprintf(&sm, "%c", v)
		}
	}

	fmt.Println("Encoded String: ", sm.String())

	return sm.String()
}

func getTheNumber(s []rune) (number int, lastIndex int) {
	log.Println(" Rune: ", string(s))

	for lastIndex < len(s) {
		if unicode.IsDigit(s[lastIndex]) {
			lastIndex++
			continue
		}
		break
	}

	number, _ = strconv.Atoi(string(s[0:lastIndex]))

	log.Println(" Number & Index ", number, lastIndex)

	return number, lastIndex
}

func RunLengthDecode(input string) string {
	var runes = []rune(input)
	var s strings.Builder
	var i int
	for i < len(runes) {
		log.Println(" Checking for ", string(runes[i]))
		v := runes[i]
		if unicode.IsDigit(v) {
			repeat, m := getTheNumber(runes[i:])
			log.Println(" Repeat & Index ", repeat, m)
			fmt.Fprintf(&s, "%s", strings.Repeat(string(runes[i+m]), repeat))
			i += (m + 1)
		} else {
			n := fmt.Sprintf("%s", string(v))
			fmt.Fprintf(&s, "%s", n)
			i++
		}
	}

	return s.String()
}

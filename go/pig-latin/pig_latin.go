package piglatin

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

var Vowel = regexp.MustCompile(`^([aeiou]|xr|yt)[a-z]*`)
var Consonant = regexp.MustCompile(`^([^aeiou]?qu|[^aeiou]+)[a-z]*`)
var YCase = regexp.MustCompile(`^([^aeiou]+y)[a-z]*`)

func AddAy(s string) string {
	return fmt.Sprintf("%say", s)
}

// Sentence
func Sentence(s string) string {
	// Fields splits the string s around each instance of one or more consecutive white space characters
	sw := strings.Fields(s)

	for i, w := range sw {
		// Convert to lowercase
		l := strings.ToLower(w)
		log.Println("String: ", l)
		if Vowel.MatchString(l) {
			sw[i] = AddAy(l)
		} else if YCase.MatchString(l) {
			sw[i] = l[strings.IndexRune(l, 'y'):] + AddAy(l[:strings.IndexRune(l, 'y')])
		} else if x := Consonant.FindStringSubmatchIndex(l); x != nil {
			sw[i] = l[x[3]:] + AddAy(l[:x[3]])
		}
	}
	return strings.Join(sw, " ")
}

package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
)

func Symbols() []rune {
	s := `!@#$%^&*(){}[]|=+:;"'<>,.?/` + "`"
	rc := make([]rune, 0, len(s))
	for _, r := range s {
		rc = append(rc, r)
	}
	return rc
}

func Numbers() []rune {
	rc := make([]rune, 0, 10)
	for r := '0'; r <= '9'; r++ {
		rc = append(rc, r)
	}
	return rc
}

func UpperLetters() []rune {
	rc := make([]rune, 0, 26)
	for r := 'A'; r <= 'Z'; r++ {
		rc = append(rc, r)
	}
	return rc
}

func LowerLetters() []rune {
	rc := make([]rune, 0, 26)
	for r := 'a'; r <= 'z'; r++ {
		rc = append(rc, r)
	}
	return rc
}

func gen(set []rune, l int) string {
	sb := strings.Builder{}
	for i := 0; i < l; i++ {
		r := rand.Intn(len(set))
		sb.WriteRune(set[r])
	}
	return sb.String()
}

func main() {
	charsets := map[string][]rune{
		"lc":  LowerLetters(),
		"uc":  UpperLetters(),
		"num": Numbers(),
		"sym": Symbols(),
	}

	rand.Seed(time.Now().UnixNano())

	sets := flag.String("s", "lc,uc,num,sym", "character sets to use")
	length := flag.Int("l", 32, "length of passwords to generate")
	count := flag.Int("c", 40, "number of passwords to generate")
	flag.Parse()

	charset := []rune(nil)
	for _, set := range strings.Split(*sets, ",") {
		chars, found := charsets[set]
		if !found {
			log.Fatalf("invalid set %q", set)
		}
		charset = append(charset, chars...)
	}

	// entropy := math.Log2(float64(len(charset)*8)) * float64(*length)
	// log.Printf("entropy = %v", entropy)
	// log.Printf("len(charset) = %v", len(charset))

	passwords := []string{}
	for i := 0; i < *count; i++ {
		passwords = append(passwords, gen(charset, *length))
	}

	columns := 40
	written := 0
	for i, password := range passwords {
		fmt.Printf("%s", password)
		written += len(password) + 1
		if i == len(passwords)-1 || written > columns {
			fmt.Printf("\n")
			written = 0
		} else {
			fmt.Printf(" ")
		}
	}
}

package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
)

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
		"lc":  []rune("abcdefghijklmnopqrstuvwxyz"),
		"uc":  []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ"),
		"num": []rune("0123456789"),
		"sym": []rune(`!@#$%^&*(){}[]|=+:;"'<>,.?/` + "`"),
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

	passwords := make([]string, *count)
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

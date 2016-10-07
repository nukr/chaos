package chaos

import (
	"math/rand"
	"strings"
	"time"
)

func basic(alphabetWidth, digitWidth, symbolWidth int) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	alphabet += strings.ToUpper(alphabet)
	digit := "1234567890"
	symbol := "!@#$%^&*()-_=+[]{}\\|;:'\"<>,./?"

	var randString []byte

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < alphabetWidth; i++ {
		length := len(alphabet)
		c := alphabet[r.Int()%length]
		randString = append(randString, c)
	}

	for i := 0; i < digitWidth; i++ {
		length := len(digit)
		c := digit[r.Int()%length]
		randString = append(randString, c)
	}

	for i := 0; i < symbolWidth; i++ {
		length := len(symbol)
		c := symbol[r.Int()%length]
		randString = append(randString, c)
	}

	shuffle(randString)
	return string(randString)
}

// Alphabet ...
func Alphabet(width int) string {
	return basic(width, 0, 0)
}

func shuffle(a []byte) {
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

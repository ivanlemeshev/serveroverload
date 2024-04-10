package password

import (
	"math/rand"
	"time"
)

// Generate generates a random password of the specified length.
// WARNING: This function is for educational purposes only.
// Do not use this function for generating passwords for real-world applications
// there is no guarantee that the generated password is secure.
func Generate(length int) string {
	chars := [4]string{
		"abcdefghijklmnopqrstuvwxyz",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"0123456789",
		"!@#$%^&*()_-+={}[/?",
	}

	source := rand.New(rand.NewSource(time.Now().UnixNano()))

	password := ""
	for n := 0; n < length; n++ {
		randNum := source.Intn(4)
		randCharNum := source.Intn(len(chars[randNum]))
		password += string(chars[randNum][randCharNum])
	}

	return password
}

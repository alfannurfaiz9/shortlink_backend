package utils

import (
	"math/rand"
	"strings"
)

func GenSlug() string {
	alphanum := "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890abcdefghijklmnopqrstuvwxyz"
	lenSlug := 6
	var slug strings.Builder

	for range lenSlug {
		slug.WriteString(string(alphanum[rand.Intn(len(alphanum))]))
	}

	return slug.String()
}

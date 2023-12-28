package utility

import "math/rand"

type Utility interface {
	GenerateSlug() string
}

type utility struct{}

func NewUtility() Utility {
	return &utility{}
}

func (u *utility) GenerateSlug() string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz1234567890")
	b := make([]rune, 6)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

package util

import (
	crand "crypto/rand"
	"math/big"
	"math/rand"
	"time"

	"github.com/sirupsen/logrus"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RandPrime(n int) *big.Int {
	p, err := crand.Prime(crand.Reader, n)
	if err != nil {
		logrus.Panic("err creating random prime. ", err.Error())
	}
	return p
}

package record

import (
	"math/rand"
	"time"
)

//NOTE: this variable and this funcion are normal camel case
//normalCamelCase is use for non-exported variables so it's inteded for internal use only
//TrueCamelCase is used for external variables

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ1234567890"

func randString(n int) string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[r1.Intn(len(letterBytes))]
	}
	return string(b)
}

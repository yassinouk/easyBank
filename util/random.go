package util

import "math/rand"

func init() {

}

/// randomInt generates a random integer between min and max
func randomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	k := len(letters)
	result := make([]byte, n)
	for i := range result {
		result[i] = letters[rand.Intn(k)]
	}
	return string(result)
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money
func RandomBalance() int64 {
	return randomInt(0, 100000)
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "MAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

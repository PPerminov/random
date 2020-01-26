package random

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestRandomInt(t *testing.T) {
	strconv.Itoa(RandomInt())
}

func TestRandomLimitedInt(t *testing.T) {
	_, err := RandomLimitedInt(200, 2)
	if err == nil {
		t.Error(err)
	}
	a, err := RandomLimitedInt(600, 600)
	if err != nil {
		t.Error(err)
	}
	a, err = RandomLimitedInt(600, 601)
	if err != nil || a < 600 || a > 601 {
		t.Error("Something wrong with Limited integer generation")
	}
	a, err = RandomLimitedInt(6000, 60001)
	if err != nil || a < 6000 || a > 60001 {
		t.Error("Something wrong with Limited integer generation")
	}

}

func TestRandomBytes(t *testing.T) {
	data, err := RandomBytes(12)
	if err != nil || len(data) != 12 {
		t.Error(err)
	}

}

func TestRandomStringAllNumbers(t *testing.T) {
	number, err := RandomString(16, Numbers)
	if err != nil {
		t.Error(err)
	}
	_, err = strconv.Atoi(number)
	if err != nil {
		t.Fatal(err)
	}

}

func TestRandomStringAllCapitalLetters(t *testing.T) {
	str, err := RandomString(16, CapitalLetters)
	if err != nil {
		t.Error(err)
	}
	bytes := []byte(str)
	for _, i := range bytes {
		if ! (65 <= i && i <= 90) {
			t.Fail()
		}
	}
}
func TestRandomStringAllLetters(t *testing.T) {
	str, err := RandomString(16, Letters)
	if err != nil {
		t.Error(err)
	}
	bytes := []byte(str)
	for _, i := range bytes {
		if ! (97 <= i && i <= 122) {
			t.Fail()
		}
	}
}

func TestRandomString(t *testing.T) {
	a := []string{}
	for it := 0; it < 1000; it++ {
		b, err := RandomString(10, 105)
		if err != nil {
			t.Error(err)
		}
		a = append(a, b)
	}
	m := map[string]bool{}
	for _, k := range a {
		m[k] = true
	}
	if len(a) != len(m) {
		t.Fail()
	}
}

package random

import (
	"errors"
	"math/rand"
	"time"
)

var MinLargerThanMax = errors.New("Minimal value is bigger or equal to maximal value")
//var MinEqualMax = errors.New("Min and Max are the same")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt() int {
	return rand.Int()
}

func RandomLimitedInt(min, max int) (int, error) {
	delta := max - min
	if delta < 0 {
		return 0, MinLargerThanMax
	}
	return rand.Intn(delta+1) + min, nil
}

func RandomBytes(size int) (rb []byte, err error) {
	rb = make([]byte, size)
	_, err = rand.Read(rb)
	return
}

func RandomString(size int, t int) (s string, err error) {
	rb := make([]byte, 0)
	sorter := Constructor(t)
	for len(rb) != size {
		rbTmp := make([]byte, 1)
		_, err = rand.Read(rbTmp)
		if sorter(z{rbTmp[0], false}).t {
			rb = append(rb, rbTmp[0])
		}
	}
	return string(rb), err
}

type z struct {
	b byte
	t bool
}

var Letters = 3
var CapitalLetters = 5
var Numbers = 7
//var Symbols = 11

func getFunctions() map[int]func(z) z {
	f := make(map[int]func(z) z, 3)
	branches := map[int][]uint8{}
	branches[Letters] = []uint8{97, 122}
	branches[CapitalLetters] = []uint8{65, 90}
	branches[Numbers] = []uint8{48, 57}
	constructor := func(n []uint8) func(z) z {
		return func(zz z) z {
			if (! zz.t) && (n[0] <= zz.b && zz.b <= n[1]) {
				zz.t = true
			}
			return zz
		}
	}
	f[Letters] = constructor(branches[Letters])
	f[CapitalLetters] = constructor(branches[CapitalLetters])
	f[Numbers] = constructor(branches[Numbers])
	return f
}

func Constructor(o int) func(z) z {
	functions := getFunctions()
	functionsToGo := []func(z) z{}
	if o%Letters == 0 {
		functionsToGo = append(functionsToGo, functions[Letters])
	}
	if o%Numbers == 0 {
		functionsToGo = append(functionsToGo, functions[Numbers])
	}
	if o%CapitalLetters == 0 {
		functionsToGo = append(functionsToGo, functions[CapitalLetters])
	}
	return func(z2 z) z {
		for _, f := range functionsToGo {
			z2 = f(z2)
		}
		return z2
	}
}

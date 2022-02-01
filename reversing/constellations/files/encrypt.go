// go run encrypt.go

package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"strings"
)

const FLAG = "SECCON{N33d_m0r3_sp33d_vo6RgykRuK8rY9r07kLO3Aj9xsfffimRWK7ferM8MU4q5qoP32yKOaPyWcmCKyJ6yIgWJOBP5eTA8lgRl7u3JinsZPqlItrjnbsTIZ5uhnLCd5KsLcsena9wmdclyV7H_Wh47_1s_y0ur_b1r7h_51gn?}"

var Constellations = []string{
	"Cancer",
	"Aquarius",
	"Pisces",
	"Aries",
	"Leo",
	"Virgo",
	"Capricorn",
	"Gemini",
	"Scorpio",
	"Sagittarius",
	"Libra",
	"Taurus",
}

func main() {
	rand.Seed(2938227336527353971)
	for i := 0; i < len(FLAG); i++ {
		x := big.NewInt(0)
		for j := 0; j < i; j++ {
			x.Mul(x, big.NewInt(2))
			x.Add(x, big.NewInt(rand.Int63n(2)))
		}
		for {
			t := new(big.Int)
			t.Mul(x, big.NewInt(5))
			t.Mod(t, big.NewInt(256))
			if t.Int64() == int64(FLAG[i]) {
				break
			}
			x.Add(x, big.NewInt(1))
		}
		v := []string{}
		for x.Sign() > 0 {
			t := new(big.Int)
			t.Mod(x, big.NewInt(12))
			v = append(v, Constellations[t.Int64()])
			x.Div(x, big.NewInt(12))
		}
		for i2, j := 0, len(v)-1; i2 < j; i2, j = i2+1, j-1 {
			v[i2], v[j] = v[j], v[i2]
		}
		fmt.Println(strings.Join(v, "_"))
	}
}

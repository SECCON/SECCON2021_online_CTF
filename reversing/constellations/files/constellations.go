//	go build constellations.go

package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main() {
	a := []string{
		"Sagittarius_Taurus",
		"Virgo_Virgo",
		"Aquarius_Aquarius_Taurus",
		"Aquarius_Aquarius_Taurus",
		"Virgo_Gemini",
		"Sagittarius_Libra",
		"Libra_Gemini",
		"Sagittarius_Libra",
		"Aries_Aries_Aries",
		"Aries_Aries_Aries",
		"Gemini_Aries_Cancer",
		"Aquarius_Gemini",
		"Aquarius_Leo_Capricorn_Aquarius",
		"Pisces_Gemini_Cancer_Cancer",
		"Capricorn_Capricorn_Scorpio_Libra",
		"Aries_Aries_Aries",
		"Capricorn_Taurus_Scorpio_Aries",
		"Pisces_Pisces_Taurus_Scorpio_Gemini",
		"Aquarius_Cancer_Pisces_Taurus_Cancer_Cancer",
		"Aquarius_Gemini_Cancer_Capricorn_Sagittarius_Taurus",
		"Pisces_Leo_Cancer_Scorpio_Gemini_Aries",
		"Virgo_Cancer_Scorpio_Cancer_Aquarius_Scorpio",
		"Aquarius_Leo_Gemini_Sagittarius_Virgo_Libra_Taurus",
		"Aquarius_Sagittarius_Virgo_Aquarius_Leo_Libra_Capricorn",
		"Leo_Leo_Cancer_Virgo_Aries_Sagittarius_Gemini",
		"Gemini_Virgo_Virgo_Sagittarius_Gemini_Taurus_Libra",
		"Aries_Virgo_Scorpio_Taurus_Sagittarius_Gemini_Capricorn",
		"Aquarius_Sagittarius_Leo_Aquarius_Gemini_Pisces_Pisces_Aries",
		"Pisces_Capricorn_Cancer_Pisces_Cancer_Aries_Leo_Virgo",
		"Sagittarius_Virgo_Virgo_Pisces_Gemini_Pisces_Capricorn_Gemini",
		"Aquarius_Pisces_Leo_Taurus_Scorpio_Aries_Gemini_Libra_Pisces",
		"Aries_Leo_Capricorn_Scorpio_Leo_Aries_Taurus_Virgo_Virgo",
		"Virgo_Aries_Libra_Capricorn_Virgo_Leo_Virgo_Virgo_Aries",
		"Virgo_Aries_Scorpio_Taurus_Aquarius_Virgo_Cancer_Pisces_Cancer",
		"Aquarius_Taurus_Taurus_Cancer_Leo_Sagittarius_Capricorn_Aquarius_Leo_Libra",
		"Leo_Leo_Aquarius_Aries_Cancer_Scorpio_Libra_Capricorn_Scorpio_Virgo",
		"Leo_Sagittarius_Aries_Cancer_Pisces_Sagittarius_Capricorn_Gemini_Leo_Virgo",
		"Pisces_Cancer_Taurus_Libra_Gemini_Libra_Cancer_Cancer_Capricorn_Pisces",
		"Aquarius_Sagittarius_Taurus_Cancer_Sagittarius_Pisces_Virgo_Aries_Gemini_Virgo_Leo",
		"Leo_Leo_Taurus_Sagittarius_Virgo_Libra_Cancer_Gemini_Aquarius_Leo_Taurus",
		"Gemini_Leo_Aquarius_Pisces_Virgo_Leo_Taurus_Capricorn_Pisces_Aquarius_Aries",
		"Aries_Aries_Scorpio_Libra_Virgo_Capricorn_Gemini_Pisces_Taurus_Aries_Scorpio",
		"Aquarius_Capricorn_Leo_Pisces_Gemini_Scorpio_Taurus_Sagittarius_Aquarius_Scorpio_Cancer_Aries",
		"Capricorn_Scorpio_Capricorn_Scorpio_Pisces_Libra_Taurus_Virgo_Gemini_Aquarius_Cancer_Gemini",
		"Leo_Sagittarius_Pisces_Leo_Libra_Virgo_Sagittarius_Taurus_Cancer_Scorpio_Taurus_Sagittarius",
		"Pisces_Capricorn_Capricorn_Taurus_Gemini_Libra_Pisces_Aquarius_Capricorn_Scorpio_Aries_Leo_Pisces",
		"Aquarius_Capricorn_Libra_Sagittarius_Capricorn_Taurus_Aries_Cancer_Virgo_Leo_Aquarius_Aquarius_Sagittarius",
		"Leo_Leo_Scorpio_Libra_Sagittarius_Gemini_Gemini_Leo_Taurus_Gemini_Aries_Aries_Leo",
		"Aquarius_Aries_Pisces_Sagittarius_Aquarius_Capricorn_Virgo_Aquarius_Taurus_Virgo_Aquarius_Pisces_Leo_Gemini",
		"Leo_Aries_Capricorn_Virgo_Leo_Virgo_Pisces_Virgo_Taurus_Gemini_Libra_Pisces_Aquarius_Pisces",
		"Virgo_Sagittarius_Aquarius_Gemini_Aries_Taurus_Capricorn_Aquarius_Aquarius_Cancer_Leo_Sagittarius_Pisces",
		"Libra_Aries_Leo_Leo_Sagittarius_Leo_Sagittarius_Pisces_Libra_Virgo_Scorpio_Leo_Sagittarius_Pisces",
		"Aquarius_Leo_Gemini_Pisces_Scorpio_Cancer_Taurus_Sagittarius_Scorpio_Cancer_Taurus_Taurus_Aries_Aries_Aquarius",
		"Aries_Taurus_Libra_Gemini_Scorpio_Leo_Scorpio_Leo_Leo_Scorpio_Leo_Sagittarius_Leo_Sagittarius",
		"Leo_Virgo_Libra_Aries_Gemini_Capricorn_Virgo_Leo_Gemini_Leo_Taurus_Cancer_Aquarius_Pisces_Pisces",
		"Aquarius_Aquarius_Leo_Libra_Leo_Gemini_Virgo_Gemini_Gemini_Aries_Capricorn_Gemini_Libra_Sagittarius_Pisces_Aries",
		"Virgo_Sagittarius_Capricorn_Pisces_Aries_Aries_Sagittarius_Taurus_Aquarius_Gemini_Virgo_Pisces_Leo_Capricorn_Gemini",
		"Leo_Gemini_Leo_Cancer_Libra_Capricorn_Leo_Libra_Cancer_Leo_Pisces_Aries_Cancer_Cancer_Cancer_Taurus",
		"Sagittarius_Capricorn_Cancer_Scorpio_Virgo_Virgo_Aquarius_Sagittarius_Virgo_Sagittarius_Gemini_Pisces_Scorpio_Libra_Aquarius_Pisces",
		"Aries_Virgo_Libra_Capricorn_Leo_Scorpio_Taurus_Libra_Gemini_Taurus_Aquarius_Capricorn_Scorpio_Aquarius_Capricorn_Sagittarius",
		"Aquarius_Leo_Scorpio_Leo_Cancer_Aries_Scorpio_Sagittarius_Capricorn_Leo_Taurus_Capricorn_Virgo_Aquarius_Scorpio_Cancer_Libra",
		"Taurus_Cancer_Sagittarius_Aries_Aquarius_Gemini_Sagittarius_Aquarius_Capricorn_Capricorn_Scorpio_Cancer_Scorpio_Sagittarius_Virgo_Pisces_Aquarius",
		"Taurus_Virgo_Cancer_Aquarius_Pisces_Cancer_Libra_Virgo_Leo_Sagittarius_Virgo_Pisces_Scorpio_Aries_Capricorn_Leo_Scorpio",
		"Aquarius_Sagittarius_Leo_Scorpio_Aries_Cancer_Sagittarius_Leo_Libra_Cancer_Virgo_Leo_Virgo_Sagittarius_Pisces_Pisces_Cancer_Sagittarius",
		"Leo_Aquarius_Sagittarius_Libra_Aquarius_Taurus_Gemini_Aquarius_Capricorn_Cancer_Capricorn_Libra_Leo_Taurus_Aries_Virgo_Cancer_Aquarius",
		"Taurus_Pisces_Sagittarius_Leo_Gemini_Sagittarius_Capricorn_Scorpio_Cancer_Sagittarius_Leo_Taurus_Leo_Gemini_Taurus_Leo_Aries_Cancer",
		"Pisces_Cancer_Scorpio_Taurus_Pisces_Pisces_Sagittarius_Pisces_Leo_Pisces_Libra_Pisces_Pisces_Leo_Gemini_Leo_Gemini_Taurus_Sagittarius",
		"Aries_Taurus_Libra_Gemini_Virgo_Gemini_Pisces_Cancer_Aries_Taurus_Sagittarius_Gemini_Pisces_Virgo_Virgo_Libra_Virgo_Pisces_Sagittarius",
		"Aquarius_Virgo_Aries_Sagittarius_Gemini_Capricorn_Scorpio_Sagittarius_Taurus_Leo_Pisces_Pisces_Scorpio_Virgo_Aquarius_Sagittarius_Libra_Pisces_Virgo",
		"Sagittarius_Leo_Leo_Capricorn_Taurus_Taurus_Virgo_Sagittarius_Leo_Capricorn_Virgo_Cancer_Virgo_Pisces_Taurus_Pisces_Pisces_Virgo_Gemini",
		"Aquarius_Sagittarius_Sagittarius_Cancer_Aries_Virgo_Taurus_Cancer_Scorpio_Virgo_Libra_Capricorn_Cancer_Scorpio_Capricorn_Scorpio_Capricorn_Capricorn_Leo_Cancer",
		"Pisces_Pisces_Pisces_Aquarius_Aquarius_Leo_Libra_Sagittarius_Scorpio_Gemini_Taurus_Aquarius_Gemini_Leo_Pisces_Aquarius_Scorpio_Scorpio_Gemini_Aries",
		"Virgo_Leo_Aquarius_Aries_Capricorn_Aquarius_Virgo_Cancer_Libra_Cancer_Libra_Sagittarius_Aries_Capricorn_Pisces_Aquarius_Aries_Aquarius_Leo_Libra",
		"Aquarius_Sagittarius_Aries_Aries_Taurus_Virgo_Leo_Scorpio_Aries_Virgo_Sagittarius_Capricorn_Capricorn_Virgo_Leo_Taurus_Cancer_Leo_Libra_Virgo_Sagittarius",
		"Aquarius_Cancer_Cancer_Aquarius_Taurus_Libra_Scorpio_Leo_Aries_Gemini_Gemini_Virgo_Aries_Sagittarius_Pisces_Taurus_Aries_Aries_Aries_Pisces_Gemini",
		"Capricorn_Taurus_Gemini_Taurus_Cancer_Aries_Pisces_Virgo_Scorpio_Capricorn_Leo_Cancer_Scorpio_Aquarius_Virgo_Leo_Libra_Pisces_Capricorn_Scorpio_Aries",
		"Sagittarius_Capricorn_Cancer_Aquarius_Pisces_Gemini_Capricorn_Capricorn_Gemini_Libra_Scorpio_Libra_Gemini_Capricorn_Cancer_Sagittarius_Libra_Scorpio_Leo_Sagittarius_Aquarius",
		"Pisces_Virgo_Gemini_Gemini_Gemini_Aries_Capricorn_Aries_Aquarius_Capricorn_Aquarius_Leo_Scorpio_Taurus_Leo_Aquarius_Gemini_Gemini_Virgo_Cancer_Cancer",
		"Aquarius_Taurus_Capricorn_Sagittarius_Pisces_Libra_Sagittarius_Aquarius_Aquarius_Aquarius_Sagittarius_Sagittarius_Taurus_Aquarius_Cancer_Scorpio_Aquarius_Cancer_Scorpio_Aries_Leo_Virgo",
		"Libra_Taurus_Aquarius_Pisces_Virgo_Leo_Scorpio_Gemini_Capricorn_Sagittarius_Taurus_Leo_Gemini_Pisces_Scorpio_Pisces_Gemini_Cancer_Aries_Pisces_Capricorn_Aries",
		"Sagittarius_Leo_Scorpio_Sagittarius_Scorpio_Libra_Capricorn_Aries_Leo_Cancer_Cancer_Gemini_Sagittarius_Cancer_Libra_Leo_Virgo_Aries_Scorpio_Pisces_Aries_Aries",
		"Aquarius_Aquarius_Aries_Leo_Cancer_Scorpio_Libra_Aquarius_Pisces_Capricorn_Scorpio_Pisces_Virgo_Libra_Virgo_Libra_Aries_Capricorn_Sagittarius_Aries_Aries_Gemini_Virgo",
		"Pisces_Aries_Pisces_Aquarius_Sagittarius_Taurus_Gemini_Pisces_Leo_Sagittarius_Aries_Scorpio_Capricorn_Aquarius_Aquarius_Aquarius_Capricorn_Virgo_Aquarius_Capricorn_Sagittarius_Aquarius_Taurus",
		"Aquarius_Aries_Gemini_Libra_Cancer_Taurus_Aries_Gemini_Virgo_Aquarius_Capricorn_Capricorn_Cancer_Sagittarius_Gemini_Cancer_Aries_Taurus_Capricorn_Pisces_Capricorn_Leo_Capricorn_Gemini",
		"Aquarius_Pisces_Gemini_Leo_Libra_Capricorn_Cancer_Libra_Aquarius_Virgo_Gemini_Aries_Scorpio_Libra_Aries_Scorpio_Gemini_Taurus_Cancer_Cancer_Aquarius_Virgo_Gemini_Aquarius",
		"Virgo_Sagittarius_Gemini_Scorpio_Sagittarius_Aquarius_Gemini_Virgo_Sagittarius_Libra_Aquarius_Pisces_Gemini_Pisces_Leo_Cancer_Aquarius_Taurus_Virgo_Aquarius_Sagittarius_Taurus_Capricorn_Libra",
		"Virgo_Cancer_Cancer_Taurus_Taurus_Gemini_Aquarius_Libra_Libra_Cancer_Pisces_Leo_Libra_Gemini_Virgo_Aquarius_Sagittarius_Libra_Aries_Capricorn_Gemini_Virgo_Aries_Libra",
		"Aquarius_Cancer_Capricorn_Aries_Capricorn_Sagittarius_Taurus_Virgo_Sagittarius_Cancer_Scorpio_Sagittarius_Aries_Pisces_Libra_Virgo_Sagittarius_Cancer_Sagittarius_Cancer_Scorpio_Sagittarius_Gemini_Leo_Virgo",
		"Aries_Aquarius_Sagittarius_Aquarius_Pisces_Gemini_Sagittarius_Libra_Sagittarius_Sagittarius_Libra_Sagittarius_Libra_Virgo_Taurus_Virgo_Aries_Taurus_Gemini_Cancer_Libra_Leo_Cancer_Sagittarius_Sagittarius",
		"Aries_Pisces_Sagittarius_Taurus_Aries_Pisces_Libra_Aries_Scorpio_Aries_Taurus_Gemini_Virgo_Scorpio_Aquarius_Scorpio_Capricorn_Libra_Gemini_Cancer_Virgo_Libra_Aries_Capricorn_Aries",
		"Aquarius_Cancer_Libra_Aries_Capricorn_Gemini_Sagittarius_Pisces_Libra_Scorpio_Capricorn_Pisces_Libra_Libra_Scorpio_Sagittarius_Gemini_Virgo_Taurus_Sagittarius_Cancer_Libra_Aries_Taurus_Leo_Taurus",
		"Taurus_Scorpio_Sagittarius_Aquarius_Leo_Capricorn_Cancer_Taurus_Taurus_Sagittarius_Aries_Taurus_Aquarius_Libra_Gemini_Gemini_Scorpio_Virgo_Aquarius_Sagittarius_Taurus_Virgo_Leo_Pisces",
		"Virgo_Gemini_Capricorn_Virgo_Scorpio_Cancer_Taurus_Libra_Aquarius_Capricorn_Scorpio_Aries_Cancer_Sagittarius_Leo_Sagittarius_Cancer_Sagittarius_Cancer_Gemini_Pisces_Capricorn_Libra_Pisces_Taurus",
		"Leo_Aquarius_Aries_Aquarius_Gemini_Sagittarius_Scorpio_Scorpio_Aries_Gemini_Pisces_Cancer_Libra_Cancer_Capricorn_Aries_Leo_Sagittarius_Aries_Libra_Libra_Aries_Scorpio_Aries_Aries_Capricorn",
		"Taurus_Libra_Sagittarius_Cancer_Libra_Leo_Libra_Gemini_Cancer_Gemini_Sagittarius_Gemini_Aries_Libra_Sagittarius_Capricorn_Leo_Libra_Sagittarius_Scorpio_Pisces_Libra_Aries_Scorpio_Capricorn_Scorpio",
		"Libra_Aries_Pisces_Aquarius_Scorpio_Aries_Libra_Pisces_Sagittarius_Scorpio_Taurus_Scorpio_Capricorn_Taurus_Scorpio_Aquarius_Leo_Gemini_Aries_Pisces_Taurus_Sagittarius_Virgo_Capricorn_Capricorn_Sagittarius",
		"Pisces_Virgo_Pisces_Libra_Libra_Libra_Scorpio_Virgo_Sagittarius_Leo_Taurus_Scorpio_Virgo_Scorpio_Cancer_Libra_Gemini_Capricorn_Scorpio_Pisces_Leo_Scorpio_Leo_Sagittarius_Pisces_Cancer_Aquarius",
		"Gemini_Capricorn_Aquarius_Libra_Gemini_Aries_Taurus_Sagittarius_Gemini_Virgo_Leo_Leo_Aries_Cancer_Virgo_Gemini_Aquarius_Sagittarius_Capricorn_Libra_Sagittarius_Aries_Pisces_Capricorn_Aries_Aquarius_Scorpio",
		"Sagittarius_Cancer_Virgo_Gemini_Gemini_Aquarius_Pisces_Capricorn_Sagittarius_Pisces_Pisces_Libra_Libra_Cancer_Gemini_Sagittarius_Capricorn_Leo_Pisces_Leo_Sagittarius_Gemini_Capricorn_Pisces_Taurus_Gemini_Sagittarius",
		"Leo_Capricorn_Libra_Virgo_Gemini_Leo_Leo_Aquarius_Taurus_Taurus_Aries_Capricorn_Sagittarius_Virgo_Aquarius_Virgo_Leo_Sagittarius_Aquarius_Gemini_Gemini_Capricorn_Aries_Scorpio_Scorpio_Aquarius_Capricorn_Cancer",
		"Gemini_Libra_Taurus_Aries_Aquarius_Virgo_Aquarius_Leo_Leo_Aries_Aquarius_Scorpio_Pisces_Leo_Leo_Libra_Leo_Aries_Aries_Aries_Pisces_Sagittarius_Leo_Leo_Aries_Virgo_Sagittarius_Cancer",
		"Aquarius_Cancer_Aquarius_Sagittarius_Sagittarius_Capricorn_Capricorn_Aries_Libra_Gemini_Pisces_Sagittarius_Taurus_Sagittarius_Aquarius_Cancer_Virgo_Leo_Taurus_Libra_Taurus_Libra_Virgo_Capricorn_Libra_Leo_Capricorn_Pisces_Aries",
		"Pisces_Aquarius_Aquarius_Cancer_Cancer_Leo_Scorpio_Pisces_Capricorn_Cancer_Scorpio_Aries_Libra_Libra_Aquarius_Sagittarius_Scorpio_Sagittarius_Gemini_Aries_Aries_Pisces_Pisces_Aquarius_Cancer_Aries_Sagittarius_Gemini_Capricorn",
		"Leo_Capricorn_Scorpio_Aquarius_Taurus_Aries_Leo_Capricorn_Capricorn_Gemini_Taurus_Taurus_Scorpio_Gemini_Gemini_Virgo_Leo_Capricorn_Virgo_Cancer_Sagittarius_Scorpio_Aquarius_Sagittarius_Sagittarius_Aquarius_Scorpio_Virgo_Cancer",
		"Libra_Cancer_Aquarius_Libra_Leo_Gemini_Scorpio_Sagittarius_Cancer_Taurus_Scorpio_Leo_Gemini_Pisces_Scorpio_Taurus_Sagittarius_Taurus_Aries_Virgo_Libra_Sagittarius_Cancer_Aries_Gemini_Taurus_Leo_Taurus_Gemini",
		"Gemini_Taurus_Cancer_Aries_Scorpio_Gemini_Leo_Cancer_Capricorn_Leo_Cancer_Aquarius_Aries_Pisces_Aries_Sagittarius_Aquarius_Scorpio_Pisces_Capricorn_Cancer_Capricorn_Leo_Capricorn_Capricorn_Libra_Cancer_Leo_Aquarius",
		"Libra_Pisces_Capricorn_Pisces_Aries_Aquarius_Virgo_Libra_Gemini_Sagittarius_Scorpio_Gemini_Cancer_Leo_Aquarius_Gemini_Pisces_Cancer_Virgo_Leo_Leo_Leo_Aquarius_Pisces_Gemini_Capricorn_Leo_Gemini",
		"Pisces_Virgo_Gemini_Taurus_Sagittarius_Libra_Leo_Libra_Gemini_Sagittarius_Capricorn_Leo_Capricorn_Aries_Aquarius_Libra_Leo_Taurus_Capricorn_Scorpio_Virgo_Sagittarius_Leo_Taurus_Virgo_Scorpio_Gemini_Aries_Capricorn_Libra",
		"Scorpio_Cancer_Taurus_Sagittarius_Sagittarius_Sagittarius_Scorpio_Leo_Scorpio_Taurus_Libra_Pisces_Gemini_Libra_Scorpio_Leo_Aquarius_Capricorn_Leo_Libra_Libra_Taurus_Libra_Aries_Capricorn_Cancer_Aquarius_Pisces_Leo_Virgo",
		"Pisces_Aquarius_Scorpio_Cancer_Gemini_Aquarius_Gemini_Taurus_Libra_Pisces_Aries_Pisces_Pisces_Capricorn_Scorpio_Virgo_Libra_Libra_Virgo_Aquarius_Taurus_Gemini_Pisces_Sagittarius_Sagittarius_Scorpio_Scorpio_Capricorn_Capricorn_Leo_Capricorn",
		"Leo_Taurus_Libra_Sagittarius_Aquarius_Gemini_Taurus_Scorpio_Scorpio_Virgo_Virgo_Cancer_Libra_Scorpio_Taurus_Libra_Capricorn_Aries_Sagittarius_Taurus_Sagittarius_Leo_Libra_Leo_Scorpio_Sagittarius_Leo_Leo_Sagittarius_Cancer_Gemini",
		"Libra_Cancer_Virgo_Leo_Pisces_Leo_Libra_Leo_Libra_Sagittarius_Leo_Libra_Scorpio_Virgo_Aries_Leo_Aries_Cancer_Pisces_Virgo_Pisces_Aquarius_Libra_Cancer_Leo_Aries_Sagittarius_Sagittarius_Taurus_Pisces_Libra",
		"Aquarius_Leo_Pisces_Leo_Taurus_Virgo_Pisces_Aries_Taurus_Virgo_Pisces_Gemini_Pisces_Pisces_Aries_Scorpio_Taurus_Taurus_Capricorn_Scorpio_Pisces_Leo_Taurus_Leo_Gemini_Leo_Libra_Aries_Cancer_Aries_Scorpio_Cancer",
		"Pisces_Gemini_Virgo_Cancer_Aries_Pisces_Libra_Aquarius_Taurus_Aquarius_Gemini_Sagittarius_Scorpio_Libra_Aquarius_Pisces_Scorpio_Taurus_Sagittarius_Capricorn_Cancer_Capricorn_Capricorn_Libra_Libra_Aquarius_Libra_Gemini_Virgo_Capricorn_Gemini_Sagittarius",
		"Taurus_Aries_Cancer_Libra_Sagittarius_Gemini_Pisces_Taurus_Virgo_Pisces_Sagittarius_Aries_Gemini_Gemini_Aquarius_Leo_Aries_Virgo_Capricorn_Leo_Leo_Aries_Cancer_Taurus_Aries_Cancer_Sagittarius_Aquarius_Aries_Aquarius_Cancer",
		"Gemini_Sagittarius_Gemini_Gemini_Taurus_Capricorn_Scorpio_Virgo_Pisces_Pisces_Taurus_Gemini_Capricorn_Scorpio_Capricorn_Capricorn_Aquarius_Libra_Gemini_Taurus_Aquarius_Aries_Capricorn_Aries_Leo_Sagittarius_Taurus_Scorpio_Sagittarius_Libra_Aquarius_Sagittarius",
		"Scorpio_Gemini_Pisces_Taurus_Leo_Leo_Libra_Cancer_Aries_Virgo_Gemini_Capricorn_Leo_Virgo_Taurus_Leo_Libra_Virgo_Libra_Pisces_Capricorn_Scorpio_Virgo_Capricorn_Aries_Scorpio_Leo_Cancer_Aquarius_Aries_Sagittarius_Scorpio",
		"Capricorn_Aries_Capricorn_Gemini_Aries_Leo_Virgo_Virgo_Aquarius_Libra_Aquarius_Capricorn_Libra_Aquarius_Taurus_Gemini_Leo_Leo_Pisces_Leo_Libra_Pisces_Pisces_Cancer_Leo_Cancer_Virgo_Aquarius_Aries_Aquarius_Libra_Pisces",
		"Gemini_Capricorn_Libra_Taurus_Sagittarius_Virgo_Taurus_Cancer_Sagittarius_Pisces_Aquarius_Gemini_Aries_Libra_Pisces_Taurus_Cancer_Gemini_Leo_Cancer_Sagittarius_Pisces_Aries_Pisces_Aries_Aquarius_Aquarius_Aquarius_Cancer_Leo_Libra_Virgo_Capricorn",
		"Aquarius_Leo_Scorpio_Capricorn_Leo_Aries_Gemini_Cancer_Libra_Cancer_Aries_Gemini_Virgo_Aquarius_Sagittarius_Capricorn_Pisces_Leo_Libra_Scorpio_Aquarius_Aries_Cancer_Gemini_Cancer_Aquarius_Scorpio_Libra_Leo_Cancer_Aries_Aquarius_Virgo_Libra",
		"Aquarius_Scorpio_Taurus_Gemini_Pisces_Sagittarius_Leo_Libra_Pisces_Aries_Leo_Gemini_Taurus_Pisces_Gemini_Leo_Cancer_Virgo_Aries_Cancer_Leo_Scorpio_Aries_Pisces_Sagittarius_Capricorn_Capricorn_Leo_Aquarius_Leo_Virgo_Taurus_Taurus_Capricorn",
		"Pisces_Taurus_Sagittarius_Sagittarius_Gemini_Sagittarius_Virgo_Gemini_Scorpio_Sagittarius_Virgo_Gemini_Sagittarius_Leo_Libra_Aquarius_Libra_Aries_Libra_Aquarius_Pisces_Aquarius_Scorpio_Cancer_Leo_Sagittarius_Aquarius_Aquarius_Aries_Virgo_Aries_Capricorn_Sagittarius_Taurus",
		"Sagittarius_Scorpio_Gemini_Sagittarius_Sagittarius_Cancer_Taurus_Leo_Sagittarius_Sagittarius_Pisces_Taurus_Virgo_Cancer_Virgo_Virgo_Virgo_Aquarius_Gemini_Gemini_Sagittarius_Aquarius_Aries_Libra_Pisces_Aquarius_Pisces_Gemini_Taurus_Virgo_Cancer_Leo_Cancer_Leo",
		"Pisces_Cancer_Aries_Pisces_Taurus_Aquarius_Leo_Scorpio_Leo_Cancer_Leo_Leo_Taurus_Aquarius_Leo_Capricorn_Sagittarius_Capricorn_Aries_Gemini_Leo_Libra_Capricorn_Virgo_Virgo_Capricorn_Scorpio_Taurus_Cancer_Scorpio_Gemini_Sagittarius_Aries_Aquarius",
		"Aries_Gemini_Cancer_Pisces_Sagittarius_Cancer_Capricorn_Gemini_Scorpio_Scorpio_Scorpio_Pisces_Pisces_Scorpio_Scorpio_Aquarius_Virgo_Aquarius_Pisces_Scorpio_Aquarius_Scorpio_Taurus_Gemini_Scorpio_Capricorn_Aries_Taurus_Taurus_Sagittarius_Aries_Libra_Leo_Capricorn_Libra",
		"Sagittarius_Taurus_Capricorn_Taurus_Virgo_Aquarius_Libra_Virgo_Cancer_Scorpio_Cancer_Virgo_Taurus_Libra_Pisces_Aquarius_Leo_Gemini_Aquarius_Aries_Libra_Libra_Gemini_Taurus_Virgo_Virgo_Virgo_Leo_Sagittarius_Cancer_Taurus_Virgo_Virgo",
		"Aries_Taurus_Leo_Taurus_Gemini_Virgo_Sagittarius_Libra_Leo_Libra_Taurus_Pisces_Sagittarius_Scorpio_Libra_Libra_Pisces_Libra_Cancer_Virgo_Cancer_Libra_Virgo_Sagittarius_Capricorn_Cancer_Aquarius_Libra_Taurus_Virgo_Taurus_Virgo_Virgo_Pisces_Sagittarius",
		"Aquarius_Capricorn_Taurus_Scorpio_Taurus_Pisces_Aquarius_Virgo_Aquarius_Scorpio_Pisces_Pisces_Leo_Scorpio_Aquarius_Scorpio_Gemini_Gemini_Gemini_Leo_Gemini_Taurus_Capricorn_Aquarius_Virgo_Aries_Cancer_Virgo_Sagittarius_Libra_Sagittarius_Capricorn_Pisces_Scorpio_Capricorn_Cancer",
		"Aquarius_Virgo_Scorpio_Virgo_Aquarius_Scorpio_Sagittarius_Scorpio_Pisces_Leo_Leo_Libra_Capricorn_Scorpio_Scorpio_Virgo_Cancer_Leo_Aquarius_Aquarius_Aquarius_Sagittarius_Virgo_Taurus_Sagittarius_Aries_Aquarius_Leo_Libra_Pisces_Taurus_Leo_Pisces_Leo_Gemini_Pisces",
		"Pisces_Libra_Capricorn_Aquarius_Leo_Cancer_Leo_Scorpio_Virgo_Aquarius_Virgo_Gemini_Scorpio_Virgo_Cancer_Libra_Leo_Sagittarius_Scorpio_Cancer_Cancer_Pisces_Taurus_Libra_Aquarius_Capricorn_Libra_Aquarius_Cancer_Leo_Gemini_Aquarius_Aries_Pisces_Libra_Leo",
		"Libra_Virgo_Capricorn_Aries_Virgo_Taurus_Capricorn_Cancer_Taurus_Pisces_Virgo_Scorpio_Libra_Libra_Sagittarius_Sagittarius_Sagittarius_Pisces_Aries_Leo_Pisces_Libra_Scorpio_Sagittarius_Capricorn_Aquarius_Virgo_Libra_Pisces_Capricorn_Aries_Sagittarius_Libra_Cancer_Aries_Aries",
		"Aquarius_Capricorn_Pisces_Gemini_Aquarius_Aquarius_Taurus_Pisces_Virgo_Gemini_Leo_Cancer_Sagittarius_Capricorn_Leo_Pisces_Capricorn_Leo_Leo_Aquarius_Cancer_Leo_Gemini_Aquarius_Scorpio_Pisces_Taurus_Aquarius_Gemini_Leo_Cancer_Cancer_Virgo_Virgo_Cancer_Gemini_Cancer",
		"Virgo_Scorpio_Cancer_Aries_Virgo_Scorpio_Leo_Cancer_Gemini_Taurus_Virgo_Scorpio_Aries_Leo_Scorpio_Leo_Capricorn_Scorpio_Aquarius_Leo_Gemini_Taurus_Libra_Aries_Capricorn_Scorpio_Sagittarius_Virgo_Virgo_Virgo_Taurus_Virgo_Aquarius_Taurus_Taurus_Cancer_Aquarius",
		"Aquarius_Cancer_Leo_Sagittarius_Gemini_Pisces_Taurus_Virgo_Cancer_Taurus_Cancer_Virgo_Leo_Virgo_Aries_Libra_Taurus_Leo_Aquarius_Libra_Taurus_Libra_Scorpio_Sagittarius_Scorpio_Sagittarius_Cancer_Scorpio_Libra_Gemini_Gemini_Capricorn_Taurus_Aries_Aquarius_Virgo_Aries",
		"Aries_Sagittarius_Taurus_Taurus_Virgo_Scorpio_Cancer_Aries_Sagittarius_Capricorn_Pisces_Sagittarius_Virgo_Cancer_Gemini_Aries_Scorpio_Scorpio_Sagittarius_Capricorn_Sagittarius_Scorpio_Capricorn_Leo_Taurus_Aries_Pisces_Pisces_Pisces_Aquarius_Gemini_Leo_Scorpio_Cancer_Virgo_Virgo_Taurus",
		"Leo_Virgo_Aries_Leo_Gemini_Aries_Leo_Scorpio_Virgo_Taurus_Aquarius_Cancer_Capricorn_Libra_Taurus_Leo_Aries_Capricorn_Leo_Libra_Gemini_Virgo_Aquarius_Sagittarius_Aquarius_Aquarius_Virgo_Capricorn_Sagittarius_Cancer_Pisces_Virgo_Taurus_Gemini_Libra_Cancer_Gemini_Scorpio",
		"Capricorn_Libra_Sagittarius_Leo_Capricorn_Sagittarius_Pisces_Aquarius_Libra_Cancer_Scorpio_Leo_Aquarius_Leo_Capricorn_Aquarius_Sagittarius_Scorpio_Aries_Aquarius_Aries_Libra_Capricorn_Aquarius_Aquarius_Taurus_Libra_Aquarius_Leo_Taurus_Taurus_Taurus_Cancer_Pisces_Leo_Cancer_Virgo_Taurus",
		"Capricorn_Leo_Aries_Aries_Leo_Taurus_Capricorn_Virgo_Capricorn_Pisces_Pisces_Virgo_Cancer_Scorpio_Cancer_Virgo_Scorpio_Cancer_Leo_Leo_Virgo_Taurus_Leo_Cancer_Capricorn_Gemini_Libra_Libra_Virgo_Libra_Cancer_Scorpio_Cancer_Libra_Cancer_Gemini_Aries_Aries",
		"Pisces_Scorpio_Leo_Gemini_Taurus_Leo_Virgo_Libra_Capricorn_Taurus_Virgo_Taurus_Pisces_Libra_Aquarius_Pisces_Capricorn_Aquarius_Aquarius_Virgo_Virgo_Capricorn_Gemini_Capricorn_Scorpio_Aries_Scorpio_Pisces_Aquarius_Sagittarius_Virgo_Gemini_Capricorn_Scorpio_Libra_Scorpio_Aquarius_Capricorn_Sagittarius",
		"Pisces_Gemini_Pisces_Gemini_Libra_Aquarius_Gemini_Pisces_Aquarius_Scorpio_Aries_Gemini_Leo_Libra_Aries_Taurus_Leo_Taurus_Libra_Aquarius_Aries_Sagittarius_Scorpio_Aquarius_Virgo_Cancer_Virgo_Pisces_Capricorn_Cancer_Capricorn_Sagittarius_Taurus_Capricorn_Libra_Sagittarius_Pisces_Leo_Capricorn",
		"Scorpio_Aquarius_Capricorn_Pisces_Aries_Capricorn_Aquarius_Gemini_Pisces_Cancer_Capricorn_Sagittarius_Taurus_Virgo_Libra_Libra_Aquarius_Aries_Capricorn_Capricorn_Pisces_Gemini_Taurus_Leo_Aquarius_Aquarius_Leo_Taurus_Cancer_Scorpio_Leo_Aquarius_Aquarius_Gemini_Capricorn_Pisces_Aquarius_Aquarius",
		"Aquarius_Scorpio_Scorpio_Virgo_Gemini_Taurus_Aries_Libra_Leo_Sagittarius_Aries_Gemini_Libra_Aries_Gemini_Pisces_Aquarius_Taurus_Taurus_Taurus_Capricorn_Aquarius_Pisces_Pisces_Libra_Libra_Virgo_Gemini_Taurus_Pisces_Capricorn_Libra_Gemini_Virgo_Aries_Aquarius_Gemini_Pisces_Virgo_Sagittarius",
		"Pisces_Cancer_Libra_Sagittarius_Cancer_Scorpio_Aquarius_Virgo_Libra_Pisces_Cancer_Virgo_Leo_Sagittarius_Taurus_Leo_Aries_Pisces_Taurus_Scorpio_Pisces_Cancer_Sagittarius_Capricorn_Aries_Cancer_Leo_Virgo_Capricorn_Aquarius_Gemini_Virgo_Cancer_Pisces_Cancer_Leo_Scorpio_Pisces_Aries_Gemini",
		"Gemini_Cancer_Capricorn_Pisces_Leo_Capricorn_Sagittarius_Capricorn_Aries_Sagittarius_Pisces_Scorpio_Taurus_Leo_Pisces_Capricorn_Scorpio_Taurus_Cancer_Capricorn_Aquarius_Aquarius_Cancer_Aquarius_Aquarius_Gemini_Virgo_Libra_Cancer_Cancer_Leo_Scorpio_Libra_Capricorn_Taurus_Leo_Gemini_Aquarius_Libra_Aquarius",
		"Virgo_Capricorn_Taurus_Virgo_Sagittarius_Aries_Libra_Sagittarius_Aquarius_Libra_Gemini_Taurus_Capricorn_Virgo_Cancer_Virgo_Sagittarius_Libra_Taurus_Libra_Aries_Sagittarius_Taurus_Gemini_Aries_Leo_Capricorn_Libra_Capricorn_Virgo_Taurus_Cancer_Libra_Scorpio_Aries_Capricorn_Taurus_Aquarius_Virgo_Scorpio",
		"Libra_Pisces_Virgo_Pisces_Cancer_Aquarius_Leo_Scorpio_Capricorn_Aries_Scorpio_Aries_Aries_Taurus_Capricorn_Leo_Scorpio_Sagittarius_Taurus_Aquarius_Sagittarius_Scorpio_Capricorn_Scorpio_Taurus_Libra_Cancer_Scorpio_Virgo_Taurus_Gemini_Libra_Capricorn_Virgo_Virgo_Aquarius_Pisces_Aquarius_Leo_Gemini",
		"Aquarius_Leo_Libra_Aries_Capricorn_Taurus_Cancer_Gemini_Capricorn_Pisces_Gemini_Libra_Virgo_Taurus_Taurus_Aquarius_Virgo_Capricorn_Cancer_Virgo_Pisces_Libra_Scorpio_Aries_Capricorn_Virgo_Gemini_Sagittarius_Gemini_Sagittarius_Pisces_Scorpio_Scorpio_Aquarius_Taurus_Scorpio_Virgo_Aries_Taurus_Aquarius_Cancer",
		"Virgo_Taurus_Cancer_Gemini_Gemini_Virgo_Capricorn_Leo_Virgo_Gemini_Virgo_Cancer_Scorpio_Cancer_Aries_Scorpio_Gemini_Aquarius_Aquarius_Aries_Cancer_Aquarius_Pisces_Pisces_Taurus_Libra_Sagittarius_Pisces_Cancer_Capricorn_Cancer_Aries_Aquarius_Libra_Taurus_Libra_Libra_Aquarius_Pisces_Cancer_Virgo",
		"Capricorn_Gemini_Pisces_Pisces_Capricorn_Leo_Scorpio_Sagittarius_Sagittarius_Cancer_Leo_Capricorn_Gemini_Cancer_Libra_Cancer_Scorpio_Gemini_Virgo_Aries_Cancer_Taurus_Libra_Pisces_Pisces_Capricorn_Capricorn_Scorpio_Cancer_Cancer_Aquarius_Aries_Scorpio_Taurus_Sagittarius_Gemini_Libra_Capricorn_Leo_Pisces_Capricorn",
		"Pisces_Aries_Pisces_Cancer_Gemini_Libra_Gemini_Taurus_Cancer_Gemini_Leo_Gemini_Libra_Pisces_Gemini_Pisces_Leo_Sagittarius_Virgo_Sagittarius_Libra_Scorpio_Virgo_Leo_Gemini_Cancer_Capricorn_Libra_Virgo_Libra_Scorpio_Virgo_Cancer_Pisces_Cancer_Leo_Capricorn_Virgo_Aquarius_Cancer_Capricorn_Aries",
		"Aquarius_Sagittarius_Cancer_Leo_Scorpio_Cancer_Scorpio_Gemini_Aries_Capricorn_Virgo_Scorpio_Cancer_Aries_Capricorn_Sagittarius_Leo_Taurus_Capricorn_Pisces_Aries_Libra_Gemini_Sagittarius_Scorpio_Sagittarius_Scorpio_Gemini_Gemini_Leo_Capricorn_Gemini_Gemini_Aquarius_Gemini_Libra_Taurus_Scorpio_Capricorn_Gemini_Leo_Scorpio",
		"Taurus_Taurus_Aquarius_Sagittarius_Capricorn_Aries_Capricorn_Capricorn_Virgo_Pisces_Taurus_Capricorn_Cancer_Taurus_Cancer_Pisces_Taurus_Libra_Libra_Leo_Leo_Pisces_Leo_Aries_Aquarius_Aries_Capricorn_Aquarius_Aries_Libra_Cancer_Taurus_Taurus_Aries_Aquarius_Sagittarius_Taurus_Taurus_Capricorn_Aquarius_Cancer_Aries",
		"Aquarius_Libra_Virgo_Virgo_Gemini_Libra_Aries_Virgo_Aries_Pisces_Capricorn_Virgo_Cancer_Aquarius_Leo_Gemini_Virgo_Aquarius_Sagittarius_Aquarius_Sagittarius_Aquarius_Sagittarius_Pisces_Gemini_Virgo_Libra_Gemini_Sagittarius_Aquarius_Capricorn_Capricorn_Sagittarius_Pisces_Taurus_Virgo_Capricorn_Sagittarius_Libra_Capricorn_Sagittarius_Pisces_Aries",
		"Aries_Libra_Aries_Leo_Capricorn_Scorpio_Taurus_Aquarius_Aries_Gemini_Aries_Leo_Aries_Aries_Virgo_Leo_Sagittarius_Capricorn_Capricorn_Sagittarius_Sagittarius_Aries_Libra_Scorpio_Aries_Scorpio_Cancer_Capricorn_Libra_Taurus_Libra_Scorpio_Capricorn_Gemini_Virgo_Aries_Leo_Pisces_Aries_Pisces_Cancer",
		"Aries_Pisces_Taurus_Pisces_Sagittarius_Libra_Cancer_Libra_Sagittarius_Capricorn_Sagittarius_Capricorn_Leo_Gemini_Capricorn_Virgo_Cancer_Libra_Capricorn_Virgo_Capricorn_Cancer_Libra_Aquarius_Cancer_Scorpio_Virgo_Gemini_Leo_Virgo_Aquarius_Libra_Gemini_Capricorn_Sagittarius_Pisces_Leo_Gemini_Aquarius_Pisces_Pisces_Cancer_Leo",
		"Leo_Aries_Leo_Sagittarius_Sagittarius_Sagittarius_Aries_Pisces_Sagittarius_Leo_Scorpio_Aries_Gemini_Aries_Scorpio_Sagittarius_Libra_Cancer_Libra_Scorpio_Gemini_Leo_Cancer_Capricorn_Sagittarius_Gemini_Gemini_Libra_Virgo_Libra_Pisces_Cancer_Aquarius_Taurus_Libra_Aquarius_Gemini_Capricorn_Libra_Libra_Taurus_Gemini_Gemini",
		"Pisces_Leo_Aries_Scorpio_Taurus_Aquarius_Leo_Aquarius_Pisces_Libra_Cancer_Scorpio_Libra_Scorpio_Taurus_Capricorn_Cancer_Leo_Taurus_Leo_Pisces_Cancer_Aquarius_Gemini_Pisces_Libra_Leo_Sagittarius_Leo_Virgo_Cancer_Aries_Gemini_Virgo_Taurus_Taurus_Leo_Aries_Virgo_Taurus_Taurus_Scorpio_Capricorn_Taurus",
		"Virgo_Sagittarius_Scorpio_Virgo_Cancer_Sagittarius_Pisces_Cancer_Gemini_Aquarius_Sagittarius_Aquarius_Scorpio_Leo_Taurus_Aries_Aries_Aquarius_Gemini_Gemini_Virgo_Taurus_Capricorn_Capricorn_Leo_Pisces_Aquarius_Aquarius_Taurus_Aquarius_Leo_Gemini_Pisces_Aries_Aquarius_Libra_Pisces_Pisces_Libra_Aquarius_Virgo_Pisces_Gemini_Sagittarius",
		"Pisces_Libra_Scorpio_Scorpio_Libra_Sagittarius_Gemini_Gemini_Pisces_Capricorn_Cancer_Aquarius_Sagittarius_Sagittarius_Cancer_Leo_Leo_Leo_Capricorn_Leo_Leo_Libra_Gemini_Gemini_Taurus_Cancer_Gemini_Aries_Pisces_Libra_Taurus_Pisces_Libra_Gemini_Capricorn_Taurus_Aquarius_Pisces_Virgo_Libra_Cancer_Cancer_Aquarius_Taurus",
		"Aquarius_Aries_Aquarius_Sagittarius_Pisces_Pisces_Virgo_Taurus_Pisces_Gemini_Capricorn_Scorpio_Cancer_Leo_Pisces_Taurus_Virgo_Pisces_Capricorn_Pisces_Scorpio_Cancer_Taurus_Scorpio_Libra_Aries_Libra_Aquarius_Cancer_Scorpio_Gemini_Leo_Leo_Capricorn_Leo_Virgo_Virgo_Taurus_Libra_Sagittarius_Virgo_Cancer_Aries_Scorpio_Aries",
		"Aries_Aquarius_Pisces_Sagittarius_Leo_Libra_Capricorn_Capricorn_Cancer_Scorpio_Taurus_Libra_Scorpio_Virgo_Cancer_Capricorn_Virgo_Libra_Leo_Gemini_Virgo_Cancer_Virgo_Aries_Libra_Aquarius_Taurus_Libra_Sagittarius_Aries_Gemini_Pisces_Libra_Virgo_Scorpio_Sagittarius_Aquarius_Aquarius_Sagittarius_Gemini_Capricorn_Leo_Capricorn_Taurus_Aquarius",
		"Scorpio_Sagittarius_Aquarius_Aries_Taurus_Gemini_Aquarius_Scorpio_Capricorn_Sagittarius_Aquarius_Libra_Pisces_Capricorn_Scorpio_Gemini_Libra_Pisces_Pisces_Libra_Virgo_Gemini_Leo_Cancer_Cancer_Capricorn_Leo_Scorpio_Leo_Aries_Scorpio_Gemini_Pisces_Leo_Taurus_Virgo_Gemini_Pisces_Pisces_Aries_Scorpio_Libra_Aquarius_Scorpio_Cancer",
		"Aquarius_Aquarius_Cancer_Leo_Sagittarius_Taurus_Capricorn_Taurus_Pisces_Pisces_Leo_Scorpio_Pisces_Cancer_Capricorn_Aquarius_Gemini_Pisces_Scorpio_Scorpio_Leo_Pisces_Libra_Scorpio_Aquarius_Cancer_Aries_Sagittarius_Taurus_Pisces_Cancer_Libra_Libra_Aquarius_Gemini_Sagittarius_Cancer_Aquarius_Taurus_Aquarius_Libra_Aries_Taurus_Pisces_Capricorn_Sagittarius",
		"Aquarius_Pisces_Pisces_Aries_Libra_Pisces_Capricorn_Capricorn_Scorpio_Cancer_Taurus_Sagittarius_Virgo_Aquarius_Pisces_Cancer_Gemini_Capricorn_Capricorn_Virgo_Virgo_Sagittarius_Taurus_Aquarius_Taurus_Pisces_Capricorn_Libra_Libra_Aquarius_Aries_Taurus_Taurus_Leo_Gemini_Capricorn_Aries_Scorpio_Capricorn_Aquarius_Cancer_Cancer_Libra_Capricorn_Scorpio_Libra",
		"Libra_Aries_Sagittarius_Aquarius_Taurus_Taurus_Virgo_Leo_Gemini_Gemini_Pisces_Gemini_Taurus_Leo_Scorpio_Scorpio_Virgo_Leo_Sagittarius_Aries_Aries_Aries_Capricorn_Aries_Leo_Taurus_Aries_Leo_Libra_Pisces_Leo_Sagittarius_Gemini_Pisces_Leo_Virgo_Virgo_Aquarius_Capricorn_Aquarius_Aquarius_Virgo_Cancer_Capricorn_Taurus",
		"Capricorn_Cancer_Scorpio_Capricorn_Cancer_Pisces_Taurus_Taurus_Leo_Sagittarius_Sagittarius_Taurus_Libra_Gemini_Capricorn_Libra_Taurus_Taurus_Pisces_Aquarius_Taurus_Capricorn_Taurus_Virgo_Capricorn_Aries_Aquarius_Leo_Sagittarius_Capricorn_Capricorn_Aries_Libra_Virgo_Cancer_Aries_Libra_Libra_Scorpio_Virgo_Capricorn_Scorpio_Gemini_Virgo_Scorpio_Libra",
		"Aquarius_Scorpio_Cancer_Libra_Pisces_Sagittarius_Leo_Capricorn_Gemini_Virgo_Pisces_Pisces_Taurus_Aries_Aries_Scorpio_Gemini_Aries_Virgo_Leo_Virgo_Libra_Scorpio_Capricorn_Leo_Cancer_Capricorn_Leo_Capricorn_Sagittarius_Capricorn_Libra_Leo_Libra_Cancer_Capricorn_Sagittarius_Sagittarius_Sagittarius_Virgo_Gemini_Leo_Capricorn_Leo_Pisces_Pisces_Virgo",
		"Leo_Aries_Gemini_Gemini_Taurus_Pisces_Gemini_Aries_Aries_Libra_Aquarius_Aquarius_Capricorn_Leo_Virgo_Leo_Aries_Aries_Scorpio_Cancer_Sagittarius_Taurus_Sagittarius_Scorpio_Capricorn_Scorpio_Aries_Aries_Sagittarius_Scorpio_Gemini_Gemini_Gemini_Sagittarius_Leo_Libra_Taurus_Cancer_Libra_Cancer_Cancer_Scorpio_Scorpio_Gemini_Gemini_Capricorn",
		"Capricorn_Sagittarius_Sagittarius_Aries_Aries_Scorpio_Scorpio_Leo_Taurus_Gemini_Libra_Sagittarius_Libra_Pisces_Aquarius_Pisces_Scorpio_Taurus_Virgo_Gemini_Aquarius_Leo_Libra_Capricorn_Aquarius_Taurus_Aries_Sagittarius_Sagittarius_Aquarius_Gemini_Capricorn_Cancer_Libra_Aries_Aquarius_Cancer_Aquarius_Capricorn_Aries_Scorpio_Aquarius_Capricorn_Cancer_Gemini_Pisces_Aries",
		"Gemini_Taurus_Scorpio_Aquarius_Aries_Capricorn_Pisces_Aquarius_Virgo_Aries_Gemini_Pisces_Sagittarius_Scorpio_Aquarius_Aquarius_Sagittarius_Aries_Leo_Cancer_Virgo_Pisces_Gemini_Libra_Aquarius_Sagittarius_Scorpio_Sagittarius_Aquarius_Capricorn_Libra_Virgo_Aquarius_Aquarius_Sagittarius_Libra_Aquarius_Gemini_Aries_Cancer_Leo_Leo_Sagittarius_Sagittarius_Taurus_Gemini_Leo",
		"Aquarius_Virgo_Scorpio_Pisces_Gemini_Gemini_Gemini_Leo_Cancer_Capricorn_Scorpio_Scorpio_Libra_Pisces_Capricorn_Virgo_Leo_Aquarius_Scorpio_Gemini_Leo_Scorpio_Virgo_Scorpio_Sagittarius_Aquarius_Sagittarius_Gemini_Taurus_Aries_Aries_Leo_Aquarius_Virgo_Libra_Capricorn_Sagittarius_Leo_Gemini_Capricorn_Sagittarius_Aquarius_Sagittarius_Aquarius_Taurus_Aries_Pisces_Taurus",
		"Aquarius_Aries_Libra_Leo_Libra_Virgo_Libra_Taurus_Virgo_Capricorn_Scorpio_Sagittarius_Aries_Capricorn_Capricorn_Aries_Sagittarius_Leo_Sagittarius_Leo_Virgo_Leo_Taurus_Leo_Pisces_Cancer_Gemini_Capricorn_Aquarius_Gemini_Taurus_Leo_Sagittarius_Aquarius_Gemini_Aries_Sagittarius_Libra_Scorpio_Leo_Leo_Scorpio_Virgo_Aquarius_Aquarius_Capricorn_Capricorn_Sagittarius",
		"Leo_Virgo_Taurus_Sagittarius_Sagittarius_Taurus_Cancer_Aries_Leo_Virgo_Libra_Capricorn_Libra_Aries_Aquarius_Cancer_Aries_Capricorn_Gemini_Pisces_Capricorn_Libra_Capricorn_Libra_Cancer_Scorpio_Sagittarius_Pisces_Leo_Virgo_Pisces_Libra_Gemini_Scorpio_Leo_Sagittarius_Capricorn_Sagittarius_Sagittarius_Gemini_Virgo_Sagittarius_Pisces_Gemini_Capricorn_Capricorn_Gemini_Sagittarius",
		"Aquarius_Sagittarius_Virgo_Scorpio_Aquarius_Pisces_Gemini_Aquarius_Gemini_Sagittarius_Aries_Aries_Cancer_Cancer_Aries_Scorpio_Scorpio_Sagittarius_Scorpio_Taurus_Cancer_Pisces_Leo_Leo_Aries_Capricorn_Sagittarius_Leo_Aquarius_Cancer_Capricorn_Gemini_Taurus_Cancer_Pisces_Aquarius_Leo_Aquarius_Virgo_Cancer_Taurus_Aries_Libra_Aquarius_Leo_Pisces_Scorpio_Libra_Aries",
		"Pisces_Capricorn_Taurus_Sagittarius_Aquarius_Scorpio_Capricorn_Libra_Aries_Taurus_Taurus_Aries_Gemini_Scorpio_Gemini_Taurus_Capricorn_Gemini_Cancer_Gemini_Leo_Virgo_Aries_Pisces_Aries_Libra_Capricorn_Gemini_Scorpio_Virgo_Aquarius_Sagittarius_Leo_Leo_Scorpio_Virgo_Pisces_Pisces_Cancer_Virgo_Aries_Cancer_Scorpio_Capricorn_Gemini_Aquarius_Pisces_Leo_Capricorn",
		"Aries_Aquarius_Leo_Leo_Virgo_Scorpio_Sagittarius_Capricorn_Libra_Aquarius_Pisces_Gemini_Aries_Pisces_Capricorn_Taurus_Aries_Leo_Aries_Gemini_Aquarius_Sagittarius_Cancer_Libra_Libra_Virgo_Capricorn_Virgo_Leo_Aries_Virgo_Cancer_Aquarius_Aries_Pisces_Virgo_Taurus_Pisces_Libra_Virgo_Gemini_Leo_Taurus_Virgo_Gemini_Virgo_Leo_Sagittarius_Gemini",
		"Capricorn_Gemini_Taurus_Libra_Cancer_Capricorn_Aquarius_Sagittarius_Pisces_Aquarius_Leo_Taurus_Aries_Leo_Leo_Sagittarius_Taurus_Gemini_Libra_Capricorn_Sagittarius_Pisces_Aquarius_Taurus_Sagittarius_Cancer_Aries_Aquarius_Capricorn_Taurus_Cancer_Aries_Virgo_Taurus_Pisces_Libra_Sagittarius_Virgo_Leo_Virgo_Aquarius_Libra_Scorpio_Libra_Capricorn_Aquarius_Taurus_Aries_Virgo",
	}
	b := []string{
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
	for c := 0; c < len(a); c++ {
		d := strings.Split(a[c], "_")
		e := big.NewInt(0)
		for f := 0; f < len(d); f++ {
			for g := 0; g < 12; g++ {
				if d[f] == b[g] {
					e.Mul(e, big.NewInt(12))
					e.Add(e, big.NewInt(int64(g)))
				}
			}
		}
		h := big.NewInt(0)
		for i := big.NewInt(0); i.Cmp(e) < 0; i = i.Add(i, big.NewInt(1)) {
			h = h.Add(h, big.NewInt(5))
		}
		h = h.Mod(h, big.NewInt(256))
		fmt.Printf("%c", h.Int64())
	}
	fmt.Println()
}

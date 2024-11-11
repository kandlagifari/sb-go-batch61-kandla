package webserver

import (
	"fmt"
	"math"
	"net/http"
)

func luasAlas(r float64) float64 {
	return math.Pi * r * r
}

func kelilingAlas(r float64) float64 {
	return 2 * math.Pi * r
}

func volumeTabung(r, t float64) float64 {
	return math.Pi * r * r * t
}

func Handler(w http.ResponseWriter, r *http.Request) {
	jariJari := 7.0
	tinggi := 10.0

	luas := luasAlas(jariJari)
	keliling := kelilingAlas(jariJari)
	volume := volumeTabung(jariJari, tinggi)

	fmt.Fprintf(w, "jariJari: %.0f, tinggi: %.0f, volume: %.2f, luas alas: %.2f, keliling alas: %.2f",
		jariJari, tinggi, volume, luas, keliling)
}

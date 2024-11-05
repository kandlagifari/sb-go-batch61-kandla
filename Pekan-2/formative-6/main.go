package main

import (
	"fmt"
	"math"
)

// soal 1
func updateLingkaran(luas *float64, keliling *float64, jariJari float64) {
	*luas = math.Pi * jariJari * jariJari

	*keliling = 2 * math.Pi * jariJari
}

// soal 2
func introduce(sentence *string, name, gender, job, age string) {
	var prefix string
	if gender == "laki-laki" {
		prefix = "Pak"
	} else if gender == "perempuan" {
		prefix = "Bu"
	}

	*sentence = fmt.Sprintf("%s %s adalah seorang %s yang berusia %s tahun", prefix, name, job, age)
}

// soal 3
func tambahBuah(buah *[]string, namaBuah string) {
	*buah = append(*buah, namaBuah)
}

// soal 4
func tambahDataFilm(title, duration, genre, year string, dataFilm *[]map[string]string) {
	film := map[string]string{
		"title":    title,
		"duration": duration,
		"genre":    genre,
		"year":     year,
	}
	*dataFilm = append(*dataFilm, film)
}

func main() {
	// soal 1
	fmt.Println("--------------------- Soal 1 ---------------------")
	var luasLingkaran float64
	var kelilingLingkaran float64

	updateLingkaran(&luasLingkaran, &kelilingLingkaran, 10)

	fmt.Printf("Luas Lingkaran: %.2f\n", luasLingkaran)
	fmt.Printf("Keliling Lingkaran: %.2f\n", kelilingLingkaran)

	fmt.Println("--------------------------------------------------")

	// soal 2
	fmt.Println("--------------------- Soal 2 ---------------------")
	var sentence string

	introduce(&sentence, "John", "laki-laki", "penulis", "30")
	fmt.Println(sentence)

	introduce(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println(sentence)

	fmt.Println("--------------------------------------------------")

	// soal 3
	fmt.Println("--------------------- Soal 3 ---------------------")
	var buah = []string{}

	tambahBuah(&buah, "Jeruk")
	tambahBuah(&buah, "Semangka")
	tambahBuah(&buah, "Mangga")
	tambahBuah(&buah, "Strawberry")
	tambahBuah(&buah, "Durian")
	tambahBuah(&buah, "Manggis")
	tambahBuah(&buah, "Alpukat")

	for i, namaBuah := range buah {
		fmt.Printf("%d. %s\n", i+1, namaBuah)
	}

	fmt.Println("--------------------------------------------------")

	// soal 4
	fmt.Println("--------------------- Soal 4 ---------------------")
	var dataFilm = []map[string]string{}

	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	for i, film := range dataFilm {
		fmt.Printf("%d. ", i+1)
		first := true
		for key, value := range film {
			if first {
				fmt.Printf("%s : %s\n", key, value)
				first = false
			} else {
				fmt.Printf("   %s : %s\n", key, value)
			}
		}
		fmt.Println()
	}

	fmt.Println("--------------------------------------------------")
}

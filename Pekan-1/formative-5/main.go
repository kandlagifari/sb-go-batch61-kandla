package main

import (
	"fmt"
	"strings"
)

// soal 1
func luasPersegiPanjang(panjang, lebar int) int {
	return panjang * lebar
}

func kelilingPersegiPanjang(panjang, lebar int) int {
	return 2 * (panjang + lebar)
}

func volumeBalok(panjang, lebar, tinggi int) int {
	return panjang * lebar * tinggi
}

// soal 2
func introduce(name, gender, job, age string) (result string) {
	var prefix string
	if gender == "laki-laki" {
		prefix = "Pak"
	} else if gender == "perempuan" {
		prefix = "Bu"
	}

	result = fmt.Sprintf("%s %s adalah seorang %s yang berusia %s tahun", prefix, name, job, age)

	return
}

// soal 3
func buahFavorit(name string, fruits ...string) (result string) {
	fruitsList := "\"" + strings.Join(fruits, "\", \"") + "\""

	result = fmt.Sprintf("halo nama saya %s dan buah favorit saya adalah %s", name, fruitsList)

	return
}

// soal 4 (closure function di dalam main function)

func main() {
	// soal 1
	fmt.Println("--------------------- Soal 1 ---------------------")
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)

	fmt.Println("--------------------------------------------------")

	// soal 2
	fmt.Println("--------------------- Soal 2 ---------------------")
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john)

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah)

	fmt.Println("--------------------------------------------------")

	// soal 3
	fmt.Println("--------------------- Soal 3 ---------------------")
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}

	var buahFavoritJohn = buahFavorit("John", buah...)

	fmt.Println(buahFavoritJohn)

	fmt.Println("--------------------------------------------------")

	// soal 4
	fmt.Println("--------------------- Soal 4 ---------------------")
	var dataFilm = []map[string]string{}

	tambahDataFilm := func(title, jam, genre, tahun string) {
		film := map[string]string{
			"title": title,
			"jam":   jam,
			"genre": genre,
			"tahun": tahun,
		}

		dataFilm = append(dataFilm, film)
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}

	fmt.Println("--------------------------------------------------")
}

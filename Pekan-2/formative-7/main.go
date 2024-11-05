package main

import (
	"fmt"
)

// soal 1
type Buah struct {
	Nama       string
	Warna      string
	AdaBijinya bool
	Harga      int
}

func (b Buah) Bijinya() string {
	if b.AdaBijinya {
		return "Ada"
	}
	return "Tidak"
}

// soal 2
type segitiga struct {
	alas, tinggi int
}

func (s segitiga) Luas() int {
	return (s.alas * s.tinggi) / 2
}

type persegi struct {
	sisi int
}

func (p persegi) Luas() int {
	return p.sisi * p.sisi
}

type persegiPanjang struct {
	panjang, lebar int
}

func (pp persegiPanjang) Luas() int {
	return pp.panjang * pp.lebar
}

// soal 3
type phone struct {
	name, brand string
	year        int
	colors      []string
}

func (p *phone) AddColor(color string) {
	p.colors = append(p.colors, color)
}

// soal 4
type movie struct {
	title    string
	genre    string
	duration int
	year     int
}

func tambahDataFilm(title string, duration int, genre string, year int, dataFilm *[]movie) {
	newMovie := movie{
		title:    title,
		genre:    genre,
		duration: duration,
		year:     year,
	}
	*dataFilm = append(*dataFilm, newMovie)
}

func tampilkanDataFilm(dataFilm []movie) {
	for i, film := range dataFilm {
		fmt.Printf("%d. title : %s\n", i+1, film.title)
		fmt.Printf("   duration : %d jam\n", film.duration/60)
		fmt.Printf("   genre : %s\n", film.genre)
		fmt.Printf("   year : %d\n\n", film.year)
	}
}

func main() {
	// soal 1
	fmt.Println("--------------------- Soal 1 ---------------------")
	nanas := Buah{
		Nama:       "Nanas",
		Warna:      "Kuning",
		AdaBijinya: false,
		Harga:      9000,
	}

	jeruk := Buah{
		Nama:       "Jeruk",
		Warna:      "Oranye",
		AdaBijinya: true,
		Harga:      8000,
	}

	semangka := Buah{
		Nama:       "Semangka",
		Warna:      "Hijau & Merah",
		AdaBijinya: true,
		Harga:      10000,
	}

	pisang := Buah{
		Nama:       "Pisang",
		Warna:      "Kuning",
		AdaBijinya: false,
		Harga:      5000,
	}

	fmt.Println("Data Buah:")
	fmt.Printf("Nama: %s, Warna: %s, Ada Bijinya: %s, Harga: %d\n", nanas.Nama, nanas.Warna, nanas.Bijinya(), nanas.Harga)
	fmt.Printf("Nama: %s, Warna: %s, Ada Bijinya: %s, Harga: %d\n", jeruk.Nama, jeruk.Warna, jeruk.Bijinya(), jeruk.Harga)
	fmt.Printf("Nama: %s, Warna: %s, Ada Bijinya: %s, Harga: %d\n", semangka.Nama, semangka.Warna, semangka.Bijinya(), semangka.Harga)
	fmt.Printf("Nama: %s, Warna: %s, Ada Bijinya: %s, Harga: %d\n", pisang.Nama, pisang.Warna, pisang.Bijinya(), pisang.Harga)

	fmt.Println("--------------------------------------------------")

	// soal 2
	fmt.Println("--------------------- Soal 2 ---------------------")
	segitiga1 := segitiga{alas: 10, tinggi: 5}
	persegi1 := persegi{sisi: 4}
	persegiPanjang1 := persegiPanjang{panjang: 8, lebar: 3}

	fmt.Println("Luas Segitiga:", segitiga1.Luas())
	fmt.Println("Luas Persegi:", persegi1.Luas())
	fmt.Println("Luas Persegi Panjang:", persegiPanjang1.Luas())

	fmt.Println("--------------------------------------------------")

	// soal 3
	fmt.Println("--------------------- Soal 3 ---------------------")
	myPhone := phone{
		name:   "Galaxy A6",
		brand:  "Samsung",
		year:   2018,
		colors: []string{"Gray", "White"},
	}

	myPhone.AddColor("Black")
	myPhone.AddColor("Blue")

	fmt.Println("Phone:", myPhone.name)
	fmt.Println("Brand:", myPhone.brand)
	fmt.Println("Year:", myPhone.year)
	fmt.Println("Colors:", myPhone.colors)

	fmt.Println("--------------------------------------------------")

	// soal 4
	fmt.Println("--------------------- Soal 4 ---------------------")
	var dataFilm = []movie{}

	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	tampilkanDataFilm(dataFilm)

	fmt.Println("--------------------------------------------------")
}

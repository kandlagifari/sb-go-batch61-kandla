package main

import (
	"fmt"
	"math"
	"strings"
)

// soal 1
type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

func (s segitigaSamaSisi) luas() int {
	return (s.alas * s.tinggi) / 2
}

func (s segitigaSamaSisi) keliling() int {
	return s.alas * 3
}

func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
	return 2 * (p.panjang + p.lebar)
}

func (t tabung) volume() float64 {
	return math.Pi * t.jariJari * t.jariJari * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
	return 2 * math.Pi * t.jariJari * (t.jariJari + t.tinggi)
}

func (b balok) volume() float64 {
	return float64(b.panjang * b.lebar * b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	return 2 * float64((b.panjang*b.lebar)+(b.panjang*b.tinggi)+(b.lebar*b.tinggi))
}

func tampilkanBangunDatar(b hitungBangunDatar) {
	fmt.Printf("Luas: %d\n", b.luas())
	fmt.Printf("Keliling: %d\n\n", b.keliling())
}

func tampilkanBangunRuang(b hitungBangunRuang) {
	fmt.Printf("Volume: %.2f\n", b.volume())
	fmt.Printf("Luas Permukaan: %.2f\n\n", b.luasPermukaan())
}

// soal 2
type phone struct {
	name, brand string
	year        int
	colors      []string
}

type phoneInfo interface {
	displayInfo() string
}

func (p phone) displayInfo() string {
	colors := strings.Join(p.colors, ", ")
	return fmt.Sprintf(
		"name : %s\nbrand : %s\nyear : %d\ncolors : %s",
		p.name, p.brand, p.year, colors,
	)
}

// soal 3
func luasPersegi(sisi int, denganPesan bool) interface{} {
	if sisi == 0 {
		if denganPesan {
			return "Maaf anda belum menginput sisi dari persegi"
		}
		return nil
	}

	luas := sisi * sisi

	if denganPesan {
		return fmt.Sprintf("luas persegi dengan sisi %d cm adalah %d cm", sisi, luas)
	}

	return luas
}

// soal 4 (<nil>)

func main() {
	// soal 1
	fmt.Println("--------------------- Soal 1 ---------------------")
	segitiga1 := segitigaSamaSisi{alas: 10, tinggi: 8}
	persegi1 := persegiPanjang{panjang: 7, lebar: 4}
	tabung1 := tabung{jariJari: 3, tinggi: 10}
	balok1 := balok{panjang: 5, lebar: 3, tinggi: 4}

	fmt.Println("# Perhitungan Bangun Datar:")
	fmt.Println("Segitiga Sama Sisi:")
	tampilkanBangunDatar(segitiga1)
	fmt.Println("Persegi Panjang:")
	tampilkanBangunDatar(persegi1)

	fmt.Println("# Perhitungan Bangun Ruang:")
	fmt.Println("Tabung:")
	tampilkanBangunRuang(tabung1)
	fmt.Println("Balok:")
	tampilkanBangunRuang(balok1)

	fmt.Println("--------------------------------------------------")

	// soal 2
	fmt.Println("--------------------- Soal 2 ---------------------")
	phone1 := phone{
		name:   "Samsung Galaxy Note 20",
		brand:  "Samsung Galaxy Note 20",
		year:   2020,
		colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}

	var info phoneInfo = phone1
	fmt.Println(info.displayInfo())

	fmt.Println("--------------------------------------------------")

	// soal 3
	fmt.Println("--------------------- Soal 3 ---------------------")
	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))

	fmt.Println("--------------------------------------------------")

	// soal 4
	fmt.Println("--------------------- Soal 4 ---------------------")
	var prefix interface{} = "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	prefixStr := prefix.(string)

	angkaPertama := kumpulanAngkaPertama.([]int)
	angkaKedua := kumpulanAngkaKedua.([]int)

	semuaAngka := append(angkaPertama, angkaKedua...)

	total := 0
	angkaStr := []string{}

	for _, angka := range semuaAngka {
		total += angka
		angkaStr = append(angkaStr, fmt.Sprintf("%d", angka))
	}

	output := fmt.Sprintf("%s%s = %d", prefixStr, strings.Join(angkaStr, " + "), total)

	fmt.Println(output)

	fmt.Println("--------------------------------------------------")
}

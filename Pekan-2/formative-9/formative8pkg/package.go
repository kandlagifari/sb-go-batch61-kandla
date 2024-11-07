package formative8pkg

import (
	"fmt"
	"math"
	"strings"
)

// soal 1
type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

type Tabung struct {
	JariJari, Tinggi float64
}

type Balok struct {
	Panjang, Lebar, Tinggi int
}

type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

func (s SegitigaSamaSisi) Luas() int {
	return (s.Alas * s.Tinggi) / 2
}

func (s SegitigaSamaSisi) Keliling() int {
	return s.Alas * 3
}

func (p PersegiPanjang) Luas() int {
	return p.Panjang * p.Lebar
}

func (p PersegiPanjang) Keliling() int {
	return 2 * (p.Panjang + p.Lebar)
}

func (t Tabung) Volume() float64 {
	return math.Pi * t.JariJari * t.JariJari * t.Tinggi
}

func (t Tabung) LuasPermukaan() float64 {
	return 2 * math.Pi * t.JariJari * (t.JariJari + t.Tinggi)
}

func (b Balok) Volume() float64 {
	return float64(b.Panjang * b.Lebar * b.Tinggi)
}

func (b Balok) LuasPermukaan() float64 {
	return 2 * float64((b.Panjang*b.Lebar)+(b.Panjang*b.Tinggi)+(b.Lebar*b.Tinggi))
}

func TampilkanBangunDatar(b HitungBangunDatar) {
	fmt.Printf("Luas: %d\n", b.Luas())
	fmt.Printf("Keliling: %d\n\n", b.Keliling())
}

func TampilkanBangunRuang(b HitungBangunRuang) {
	fmt.Printf("Volume: %.2f\n", b.Volume())
	fmt.Printf("Luas Permukaan: %.2f\n\n", b.LuasPermukaan())
}

// soal 2
type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

type PhoneInfo interface {
	DisplayInfo() string
}

func (p Phone) DisplayInfo() string {
	colors := strings.Join(p.Colors, ", ")
	return fmt.Sprintf(
		"name : %s\nbrand : %s\nyear : %d\ncolors : %s",
		p.Name, p.Brand, p.Year, colors,
	)
}

// soal 3
func LuasPersegi(sisi int, denganPesan bool) interface{} {
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

// soal 4
func HasilPenjumlahan(prefix interface{}, kumpulanAngkaPertama interface{}, kumpulanAngkaKedua interface{}) string {
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

	return fmt.Sprintf("%s%s = %d", prefixStr, strings.Join(angkaStr, " + "), total)
}

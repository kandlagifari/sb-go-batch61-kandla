package main

import (
	"fmt"

	"github.com/kandlagifari/sb-go-batch61-kandla/Pekan-2/formative-9/formative8pkg"
)

func main() {
	// soal 1
	fmt.Println("--------------------- Soal 1 ---------------------")
	segitiga1 := formative8pkg.SegitigaSamaSisi{Alas: 10, Tinggi: 8}
	persegi1 := formative8pkg.PersegiPanjang{Panjang: 7, Lebar: 4}
	tabung1 := formative8pkg.Tabung{JariJari: 3, Tinggi: 10}
	balok1 := formative8pkg.Balok{Panjang: 5, Lebar: 3, Tinggi: 4}

	fmt.Println("# Perhitungan Bangun Datar:")
	fmt.Println("Segitiga Sama Sisi:")
	formative8pkg.TampilkanBangunDatar(segitiga1)
	fmt.Println("Persegi Panjang:")
	formative8pkg.TampilkanBangunDatar(persegi1)

	fmt.Println("# Perhitungan Bangun Ruang:")
	fmt.Println("Tabung:")
	formative8pkg.TampilkanBangunRuang(tabung1)
	fmt.Println("Balok:")
	formative8pkg.TampilkanBangunRuang(balok1)

	fmt.Println("--------------------------------------------------")

	// soal 2
	fmt.Println("--------------------- Soal 2 ---------------------")
	phone1 := formative8pkg.Phone{
		Name:   "Samsung Galaxy Note 20",
		Brand:  "Samsung Galaxy Note 20",
		Year:   2020,
		Colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}

	var info formative8pkg.PhoneInfo = phone1
	fmt.Println(info.DisplayInfo())

	fmt.Println("--------------------------------------------------")

	// soal 3
	fmt.Println("--------------------- Soal 3 ---------------------")
	fmt.Println(formative8pkg.LuasPersegi(4, true))
	fmt.Println(formative8pkg.LuasPersegi(8, false))
	fmt.Println(formative8pkg.LuasPersegi(0, true))
	fmt.Println(formative8pkg.LuasPersegi(0, false))

	fmt.Println("--------------------------------------------------")

	// soal 4
	fmt.Println("--------------------- Soal 4 ---------------------")
	var prefix interface{} = "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	output := formative8pkg.HasilPenjumlahan(prefix, kumpulanAngkaPertama, kumpulanAngkaKedua)
	fmt.Println(output)

	fmt.Println("--------------------------------------------------")
}

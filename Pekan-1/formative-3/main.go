package main

import (
	"fmt"
	"strconv"
)

func main() {
	// soal 1
	fmt.Println("--------------------- Soal 1 ---------------------")

	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"
	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	intPanjangPersegiPanjang, _ := strconv.Atoi(panjangPersegiPanjang)
	intLebarPersegiPanjang, _ := strconv.Atoi(lebarPersegiPanjang)
	intAlasSegitiga, _ := strconv.Atoi(alasSegitiga)
	intTinggiSegitiga, _ := strconv.Atoi(tinggiSegitiga)

	var luasPersegiPanjang int = intPanjangPersegiPanjang * intLebarPersegiPanjang
	var kelilingPersegiPanjang int = 2 * (intPanjangPersegiPanjang + intLebarPersegiPanjang)
	var luasSegitiga int = (intAlasSegitiga * intTinggiSegitiga) / 2

	fmt.Println("Luas Persegi Panjang:", luasPersegiPanjang)
	fmt.Println("Keliling Persegi Panjang:", kelilingPersegiPanjang)
	fmt.Println("Luas Segitiga:", luasSegitiga)

	fmt.Println("--------------------------------------------------")

	// soal 2
	fmt.Println("--------------------- Soal 2 ---------------------")
	var nilaiJohn = 80
	var nilaiDoe = 50

	var indeksNilaiJohn string
	if nilaiJohn >= 80 {
		indeksNilaiJohn = "A"
	} else if nilaiJohn >= 70 && nilaiJohn < 80 {
		indeksNilaiJohn = "B"
	} else if nilaiJohn >= 60 && nilaiJohn < 70 {
		indeksNilaiJohn = "C"
	} else if nilaiJohn >= 50 && nilaiJohn < 60 {
		indeksNilaiJohn = "D"
	} else {
		indeksNilaiJohn = "E"
	}

	var indeksNilaiDoe string
	if nilaiDoe >= 80 {
		indeksNilaiDoe = "A"
	} else if nilaiDoe >= 70 && nilaiDoe < 80 {
		indeksNilaiDoe = "B"
	} else if nilaiDoe >= 60 && nilaiDoe < 70 {
		indeksNilaiDoe = "C"
	} else if nilaiDoe >= 50 && nilaiDoe < 60 {
		indeksNilaiDoe = "D"
	} else {
		indeksNilaiDoe = "E"
	}

	fmt.Println("Indeks Nilai John:", indeksNilaiJohn)
	fmt.Println("Indeks Nilai Doe:", indeksNilaiDoe)

	fmt.Println("--------------------------------------------------")

	// soal 3
	fmt.Println("--------------------- Soal 3 ---------------------")
	var tanggal = 16
	var bulan = 8
	var tahun = 2000

	var namaBulan string

	switch bulan {
	case 1:
		namaBulan = "Januari"
	case 2:
		namaBulan = "Februari"
	case 3:
		namaBulan = "Maret"
	case 4:
		namaBulan = "April"
	case 5:
		namaBulan = "Mei"
	case 6:
		namaBulan = "Juni"
	case 7:
		namaBulan = "Juli"
	case 8:
		namaBulan = "Agustus"
	case 9:
		namaBulan = "September"
	case 10:
		namaBulan = "Oktober"
	case 11:
		namaBulan = "November"
	case 12:
		namaBulan = "Desember"
	default:
		namaBulan = "Angka pada bulan harus dalam range 1-12"
	}

	fmt.Printf("%d %s %d\n", tanggal, namaBulan, tahun)

	fmt.Println("--------------------------------------------------")

	// soal 4
	fmt.Println("--------------------- Soal 4 ---------------------")
	var tahunKelahiran = 2000
	var generasi string

	if tahunKelahiran >= 1944 && tahunKelahiran <= 1964 {
		generasi = "Baby Boomer"
	} else if tahunKelahiran >= 1965 && tahunKelahiran <= 1979 {
		generasi = "Generasi X"
	} else if tahunKelahiran >= 1980 && tahunKelahiran <= 1994 {
		generasi = "Generasi Y (Millenials)"
	} else if tahunKelahiran >= 1995 && tahunKelahiran <= 2015 {
		generasi = "Generasi Z"
	} else {
		generasi = "Generasi tidak dikenali (YNTKTS)"
	}

	fmt.Println("Anda termasuk:", generasi)

	fmt.Println("--------------------------------------------------")
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// soal 1
	firstWord := "Bootcamp"
	secondWord := "Digital"
	thirdWord := "Skill"
	fourthWord := "Sanbercode"
	fifthWord := "Golang"

	sentences := firstWord + " " + secondWord + " " + thirdWord + " " + fourthWord + " " + fifthWord

	fmt.Println(sentences)

	// soal 2
	halo := "Halo Dunia"
	cari := "Dunia"
	ubah := "Golang"

	hasil := strings.Replace(halo, cari, ubah, 1)

	fmt.Println(hasil)

	// soal 3
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"

	kataKedua = strings.Title(kataKedua)
	kataKetiga = kataKetiga[:len(kataKetiga)-1] + strings.ToUpper(string(kataKetiga[len(kataKetiga)-1]))
	kataKeempat = strings.ToUpper(kataKeempat)

	modifikasi := kataPertama + " " + kataKedua + " " + kataKetiga + " " + kataKeempat

	fmt.Println(modifikasi)

	// soal 4
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	intAngkaPertama, _ := strconv.Atoi(angkaPertama)
	intAngkaKedua, _ := strconv.Atoi(angkaKedua)
	intAngkaKetiga, _ := strconv.Atoi(angkaKetiga)
	intAngkaKeempat, _ := strconv.Atoi(angkaKeempat)

	total := intAngkaPertama + intAngkaKedua + intAngkaKetiga + intAngkaKeempat

	fmt.Println(total)

	// soal 5
	kalimat := "halo halo bandung"
	angka := 2021

	kalimat = strings.Replace(kalimat, "halo halo", "Hi Hi", 1)

	kalimatBaru := `"` + kalimat + `"` + " - " + fmt.Sprint(angka)

	fmt.Println(kalimatBaru)
}

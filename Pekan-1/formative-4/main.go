package main

import (
	"fmt"
	"strings"
)

func main() {
	// soal 1
	fmt.Println("--------------------- Soal 1 ---------------------")

	for i := 1; i <= 20; i++ {
		if i%3 == 0 && i%2 != 0 {
			fmt.Println(i, "- I Love Coding")
		} else if i%2 == 0 {
			fmt.Println(i, "- Berkualitas")
		} else {
			fmt.Println(i, "- Santai")
		}
	}

	fmt.Println("--------------------------------------------------")

	// soal 2
	fmt.Println("--------------------- Soal 2 ---------------------")
	for i := 1; i <= 7; i++ {
		fmt.Println(strings.Repeat("#", i))
	}

	fmt.Println("--------------------------------------------------")

	// soal 3
	fmt.Println("--------------------- Soal 3 ---------------------")
	var kalimat = [...]string{
		"aku",
		"dan",
		"saya",
		"sangat",
		"senang",
		"belajar",
		"golang",
	}

	var hasil = kalimat[2:7]

	fmt.Println(hasil)

	fmt.Println("--------------------------------------------------")

	// soal 4
	fmt.Println("--------------------- Soal 4 ---------------------")
	var sayuran = []string{}

	sayuran = append(
		sayuran,
		"Bayam",
		"Buncis",
		"Kangkung",
		"Kubis",
		"Seledri",
		"Tauge",
		"Timun",
	)

	for i, item := range sayuran {
		fmt.Printf("%d. %s\n", i+1, item)
	}

	fmt.Println("--------------------------------------------------")

	// soal 5
	fmt.Println("--------------------- Soal 5 ---------------------")
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}

	for key, value := range satuan {
		fmt.Printf("%s = %d\n", key, value)
	}

	volumeBalok := satuan["panjang"] * satuan["lebar"] * satuan["tinggi"]

	fmt.Printf("Volume Balok = %d\n", volumeBalok)

	fmt.Println("--------------------------------------------------")
}

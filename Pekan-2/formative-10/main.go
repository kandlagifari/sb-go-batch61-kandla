package main

import (
	"errors"
	"fmt"
	"sort"
	"sync"
	"time"
)

// soal 1
func tampilkanPesan(kalimat string, tahun int) {
	fmt.Printf("%s %d\n", kalimat, tahun)
}

func soal1() {
	// Dikeluarkan dari main function untuk menyesuaikan dengan pembatas (---)
	defer tampilkanPesan("Golang Backend Development", 2021)

	fmt.Println("Kode ini dieksekusi terlebih dahulu.")
	fmt.Println("Program selesai.")
}

// soal 2
func kelilingSegitigaSamaSisi(sisi int, denganPesan bool) (string, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Terjadi panic dan berhasil di-recover:", r)
		}
	}()

	if sisi == 0 {
		err := errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
		if !denganPesan {
			panic(err)
		}
		return "", err
	}

	keliling := sisi * 3

	if denganPesan {
		return fmt.Sprintf("keliling segitiga sama sisinya dengan sisi %d cm adalah %d cm", sisi, keliling), nil
	}
	return fmt.Sprintf("%d", keliling), nil
}

// soal 3
func tambahAngka(angkaTambahan int, angka *int) {
	*angka += angkaTambahan
}

func cetakAngka(angka *int) {
	fmt.Println("Total angka:", *angka)
}

func soal3() {
	// Dikeluarkan dari main function untuk menyesuaikan dengan pembatas (---)
	angka := 1

	defer cetakAngka(&angka)

	tambahAngka(7, &angka)

	tambahAngka(6, &angka)

	tambahAngka(-1, &angka)

	tambahAngka(9, &angka)
}

// soal 4
func tambahPhone(phones *[]string, phone string) {
	*phones = append(*phones, phone)
}

func tampilkanPhones(phones []string) {
	sort.Strings(phones)

	for i, phone := range phones {
		time.Sleep(1 * time.Second)

		fmt.Printf("%d. %s\n", i+1, phone)
	}
}

// soal 5
func tampilkanPhoneDenganUrutan(index int, phone string, wg *sync.WaitGroup, ch chan<- string) {
	defer wg.Done()

	time.Sleep(1 * time.Second)

	ch <- fmt.Sprintf("%d. %s", index, phone)
}

func soal5() {
	// Dikeluarkan dari main function karena ada nama variable yang sama (`phones`)
	var phones = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}

	sort.Strings(phones)

	ch := make(chan string, len(phones))

	var wg sync.WaitGroup

	for i, phone := range phones {
		wg.Add(1)
		go tampilkanPhoneDenganUrutan(i+1, phone, &wg, ch)
		wg.Wait()
	}

	close(ch)

	var results []string
	for res := range ch {
		results = append(results, res)
	}

	for _, res := range results {
		fmt.Println(res)
	}
}

// soal 6
func getMovies(moviesChannel chan string, movies ...string) {
	for _, movie := range movies {
		moviesChannel <- movie
	}

	close(moviesChannel)
}

func main() {
	// soal 1
	fmt.Println("--------------------- Soal 1 ---------------------")
	soal1()

	fmt.Println("--------------------------------------------------")

	// soal 2
	fmt.Println("--------------------- Soal 2 ---------------------")
	fmt.Println(kelilingSegitigaSamaSisi(4, true))
	fmt.Println(kelilingSegitigaSamaSisi(8, false))
	fmt.Println(kelilingSegitigaSamaSisi(0, true))
	fmt.Println(kelilingSegitigaSamaSisi(0, false))

	fmt.Println("--------------------------------------------------")

	// soal 3
	fmt.Println("--------------------- Soal 3 ---------------------")
	soal3()

	fmt.Println("--------------------------------------------------")

	// soal 4
	fmt.Println("--------------------- Soal 4 ---------------------")
	var phones []string

	tambahPhone(&phones, "Xiaomi")
	tambahPhone(&phones, "Asus")
	tambahPhone(&phones, "IPhone")
	tambahPhone(&phones, "Samsung")
	tambahPhone(&phones, "Oppo")
	tambahPhone(&phones, "Realme")
	tambahPhone(&phones, "Vivo")

	tampilkanPhones(phones)

	fmt.Println("--------------------------------------------------")

	// soal 5
	fmt.Println("--------------------- Soal 5 ---------------------")
	soal5()

	fmt.Println("--------------------------------------------------")

	// soal 6
	fmt.Println("--------------------- Soal 6 ---------------------")
	var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}

	moviesChannel := make(chan string)

	go getMovies(moviesChannel, movies...)

	for value := range moviesChannel {
		fmt.Println(value)
	}

	fmt.Println("--------------------------------------------------")
}

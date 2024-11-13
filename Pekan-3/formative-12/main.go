package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type NilaiMahasiswa struct {
	Nama, MataKuliah, IndeksNilai string
	Nilai, ID                     uint
}

var nilaiNilaiMahasiswa = []NilaiMahasiswa{}
var mu sync.Mutex

var lastID uint = 0

func hitungIndeksNilai(nilai uint) string {
	if nilai >= 80 {
		return "A"
	} else if nilai >= 70 {
		return "B"
	} else if nilai >= 60 {
		return "C"
	} else if nilai >= 50 {
		return "D"
	} else {
		return "E"
	}
}

func postNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if !ok || username != "admin" || password != "password123" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "Unauthorized")
		return
	}

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		Nama, MataKuliah string
		Nilai            uint
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if input.Nilai > 100 {
		http.Error(w, "Nilai tidak boleh lebih dari 100", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	lastID++
	indeksNilai := hitungIndeksNilai(input.Nilai)

	nilaiMahasiswa := NilaiMahasiswa{
		Nama:        input.Nama,
		MataKuliah:  input.MataKuliah,
		Nilai:       input.Nilai,
		ID:          lastID,
		IndeksNilai: indeksNilai,
	}

	nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, nilaiMahasiswa)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nilaiMahasiswa)
}

func getNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nilaiNilaiMahasiswa)
}

func main() {
	mu.Lock()
	lastID++
	nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, NilaiMahasiswa{
		ID:          lastID,
		Nama:        "Zoro",
		MataKuliah:  "Pedang",
		Nilai:       85,
		IndeksNilai: hitungIndeksNilai(85),
	})
	lastID++
	nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, NilaiMahasiswa{
		ID:          lastID,
		Nama:        "Nami",
		MataKuliah:  "Uang",
		Nilai:       78,
		IndeksNilai: hitungIndeksNilai(78),
	})
	lastID++
	nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, NilaiMahasiswa{
		ID:          lastID,
		Nama:        "Sanji",
		MataKuliah:  "Okama",
		Nilai:       65,
		IndeksNilai: hitungIndeksNilai(65),
	})
	mu.Unlock()

	http.HandleFunc("/nilai-mahasiswa", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			postNilaiMahasiswa(w, r)
		} else if r.Method == http.MethodGet {
			getNilaiMahasiswa(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Starting web server at http://localhost:30069")
	log.Fatal(http.ListenAndServe(":30069", nil))
}

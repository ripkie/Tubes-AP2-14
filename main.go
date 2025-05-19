package main

import (
	"fmt"
	"strings"
)

const NMAX = 100

type Komentar struct {
	Teks      string
	Sentimen  string
	Moderator string
}

type Akun struct {
	username string
	password string
}

var akunList [NMAX]Akun
var komentarList [NMAX]Komentar
var jumlahAkun, jumlahKomentar int

// ====== Data Dummy ======
func inisialisasiDataDummy() {
	komentarList[0] = Komentar{"Saya sangat senang dengan fitur ini", "", "admin"}
	komentarList[1] = Komentar{"Aplikasi ini buruk dan membingungkan", "", "admin"}
	komentarList[2] = Komentar{"Cukup membantu, tapi masih kurang", "", "admin"}
	jumlahKomentar = 3

	for i := 0; i < jumlahKomentar; i++ {
		komentarList[i].Sentimen = analisisSentimen(komentarList[i].Teks)
	}
}

// ===== Register & Login =====
func register(username, password string) string {
	for i := 0; i < jumlahAkun; i++ {
		if akunList[i].username == username {
			return "Username sudah digunakan."
		}
	}
	akunList[jumlahAkun] = Akun{username, password}
	jumlahAkun++
	return "Registrasi berhasil."
}

func login(username, password string) bool {
	for i := 0; i < jumlahAkun; i++ {
		if akunList[i].username == username && akunList[i].password == password {
			return true
		}
	}
	return false
}

// ===== Analisis Sentimen =====
func analisisSentimen(teks string) string {
	positif := [5]string{"senang", "bagus", "baik", "suka", "mantap"}
	negatif := [5]string{"buruk", "jelek", "benci", "tidak", "kurang"}

	teks = strings.ToLower(teks)
	for _, kata := range positif {
		if strings.Contains(teks, kata) {
			return "positif"
		}
	}
	for _, kata := range negatif {
		if strings.Contains(teks, kata) {
			return "negatif"
		}
	}
	return "netral"
}

// ===== CRUD Komentar =====
func tambahKomentar(teks, moderator string) {
	komentarList[jumlahKomentar] = Komentar{teks, analisisSentimen(teks), moderator}
	jumlahKomentar++
}

func ubahKomentar(index int, teks string) {
	if index >= 0 && index < jumlahKomentar {
		komentarList[index].Teks = teks
		komentarList[index].Sentimen = analisisSentimen(teks)
	}
}

func hapusKomentar(index int) {
	if index >= 0 && index < jumlahKomentar {
		for i := index; i < jumlahKomentar-1; i++ {
			komentarList[i] = komentarList[i+1]
		}
		jumlahKomentar--
	}
}

// ===== Pencarian =====
func sequentialSearch(keyword string) {
	keyword = strings.ToLower(keyword)
	for i := 0; i < jumlahKomentar; i++ {
		if strings.Contains(strings.ToLower(komentarList[i].Teks), keyword) {
			tampilkanKomentar(i)
		}
	}
}

func binarySearch(keyword string) {
	insertionSortByTeks()
	low := 0
	high := jumlahKomentar - 1
	for low <= high {
		mid := (low + high) / 2
		if strings.Contains(komentarList[mid].Teks, keyword) {
			tampilkanKomentar(mid)
			return
		} else if komentarList[mid].Teks < keyword {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	fmt.Println("Komentar tidak ditemukan.")
}

// ===== Pengurutan =====
func selectionSortBySentimen() {
	for i := 0; i < jumlahKomentar-1; i++ {
		min := i
		for j := i + 1; j < jumlahKomentar; j++ {
			if komentarList[j].Sentimen < komentarList[min].Sentimen {
				min = j
			}
		}
		komentarList[i], komentarList[min] = komentarList[min], komentarList[i]
	}
}

func insertionSortByTeks() {
	for i := 1; i < jumlahKomentar; i++ {
		temp := komentarList[i]
		j := i - 1
		for j >= 0 && komentarList[j].Teks > temp.Teks {
			komentarList[j+1] = komentarList[j]
			j--
		}
		komentarList[j+1] = temp
	}
}

// ===== Statistik =====
func tampilkanStatistik() {
	pos, net, neg := 0, 0, 0
	for i := 0; i < jumlahKomentar; i++ {
		switch komentarList[i].Sentimen {
		case "positif":
			pos++
		case "netral":
			net++
		case "negatif":
			neg++
		}
	}
	fmt.Println("Statistik Sentimen:")
	fmt.Println("Positif:", pos)
	fmt.Println("Netral :", net)
	fmt.Println("Negatif:", neg)
}

// ===== Tampilan Komentar =====
func tampilkanKomentar(index int) {
	k := komentarList[index]
	fmt.Printf("[%d] (%s) %s -> %s\n", index, k.Moderator, k.Teks, k.Sentimen)
}

func tampilkanSemuaKomentar() {
	for i := 0; i < jumlahKomentar; i++ {
		tampilkanKomentar(i)
	}
}

// ===== Main Program =====
func main() {
	inisialisasiDataDummy()

	var u, p string
	fmt.Print("Username: ")
	fmt.Scan(&u)
	fmt.Print("Password: ")
	fmt.Scan(&p)

	if !login(u, p) {
		fmt.Println(register(u, p))
	}

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Tambah Komentar")
		fmt.Println("2. Ubah Komentar")
		fmt.Println("3. Hapus Komentar")
		fmt.Println("4. Tampilkan Semua Komentar")
		fmt.Println("5. Cari (Sequential)")
		fmt.Println("6. Cari (Binary Search)")
		fmt.Println("7. Urutkan (Sentimen)")
		fmt.Println("8. Urutkan (Teks)")
		fmt.Println("9. Statistik Sentimen")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih: ")

		var pilih int
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			fmt.Print("Komentar baru: ")
			var teks string
			fmt.Scanln()
			getline := ""
			for getline == "" {
				fmt.Scanln(&getline)
			}
			tambahKomentar(getline, u)
		case 2:
			fmt.Print("Index komentar yang diubah: ")
			var i int
			fmt.Scan(&i)
			fmt.Print("Komentar baru: ")
			var teks string
			fmt.Scanln()
			getline := ""
			for getline == "" {
				fmt.Scanln(&getline)
			}
			ubahKomentar(i, getline)
		case 3:
			fmt.Print("Index komentar yang dihapus: ")
			var i int
			fmt.Scan(&i)
			hapusKomentar(i)
		case 4:
			tampilkanSemuaKomentar()
		case 5:
			fmt.Print("Kata kunci: ")
			var k string
			fmt.Scan(&k)
			sequentialSearch(k)
		case 6:
			fmt.Print("Kata kunci: ")
			var k string
			fmt.Scan(&k)
			binarySearch(k)
		case 7:
			selectionSortBySentimen()
			fmt.Println("Diurutkan berdasarkan sentimen.")
		case 8:
			insertionSortByTeks()
			fmt.Println("Diurutkan berdasarkan teks.")
		case 9:
			tampilkanStatistik()
		case 0:
			fmt.Println("Keluar...")
			return
		}
	}
}

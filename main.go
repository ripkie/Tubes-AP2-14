package main

import "fmt"

const NMAX int = 100

type arrkata [NMAX]string

type komentar struct {
	isi        arrkata
	jumlahKata int
}

type arrKalimat [NMAX]komentar

func main() {
	var daftarKomentar arrKalimat
	var jumlahKomentar int
	var pilihan int

	// Data dummy contoh (3 komentar)
	daftarKomentar[0].isi[0] = "saya"
	daftarKomentar[0].isi[1] = "sangat"
	daftarKomentar[0].isi[2] = "senang"
	daftarKomentar[0].jumlahKata = 3

	daftarKomentar[1].isi[0] = "aplikasi"
	daftarKomentar[1].isi[1] = "ini"
	daftarKomentar[1].isi[2] = "buruk"
	daftarKomentar[1].jumlahKata = 3

	daftarKomentar[2].isi[0] = "fitur"
	daftarKomentar[2].isi[1] = "mantap"
	daftarKomentar[2].jumlahKata = 2

	jumlahKomentar = 3

	for {
		tampilkanMenu()
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			tambahKomentar(&daftarKomentar, &jumlahKomentar)
		case 2:
			ubahKomentar(&daftarKomentar, jumlahKomentar)
		case 3:
			hapusKomentar(&daftarKomentar, &jumlahKomentar)
		case 4:
			tampilkanKomentar(daftarKomentar, jumlahKomentar)
		case 5:
			tampilkanAnalisisKomentar(daftarKomentar, jumlahKomentar)
		case 6:
			cariKomentarBerdasarkanKategori(daftarKomentar, jumlahKomentar)
		case 7:
			urutkanKomentar(&daftarKomentar, jumlahKomentar)
		case 8:
			tampilkanStatistik(daftarKomentar, jumlahKomentar)

		case 9:
			fmt.Println("╔═══════════════════════════════════════╗")
			fmt.Println("║     TERIMA KASIH TELAH MENGGUNAKAN    ║")
			fmt.Println("║  APLIKASI ANALISIS SENTIMEN KOMENTAR! ║")
			fmt.Println("╠═══════════════════════════════════════╣")
			fmt.Println("║         Semoga HARI-HARI Anda         ║")
			fmt.Println("║              menyenangkan             ║")
			fmt.Println("╚═══════════════════════════════════════╝")
			return
		default:
			fmt.Println("❌ Pilihan tidak valid.")
		}
	}
}

func tampilkanMenu() {
	fmt.Println("╔════════════════════════════════════════════════╗")
	fmt.Println("║             MENU ANALISIS KOMENTAR             ║")
	fmt.Println("╠════════════════════════════════════════════════╣")
	fmt.Println("║ 1. Tambah Komentar                             ║")
	fmt.Println("║ 2. Ubah Komentar                               ║")
	fmt.Println("║ 3. Hapus Komentar                              ║")
	fmt.Println("║ 4. Tampilkan Seluruh Komentar                  ║")
	fmt.Println("║ 5. Tampilkan Anlisis Komentar                  ║")
	fmt.Println("║ 6. Cari Komentar (Binary & Sequential Search)  ║")
	fmt.Println("║ 7. Urutkan Komentar                            ║")
	fmt.Println("║ 8. Tampilkan Statistik Komentar                ║")
	fmt.Println("║ 9. Keluar                                      ║")
	fmt.Println("╚════════════════════════════════════════════════╝")
	fmt.Println("▶ Pilih menu: ")
}

func tambahKomentar(daftarKomentar *arrKalimat, idxKomentar *int) {
	var jumKata int
	if *idxKomentar >= NMAX {
		fmt.Println("❌ Maksimum komentar telah tercapai.")
		return
	}

	fmt.Print("masukan jumlah kata dalam komentar: ")
	fmt.Scanln(&jumKata)

	fmt.Println("Ketik komentar (kata-kata dipisah spasi):")

	daftarKomentar[*idxKomentar].jumlahKata = 0 // reset jumlah kata sebelum input

	for i := 0; i < jumKata; i++ {
		fmt.Scan(&daftarKomentar[*idxKomentar].isi[i])
		daftarKomentar[*idxKomentar].jumlahKata++
	}

	*idxKomentar = *idxKomentar + 1
	fmt.Println("✅ Komentar berhasil ditambahkan.")

	// Stop / pause sebelum lanjut
	fmt.Print("Tekan ENTER untuk kembali ke menu...")
	fmt.Scanln() // baca newline yang tertinggal
	fmt.Scanln() // tunggu user tekan enter

	tampilkanMenu() // panggil menu selanjutnya
}

func ubahKomentar(daftarKomentar *arrKalimat, jumlahKomentar int) {
	var indeks int
	var jumKata int

	if jumlahKomentar == 0 {
		fmt.Println("❗ Belum ada komentar yang dapat diubah.")
		return
	}

	fmt.Println("Daftar Komentar:")
	for i := 0; i < jumlahKomentar; i++ {
		fmt.Printf("%d. ", i+1)
		for j := 0; j < daftarKomentar[i].jumlahKata; j++ {
			fmt.Print(daftarKomentar[i].isi[j], " ")
		}
		fmt.Println()
	}

	fmt.Print("Masukkan nomor komentar yang ingin diubah: ")
	fmt.Scanln(&indeks)

	indeks-- // Konversi ke indeks array
	if indeks < 0 || indeks >= jumlahKomentar {
		fmt.Println("❌ Indeks komentar tidak valid.")
		return
	}

	fmt.Print("Berapa jumlah kata dalam komentar baru? ")
	fmt.Scanln(&jumKata)

	daftarKomentar[indeks].jumlahKata = 0
	for i := 0; i < jumKata; i++ {
		fmt.Scan(&daftarKomentar[indeks].isi[i])
		daftarKomentar[indeks].jumlahKata++
	}

	fmt.Println("✅ Komentar berhasil diubah.")

	// Pause sebelum lanjut, misalnya kembali ke menu
	fmt.Print("Tekan ENTER untuk kembali ke menu...")
	fmt.Scanln() // menangkap newline leftover
	fmt.Scanln() // menunggu enter dari user

	tampilkanMenu() // panggil menu utama jika ada
}

func hapusKomentar(daftarKomentar *arrKalimat, jumlahKomentar *int) {
	var indeks int

	if *jumlahKomentar == 0 {
		fmt.Println("❗ Belum ada komentar yang dapat dihapus.")
		return
	}

	fmt.Println("Daftar Komentar:")
	for i := 0; i < *jumlahKomentar; i++ {
		fmt.Printf("%d. ", i+1)
		for j := 0; j < daftarKomentar[i].jumlahKata; j++ {
			fmt.Print(daftarKomentar[i].isi[j], " ")
		}
		fmt.Println()
	}

	fmt.Print("Masukkan nomor komentar yang ingin dihapus: ")
	fmt.Scanln(&indeks)

	indeks-- // Konversi ke indeks array
	if indeks < 0 || indeks >= *jumlahKomentar {
		fmt.Println("❌ Indeks komentar tidak valid.")
		return
	}

	// Geser semua komentar setelah indeks ke kiri
	for i := indeks; i < *jumlahKomentar-1; i++ {
		daftarKomentar[i] = daftarKomentar[i+1]
	}

	*jumlahKomentar = *jumlahKomentar - 1
	fmt.Println("✅ Komentar berhasil dihapus.")
}

func tampilkanKomentar(daftarKomentar arrKalimat, jumlahKomentar int) {
	if jumlahKomentar == 0 {
		fmt.Println("Belum ada komentar.")
		return
	}

	fmt.Println("📋 Daftar Komentar:")
	for i := 0; i < jumlahKomentar; i++ {
		fmt.Printf("- Komentar %d: ", i+1)
		for j := 0; j < daftarKomentar[i].jumlahKata; j++ {
			fmt.Print(daftarKomentar[i].isi[j], " ")
		}
		fmt.Println()
	}
}

func tampilkanAnalisisKomentar(daftarKomentar arrKalimat, jumlahKomentar int) {
	var positif = [5]string{"baik", "bagus", "senang", "mantap", "puas"}
	var negatif = [5]string{"buruk", "jelek", "kecewa", "parah", "tidak"}

	if jumlahKomentar == 0 {
		fmt.Println("❗ Belum ada komentar yang dapat dianalisis.")
		return
	}

	fmt.Println("📊 Hasil Analisis Sentimen Komentar:")
	for i := 0; i < jumlahKomentar; i++ {
		var jumlahPositif int = 0
		var jumlahNegatif int = 0

		for j := 0; j < daftarKomentar[i].jumlahKata; j++ {
			kata := daftarKomentar[i].isi[j]

			// Cek kata positif satu per satu
			for k := 0; k < 5; k++ {
				if kata == positif[k] {
					jumlahPositif = jumlahPositif + 1
				}
			}

			// Cek kata negatif satu per satu
			for k := 0; k < 5; k++ {
				if kata == negatif[k] {
					jumlahNegatif = jumlahNegatif + 1
				}
			}
		}

		var hasil string = "Netral"
		if jumlahPositif > jumlahNegatif {
			hasil = "Positif"
		} else if jumlahNegatif > jumlahPositif {
			hasil = "Negatif"
		}

		// Tampilkan hasil komentar dan analisisnya
		fmt.Printf("%d. ", i+1)
		for j := 0; j < daftarKomentar[i].jumlahKata; j++ {
			fmt.Print(daftarKomentar[i].isi[j], " ")
		}
		fmt.Printf("➡️  [%s]\n", hasil)
	}
}

func cariKomentarBerdasarkanKategori(daftarKomentar arrKalimat, jumlahKomentar int) {
	var pilihan int
	fmt.Println("╔════════════════════════════════════════════╗")
	fmt.Println("║        PILIH JENIS PENCARIAN KOMENTAR      ║")
	fmt.Println("╠════════════════════════════════════════════╣")
	fmt.Println("║ 1. Cari komentar berdasarkan sentimen      ║")
	fmt.Println("║ 2. Cari komentar berdasarkan kata          ║")
	fmt.Println("╚════════════════════════════════════════════╝")
	fmt.Print("▶ Pilih jenis pencarian: ")
	fmt.Scanln(&pilihan)

	// sequential search
	if pilihan == 1 {
		var kategori string
		fmt.Print("Masukkan kategori yang ingin dicari (positif/negatif/netral): ")
		fmt.Scanln(&kategori)

		var hasilDitemukan bool = false
		fmt.Println("\n🔍 Hasil Pencarian (Sequential Search):")

		for i := 0; i < jumlahKomentar; i++ {
			var jumlahPositif, jumlahNegatif int

			for j := 0; j < daftarKomentar[i].jumlahKata; j++ {
				kata := daftarKomentar[i].isi[j]

				if kata == "baik" || kata == "bagus" || kata == "senang" || kata == "mantap" || kata == "puas" {
					jumlahPositif++
				}
				if kata == "buruk" || kata == "jelek" || kata == "kecewa" || kata == "parah" || kata == "tidak" {
					jumlahNegatif++
				}
			}

			var hasil string = "netral"
			if jumlahPositif > jumlahNegatif {
				hasil = "positif"
			} else if jumlahNegatif > jumlahPositif {
				hasil = "negatif"
			}

			if hasil == kategori {
				hasilDitemukan = true
				fmt.Printf("- Komentar ke-%d: ", i+1)
				for j := 0; j < daftarKomentar[i].jumlahKata; j++ {
					fmt.Print(daftarKomentar[i].isi[j], " ")
				}
				fmt.Println()
			}
		}

		if !hasilDitemukan {
			fmt.Println("Tidak ada komentar dengan kategori tersebut.")
		}

		// binary search
	} else if pilihan == 2 {
		fmt.Print("Masukkan kata yang ingin dicari di komentar: ")
		var keyword string
		fmt.Scanln(&keyword)

		var ditemukan bool = false
		fmt.Println("\n🔍 Hasil Pencarian (Sequential Search):")
		for i := 0; i < jumlahKomentar; i++ {
			var komentarDicetak bool = false
			for j := 0; j < daftarKomentar[i].jumlahKata; j++ {
				if daftarKomentar[i].isi[j] == keyword {
					ditemukan = true
					if !komentarDicetak {
						fmt.Printf("- Komentar ke-%d: ", i+1)
						for k := 0; k < daftarKomentar[i].jumlahKata; k++ {
							fmt.Print(daftarKomentar[i].isi[k], " ")
						}
						fmt.Println()
						komentarDicetak = true
					}
				}
			}
		}

		if !ditemukan {
			fmt.Println("Komentar dengan kata tersebut tidak ditemukan.")
		}

	} else {
		fmt.Println("❌ Pilihan tidak valid.")
	}
}

func urutkanKomentar(data *arrKalimat, jumlah int) {
	var pilihan int
	fmt.Println("╔═════════════════════════════════════════════════════╗")
	fmt.Println("║             PILIH JENIS PENGURUTAN KOMENTAR         ║")
	fmt.Println("╠═════════════════════════════════════════════════════╣")
	fmt.Println("║ 1. Berdasarkan jumlah kata (Selection Sort)         ║")
	fmt.Println("║ 2. Berdasarkan sentimen Positif → Netral → Negatif  ║")
	fmt.Println("╚═════════════════════════════════════════════════════╝")
	fmt.Print("▶ Pilih jenis pengurutan: ")
	fmt.Scanln(&pilihan)

	if pilihan == 1 {
		// Selection Sort berdasarkan jumlah kata
		for i := 0; i < jumlah-1; i++ {
			minIdx := i
			for j := i + 1; j < jumlah; j++ {
				if data[j].jumlahKata < data[minIdx].jumlahKata {
					minIdx = j
				}
			}
			if minIdx != i {
				data[i], data[minIdx] = data[minIdx], data[i]
			}
		}

		fmt.Println("✅ Komentar berhasil diurutkan berdasarkan jumlah kata.")
		fmt.Println("📋 Daftar Komentar:")
		for i := 0; i < jumlah; i++ {
			fmt.Printf("- Komentar %d: ", i+1)
			for j := 0; j < data[i].jumlahKata; j++ {
				fmt.Print(data[i].isi[j], " ")
			}
			fmt.Println()
		}
	} else if pilihan == 2 {
		// Insertion Sort berdasarkan sentimen (positif < netral < negatif)
		for i := 1; i < jumlah; i++ {
			temp := data[i]
			j := i - 1

			for j >= 0 && bandingkanSentimen(data[j]) > bandingkanSentimen(temp) {
				data[j+1] = data[j]
				j--
			}
			data[j+1] = temp
		}

		fmt.Println("✅ Komentar berhasil diurutkan berdasarkan sentimen.")
		fmt.Println("📋 Daftar Komentar:")
		for i := 0; i < jumlah; i++ {
			fmt.Printf("- Komentar %d: ", i+1)
			for j := 0; j < data[i].jumlahKata; j++ {
				fmt.Print(data[i].isi[j], " ")
			}
			fmt.Println()
		}

	} else {
		fmt.Println("❌ Pilihan tidak valid.")
	}
}

// func pendukung untuk sorting
func bandingkanSentimen(k komentar) int {
	var jumlahPositif, jumlahNegatif int

	for i := 0; i < k.jumlahKata; i++ {
		if k.isi[i] == "baik" || k.isi[i] == "bagus" || k.isi[i] == "senang" || k.isi[i] == "mantap" || k.isi[i] == "puas" {
			jumlahPositif++
		} else if k.isi[i] == "buruk" || k.isi[i] == "jelek" || k.isi[i] == "kecewa" || k.isi[i] == "parah" || k.isi[i] == "tidak" {
			jumlahNegatif++
		}
	}

	if jumlahPositif > jumlahNegatif {
		return 0 // Positif
	} else if jumlahNegatif > jumlahPositif {
		return 2 // Negatif
	} else {
		return 1 // Netral
	}
}

func tampilkanStatistik(daftarKomentar arrKalimat, jumlahKomentar int) {
	if jumlahKomentar == 0 {
		fmt.Println("❗ Belum ada komentar untuk ditampilkan statistiknya.")
		return
	}

	var jumlahPositif, jumlahNegatif, jumlahNetral int

	for i := 0; i < jumlahKomentar; i++ {
		positif := 0
		negatif := 0

		for j := 0; j < daftarKomentar[i].jumlahKata; j++ {
			kata := daftarKomentar[i].isi[j]

			// Cek apakah kata termasuk positif
			if kata == "baik" || kata == "bagus" || kata == "senang" || kata == "mantap" || kata == "puas" {
				positif++
			}

			// Cek apakah kata termasuk negatif
			if kata == "buruk" || kata == "jelek" || kata == "kecewa" || kata == "parah" || kata == "tidak" {
				negatif++
			}
		}

		if positif > negatif {
			jumlahPositif++
		} else if negatif > positif {
			jumlahNegatif++
		} else {
			jumlahNetral++
		}
	}

	// Hitung persentase
	total := jumlahPositif + jumlahNetral + jumlahNegatif
	persenPositif := float64(jumlahPositif) / float64(total) * 100
	persenNetral := float64(jumlahNetral) / float64(total) * 100
	persenNegatif := float64(jumlahNegatif) / float64(total) * 100

	// Tampilkan laporan statistik
	fmt.Println("╔═══════════════════════════════════════╗")
	fmt.Println("║       STATISTIK SENTIMEN KOMENTAR     ║")
	fmt.Println("╠═══════════════════════════════════════╣")
	fmt.Printf("║ Total Komentar : %-21d║\n", total)
	fmt.Printf("║ Positif        : %-3d (%.1f%%)          ║\n", jumlahPositif, persenPositif)
	fmt.Printf("║ Netral         : %-3d (%.1f%%)           ║\n", jumlahNetral, persenNetral)
	fmt.Printf("║ Negatif        : %-3d (%.1f%%)          ║\n", jumlahNegatif, persenNegatif)
	fmt.Println("╚═══════════════════════════════════════╝")
}

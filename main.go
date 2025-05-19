package main

import "fmt"

type Akun struct {
	username string
	password string
}

var akunList [100]Akun
var jumlahAkun int

func register(username string, password string) string {
	// Cek apakah username sudah ada
	ketemu := false
	for i := 0; i < jumlahAkun; i++ {
		ketemu = ketemu || (akunList[i].username == username)
	}
	return map[bool]string{
		true:  "Username sudah digunakan",
		false: tambahAkun(username, password),
	}[ketemu]
}

func tambahAkun(username string, password string) string {
	akunList[jumlahAkun] = Akun{username, password}
	jumlahAkun++
	return "Registrasi berhasil"
}

func login(username string, password string) string {
	sukses := false
	for i := 0; i < jumlahAkun; i++ {
		sukses = sukses || (akunList[i].username == username && akunList[i].password == password)
	}
	return map[bool]string{
		true:  "Login berhasil",
		false: "Login gagal",
	}[sukses]
}

func main() {
	for {
		var pilih int
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilih)

		pesan := map[int]string{
			1: jalankanRegister(),
			2: jalankanLogin(),
			3: "Keluar",
		}[pilih]

		fmt.Println(pesan)
		if pilih == 3 {
			break
		}
	}
}

func jalankanRegister() string {
	var u, p string
	fmt.Print("Username baru: ")
	fmt.Scan(&u)
	fmt.Print("Password baru: ")
	fmt.Scan(&p)
	return register(u, p)
}

func jalankanLogin() string {
	var u, p string
	fmt.Print("Masukkan username: ")
	fmt.Scan(&u)
	fmt.Print("Masukkan password: ")
	fmt.Scan(&p)
	return login(u, p)
}

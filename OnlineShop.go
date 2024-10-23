package main

// Import Fungsi Yang akan Digunakan
import (
	"fmt"
	"os"
)

// Struct untuk User
type User struct {
	ID       int
	Username string
	Password string
	Role     string // "admin", "pemilik", "pembeli"
	Status   string // "pending", "approved", "rejected"
}

// Struct untuk Barang
type Barang struct {
	ID    int
	Nama  string
	Harga int
	Stock int
}

// Struct untuk Transaksi
type Transaksi struct {
	ID         int
	PembeliID  int
	BarangID   int
	Jumlah     int
	TotalHarga int
}

// Deklarasi variabel global
const maxUsers = 100
const maxBarangs = 100
const maxTransaksis = 100

var users [maxUsers]User
var barangs [maxBarangs]Barang
var transaksis [maxTransaksis]Transaksi
var currentUser User
var nextUserID = 1
var nextBarangID = 1
var nextTransaksiID = 1
var userCount = 0
var barangCount = 0
var transaksiCount = 0

func main() {
	// Tambahkan user admin default saat aplikasi dimulai
	users[userCount] = User{ID: nextUserID, Username: "admin", Password: "admin", Role: "admin", Status: "approved"}
	nextUserID++
	userCount++

	// Loop utama untuk menampilkan menu utama secara terus menerus
	for {
		showMainMenu()
	}
}

// Menampilkan menu utama
func showMainMenu() {
	fmt.Println("Selamat datang di Aplikasi Toko Online")
	fmt.Println("1. Login")
	fmt.Println("2. Registrasi")
	fmt.Println("3. Keluar")
	fmt.Print("Pilih menu: ")
	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		login()
	case 2:
		registrasi()
	case 3:
		os.Exit(0)
	default:
		fmt.Println("Pilihan tidak valid")
	}
}

// Fungsi untuk login
func login() {
	fmt.Print("Username: ")
	var username string
	fmt.Scan(&username)
	fmt.Print("Password: ")
	var password string
	fmt.Scan(&password)

	// Cari user yang cocok dengan username dan password
	for i := 0; i < userCount; i++ {
		if users[i].Username == username && users[i].Password == password {
			if users[i].Status != "approved" {
				fmt.Println("Akun Anda belum disetujui oleh admin.")
				return
			}
			currentUser = users[i]
			switch currentUser.Role {
			case "admin":
				showAdminMenu()
			case "pemilik":
				showPemilikMenu()
			case "pembeli":
				showPembeliMenu()
			default:
				fmt.Println("Peran tidak dikenal")
			}
			return
		}
	}
	fmt.Println("Username atau password salah.")
}

// Fungsi untuk registrasi
func registrasi() {
	fmt.Println("Registrasi Akun")
	fmt.Print("Username: ")
	var username string
	fmt.Scan(&username)
	fmt.Print("Password: ")
	var password string
	fmt.Scan(&password)
	fmt.Print("Role (pemilik/pembeli): ")
	var role string
	fmt.Scan(&role)

	// Tambahkan user baru ke daftar users
	newUser := User{ID: nextUserID, Username: username, Password: password, Role: role, Status: "pending"}
	users[userCount] = newUser
	nextUserID++
	userCount++
	fmt.Println("Registrasi berhasil, menunggu persetujuan admin.")
}

// Menampilkan menu admin
func showAdminMenu() {
	for {
		fmt.Println("Menu Admin")
		fmt.Println("1. Setujui/Tolak Registrasi Akun")
		fmt.Println("2. Logout")
		fmt.Print("Pilih menu: ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			approveRejectRegistrations()
		case 2:
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

// Fungsi untuk menyetujui atau menolak registrasi
func approveRejectRegistrations() {
	fmt.Println("Daftar Registrasi")
	// Tampilkan daftar user yang statusnya pending
	for i := 0; i < userCount; i++ {
		if users[i].Status == "pending" {
			fmt.Printf("ID: %d, Username: %s, Role: %s\n", users[i].ID, users[i].Username, users[i].Role)
		}
	}
	fmt.Print("Masukkan ID user untuk disetujui/ditolak: ")
	var userID int
	fmt.Scan(&userID)
	fmt.Print("Setujui (y/n): ")
	var approve string
	fmt.Scan(&approve)

	// Setujui atau tolak user berdasarkan ID
	for i := 0; i < userCount; i++ {
		if users[i].ID == userID {
			if approve == "y" {
				users[i].Status = "approved"
				fmt.Println("Akun disetujui.")
			} else {
				users[i].Status = "rejected"
				fmt.Println("Akun ditolak.")
			}
			return
		}
	}
	fmt.Println("User tidak ditemukan.")
}

// Menampilkan menu pemilik toko
func showPemilikMenu() {
	for {
		fmt.Println("Menu Pemilik Toko")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Ubah Barang")
		fmt.Println("3. Hapus Barang")
		fmt.Println("4. Cetak Transaksi")
		fmt.Println("5. Lihat Barang")
		fmt.Println("6. Logout")
		fmt.Print("Pilih menu: ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			tambahBarang()
		case 2:
			ubahBarang()
		case 3:
			hapusBarang()
		case 4:
			cetakTransaksi()
		case 5:
			lihatBarang()
		case 6:
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

// Fungsi untuk menambah barang
func tambahBarang() {
	fmt.Print("Nama Barang: ")
	var nama string
	fmt.Scan(&nama)
	fmt.Print("Harga Barang: ")
	var harga int
	fmt.Scan(&harga)
	fmt.Print("Stock Barang: ")
	var stock int
	fmt.Scan(&stock)

	// Tambahkan barang baru ke daftar barangs
	newBarang := Barang{ID: nextBarangID, Nama: nama, Harga: harga, Stock: stock}
	barangs[barangCount] = newBarang
	nextBarangID++
	barangCount++
	fmt.Println("Barang berhasil ditambahkan.")
}

// Fungsi untuk mengubah barang
func ubahBarang() {
	fmt.Print("ID Barang: ")
	var barangID int
	fmt.Scan(&barangID)

	// Cari barang yang sesuai dengan ID dan ubah datanya
	for i := 0; i < barangCount; i++ {
		if barangs[i].ID == barangID {
			fmt.Print("Nama Barang: ")
			fmt.Scan(&barangs[i].Nama)
			fmt.Print("Harga Barang: ")
			fmt.Scan(&barangs[i].Harga)
			fmt.Print("Stock Barang: ")
			fmt.Scan(&barangs[i].Stock)
			fmt.Println("Barang berhasil diubah.")
			return
		}
	}
	fmt.Println("Barang tidak ditemukan.")
}

// Fungsi untuk menghapus barang
func hapusBarang() {
	fmt.Print("ID Barang: ")
	var barangID int
	fmt.Scan(&barangID)

	// Cari barang yang sesuai dengan ID dan hapus dari daftar barangs
	for i := 0; i < barangCount; i++ {
		if barangs[i].ID == barangID {
			for j := i; j < barangCount-1; j++ {
				barangs[j] = barangs[j+1]
			}
			barangCount--
			fmt.Println("Barang berhasil dihapus.")
			return
		}
	}
	fmt.Println("Barang tidak ditemukan.")
}

// Fungsi untuk mencetak transaksi
func cetakTransaksi() {
	fmt.Println("Daftar Transaksi")
	// Tampilkan semua transaksi
	for i := 0; i < transaksiCount; i++ {
		fmt.Printf("ID: %d, PembeliID: %d, BarangID: %d, Jumlah: %d, TotalHarga: %d\n", transaksis[i].ID, transaksis[i].PembeliID, transaksis[i].BarangID, transaksis[i].Jumlah, transaksis[i].TotalHarga)
	}
}

// Fungsi untuk melihat daftar barang yang diurutkan berdasarkan stok
func lihatBarang() {
	fmt.Println("Daftar Barang")

	// Lakukan pengurutan barang berdasarkan stok (bubble sort)
	for i := 0; i < barangCount-1; i++ {
		for j := 0; j < barangCount-i-1; j++ {
			if barangs[j].Stock > barangs[j+1].Stock {
				// Tukar posisi barang
				barangs[j], barangs[j+1] = barangs[j+1], barangs[j]
			}
		}
	}

	// Tampilkan daftar barang yang telah diurutkan
	for i := 0; i < barangCount; i++ {
		fmt.Printf("ID: %d, Nama: %s, Harga: %d, Stock: %d\n", barangs[i].ID, barangs[i].Nama, barangs[i].Harga, barangs[i].Stock)
	}
}

// Menampilkan menu pembeli
func showPembeliMenu() {
	for {
		fmt.Println("Menu Pembeli")
		fmt.Println("1. Beli Barang")
		fmt.Println("2. Logout")
		fmt.Print("Pilih menu: ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			beliBarang()
		case 2:
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

// Fungsi untuk membeli barang
func beliBarang() {
	fmt.Println("Daftar Barang")
	// Tampilkan semua barang yang tersedia
	for i := 0; i < barangCount; i++ {
		fmt.Printf("ID: %d, Nama: %s, Harga: %d, Stock: %d\n", barangs[i].ID, barangs[i].Nama, barangs[i].Harga, barangs[i].Stock)
	}
	fmt.Print("Masukkan ID Barang: ")
	var barangID int
	fmt.Scan(&barangID)
	fmt.Print("Masukkan Jumlah: ")
	var jumlah int
	fmt.Scan(&jumlah)

	// Cari barang yang sesuai dengan ID dan lakukan transaksi
	for i := 0; i < barangCount; i++ {
		if barangs[i].ID == barangID {
			if barangs[i].Stock < jumlah {
				fmt.Println("Stock barang tidak mencukupi.")
				return
			}
			totalHarga := barangs[i].Harga * jumlah
			transaksi := Transaksi{ID: nextTransaksiID, PembeliID: currentUser.ID, BarangID: barangID, Jumlah: jumlah, TotalHarga: totalHarga}
			transaksis[transaksiCount] = transaksi
			nextTransaksiID++
			transaksiCount++
			barangs[i].Stock -= jumlah
			fmt.Println("Barang berhasil dibeli.")
			return
		}
	}
	fmt.Println("Barang tidak ditemukan.")
}

package main

import (
	"fmt"
)

type Barang struct {
	Kode  string
	Nama  string
	Stok  int
	Harga int
}

var gudang []Barang

// ================= TAMBAH BARANG =================
func tambahBarang() {
	var b Barang

	fmt.Print("Masukkan kode barang : ")
	fmt.Scan(&b.Kode)

	fmt.Print("Masukkan nama barang : ")
	fmt.Scan(&b.Nama)

	fmt.Print("Masukkan stok barang : ")
	fmt.Scan(&b.Stok)

	fmt.Print("Masukkan harga barang : ")
	fmt.Scan(&b.Harga)

	gudang = append(gudang, b)

	fmt.Println("Barang berhasil ditambahkan.")
}

// ================= TAMPILKAN BARANG =================
func tampilkanBarang() {
	if len(gudang) == 0 {
		fmt.Println("Data barang kosong.")
		return
	}

	fmt.Println("\n===== DATA BARANG =====")

	for i, b := range gudang {
		fmt.Printf("%d. Kode: %s | Nama: %s | Stok: %d | Harga: %d\n",
			i+1, b.Kode, b.Nama, b.Stok, b.Harga)
	}
}

// ================= UBAH BARANG =================
func ubahBarang() {
	var kode string

	fmt.Print("Masukkan kode barang yang ingin diubah : ")
	fmt.Scan(&kode)

	for i := 0; i < len(gudang); i++ {
		if gudang[i].Kode == kode {

			fmt.Print("Masukkan nama baru : ")
			fmt.Scan(&gudang[i].Nama)

			fmt.Print("Masukkan stok baru : ")
			fmt.Scan(&gudang[i].Stok)

			fmt.Print("Masukkan harga baru : ")
			fmt.Scan(&gudang[i].Harga)

			fmt.Println("Data barang berhasil diubah.")
			return
		}
	}

	fmt.Println("Barang tidak ditemukan.")
}

// ================= HAPUS BARANG =================
func hapusBarang() {
	var kode string

	fmt.Print("Masukkan kode barang yang ingin dihapus : ")
	fmt.Scan(&kode)

	for i := 0; i < len(gudang); i++ {
		if gudang[i].Kode == kode {

			gudang = append(gudang[:i], gudang[i+1:]...)

			fmt.Println("Barang berhasil dihapus.")
			return
		}
	}

	fmt.Println("Barang tidak ditemukan.")
}

// ================= BARANG MASUK =================
func barangMasuk() {
	var kode string
	var jumlah int

	fmt.Print("Masukkan kode barang : ")
	fmt.Scan(&kode)

	fmt.Print("Masukkan jumlah barang masuk : ")
	fmt.Scan(&jumlah)

	for i := 0; i < len(gudang); i++ {
		if gudang[i].Kode == kode {

			gudang[i].Stok += jumlah

			fmt.Println("Stok barang berhasil ditambah.")
			return
		}
	}

	fmt.Println("Barang tidak ditemukan.")
}

// ================= BARANG KELUAR =================
func barangKeluar() {
	var kode string
	var jumlah int

	fmt.Print("Masukkan kode barang : ")
	fmt.Scan(&kode)

	fmt.Print("Masukkan jumlah barang keluar : ")
	fmt.Scan(&jumlah)

	for i := 0; i < len(gudang); i++ {
		if gudang[i].Kode == kode {

			if gudang[i].Stok >= jumlah {
				gudang[i].Stok -= jumlah
				fmt.Println("Stok barang berhasil dikurangi.")
			} else {
				fmt.Println("Stok tidak mencukupi.")
			}

			return
		}
	}

	fmt.Println("Barang tidak ditemukan.")
}

// ================= CARI BERDASARKAN NAMA =================
func cariNamaBarang() {
	var nama string

	fmt.Print("Masukkan nama barang : ")
	fmt.Scan(&nama)

	for _, b := range gudang {
		if b.Nama == nama {

			fmt.Println("\nBarang ditemukan")
			fmt.Println("Kode  :", b.Kode)
			fmt.Println("Nama  :", b.Nama)
			fmt.Println("Stok  :", b.Stok)
			fmt.Println("Harga :", b.Harga)

			return
		}
	}

	fmt.Println("Barang tidak ditemukan.")
}

// ================= URUTKAN KODE BARANG =================
func urutkanKodeBarang() {
	n := len(gudang)

	for i := 0; i < n-1; i++ {

		min := i

		for j := i + 1; j < n; j++ {
			if gudang[j].Kode < gudang[min].Kode {
				min = j
			}
		}

		gudang[i], gudang[min] = gudang[min], gudang[i]
	}
}

// ================= CARI BERDASARKAN KODE =================
func cariKodeBarang() {
	var kode string

	urutkanKodeBarang()

	fmt.Print("Masukkan kode barang : ")
	fmt.Scan(&kode)

	kiri := 0
	kanan := len(gudang) - 1

	for kiri <= kanan {

		tengah := (kiri + kanan) / 2

		if gudang[tengah].Kode == kode {

			fmt.Println("\nBarang ditemukan")
			fmt.Println("Kode  :", gudang[tengah].Kode)
			fmt.Println("Nama  :", gudang[tengah].Nama)
			fmt.Println("Stok  :", gudang[tengah].Stok)
			fmt.Println("Harga :", gudang[tengah].Harga)

			return
		}

		if gudang[tengah].Kode < kode {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	fmt.Println("Barang tidak ditemukan.")
}

// ================= URUTKAN STOK TERKECIL =================
func urutkanStokTerkecil() {
	n := len(gudang)

	for i := 0; i < n-1; i++ {

		min := i

		for j := i + 1; j < n; j++ {
			if gudang[j].Stok < gudang[min].Stok {
				min = j
			}
		}

		gudang[i], gudang[min] = gudang[min], gudang[i]
	}

	fmt.Println("Data berhasil diurutkan dari stok terkecil.")
}

// ================= URUTKAN STOK TERBESAR =================
func urutkanStokTerbesar() {

	for i := 1; i < len(gudang); i++ {

		key := gudang[i]
		j := i - 1

		for j >= 0 && gudang[j].Stok < key.Stok {
			gudang[j+1] = gudang[j]
			j--
		}

		gudang[j+1] = key
	}

	fmt.Println("Data berhasil diurutkan dari stok terbesar.")
}

// ================= STATISTIK GUDANG =================
func statistikGudang() {

	if len(gudang) == 0 {
		fmt.Println("Data barang kosong.")
		return
	}

	totalNilai := 0
	barangSedikit := gudang[0]

	for _, b := range gudang {

		totalNilai += b.Stok * b.Harga

		if b.Stok < barangSedikit.Stok {
			barangSedikit = b
		}
	}

	fmt.Println("\n===== STATISTIK TOTAL NILAI GUDANG =====")
	fmt.Println("Total Nilai Gudang :", totalNilai)

	fmt.Println("\nBarang dengan stok paling sedikit")
	fmt.Println("Kode :", barangSedikit.Kode)
	fmt.Println("Nama :", barangSedikit.Nama)
	fmt.Println("Stok :", barangSedikit.Stok)
}

// ================= MENU =================
func menu() {

	var pilih int

	for {

		fmt.Println("\n===== SISTEM INVENTARIS GUDANG =====")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Tampilkan Barang")
		fmt.Println("3. Ubah Barang")
		fmt.Println("4. Hapus Barang")
		fmt.Println("5. Barang Masuk")
		fmt.Println("6. Barang Keluar")
		fmt.Println("7. Mencari Data Barang Berdasarkan Nama")
		fmt.Println("8. Mencari Data Barang Berdasarkan Kode")
		fmt.Println("9. Mengurutkan Data Stok dari Jumlah Terkecil")
		fmt.Println("10. Mengurutkan Data Stok dari Jumlah Terbesar")
		fmt.Println("11. Menampilkan Statistik Total Nilai Aset Gudang")
		fmt.Println("0. Keluar")

		fmt.Print("Pilih menu : ")
		fmt.Scan(&pilih)

		switch pilih {

		case 1:
			tambahBarang()

		case 2:
			tampilkanBarang()

		case 3:
			ubahBarang()

		case 4:
			hapusBarang()

		case 5:
			barangMasuk()

		case 6:
			barangKeluar()

		case 7:
			cariNamaBarang()

		case 8:
			cariKodeBarang()

		case 9:
			urutkanStokTerkecil()

		case 10:
			urutkanStokTerbesar()

		case 11:
			statistikGudang()

		case 0:
			fmt.Println("Program selesai.")
			return

		default:
			fmt.Println("Pilihan tidak tersedia.")
		}
	}
}

// ================= MAIN =================
func main() {
	menu()
}

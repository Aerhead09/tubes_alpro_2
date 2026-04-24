/*
	TUBES NO.21: Aplikasi Pengelolaan Data Crowdfunding Sederhana
	Deskripsi:
	Aplikasi ini memungkinkan pengguna untuk melihat dan mencatat kontribusi terhadap proyek crowdfunding.
	Data utama yang digunakan  adalah daftar proyek, target dana, dan jumlah kontribusi.
	Pengguna aplikasi adalah individu yang ingin berpartisipasi dalam proyek crowdfunding atau pemilik proyek yang mencari pendanaan.
	Spesifikasi:
	a. Pengguna dapat menambahkan, mengubah, dan menghapus proyek crowdfunding yang mereka buat.
	b. Sistem mencatat jumlah dana yang terkumpul dan jumlah donatur.
	c. Pengguna dapat mencari proyek berdasarkan nama atau kategori menggunakan Sequential dan Binary Search.
	d. Pengguna dapat mengurutkan proyek berdasarkan total dana terkumpul atau jumlah donatur menggunakan Selection dan Insertion Sort.
	e. Sistem menampilkan daftar proyek yang sudah mencapai target pendanaan.

	Nama Anggota:
	Jean Yudhistira Diva Waluyo  103012400113
	Rifky Syaputra  103012400113

	Apa yang sudah ditambahkan:
	Menambahkan fungsi binary search, dimana fungsi binary search ini akan dipakai untuk melakukan pencarian kontributor berdasarkan anggaran yang sudah di sorting
	pada menu sorting, sehingga pengguna bisa mencari tau kotributor dari anggaran nya. fungsi binary search nya sendiri akan mengecek apakah datanya disorting secara ascending
	maupun decending, sehingga tidak melakukan perhitungan yang salah. selain dari itu fungsinya sama seperti fungsi binary search pada umumnya. Selain itu juga mengganti
	contanta NMAX yang awalnya 100 menjadi 1234 karena dianggap terlalu kecil. Untuk fitur lain tidak ada yang mau ditambahkan, karena menurut kami ini sudah memenuhi
	spesifikasi yang diminta serta memenuhi semua materi yang telah dipelajari.
*/

package main

import (
	"fmt"
)

const NMAX int = 1234

/*struct untuk memanggil tabel dengan field ID, nama proyek, anggaran proyek, dan banyak kontributor*/
type proyek struct {
	ID              string
	NAMA_PROYEK     string
	ANGGARAN_PROYEK int
	KONTRIBUTOR     int
}

/*type untuk memanggil struct proyek yang dibatasi array sebanyak Nmax, dimana array cuma bisa menampung 100 tabel*/
type tabInt [NMAX]proyek

/*prosedur untuk menampikan menu sesuai dengan namanya masing - masing */
func menuWelcome() {
	fmt.Println("-------------------------------------------------")
	fmt.Println("          SELAMAT DATANG DI APLIKASI          ")
	fmt.Println("   LIST DATA CROWDFUNDING UNTUK*KITA*BERSAMA     ")
}

func menuUtama() {
	fmt.Println("----------------------------------------------")
	fmt.Println("ADA YANG BISA KAMI BANTU? ")
	fmt.Println("1. SAYA INGIN MELIHAT DATA CROWDFUNDING")
	fmt.Println("2. SAYA INGIN MENGATUR DATA CROWDFUNDING")
	fmt.Println("0. SAYA INGIN KELUAR DARI APLIKASI")
	fmt.Print("KETIK PILIHAN ANDA DISINI (0-2): ")
}

func menuGoodbye() {
	fmt.Println("----------------------------------------------")
	fmt.Println("BAIKLAH, TERIMA KASIH SUDAH MENGGUNAKAN APLIKASI KAMI")
	fmt.Println("SEE YOU NEXT TIME, COMRADE")
}

func menuMengaturData() {
	fmt.Println("-------------------------------------------------------------------")
	fmt.Println("ANDA INGIN MENGATUR DATA CROWDFUNDING? APA YANG INGIN ANDA LAKUKAN?")
	fmt.Println("1.MENAMBAHKAN PROYEK")
	fmt.Println("2.MENGUBAH PROYEK")
	fmt.Println("3.MENGHAPUS PROYEK")
	fmt.Println("4.RESET SEMUA DATA")
	fmt.Println("0.KELUAR DARI MENU")
	fmt.Print("KETIK PILIHAN ANDA DISINI (0-4): ")
}

func menuMengubahProyek() {
	fmt.Println("--------------------------")
	fmt.Println("APA YANG INGIN ANDA GANTI?")
	fmt.Println("1.NAMA PROYEK")
	fmt.Println("2.ANGGARAN PROYEK")
	fmt.Println("3.JUMLAH KONTRIBUTOR PROYEK")
	fmt.Println("0.TIDAK INGIN MENGGANTI APA - APA / KELUAR DARI MENU")
	fmt.Print("KETIK PILIHAN ANDA DISINI (0-3): ")
}

func menuMenghapusProyek() {
	fmt.Println("------------------------------------------")
	fmt.Println("APA ANDA YAKIN UNTUK MENGHAPUS PROYEK INI?")
	fmt.Print("KETIK PILIHAN ANDA DISINI (YA/TIDAK): ")
}

func menuResetProyek() {
	fmt.Println("--------------------------------------------")
	fmt.Println("APAKAH ANDA YAKIN UNTUK MERESET DATA PROYEK?")
	fmt.Println("KETIK PILIHAN ANDA DISINI (YA/TIDAK): ")
}

func menuMelihatData() {
	fmt.Println("-------------------------------------------------------------------")
	fmt.Println("INI ADALAH DATA PROYEK YANG TERSEDIA. APA YANG INGIN ANDA LAKUKAN?")
	fmt.Println("1.SORTING LIST DATA DARI DANA TERBANYAK HINGGA TERKECIL DAN TAMPILKAN")
	fmt.Println("2.SORTING LIST DATA DARI DANA TERKECIL HINGGA TERBANYAK DAN TAMPILKAN")
	fmt.Println("3.MENCARI KONTRIBUTOR BERDASARKAN ANGGARAN DARI DATA YANG SUDAH DI SORTING")
	fmt.Println("0.KELUAR DARI MENU")
	fmt.Print("KETIK PILIHAN ANDA DISINI (0-3): ")
}

/*prosedur untuk mengurutkan array secara decending menggunakan insertion sort*/
func decendingInsertionSort(A *tabInt, jproyek int) {
	var i, j int
	var key proyek
	for i = 1; i < jproyek; i++ {
		key = A[i]
		j = i - 1
		for j >= 0 && key.ANGGARAN_PROYEK > A[j].ANGGARAN_PROYEK {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = key
	}
}

/*prosedur untuk mengurutkan array secara decending menggunakan insertion sort*/
func ascendingSelectionSort(A *tabInt, jproyek int) {
	var i, j, min int
	var temp proyek
	for i = 0; i < jproyek-1; i++ {
		min = i
		for j = i + 1; j < jproyek; j++ {
			if A[j].ANGGARAN_PROYEK < A[min].ANGGARAN_PROYEK {
				min = j
			}
		}
		temp = A[i]
		A[i] = A[min]
		A[min] = temp
	}
}

/*fungsi untuk mencari kontributor berdasarkan anggaran yang diinput*/
func binarySearchKontributor(A tabInt, jproyek int, anggaranTarget int) int {
	var left, right, middle int
	var found bool
	var ascending bool
	left = 0
	right = jproyek - 1
	found = false
	ascending = false

	if A[left].ANGGARAN_PROYEK < A[right].ANGGARAN_PROYEK {
		ascending = true
	}

	if ascending {
		for left <= right && !found {
			middle = (left + right) / 2
			if A[middle].ANGGARAN_PROYEK < anggaranTarget {
				left = middle + 1
			} else if A[middle].ANGGARAN_PROYEK > anggaranTarget {
				right = middle - 1
			} else {
				found = true
			}
		}
	} else {
		for left <= right && !found {
			middle = (left + right) / 2
			if A[middle].ANGGARAN_PROYEK > anggaranTarget {
				left = middle + 1
			} else if A[middle].ANGGARAN_PROYEK < anggaranTarget {
				right = middle - 1
			} else {
				found = true
			}
		}
	}

	if found {
		return A[middle].KONTRIBUTOR
	} else {
		return -1
	}
}

/*prosedur untuk print tabel array*/
func printDataProyek(A tabInt, jproyek int) {
	var i int
	fmt.Printf("|%6s|%30s|%15s|%12s|\n", "ID", "NAMA PROYEK", "ANGGARAN PROYEK", "KONTRIBUTOR")
	for i = 0; i < jproyek; i++ {
		fmt.Printf("|%6s|%30s|%15d|%12d|\n", A[i].ID, A[i].NAMA_PROYEK, A[i].ANGGARAN_PROYEK, A[i].KONTRIBUTOR)
	}
}

/*prosedur untuk scan tabel array*/
func scanDataProyek(A *tabInt, jProyek, n int) {
	var i int
	for i = jProyek; i < jProyek+n; i++ {
		fmt.Scan(&A[i].ID, &A[i].NAMA_PROYEK, &A[i].ANGGARAN_PROYEK, &A[i].KONTRIBUTOR)
	}
}

/*fungsi untuk searching idx dari array tabel yang ingin dicari berdasarkan ID */
func sequentialSearchIdx(A tabInt, jProjek int, idTarget string) int {
	var i int
	var idxTarget int
	idxTarget = -1
	i = 0
	for i < jProjek && idxTarget == -1 {
		if idTarget == A[i].ID {
			idxTarget = i
		}
		i++
	}
	return idxTarget
}

/*prosedur untuk menghapus array tabel sesuai dengan id yang dipilih*/
func hapusDataProyek(A *tabInt, jproyek *int, idxTarget int) {
	var i int
	if *jproyek > idxTarget {
		for i = idxTarget; i < *jproyek-1; i++ {
			A[i] = A[i+1]
		}
	}
	*jproyek = *jproyek - 1
}

func main() {
	var Data tabInt                                     /*variabel yang memanggil tipe data tabInt dimana tabInt memanggil tipe data proyek yang dimana proyek memiliki 4 field*/
	var choiceUtama, choiceMengatur, choiceMelihat int  /*variabel untuk memasukkan pilihan berdasarkan nomor di menu koresponden dengan namanya*/
	var validSort bool                                  /*variabel untuk mengecek status sorting dimana sorting harus dilakukan sehingga bisa mencari ID kontributor target*/
	var sortAnggaranPilihan int                         /*variabel untuk menyimpan anggaran yang ingin kita cari*/
	var choiceMengubahProyek int                        /*sama dengan variabel diatas*/
	var choiceMenghapusProyek, choiceResetProyek string /*sama dengan variabel diatas, namun yang dicek adalah string berdasarkan pilihan*/
	var validID bool                                    /*variabel untuk mengecek apakah ID yang kita input benar*/
	var idPilihan string                                /*variabel untuk memilih ID berdasarkan tabel*/
	var idxID int                                       /*variabel untuk mencatat index dari id pilihan*/
	var jumlahProyek, tambahProyek int                  /*jumlah proyek adalah jumlah data pada tabel, sedangkan tambah proyek digunakan untuk menambahkan jumlah proyek*/

	jumlahProyek = 0

	menuWelcome()
	menuUtama()
	fmt.Scan(&choiceUtama)
	for choiceUtama != 0 {
		switch choiceUtama {
		case 1:
			/*case 1 digunakan untuk memperlihatkan array tabel, serta memberikan opsi untuk melakukan sorting data*/
			if jumlahProyek == 0 {
				fmt.Println("MAAF, DATA MASIH KOSONG. MOHON DIISI DULU DATANYA")
			} else {
				validSort = false
				printDataProyek(Data, jumlahProyek)
				menuMelihatData()
				fmt.Scan(&choiceMelihat)
				for choiceMelihat != 0 {
					switch choiceMelihat {
					case 1:
						/*data diurutkan secara decending*/
						if jumlahProyek < 3 {
							fmt.Println("-----------------------------------------------------------------------------------------")
							fmt.Println("DATA ANDA TIDAK CUKUP BANYAK UNTUK DIURUTKAN. MINIMAL HARUS 3 DATA UNTUK BISA MENGURUTKAN")
							fmt.Println("KEMBALI KE MENU")
						} else {
							fmt.Println("--------------------------")
							fmt.Println("BAIKLAH DATA AKAN DISORTIR")
							decendingInsertionSort(&Data, jumlahProyek)
							fmt.Println("******************************************")
							fmt.Println("DATA SUDAH DISORTIR, DATA AKAN DITAMPILKAN")
							printDataProyek(Data, jumlahProyek)
							validSort = true
						}
					case 2:
						/*data diurutkan secara ascending*/
						if jumlahProyek < 3 {
							fmt.Println("-----------------------------------------------------------------------------------------")
							fmt.Println("DATA ANDA TIDAK CUKUP BANYAK UNTUK DIURUTKAN. MINIMAL HARUS 3 DATA UNTUK BISA MENGURUTKAN")
							fmt.Println("KEMBALI KE MENU")
						} else {
							fmt.Println("--------------------------")
							fmt.Println("BAIKLAH DATA AKAN DISORTIR")
							ascendingSelectionSort(&Data, jumlahProyek)
							fmt.Println("******************************************")
							fmt.Println("DATA SUDAH DISORTIR, DATA AKAN DITAMPILKAN")
							printDataProyek(Data, jumlahProyek)
							validSort = true
						}
					case 3:
						/*mencari kontributor berdasarkan anggaran*/
						if validSort {
							fmt.Println("-------------------------------")
							fmt.Print("MASUKKAN ANGGARAN YANG INGIN DICARI: ")
							fmt.Scan(&sortAnggaranPilihan)
							fmt.Println()
							fmt.Println("KONTRIBUTOR DARI ANGGARAN TARGET YANG DICARI(-1 KETIKA TIDAK ADA HASIL): ", binarySearchKontributor(Data, jumlahProyek, sortAnggaranPilihan))
							fmt.Println("KEMBALI KE MENU")
						} else {
							fmt.Println("--------------------------------------------------------")
							fmt.Println("DATA BELUM DISORTING, MOHON SORTING DATA TERLEBIH DAHULU")
						}
					default:
						fmt.Println("PILIHAN TIDAK VALID, MOHON PILIH SESUAI PILIHAN YANG TERSEDIA")
						fmt.Println("-------------------------------------------------------------")
					}
					menuMelihatData()
					fmt.Scan(&choiceMelihat)
				}
			}
		case 2:
			menuMengaturData()
			fmt.Scan(&choiceMengatur)
			for choiceMengatur != 0 {
				switch choiceMengatur {
				case 1:
					/*case 1 digunakan untuk menambahkan proyek sesuai dengan input*/
					fmt.Println("-----------------------------------")
					fmt.Println("BERAPA PROYEK YANG INGIN TAMBAHKAN?")
					fmt.Println("TULIS DISINI: ")
					fmt.Scan(&tambahProyek)
					fmt.Println("MASUKKAN ID PROYEK, NAMA PROYEK [UNTUK SPASI GUNAKAN LAMBANG '_'], ANGGARAN PROYEK, JUMLAH KONTRIBUTOR: ")
					scanDataProyek(&Data, jumlahProyek, tambahProyek)
					jumlahProyek += tambahProyek
					fmt.Println("BAIKLAH, DATA SUDAH DIMASUKKAN KE DALAM LIST. BALIK KE MENU SEBELUMNYA")
				case 2:
					/*case 2 digunakan untuk mencari sebuah array berdasarkan ID untuk diubah nama proyeknya, anggarannya, dan kontributornya*/
					if jumlahProyek == 0 {
						fmt.Println("---------------------------------------------------------------------------------------")
						fmt.Println("MAAF, BELUM ADA DATA YANG BISA DIGANTI, MOHON ISI DATA TERLEBIH DAHULU. KEMBALI KE MENU")
					} else {
						printDataProyek(Data, jumlahProyek)
						validID = false
						fmt.Println("---------------------------------")
						fmt.Print("ID YANG INGIN ANDA UBAH PROYEKNYA: ")
						fmt.Scan(&idPilihan)
						idxID = sequentialSearchIdx(Data, jumlahProyek, idPilihan)
						for validID == false {
							if idxID == -1 {
								fmt.Println("ID TIDAK ADA DI LIST. MOHON ISI LAGI")
								fmt.Println("---------------------------------")
								fmt.Print("ID YANG INGIN ANDA UBAH PROYEKNYA: ")
								fmt.Scan(&idPilihan)
								idxID = sequentialSearchIdx(Data, jumlahProyek, idPilihan)
							} else {
								validID = true
							}
						}
						menuMengubahProyek()
						fmt.Scan(&choiceMengubahProyek)
						for choiceMengubahProyek != 0 {
							switch choiceMengubahProyek {
							case 1:
								/*case 1 digunakan untuk mengganti nama proyek*/
								fmt.Println("-------------------")
								fmt.Println("GANTI NAMA PROYEK: ")
								fmt.Scan(&Data[idxID].NAMA_PROYEK)
								fmt.Println("NAMA PROYEK SUDAH DIGANTI. KEMBALI KE MENU")
							case 2:
								/*case 2 digunakan untuk mengganti anggaran proyek*/
								fmt.Println("-----------------------")
								fmt.Println("GANTI ANGGARAN PROYEK: ")
								fmt.Scan(&Data[idxID].ANGGARAN_PROYEK)
								fmt.Println("ANGGARAN PROYEK SUDAH DIGANTI. KEMBALI KE MENU")
							case 3:
								/*case 3 digunakan untuk mengganti banyak kontributor proyek*/
								fmt.Println("--------------------------")
								fmt.Println("GANTI BANYAK KONTRIBUTOR: ")
								fmt.Scan(&Data[idxID].KONTRIBUTOR)
								fmt.Println("BANYAK KONTRIBUTOR SUDAH DIGANTI. KEMBALI KE MENU")
							default:
								fmt.Println("PILIHAN TIDAK VALID, MOHON PILIH SESUAI PILIHAN YANG TERSEDIA")
								fmt.Println("-------------------------------------------------------------")
							}
							menuMengubahProyek()
							fmt.Scan(&choiceMengubahProyek)
						}
					}
				case 3:
					/*case 3 digunakan untuk menghapus 1 array tabel*/
					if jumlahProyek == 0 {
						fmt.Println("---------------------------------------------------------------------------------------")
						fmt.Println("MAAF, BELUM ADA DATA YANG BISA DIHAPUS, MOHON ISI DATA TERLEBIH DAHULU. KEMBALI KE MENU")
					} else {
						printDataProyek(Data, jumlahProyek)
						validID = false
						fmt.Println("---------------------------------")
						fmt.Print("ID YANG INGIN ANDA HAPUS PROYEKNYA: ")
						fmt.Scan(&idPilihan)
						idxID = sequentialSearchIdx(Data, jumlahProyek, idPilihan)
						for validID == false {
							if idxID == -1 {
								fmt.Println("ID TIDAK ADA DI LIST. MOHON ISI LAGI")
								fmt.Println("---------------------------------")
								fmt.Print("ID YANG INGIN ANDA HAPUS PROYEKNYA: ")
								fmt.Scan(&idPilihan)
								idxID = sequentialSearchIdx(Data, jumlahProyek, idPilihan)
							} else {
								validID = true
							}
						}
						menuMenghapusProyek()
						fmt.Scan(&choiceMenghapusProyek)
						if choiceMenghapusProyek == "YA" {
							hapusDataProyek(&Data, &jumlahProyek, idxID) /*prosedur yang digunakan untuk menghapus array tabel*/
							fmt.Println("DATA YANG ANDA PILIH SUDAH DIHAPUS. KEMBALI KE MENU")
						} else if choiceMenghapusProyek == "TIDAK" {
							fmt.Println("DATA YANG ANDA PILIH TIDAK JADI DIHAPUS. KEMBALI KE MENU")
						}
					}
				case 4:
					/*case 4 digunakan untuk mereset semua data tabel dengan cara mereset nilai jumlah proyek ke 0
					namun array secara langsung belum terhapus, namun array akan tertimpa dengan array baru ketika jumlah proyek diisi lagi*/
					if jumlahProyek == 0 {
						fmt.Println("---------------------------------------------------------------------------------------")
						fmt.Println("MAAF, BELUM ADA DATA YANG BISA DIHAPUS, MOHON ISI DATA TERLEBIH DAHULU. KEMBALI KE MENU")
					} else {
						menuResetProyek()
						fmt.Scan(&choiceResetProyek)
						if choiceResetProyek == "YA" {
							jumlahProyek = 0
							fmt.Println("DATA ANDA SUDAH DIRESET. KEMBALI KE MENU")
						} else if choiceResetProyek == "TIDAK" {
							fmt.Println("DATA ANDA TIDAK JADI DIRESET. KEMBALI KE MENU")
						}
					}
				default:
					fmt.Println("PILIHAN TIDAK VALID, MOHON PILIH SESUAI PILIHAN YANG TERSEDIA")
					fmt.Println("-------------------------------------------------------------")
				}
				menuMengaturData()
				fmt.Scan(&choiceMengatur)
			}
		default:
			fmt.Println("PILIHAN TIDAK VALID, MOHON PILIH SESUAI PILIHAN YANG TERSEDIA")
			fmt.Println("-------------------------------------------------------------")
		}
		menuUtama()
		fmt.Scan(&choiceUtama)
	}
	menuGoodbye()
}

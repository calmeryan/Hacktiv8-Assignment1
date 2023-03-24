package main

import (
	"fmt"
	"os"
)

type Teman struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {

	listTeman := [5]Teman{}
	listTeman[0] = Teman{nama: "Johanes", alamat: "Jl. Angkur 15 No. 6, Jakarta Utara", pekerjaan: "Programmer", alasan: "Tertarik dengan bahasa Go"}
	listTeman[1] = Teman{nama: "Ian", alamat: "Jl. Melati Blok C No. 16, Jakarta Timur", pekerjaan: "Guru", alasan: "Belajar bahasa pemrograman baru adalah hobi saya"}
	listTeman[2] = Teman{nama: "Ryan", alamat: "Jl. Pelangi Raya III No. 19, Jakarta Barat", pekerjaan: "Programmer", alasan: "Kebutuhan dari pekerjaan"}
	listTeman[3] = Teman{nama: "Raja", alamat: "Jl. Gading Timur A3 No. 25, Jakarta Selatan", pekerjaan: "Pelajar", alasan: "Mempersiapkan diri untuk perkuliahan di semester depan"}
	listTeman[4] = Teman{nama: "Nathan", alamat: "Jl. Maritim Abadi No. 3C, Jakarta Pusat", pekerjaan: "Programmer", alasan: "Memperluas skill yang dimiliki"}

	fmt.Println(cari(os.Args[1], listTeman))
}

func cari(namaTeman string, daftarTeman [5]Teman) string {
	for id, teman := range daftarTeman {
		if teman.nama == namaTeman {
			return fmt.Sprintf("ID : %d \nnama : %s, \nalamat : %s, \npekerjaan : %s, \nalasan : %s", id, teman.nama, teman.alamat, teman.pekerjaan, teman.alasan)
		}
	}
	return fmt.Sprintf("Tidak ditemukan teman dengan nama %s", namaTeman)
}

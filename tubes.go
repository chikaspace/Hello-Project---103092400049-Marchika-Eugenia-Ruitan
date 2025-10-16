package main

import "fmt"

const lim int = 1000

type data struct {
	nama, penyakit, waktu, paket string
}
type arBp [lim]data

func main() { // ini func mainnya kayaknya kepanjangan, belum tau ini udah bener bagian ini yang jadi func mainnya atau bukan
	var nambar, sandi1, sandi2, user1, user2 string
	var pilihan, nData int
	var p arBp

	fmt.Println("~~~~~Medical Check Up~~~~~")

	//ini buat awalnya sebelum masuk kan, tapi kemaren kata masnya kalo ga disuruh di kasusnya gausah dibuat gapapa
	fmt.Println("=== Daftar ===")
	fmt.Print("Masukkan Username: ")
	fmt.Scanln(&user1)
	fmt.Print("Masukkan Password: ")
	fmt.Scanln(&sandi1)

	fmt.Println("\n=== Masuk ===")
	fmt.Print("Username : ")
	fmt.Scanln(&user2)
	fmt.Print("Password : ")
	fmt.Scanln(&sandi2)

	fmt.Println()

	if masuk(sandi1, sandi2, user1, user2) {
		fmt.Println("~~Selamat Datang di Aplikasi Medical Check Up~~")
		for {
			fmt.Println()
			fmt.Println("1. Mengisi Data Pasien")
			fmt.Println("2. Mengedit Data Pasien")
			fmt.Println("3. Laporan Pemasukan Layanan MCU")
			fmt.Println("4. Rekap Data Pasien")
			fmt.Println("4. Daftar Pasien")
			fmt.Println("5. Keluar")
			fmt.Println()
			fmt.Print("Mau Apa Hari Ini (1/2/3/4/5) ? ")
			fmt.Scan(&pilihan)

			switch pilihan {
			case 1:
				fmt.Print("Masukan Jumlah Data Pasien : ")
				fmt.Scan(&nData)
				isiData(nData, &p)
				fmt.Println()

			case 2: // ini buat manggil si edit
				editData(&p, &nData, nambar)
			}

		}
	}
}

func masuk(sandi1, sandi2, user1, user2 string) bool { // ini buat username dan sandinya, aku juga ga nyangka dia sesingkat ini

	if sandi1 == sandi2 && user1 == user2 {
		fmt.Println("$$$$$$$$$$$$$$BERHASIL MASUK$$$$$$$$$$$$$$")
	} else {
		fmt.Println("xxxxxxxxxxxDITOLAKxxxxxxxxxxx")
	}
	return sandi1 == sandi2 && user1 == user2
}

func isiData(n int, p *arBp) {
	var m, i int
	m = 1

	for i = 0; i < n; i++ {

		fmt.Printf("Pasien %d\n", m) //ini sekedar tambahan aja sih biar keluarannay tuh ada nomernyaa gituuu, pasien 1, pasien 2 gituu
		fmt.Print("Nama          : ")
		fmt.Scan(&p[i].nama)
		fmt.Print("Penyakit      : ")
		fmt.Scan(&p[i].penyakit)
		fmt.Print("Waktu MCU     : ")
		fmt.Scan(&p[i].waktu)
		fmt.Print("Paket Layanan : ")
		fmt.Scan(&p[i].paket)
		fmt.Println()
		m++
	}

}

func editData(p *arBp, n *int, nambar string) {
	var cari int
	var edit string
	//var editap string
	fmt.Printf("Nama Pasien yang ingin di edit : ")
	fmt.Scan(&nambar)
	cari = cariEdit(*p, *n, nambar)

	if cari != -1 {
		fmt.Print("Data Pasien ditemukan\n")
		fmt.Printf("Nama          : %s\n", p[cari].nama)
		fmt.Printf("Penyakit      : %s\n", p[cari].penyakit)
		fmt.Printf("Waktu MCU     : %s\n", p[cari].waktu)
		fmt.Printf("Paket Layanan : %s\n", p[cari].paket)

		// bagian ini yang belum sesuai ekspektasi keluarnya soalnya hasil editannya masih belum muncull, keluarannya baru sampe di apa yang mau diedit gitu
		fmt.Print("Apa yang ingin anda ubah (nama/penyakit/waktu/paket) ?")
		fmt.Scan(&edit)
		switch edit {
		case "nama":
			fmt.Print("Nama : ")
			fmt.Scan(&nambar)
			p[cari].nama = nambar
		case "penyakit":
			fmt.Print("penyakit : ")
			fmt.Scan(&nambar)
			p[cari].penyakit = nambar
		case "waktu":
			fmt.Print("waktu : ")
			fmt.Scan(&nambar)
			p[cari].waktu = nambar
		case "paket":
			fmt.Print("paket : ")
			fmt.Scan(&nambar)
			p[cari].paket = nambar
		}

		fmt.Printf("Nama          : %s\n", p[cari].nama)
		fmt.Printf("Penyakit      : %s\n", p[cari].penyakit)
		fmt.Printf("Waktu MCU     : %s\n", p[cari].waktu)
		fmt.Printf("Paket Layanan : %s\n", p[cari].paket)

		fmt.Print("Data berhasil diedit")
	} else {
		fmt.Print("Pasien tidak ditemukan")

	}
}

func cariEdit(p arBp, n int, nam string) int {
	var i, cari int

	i = 0
	cari = -1
	for i < n && cari == -1 {
		if p[i].nama == nam {
			cari = i
		}
		i++
	}
	return cari

}

//ini aku ga sempet ubah ubah lagi, soalnya kemarin langsung tewas sampe dicariin satu keluarga soalnya aku ga aktif :), hehe
//soal koding lagi, aku tuh ngasal, aku kurang yakin mana yang harusnya prosedur, pokoknya yang penting "func" dulu ya kan;)
//tapi kayaknya udah cocok ga sih, kayaknya aku mulai nangkep mana prosedur mana fungsi

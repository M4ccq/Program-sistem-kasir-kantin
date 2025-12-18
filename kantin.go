// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	var menu int
	var jumlah int
	var total int
	var lanjut string
	var bayar int
	var kembalian int
	var totalasli int
	for {
		fmt.Println("=== KASIR KANTIN SEDERHANA ===")
		fmt.Println("1. Nasi Goreng (12000)")
		fmt.Println("2. Mie Ayam    (10000)")
		fmt.Println("3. Es Teh      (5000)")
		fmt.Println("4. Es Jeruk    (6000)")
		fmt.Println("===========================")
		fmt.Print("Pilih menu (1-4): ")
		fmt.Scan(&menu)

		if menu == 1 {
			total += 12000
			fmt.Println("Anda membeli Nasi Goreng")
			fmt.Println("===========================")
		} else if menu == 2 {
			total += 10000
			fmt.Println("Anda membeli Mie Ayam")
			fmt.Println("===========================")
		} else if menu == 3 {
			total += 5000
			fmt.Println("Anda membeli Es Teh")
			fmt.Println("===========================")
		} else if menu == 4 {
			total += 6000
			fmt.Println("Anda membeli Es Jeruk")
			fmt.Println("===========================")
		} else {
			fmt.Println("Menu tidak tersedia")
			fmt.Println("===========================")
		}

		fmt.Print("Jumlah beli: ")
		fmt.Scan(&jumlah)

		totalasli = total * jumlah

		fmt.Println("Total 		 :", totalasli)

		fmt.Print("Tambah pesanan? (y/n): ")
		fmt.Scan(&lanjut)

		if lanjut != "y" && lanjut != "Y" {
			break
		}

		fmt.Println()
	}
	for {
		fmt.Println("===========================")
		fmt.Println("Total:", totalasli)

		fmt.Print("Bayar: ")
		fmt.Scan(&bayar)

		kembalian = bayar - totalasli
		if kembalian >= 1 {
			fmt.Println("Kembalian	:", kembalian)
			fmt.Println("TERIMAKASIH TELAH BELANJA")
			break
		} else if kembalian == 0 {
			fmt.Println("UANG ANDA PAS")
			fmt.Println("TERIMAKASIH TELAH BELANJA")
			break
		} else if bayar <= totalasli {
			fmt.Println("Jumlah tidak sesuai")

		} else {
			fmt.Println("ERROR")
			break
		}
	}

}

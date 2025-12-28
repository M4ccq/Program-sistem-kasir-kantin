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
	totalasli = 12000 * jumlah
	for {
		for {
			fmt.Println("=== KASIR KANTIN SEDERHANA ===")
			fmt.Println("1. Nasi Goreng (12000)")
			fmt.Println("2. Mie Ayam    (10000)")
			fmt.Println("3. Es Teh      (5000)")
			fmt.Println("4. Es Jeruk    (6000)")
			fmt.Println("===========================")
			fmt.Print("Pilih menu (1-4): ")
			fmt.Scan(&menu)
			fmt.Print("Jumlah beli: ")
			fmt.Scan(&jumlah)
			fmt.Println("===========================")

			if menu == 1 {
				total += 12000 * jumlah
				fmt.Println("Anda membeli", jumlah, "Nasi Goreng")
				fmt.Println("===========================")
				break
			} else if menu == 2 {
				total += 10000 * jumlah
				fmt.Println("Anda membeli Mie Ayam")
				fmt.Println("===========================")
				break
			} else if menu == 3 {
				total += 5000 * jumlah
				fmt.Println("Anda membeli Es Teh")
				fmt.Println("===========================")
				break
			} else if menu == 4 {
				total += 6000 * jumlah
				fmt.Println("Anda membeli Es Jeruk")
				fmt.Println("===========================")
				break
			} else {
				fmt.Println("Menu tidak tersedia")
				fmt.Println("===========================")
			}
		}

		if jumlah > 0 {
			totalasli = total
			fmt.Println("Total 		 :", totalasli)

			fmt.Print("Tambah pesanan? (y/n): ")
			fmt.Scan(&lanjut)

			if lanjut == "n" || lanjut == "N" {
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
				break
			} else if lanjut == "y" || lanjut == " Y" {

			} else {
				fmt.Println("===========================")
				fmt.Println("Pilih hanya antara Y atau N")
				total = 0
				fmt.Println("HARGA TOTAL DIRISET KE 0 RUPIAH")
				fmt.Println("===========================")
			}

		} else {
			total = 0
			fmt.Println("TERIMAKASIH TELAH BERKUNJUNG")
			fmt.Println("===========================")
		}

	}
}


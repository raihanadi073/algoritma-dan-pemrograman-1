package main

import "fmt"

func main() {
	var angka int
	fmt.Print("masukkan angka : ")
	fmt.Scan(&angka)
	// Menghasilkan false jika bilangan ganjil, true jika genap
	fmt.Println(angka%2 == 0)
}

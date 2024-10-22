package main

import "fmt"

func hitungSkor() (int, int) {
    jumlahDiselesaikan := 0
    totalWaktu := 0
    var waktu int

    for i := 0; i < 8; i++ {
        fmt.Scan(&waktu)
        if waktu < 301 { 
            jumlahDiselesaikan++
            totalWaktu += waktu
        }
    }

    return jumlahDiselesaikan, totalWaktu
}

func main() {
    var nama1, nama2 string
    var jumlahDiselesaikan1, totalWaktu1, jumlahDiselesaikan2, totalWaktu2 int

    fmt.Scan(&nama1)
    jumlahDiselesaikan1, totalWaktu1 = hitungSkor()

    fmt.Scan(&nama2)
    jumlahDiselesaikan2, totalWaktu2 = hitungSkor()

    if jumlahDiselesaikan1 > jumlahDiselesaikan2 || (jumlahDiselesaikan1 == jumlahDiselesaikan2 && totalWaktu1 < totalWaktu2) {
		fmt.Printf("%s %d %d\n", nama1, jumlahDiselesaikan1, totalWaktu1)
	} else {
		fmt.Printf("%s %d %d\n", nama2, jumlahDiselesaikan2, totalWaktu2)
	}
	
}

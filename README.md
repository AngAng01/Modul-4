# MODUL EMPAT
 ## SOAL 1
   Program di atas adalah program yang menghitung permutasi dan kombinasi dari dua pasang bilangan yang dimasukkan oleh pengguna.
   
   ## Overview
      Program ini terdiri dari satu file bernama 'main.go' dan mencakup komponen-komponen utama berikut:
      - Pernyataan 'package main', yang mendefinisikan paket untuk program yang dapat dieksekusi.
      - Pernyataan 'import', yang digunakan untuk menyertakan paket-paket yang diperlukan (dalam hal ini, 'fmt').
      - Fungsi 'main()', yang merupakan titik awal dari setiap program Go.
      - Fungsi 'factorial(n int, hasil *int)', Fungsi ini menghitung faktorial dari bilangan n dan menyimpan hasilnya dalam variabel yang ditunjuk oleh pointer hasil. Faktorial dihitung menggunakan loop dari n hingga 1.
      - Fungsi 'permutasi(n, r int) int', Fungsi ini menghitung permutasi dari n elemen yang diambil sebanyak r. Rumus yang digunakan adalah P(n,r) = n!/(n-r)!, dimana faktorial dari n dan n-r dihitung menggunakan fungsi factorial.
      - Fungsi 'kombinasi(n, r int) int', Fungsi ini menghitung kombinasi dari n elemen yang diambil sebanyak r. Rumus yang digunakan adalah C(n,r) = n!/r!(n-r)!, di mana faktorial dari n, r, dan n-r dihitung menggunakan fungsi factorial.
      
   ## Code Explanation
      ```go
          func factorial(n int, hasil *int) {
	            *hasil = 1 
	            for i := n; i >= 1; i-- {
		             *hasil *= i
	            }
          }
      ```
   Kode di atas adalah kode fungsi yang digunakan untuk menghitung faktorial dari sebuah bilangan bulat n. 
   
   ```go
      func permutasi(n, r int) int {
      	var pembilang, penyebut int
      	factorial(n, &pembilang)
      	factorial(n-r, &penyebut)
      	return pembilang / penyebut
      }
   ```
   Kode di atas adalah Fungsi permutasi yang digunakan untuk menghitung permutasi dari n elemen yang diambil sebanyak r. Fungsi ini pertama-tama menghitung faktorial dari n (disimpan dalam variabel pembilang) dan faktorial dari n - r (disimpan dalam variabel penyebut) dengan memanfaatkan fungsi factorial. 
   
   ```go
      func kombinasi(n, r int) int {
      	var pembilang, penyebut1, penyebut2 int
      	factorial(n, &pembilang)               
      	factorial(r, &penyebut1)               
      	factorial(n-r, &penyebut2)             
      	return pembilang / (penyebut1 * penyebut2) 
      }
   ```
   Kode di atas adalah fungsi yang digunakan untuk menghitung kombinasi dari n elemen yang diambil sebanyak r. 
   
   ```go
      var a, b, c, d int
   ```
   kode diatas adalah kode yang digunakan untuk mendeklarasikan 4 variabel int yang digunakan sebagai inputan.
   
   ```go
      fmt.Scan(&a, &b, &c, &d)
   ```
   kode diatas adalah kode yang digunakan untuk pengguna menginputkan suatu nilai untuk dieksekusi.
   
   ```go
      fmt.Println(permutasi(a, c), kombinasi(a, c))
	  fmt.Println(permutasi(b, d), kombinasi(b, d))
   ```
   Kode di atas adalah bagian dari fungsi main() dalam program Go yang berfungsi untuk mencetak hasil perhitungan permutasi dan kombinasi dari dua pasangan bilangan (a, c dan b, d) menggunakan fungsi fmt.Println().
   
  
  ## SOAL 2
   Program di atas adalah program yang digunakan untuk menghitung dan membandingkan skor dari dua peserta berdasarkan jumlah tugas yang diselesaikan dan total waktu yang digunakan dalam menyelesaikan tugas tersebut. 
   
   ## Overview
      Program ini terdiri dari satu file bernama 'main.go' dan mencakup komponen-komponen utama berikut:
      - Pernyataan 'package main', yang mendefinisikan paket untuk program yang dapat dieksekusi.
      - Pernyataan 'import', yang digunakan untuk menyertakan paket-paket yang diperlukan (dalam hal ini, 'fmt').
      - Fungsi 'main()', yang merupakan titik awal dari setiap program Go.
      - Fungsi 'hitungSkor()', digunakan untuk menghitung jumlah tugas yang diselesaikan oleh peserta dan total waktu yang dibutuhkan untuk menyelesaikan tugas-tugas tersebut.
      - Program membandingkan jumlah tugas yang diselesaikan oleh masing-masing peserta. Peserta dengan jumlah tugas lebih banyak dinyatakan menang.
      
   ## Code Explanation
      ```go
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
      ```
   Kode di atas adalah fungsi yang bernama hitungSkor(), yang menghitung jumlah tugas yang diselesaikan dan total waktu yang digunakan oleh seorang peserta. Fungsi ini menginisialisasi dua variabel, jumlahDiselesaikan dan totalWaktu, kemudian menggunakan loop untuk meminta input waktu selama 8 iterasi. Jika waktu untuk suatu tugas kurang dari 301 detik, maka tugas tersebut dianggap selesai, dan jumlahDiselesaikan serta totalWaktu diperbarui. Setelah itu, fungsi mengembalikan jumlah tugas yang diselesaikan dan total waktu yang dihabiskan.
   
   ```go
      var nama1, nama2 string
      var jumlahDiselesaikan1, totalWaktu1, jumlahDiselesaikan2, totalWaktu2 int
   ```
   Kode di atas adalah deklarasi beberapa variabel, 'nama1' dan 'nama2' bertipe string untuk menyimpan nama dua peserta, serta 'jumlahDiselesaikan1', 'totalWaktu1', 'jumlahDiselesaikan2', dan 'totalWaktu2' bertipe int untuk menyimpan jumlah tugas yang diselesaikan dan total waktu yang dihabiskan oleh masing-masing peserta. Variabel-variabel ini digunakan untuk membandingkan kinerja kedua peserta.
   
   ```go
      fmt.Scan(&nama1)
      jumlahDiselesaikan1, totalWaktu1 = hitungSkor()
   ```
   Kode di atas digunakan untuk menerima input nama peserta pertama dengan fmt.Scan(&nama1) dan kemudian memanggil fungsi hitungSkor() untuk menghitung jumlah tugas yang diselesaikan dan total waktu peserta pertama, menyimpan hasilnya dalam jumlahDiselesaikan1 dan totalWaktu1.
   
   ```go
      fmt.Scan(&nama2)
      jumlahDiselesaikan2, totalWaktu2 = hitungSkor()
   ```
   Kode di atas berfungsi untuk menerima input nama peserta kedua dengan fmt.Scan(&nama2) dan kemudian memanggil fungsi hitungSkor() untuk menghitung jumlah tugas yang diselesaikan dan total waktu peserta kedua, menyimpan hasilnya dalam jumlahDiselesaikan2 dan totalWaktu2.
   
   ```go
      if jumlahDiselesaikan1 > jumlahDiselesaikan2 || (jumlahDiselesaikan1 == jumlahDiselesaikan2 && totalWaktu1 < totalWaktu2) {
        fmt.Printf("%s %d %d\n", nama1, jumlahDiselesaikan1, totalWaktu1)
      } else {
        fmt.Printf("%s %d %d\n", nama2, jumlahDiselesaikan2, totalWaktu2)
      }
   ```
   kode diatas adalah Kode membandingkan skor dua peserta. Jika peserta pertama memiliki lebih banyak tugas atau waktu lebih sedikit dengan jumlah tugas yang sama, namanya ditampilkan. jika tidak, nama peserta kedua yang ditampilkan.


   ## SOAL 3
   Program di atas adalah program yang mencetak deret bilangan menurut aturan Collatz. Fungsi cetakDeret(n int) menerima input n dan mencetak nilai n hingga mencapai 1, membagi n dengan dua jika genap dan mengubahnya menjadi 3n + 1 jika ganjil. 
   
   ## Overview
      Program ini terdiri dari satu file bernama 'main.go' dan mencakup komponen-komponen utama berikut:
      - Pernyataan 'package main', yang mendefinisikan paket untuk program yang dapat dieksekusi.
      - Pernyataan 'import', yang digunakan untuk menyertakan paket-paket yang diperlukan (dalam hal ini, 'fmt').
      - Fungsi 'main()', yang merupakan titik awal dari setiap program Go.
      - Fungsi cetakDeret(n int), Fungsi ini menerima integer n dan mencetak deret bilangan berdasarkan aturan Collatz hingga n mencapai 1, dengan memproses n sebagai genap atau ganjil.
      - Struktur Perulangan, Menggunakan loop for untuk terus menjalankan proses hingga n sama dengan 1, mencetak setiap nilai n selama proses.
      - Kondisi Genap dan Ganjil, Menggunakan kondisi if untuk memeriksa apakah n genap atau ganjil, kemudian mengubah nilai n sesuai aturan: membagi dua jika genap dan menghitung 3n + 1 jika ganjil
      
   ## Code Explanation
      ```go
          func cetakDeret(n int) {
              for n != 1 {
                  fmt.Printf("%d ", n) 
                  if n%2 == 0 {
                      n = n / 2 
                  } else {
                      n = 3*n + 1 
                  }
              }
              fmt.Println(n) 
          }
      ```
  Kode di atas adalah fungsi dalam bahasa Go bernama cetakDeret(n int), yang mencetak deret bilangan berdasarkan aturan Collatz. Fungsi ini menerima parameter integer n dan menggunakan loop for untuk mencetak nilai n hingga mencapai 1. Dalam setiap iterasi, fungsi memeriksa apakah n genap atau ganjil. Jika n genap, nilainya dibagi dua; jika ganjil, nilainya diubah menjadi 3n + 1. Proses ini terus berlanjut sampai n sama dengan 1, setelah itu nilai 1 dicetak dan fungsi selesai.
   
   ```go
      var n int
   ```
   kode diatas adalah kode yang digunakan untuk mendeklarasikan 1 variabel int yang digunakan sebagai inputan.
   
   ```go
      fmt.Scan(&n)
   ```
   kode diatas adalah kode yang digunakan untuk pengguna menginputkan suatu nilai untuk dieksekusi.
   
   ```go
      cetakDeret(n)
   ```
   Kode di atas adalah pemanggilan fungsi cetakDeret(n), yang digunakan untuk menjalankan fungsi tersebut dengan argumen n.
   
  
   

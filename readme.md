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
            func factorial(n int) int {
                if n == 0 || n == 1 {
                    return 1
                }
                result := 1
                for i := 2; i <= n; i++ {
                    result *= i
                }
                return result
            }
      ```
   Kode di atas adalah implementasi fungsi factorial dalam bahasa Go yang digunakan untuk menghitung faktorial dari suatu bilangan bulat positif. Faktorial dari sebuah bilangan n, ditulis sebagai n!, didefinisikan sebagai hasil perkalian semua bilangan bulat positif dari 1 hingga n.
   
   ```go
      func permutation(a, c int) int {
          return factorial(a) / factorial(a-c)
      }
   ```
   Kode di atas adalah implementasi fungsi permutation yang digunakan untuk menghitung permutasi dari dua bilangan a dan c. Permutasi digunakan untuk menghitung berapa banyak cara untuk mengatur c objek yang dipilih dari a objek secara berurutan.

   ```go
      func combination(a, c int) int {
          return factorial(a) / (factorial(c) * factorial(a-c))
      }
   ```
   Kode di atas adalah implementasi fungsi combination yang digunakan untuk menghitung kombinasi dari dua bilangan a dan c. Kombinasi menghitung berapa banyak cara untuk memilih c objek dari a objek tanpa memperhatikan urutan.

   ```go
      var a, b, c, d int
   ```
   kode diatas adalah kode yang digunakan untuk mendeklarasikan 4 variabel int yang digunakan sebagai inputan.

   ```go
      fmt.Println("Masukkan 4 angka (a, b, c, d): ")
      fmt.Scan(&a, &b, &c, &d)
   ```
   kode diatas adalah kode yang digunakan untuk menampilkan pesan kepada pengguna agar pengguna menginputkan suatu nilai untuk dieksekusi.

   ```go
      permAC := permutation(a, c)
      combAC := combination(a, c)
  
      permBD := permutation(b, d)
      combBD := combination(b, d)
   ```
   Kode di atas adalah bagian kode yang digunakan untuk menghitung permutasi dan kombinasi untuk dua pasang bilangan input, yaitu (a, c) dan (b, d), dengan memanfaatkan fungsi permutation dan combination.

   ```go
      fmt.Printf("P(%d, %d) = %d\n", a, c, permAC)
      fmt.Printf("C(%d, %d) = %d\n", a, c, combAC)
      fmt.Printf("P(%d, %d) = %d\n", b, d, permBD)
      fmt.Printf("C(%d, %d) = %d\n", b, d, combBD)
   ```
   Kode di atas digunakan untuk menampilkan hasil perhitungan permutasi dan kombinasi ke layar dengan format tertentu. Fungsi fmt.Printf digunakan untuk mencetak string yang diformat.


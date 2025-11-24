package main

import (
	"fmt"

)

func main() {
	// (ctrl + /)inline comentar

	/* (ctrl + shift + A)
	komentar panjang
	*/

	/* Program pertama, menampilkan
	Halo Dunia */
	fmt.Printf("Halo Dunia \n")

	// A.9. Variabel
	/* A.9.1. Deklarasi Variabel Beserta Tipe Data */
	var firstName1 string = "john"
	var lastName2 string = "wick"

	fmt.Printf("haloo %s %s!\n", firstName1, lastName2)

	/* A.9.2. Deklarasi Variabel Menggunakan Keyword  */

	var lastName1 string = "Tor "
	var firstName2 string = "john"
	fmt.Printf("haloo %s %s!\n", firstName2, lastName1+"87")

	/* A.9.3. Deklarasi Variabel Tanpa Tipe Data */
	//  %s untuk memanggil variabel yang dibuat
	var firstName string = "John"
	lastName := "Wick"    /* diubah menjadi "hutabara" */
	lastName = "Hutabara" /* dibuah menjadi p */
	lastName = "p"        /* output terakhit yang keluar */

	fmt.Printf("halo %s %s! \n", firstName, lastName)

	/* A.9.4. Deklarasi Multi Variabel */
		var first, second, third string
		first, second, third = "satu", "dua", "tiga"

		var four, five, six string = "empat", "lima", "enam"

		seventh, eight, ninth := "tujuh", "delapan", "sembilan"

		one, isFriday, twoPointTwo, say := 1, true, 2.2, "Halo"

		fmt.Println(first, second, third, four, five, six, seventh, eight, ninth)
		fmt.Printf("%v %v %v %v \n", one, isFriday, twoPointTwo, say)
	/*
		- tidak harus mengginakan "%s" untuk memanggil value,
			Itu hanya diperlukan saat menggunakan fmt.Printf atau fmt.Sprintf,
			karena kedua fungsi itu memakai format string.
		- %v adalah verb "default" di package fmt — mencetak nilai dalam bentuk wajar
			sesuai tipenya. Cocok ketika kamu tidak ingin menentukan format khusus.
	*/

	/* A.9.5. Variabel Underscore (_) */
	_ = "Belajar Golang"   // variabel underscore untuk menghindari error variabel tidak terpakai
	_ = "Golang Itu Mudah" //isi variabel underscore tidak dapat ditampilkan. Data yang sudah masuk variabel tersebut akan hilang.
	nama, _ := "John", "Wick"
	fmt.Println(nama) // variabel underscore untuk menghindari error variabel tidak terpakai

	/* A.9.6. Deklarasi Variabel Menggunakan Keyword (new) */
		name := new(string)

		fmt.Println(name)  // output: 0xc000010230 (alamat memory)
		fmt.Println(*name) // output: "" (nilai default tipe data string adalah "")


	/* A.10. Tipe Data */
	/* A.10.1. Tipe Data Numerik Non-Desimal */
		/* 
		- uint, tipe data untuk bilangan cacah (bilangan positif).
		- int, tipe data untuk bilangan bulat (bilangan negatif dan positif).
		*/

		var positiveNumber uint8 = 89
		var negativeNumber = -1243423644 

		fmt.Printf("bilangan positif : %d \n",positiveNumber)
		fmt.Printf("bilangan negatif : %d \n",negativeNumber)

	/* A.10.2. Tipe Data Numerik Desimal */
	 var decimalNumber = 2.62

	 fmt.Printf("bilangan desimal : %f \n", decimalNumber) 
	 // %f untuk memanggil tipe data desimal
	 
	 fmt.Printf("bilangan desimal : %.3f \n", decimalNumber) 
	 // %.3f untuk membatasi angka dibelakang koma menjadi 3 angka

	 /* A.10.3. Tipe Data bool (Boolean) */
	 var exist bool = true 
	 // tipe data boolean hanya memiliki dua nilai yaitu true dan false

	 fmt.Printf("Exist? %t \n", exist)
	  // %t untuk memanggil tipe data boolean


	/* A.10.4. Tipe Data string */  
		var message string = "error" 
		// tipe data string untuk menyimpan karakter atau teks

		fmt.Printf("pesan saat ini = %s \n", message) 
		// %s untuk memanggil tipe data string

		var longmessage string = `Nama Saya "John Wick" 
		Salam Kenal. 
		Mari belajar "Golang"` 
		// tipe data string dengan backtick untuk membuat string multi line

		fmt.Println(longmessage) 
		// menampilkan string multi line

	/* A.10.5. Nilai nil & Zero Value */
		var namePointer *string = nil
		// nilai nil pada pointer berarti tidak
		fmt.Println(namePointer) 
		// output: <nil>



	/* A.11. Konstanta */
		fmt.Println("A.11. Konstanta")

		/* A.11.1. Penggunaan Konstanta */
		const firstNameConstanta string = "john" 
		// deklarasi konstanta dengan tipe data string
		fmt.Print("haloo ", firstNameConstanta, "\n") 
		// %s untuk memanggil variabel yang dibuat

		const lastNameConstanta = "wick" 
		// deklarasi konstanta tanpa tipe data
		fmt.Print("Nice to meet you ", lastNameConstanta,"\n") 
		// %s untuk memanggil variabel yang dibuat

		/* ◉ Penggunaan Fungsi fmt.Print() */
		fmt.Println("John Wick")
		fmt.Println("John", "Wick")

		fmt.Print("John Wick \n")
		fmt.Print("John ", "Wick \n")
		fmt.Print("John"," ", "Wick \n")


	/* A.11.2. Deklarasi Multi Konstanta */
		const (
			square	= "kotak" // konstanta bertipe data string
			isToday bool = true // konstanta bertipe data boolean
			numeric uint8 = 8 // konstanta bertipe data uint8
			floatNum = 2.2 // konstanta bertipe data float64	
		)
		fmt.Println(square, isToday, numeric, floatNum) // menampilkan semua konstanta

		/* Deklarasi konstanta dengan nilai yang sama */
		const (
			a = "Konstanta"
			b
			c
		)
		fmt.Println(a, b, c) // menampilkan semua konstanta

		/* contoh gabungan dari keduanya: */
		const (
			today string = "Jumat" // konstanta bertipe data string
			sekarang // konstanta dengan tipe data yang sama
			isToday2 = true // konstanta bertipe data boolean
		)
		fmt.Println(today, sekarang, isToday2)

		/* contoh deklrasi multiple konstanta dalam satu baris: */
		const satu, dua = 1, 2 // konstanta bertipe data integer
		const tiga, empat = 3, 4 // konstanta bertipe data integer
		fmt.Println(satu, dua, tiga, empat) // menampilkan semua konstanta

	/* A.12. Operator */
		/* A.12.1. Operator Aritmatika */
		/* 
			- + (penjumlahan)
			- - (pengurangan)
			- * (perkalian)
			- / (pembagian)
			- % (modulus/sisa bagi)
		*/
		var value = (((2+6) % 3 ) * 4 - 2 ) / 3
		fmt.Println("Hasil Value =", value) // %d untuk memanggil tipe data integer

		/* A.12.2. Operator Perbandingan */
		/* 
			- == (sama dengan)
			- != (tidak sama dengan)
			- > (lebih besar dari)
			- < (lebih kecil dari)
			- >= (lebih besar dari atau sama dengan)
			- <= (lebih kecil dari atau sama dengan)
		*/
		var value1 = (((2+6) % 3 ) * 4 - 2 ) / 3 // hasilnya adalah 2
		var isEqual = (value1 == 2) // operator perbandingan sama dengan (==)
		fmt.Printf("Hasilnya = %d %t \n", value1, isEqual) // %d untuk memanggil tipe data integer, %t untuk memanggil tipe data boolean

		/* A.12.3. Operator Logika */
		/*
			- && (AND)
			- || (OR)
			- ! (NOT) 
			Template \t digunakan untuk menambahkan indent tabulasi.
			Biasa dimanfaatkan untuk merapikan tampilan output pada console.
		 */
		var left = false
		var right = true

		var leftAndRight = left && right // operator logika AND (&&)
		fmt.Printf("left && right \t (%t) \n", leftAndRight) // operator logika AND (&&)

		var leftOrRight = left || right // operator logika OR (||)
		fmt.Printf("left || right \t (%t) \n", leftOrRight) // operator logika OR (||)

		var leftReverse = !left // negasi nilai left
		fmt.Printf("!left \t\t (%t) \n", leftReverse) // operator logika NOT (!)

	/* A.13. Seleksi Kondisi */
		/* A.13.1. Seleksi Kondisi Menggunakan Keyword if, else if, & else */
		var point = 8

		if point == 10 {
			fmt.Println("Lulus dengan sempurna")
		} else if point > 5 {
			fmt.Println("Lulus")
		} else if point == 4 {
			fmt.Println("Hampir lulus")
		} else {
			fmt.Printf("Tidak Lulus. Nilai anda %d\n", point)
		}


		/* A.13.2. Variabel Temporary Pada if - else */
		var point1 = 8840.0 // nilai awal point1

		if percent  := point1 / 100; percent >= 100 { // mendeklarasikan variabel temporary percent
			fmt.Printf("%1.f%s perfect! \n", percent, "%") // menampilkan nilai percent dengan format 1 angka di depan koma
		} else if percent >=70 { // menggunakan variabel temporary percent
			fmt.Printf("%1.f%s good \n", percent, "%") // menampilkan nilai percent dengan format 1 angka di depan koma
		} else { // menggunakan variabel temporary percent
			fmt.Printf("%1.f%s not bad \n", percent, "%") // menampilkan nilai percent dengan format 1 angka di depan koma
		}
			
		/* A.13.3. Seleksi Kondisi Menggunakan Keyword switch - case */
		var point2 = 6
		switch point2 {
		case 8 :
			fmt.Println("perfect")
		case 7 :
			fmt.Println("awesome")
		case 6 :
			fmt.Println("not bad")
		}
		
		/* A.13.4. Pemanfaatan case Untuk Banyak Kondisi */
		var point3 = 7 // nilai awal point3
		switch point3 { // memulai seleksi kondisi switch
		case 8: // kondisi pertama
			fmt.Println("perfect") // aksi untuk kondisi pertama
		case 7,6,5,4 : // kondisi kedua dengan banyak pilihan
			fmt.Println("awesome") // aksi untuk kondisi kedua
		default : // kondisi default
			fmt.Println("not bad") // aksi untuk kondisi default
		}

		/* A.13.5. Kurung Kurawal Pada Keyword case & default */
		var point4 = 6 // nilai awal point4

		switch point4 { // memulai seleksi kondisi switch
		case 8 : // kondisi pertama
			fmt.Println("perfect") // aksi untuk kondisi pertama
		case 7,6,5,4 : // kondisi kedua dengan banyak pilihan
			fmt.Println("awesome") // aksi untuk kondisi kedua
		default : // kondisi default
			// penggunaan kurung kurawal pada default
			// untuk menampung banyak aksi
			{
				fmt.Println("not bad") // aksi pertama pada default
				fmt.Println("you can better!") // aksi kedua pada default
			}
		}

		/* A.13.6. Switch Dengan Gaya if - else */
		var point5 = 8 // nilai awal point

		switch { // memulai seleksi kondisi switch
		case point5 == 8 : // kondisi pertama
			fmt.Println("perfect") // aksi untuk kondisi pertama
		case (point5 < 8) && (point5 > 3) : // kondisi kedua
			fmt.Println("awesome") // aksi untuk kondisi kedua

		default : // kondisi default
			{
				fmt.Println("not bad") // aksi pertama pada default
				fmt.Println("you can better!") // aksi kedua pada default
			}

		}

		/* A.13.7. Penggunaan Keyword fallthrough Dalam switch */
		var point6 = 6

		switch {
		case point6 == 8:
			fmt.Println("perfect")
		case (point6 <8 ) && (point6 > 3):
			fmt.Println("asweome")
			fallthrough
		case point6 <5:
			fmt.Println("you need to learn more")
		default:
			{
				fmt.Println("not bad")
				fmt.Println("you can better!")
			}
		}
		
		/* A.13.8. Seleksi Kondisi Bersarang */
		var point7 = 10
		
		if point7 < 7 {
			switch point7 {
			case 10 :
				fmt.Println("perfect")
			default :
				fmt.Println("nice")
			}
			} else {
				switch point7 {
case 5:
					fmt.Println("not bad")
			case 3:
					fmt.Println("keep trying")
			default:
				fmt.Println("you can do it")
				if point7 == 0 {
					fmt.Println("try harder!")
				}
			}
		}

	/* A.14. Perulangan */
		/* A.14.1. Perulangan Menggunakan Keyword for */
		/* for i :=0; i < 5; i++ { //perulangan dari 0 sampai 4
			fmt.Printf("Perulangan ke-%d \n", i) // menampilkan nilai i
		}
 */
		/* A.14.2. Penggunaan Keyword for Dengan Argumen Hanya Kondisi */
		/* var i = 1

		for i < 5 {
			fmt.Println("Angka", i)
			i++
		} */

		/* A.14.3. Penggunaan Keyword for Tanpa Argumen */
		/* var d = 0
		for {
			fmt.Println("angka", d)

			d++
			if d == 10 {
				break
			}
		} */

		/* A.14.4. Penggunaan Keyword for - range */
		/* var xs = "123"  // string dengan value "123"

		for i, v := range xs {
			fmt.Println("index =", i, "Value =", v)
		}

		var ys = [5]int{10,20,30,40,50} // array of int dengan 5 elemen

		for _, v := range ys {
			fmt.Println("Value =", v)
		}

		var zs = ys[0:2] // slice dari array ys, dimulai dari index 0 sampai 2 (tidak termasuk index 2)

		for _, v := range zs {
			fmt.Println("value =", v)
		}


		var kvs = map[byte]int{'a':0, 'b':1, 'c':2} //map dengan key bertipe byte dan value bertipe int
		for k, v := range kvs {
			fmt.Println("key=", k, "value=", v)
		}

		// boleh juga baik k dan atau v nya diabaikan
		for range kvs {
			fmt.Println("Done")
		}

		// selain itu, bisa juga dengan cukup menentukan nilai numerik perulangan
		for ii := range 5 {
			fmt.Print(ii)
		} */
			
		/* A.14.5. Penggunaan Keyword break & continue */
		for i := 1; i <= 10; i++ {
			if i % 2 == 1 {
				continue // melewati perulangan jika i adalah bilangan ganjil
			}
			if i < 8 {
				break // menghentikan perulangan jika i kurang dari 8
			}
			fmt.Println("Angka", i)
		}



}
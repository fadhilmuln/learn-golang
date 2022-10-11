package variabel

import "fmt"

func manifestTyping() {
	//9. Variabel
	/*
		Go Mengadopsi 2 jenis variabel, yaitu :
		1. Yang dituliskan tipe datanya (Manifest Typing)
		2. Yang tidak dituliskan tipe datanya (Type Inference)

		Cara Mendeklarasi Variabel Beserta Tipe Data atau yang bisa disebut dengan Manifest Typing
		Barikut adalah contohnya :
	*/
	fmt.Println("A.9.1 dan A.9.2 Deklarasi Variabel Beserta Tipe Data dan menggunakan Keyword var")
	var firstName string = "John"
	var lastName string
	var nilai uint8 = 225

	lastName = "Wick"

	/*Skema penggunaan Keyword var "var <nama-variabel> <tipe-data> = <nilai> (Ini opsional bisa langsung dideklarasikan atau bisa setelah mendeklarasikan variabelnya)*/
	fmt.Printf("Hello %s %s %d!\n", firstName, lastName, nilai)
	/*Fungsi "Printf" dalam package fmt digunakan untuk menampilkan sebuah isi dari variabel dengan karakter %s sebagai pengganti data string dan %d sebagai pengganti desimal*/
	fmt.Println("Hallo", firstName, ",", lastName, "!")
	fmt.Println()
}

func typeInference() {
	/* 9.3 Deklarasi Variabel tanpa Tipe Data
	Selain manifest typing, Go juga mengadopsi konsep type inferance, yaitu metode
	deklarasi variabel yang tipe datanya ditentukan oleh tipe data nilainya, cara kontradiktif
	jika dibandingkan dengan cara pertama. Dengan metode jenis keyword var dan tipe data tidak pernah ditulis

	Sebagai Contoh :
	*/
	fmt.Println("A.9.3 Deklarasi Variabel tanpa Tipe Data")
	/*Variabel yang ditulis dengan type inferance dengan keyword var saja*/
	var firstName = "Fadhil"

	/*Variabel yang ditulis dengan type inferance
	Ciri dari type interface ini adalah menggantikan operator = dengan := dan keyword var dihilangkan

	Perlu diketahui, deklarasi menggunakan := hanya bisa dilakukan didalam blok fungsi*/
	lastName := "Maulana"
	fmt.Println("Lastname Saat Ini : ", lastName)

	/*Untuk deklarasi nilai selanjutnya cukup dengan menggunakan operator = tanpa menggunakan := lagi*/
	lastName = "Skuyy"

	fmt.Printf("Selamat Pagi, %s %s\n", firstName, lastName)
	fmt.Println()
}

func multiVariabel() {
	// 9.4 Deklarasi Multi Variabel
	fmt.Println("A.9.4 Deklarasi Multi Variabel")
	/*
		Go mendukung metode deklarasi banyak variabel secara bersamaan, caranya dengan mennuliskan variabel-variabelnya dengan pembatas
		tanda koma(,). Untuk pengisian nilainya-pun diperbolehkan secara bersamaan.

		Sebagai contoh :
	*/

	/*Ini menggunakan manifest typing*/
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"
	fmt.Printf("%s %s %s\n", first, second, third)

	/*Ini menggunakan type inferance*/
	isiFirst, isiSecond, isiThird := 598, "two", "three"
	fmt.Printf("%d %s %s\n", isiFirst, isiSecond, isiThird)
	fmt.Println()

}

func variabelUnderscore() {
	// 9.5 Variabel Underscore _
	fmt.Println("A.9.5 Variabel Underscore")
	/*
		Go memiliki aturan unik yang jarang dimiliki bahasa lain, yaitu tidak boleh ada satupun variabel yang menganggur. Artinya, semua
		variabel yang dideklarasikan harus digunakan. Jika ada variabel yang tidak digunakan tapi dideklarasikan, error akan muncul pada
		saat kompilkasi dan program tidak akan bisa di-run

		Underscore(_) adalah reserved variable yang bisa dimanfaatkan untuk menampung nilai yang tidak dipakai. bisa dibilang variabel
		ini merupakan keranjang sampah.

		Sebagai Contoh
	*/

	_ = "Belajar Golang"
	_ = "Golang itu mudah"
	name, _ := "John", "wick"
	fmt.Printf("%s\n", name)

	/*
		Pada contoh diatas variabel name akan berisikan text john, sedangkan nilai wick ditampung oleh variabel underscore, menandakan bahwa
		nilai tersebut tidak akan digunakan.

		Variabel underscore adalah predefined(sudah ditentukan sebelumnya), jadi tidak perlu menggunakan := untuk pengisian nilai, cukup dengan = saja.
		Namun khusus untuk pengisian nilai multi variabel yang menggunakan metode type inferance , boleh didalamnya terdapat variabel underscore.

		Biasanya variabel underscore sering dimanfaatkan untuk menampung nilai balik fungsi yang tidak digunakan dan isi dari variabel ini tidak dapat ditampilkan.
	*/
	fmt.Println()
}

func keywordNew() {
	// 9.6 Deklarasi Variabel menggunakan keyword new
	fmt.Println("A.9.6 Deklarasi Variabel menggunakan keyword new")
	/*
		Keyword new digunakan untuk membuat variabel pointer dengan tipe data tertentu. Nilai data default-nya akan menyesuaikan tipe datanya
		sebagai contoh :
	*/
	nama := new(string)
	fmt.Println(nama)
	*nama = "fadhil"
	fmt.Println(*nama)
	fmt.Println()
	/*
		Variabel nama menampung data bertipe pointer string. jika ditampilkan yang muncul bukanlah nilainya melainkan alamat memori nilai tersebut
		(Dalam bentuk notasi heksadesimal), Untuk menampung nilai aslinya, variabel tersebut perlu di-dereference terlebih dahulu menggunakan tanda asterisk(*)
	*/
}

func keywordMake() {
	// 9.7 Deklarasi variabel menggunakan keyword make
	fmt.Println("A.9.7 Deklarasi variabel menggunakan keyword make")
	/*
		Keyword ini hanya bisa digunakan untuk pembuatan beberapa jenis variabel saja, yaitu:
		1. channel
		2. Slice
		3. map
		Dan untuk pembahasannya ada didalam poin-poin selanjutnya
	*/

	/*Perbedaan Fungsi Print dan Println*/
	fmt.Print("john", "wick\n")
	fmt.Println("john", "wick")
}

func InitVariabel() {
	/*Manifest Typing (Deklarasi Variabel dengan Tipe Data)*/
	manifestTyping()

	/*Type Inference (Deklarasi Variabel tanpa mendeklarasikan tipe datanya)*/
	typeInference()

	/*Multi Variabel*/
	multiVariabel()

	/*Variabel Underscore*/
	variabelUnderscore()

	/*Keyword New*/
	keywordNew()

	/*Keyword Make*/
	keywordMake()
}

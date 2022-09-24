package main

import (
	"fmt"
)

func main() {
	// ConditionalStatement()
	MiniQuiz()
}

func ConditionalStatement() {
	var cuacaHariIni string = "cerah"

	cuacaHujan, cuacaCerah := "hujan", "cerah"

	// hanya mengeksekusi di dalam if ketika bernilai true
	// cuaca hari ini cerah == hujan? false
	if cuacaHariIni == cuacaHujan {
		fmt.Println("kerja dari rumah")
	}

	if cuacaHariIni == cuacaCerah {
		fmt.Println("kerja dari kantor")
	}

	// ketika angka 1
	// > 0 dan < 6 -> ok
	// >= 6 dan <= 10 -> not ok
	// >= 10 -> maybe ok

	angka1 := 10

	switch cuacaHariIni {
	case cuacaHujan:
		fmt.Println("ok")
	case cuacaCerah:
		fmt.Println("very ok")
	}

	switch {
	case angka1 > 0 && angka1 < 6:
		fmt.Println("ok")
	case angka1 >= 6 && angka1 <= 10:
		fmt.Println("not ok")
	case angka1 >= 10:
		fmt.Println("maybe ok")
	default:
		fmt.Println("worst")
	}

	//best practice
	var lokasiKerja1 string
	if cuacaHariIni == cuacaHujan {
		lokasiKerja1 = "wfh"
	} else if cuacaHariIni == cuacaCerah {
		lokasiKerja1 = "wfo"
	}

	var lokasiKerja2 string = "wfo"
	if cuacaHariIni == cuacaHujan {
		lokasiKerja2 = "wfh"
	}

	fmt.Println(lokasiKerja1, lokasiKerja2)

	//----- if scope
	if cuaca := "hujan"; cuaca == cuacaHujan {
		fmt.Println(cuaca)
	}
}

/*
penjelasan fungsi
*/
func Loop() {
	// looping di program => akan digunakan ketika:
	// 1. mengulang suatu proses hingga hitungan tertentu
	// 2. kita harus mengakses semua isi slice / map

	// for loop
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}

	k := 0
	for k < 100 {
		fmt.Println(k)
		k++
	}

	//break
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 6 {
			break
		}
	}

	//continue
	for i := 0; i < 10; i++ {
		if i == 6 {
			fmt.Println("invalid")
			continue
		}
		fmt.Println("valid")
	}

	// infinite for loop
	j := 0

	for {
		j += 10
		if j == 1000 {
			break
		}
	}

	fmt.Println("intinite loop executed", j)
}

func ArrayAndSlice() {
	// assign
	var arrInt []int
	arrInt2 := []int{1, 2, 3, 4}

	//add & get data

	// append -> memasukkan element ke dalam suatu array
	arrInt = append(arrInt2, 2)

	fmt.Println(arrInt)
	fmt.Println(arrInt2)

	fmt.Println(arrInt[2])
	fmt.Println(arrInt[1:3])

	// jika ada array dengan jumlah elm 5,
	// apa yang eterjadi kalau aku akses [6]
	// out of range, panic

	// loop over
	// print element yang ada di array

	// pendekatan pertama
	// len -> untuk mendapatkan jumlah elm dalam array
	fmt.Println(len(arrInt))
	for i := 0; i < len(arrInt); i++ {
		fmt.Println(arrInt[i])
	}

	for i, v := range arrInt {
		// i -> mengindikasikan index keberapa
		// v -> mengindikasikan value di index tersebut
		fmt.Println(i, v)
	}

	// quiz
	// [1, 2, 3, 4]
	// [5, 6, 7, 8]
	// expected: [1, 2, 3, 4, 5, 6, 7, 8]

	arr1 := []int{1, 2, 3, 4}
	arr2 := []int{5, 6, 7, 8}

	// jawaban 1
	arr3 := append(arr1, arr2...)
	fmt.Println(arr3)

	// jawaban 2
	var arr4 []int
	for i := 0; i < len(arr1); i++ {
		arr4 = append(arr4, arr1[i])
	}
	arr4 = append(arr4, arr2...)
	fmt.Println(arr4)

	// mutable type
	arrMut2 := []int{1, 2, 3, 4}
	arrMut3 := arrMut2[:3]
	arrMut3[1] = 10
	fmt.Println(arrMut2, arrMut3)

	// multidimensional
	arrMD1 := [][]int{{1, 2, 3, 4}, {1, 2, 3, 4}}
	fmt.Println(arrMD1)

	// struct
	type Struct1 struct {
		ArrString []string
		ArrInt    []int
	}

	arrMD2 := []Struct1{
		{[]string{"a", "b", "c"}, []int{1, 2, 3}},
		{[]string{"d", "e", "f"}, []int{4, 5, 6}},
	}
	fmt.Println(arrMD2)
}

func HashMap() {
	// assign

	map1 := map[string]int{"key1": 1, "key2": 2}
	map2 := make(map[string]int, 0)
	var map3 map[string]int // belum assign memory (nil)

	fmt.Println(map1, map2, map3)

	// get data
	fmt.Println(map1["key1"])

	//validation
	val, ok := map1["key3"]
	if ok {
		fmt.Println(val)
	}
	fmt.Println(val, ok)

	val, ok = map1["key2"]
	if ok {
		fmt.Println(val)
	}
	fmt.Println(val, ok)

	//assign value ke map
	fmt.Println(map3 == nil)
	map2["key1"] = 1

	map3 = map2
	map3["key1"] = 1

	fmt.Println(map2, map3)
	// delete
	delete(map3, "key1")

	fmt.Println(map2, map3)
	// kenapa map2 juga jadi 0
	// -> mengakses memory yang sama

	// loop over
	for k, v := range map1 {
		fmt.Println(k, v)
	}
}

func Alias() {
	// menghindari ambigu specific type

	// assign alias
	// immutable: int, string, dll
	// mutable: map, arr, dll

	type TestVariable string
	var1 := []TestVariable{"test", "ing"}
	fmt.Println(var1)

	age := 1
	// Bar(age) // tidak bisa digunakan, karena beda type alias

	// casting alias
	// -> mengubah type dari variable
	Bar(Foo(age))
}

// alias example
type Foo int

func Bar(age Foo) {
	fmt.Println(age)
}

func StringInDepth() {
	// string -> kumpulan dari rune / ascii / byte
	str := "test" // t -> rune sendiri, e -> rune sendiri dst

	// ascii / rune /byte
	for _, val := range str {
		// val -> angka ascii / rune dari suatu karakter
		fmt.Println(val)
		fmt.Println(string(val))
	}

	// typecast string -> array of bytes
	by := []byte(str)
	for _, v := range by {
		fmt.Println(v)
	}

	// loop over
}

func MiniQuiz() {
	// quiz 1
	// 1. ada array of alias Int -> alias dari int
	// 2. aku mau menghitung berapa jumlah Int yang sama dalam 1 array
	// 3. aku ingin ngeprint value Int beserta jumlahnya

	// [1, 2, 3, 4, 1, 1, 1, 2, 3, 4, 5, 6, 7]Int
	// hasil -> 1:4, 2:2, 3:2, 4:2, 5:1, 6:1, 7:1

	type Int int
	arrInt := []Int{1, 2, 3, 4, 1, 1, 1, 2, 3, 4, 5, 6, 7}
	mapInt := map[Int]int{}

	for _, v := range arrInt {
		mapInt[v]++
	}

	for k, v := range mapInt {
		fmt.Println(k, v)
	}

	// quiz 2
	// 1. aku ada array of string
	// 2. aku akan membuat array string baru
	// - jika jumlah element dalam array tsb >= 3

	// [cal, cal, cal, man, man, tar, tar, ra, ra, ra]
	// [cal, ra]
	str := []string{"cal", "cal", "cal", "man", "man", "tar", "tar", "ra", "ra", "ra"}
	mapSum := map[string]int{}
	mapDuplicate := map[string]bool{}

	var newArr []string
	for _, v := range str {
		mapSum[v]++
		if (mapSum[v] >= 3) && !mapDuplicate[v] {
			newArr = append(newArr, v)
			mapDuplicate[v] = true
		}
	}
	fmt.Println(newArr)
}

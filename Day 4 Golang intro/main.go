package main

import "fmt"

// func swap(x, y string) (string, string){
// 	return y, x
// }

//----Struct----
// type person struct{
// 	name string
// 	age int
// }

// type student struct{
// 	person
// 	grade int
// }

func main() {
	// for func swap
	// a, b := swap("hello", "wolrd")
	// fmt.Println(a, b)
	// fmt.Println("hellow");
	// var decimal_number = 2.62
	// fmt.Printf("bilangan desimal : %f\n", decimal_number);
	// fmt.Printf("bilangan desimal : %.3f\n", decimal_number);

	// var exist bool = true
	// fmt.Printf("exist? %t \n", exist)

	// nama := "Sanding Adhieguna Rachmat yasin";
	// fmt.Printf("nama saya : %s", nama)

	// message := ` Nama saya "John Wick".
	// Salam Kenal.
	// mari belajar "Golang".`

	// fmt.Println(message);

	// const umur int = 22;
	// fmt.Printf("umur saya :", umur)

	// var point = 8;

	// if point == 10{
	// 	fmt.Println("Lulu sempurna")
	// } else if point > 5 {
	// 	fmt.Println("lulus")
	// } else if point == 4 {
	// 	fmt.Println("hampir lulu")
	// } else {
	// 	fmt.Printf("tidak lulu, nilai anda %d\n", point)
	// }

	// var fruits = [4]string{"apple", "grape", "banana", "melon"}

	// for i, fruit := range fruits {
	// 	fmt.Printf("elemen %d : %s\n", i, fruit)
	// }

	// for i := 1; i<= 10; i++ {
	// 	if i % 2 == 1 {
	// 		continue
	// 	}
	// 	if i > 8 {
	// 		break
	// 	}
	// 	fmt.Println("angka", i)
	// }

	//------Struct-----
	// var s1 = student{}
	// s1.name = "uzumaki naruto"
	// s1.grade = 2
	// s1.age = 22

	// fmt.Println("name :", s1.name)
	// fmt.Println("age : ", s1.person.age)
	
	// //---combine array with struct---
	// var allStudent = []person{
	// 	{name : "Naruto", age: 22},
	// 	{name: "sasuke", age: 23},
	// 	{name: "sakura", age: 24},
	// }

	// for _, student := range allStudent {
	// 	fmt.Println(student.name, "age is", student.age)
	// }

	//----data structure : Map ------ 
	// var chicken map[string]int
	// chicken = map[string]int{}

	// chicken["Januari"] = 50
	// chicken["feb"] = 10

	// fmt.Println("Januari", chicken["Januari"])
	// fmt.Println("mei", chicken["mei"])

	//---slice---
	// var fruits = []string{"apple", "grape", "banana", "melon"}
	// fmt.Println(fruits[0])

	//---defer---mengakhirkan eksekusi sebuah statement tepat sebelum blok fungsi selesai.
	// defer fmt.Println("hallo defer")
	// fmt.Println("Welcome")

	//---Pointer---

	numberA := 4
	var numberB *int = &numberA

	fmt.Println("numberA (value) : ", numberA)
	fmt.Println("numberA (address) : ", &numberA)

	fmt.Println("numberB (value) : ", *numberB)
	fmt.Println("numberB (address) : ", &numberB)
	


}
package main

import (
	"fmt"
	"strings"
)

func multiply(a int, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func nakedReturn(name string) (length int, upperCase string) {
	defer fmt.Println("I am printing this after nakeReturn returned!")
	defer fmt.Println("I am also printing this after nakeReturn returned and first defer printed!")
	length = len(name)
	upperCase = strings.ToUpper(name)
	return
}

func loopRange(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func loopFor(numbers ...int) int {
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	return 1
}

func canIDrink(age int) bool {
	//We can create a variable right inside of if statement
	if koreanAge := age + 2; koreanAge < 18 {
		fmt.Println("korean age: ", koreanAge)
		return false
	}
	return true
}

func main() {
	name := "sang"
	var name2 string = "sang2"
	fmt.Println(name)
	fmt.Println(name2)
	fmt.Println(multiply(2, 4))
	totalLength, upperName := lenAndUpper(name)
	fmt.Println(upperName)
	fmt.Println(totalLength)
	repeatMe("hi ", "my ", "name ", "is ", "Josh")

	_, up := nakedReturn(name)
	fmt.Println(up)
	fmt.Println(loopRange(7, 8, 9, 10, 11, 12))
	loopFor(7, 8, 9, 10, 11, 12)
	canIDrink(12)

	//pointer thing
	//& means address
	//* means see through
	// *b = &a means change the a's value
	a := 2
	b := a
	a = 10
	fmt.Println("a's memory address is: ", &a, " b's memory address is: ", &b)

	c := 2
	d := &c
	c = 5
	fmt.Println("c's memory address is: ", &c, " d's value is looking for c's memory address: ", d)
	fmt.Println("d's value's, which is c's addres, value is : ", *d)
	*d = 10
	fmt.Println("Updated c's value from *d is : ", c)

	//array
	array := [5]string{"sang", "jeong", "jake"}
	fmt.Println(array)
	array[3] = "josh"
	array[4] = "chloe"
	fmt.Println(array)

	sliceArray := []string{"No", "Size", "is"}
	fmt.Println(sliceArray)
	sliceArray = append(sliceArray, "required") //append() is returned a modified array, which is added the new element
	fmt.Println(sliceArray)

	//map
	josh := map[string]int{"height": 173, "age": 12}
	fmt.Println(josh)
	for key, value := range josh {
		fmt.Println(key, value)
	}

	favoriteFood := []string{"kimchi", "ramen"}
	//building map with diverse type of values
	joshRhee := person{name: "josh", age: 18, favFood: favoriteFood}
	fmt.Println(joshRhee)
}

//struct
type person struct {
	name    string
	age     int
	favFood []string
}

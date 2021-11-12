package main

import (
	"fmt"
	"log"
	"simplelogic/example/week"
)

func main() {
	num, num1 := 10, 12
	sum := sumOfMultiplesOf3And5Below1000(num, num1)
	fmt.Println(sum)

	biger1, biger2 := 10, 12
	bigg := bigger(biger1, biger2)
	fmt.Println("the bigger num is ", bigg)

	no1, no2 := 10, 999
	total := pengurangan(no1, no2)
	fmt.Println("the total is ", total, "because the number x is ", no1, " so ", no1, " - ", no2)

	avg1, avg2 := 10, 12
	totalav := average(avg1, avg2)
	fmt.Println("the average is ", totalav)

	swap1, swap2 := 10, 999
	totalSwap1, totalSwap2 := Swap(swap1, swap2)
	fmt.Println("the swap is ", totalSwap1, " and ", totalSwap2)

	theName1, theName2 := "John", "Doe"
	Calling1, calling2 := callName(theName1, theName2)
	fmt.Println("the name is ", Calling1, " and ", calling2)
	printMultiplicationTable1(10)


	//this about large number in looping
	squance := []int{20,11,12,33,22,14}
	fmt.Println("the largest number is ", findLargestNumber(squance...))

	//cari data yang lebih besar
	listnumber := []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println("this is the list number", listnumber)
	fmt.Println("====================================")
	fmt.Println("the largest number is ", carilebihbesar(listnumber...))

	//call date function
	fmt.Println("====================================")
	fmt.Print("input year")
	var y, m , d int
	if _, err := fmt.Scan(&y,&m,&d); err != nil {
		log.Fatal("scanner failed")
	}
	weekday := week.Weekday(y, m, d)
	fmt.Println("the date is ", weekday)

}

//sum of all the multiples of 3 or 5 below 1000
func sumOfMultiplesOf3And5Below1000(x, y int) int {
	sum := x + y
	return sum
}

func bigger(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func pengurangan(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func average(x, y int) (avg float32) {
	avg = float32(x+y) / 2.0
	return
}

func Swap(x, y int) (int, int) {
	return x, y
}
 
func callName(name1, name2 string) (string, string){
	return name2, name1
}

//create print multiplicationtable function
func printMultiplicationTable1(number int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", number, i, number*i)
	}
}

//find largest number in a list
func findLargestNumber(numbers ...int) int {
	var largestNumber int
	for _, number := range numbers {
		if number > largestNumber {
			largestNumber = number
		}
	}
	return largestNumber
}


func carilebihbesar(numbers ...int) int {
	var maxNum int
	for _, number := range numbers {
		if number > maxNum {
			maxNum = number
		}
	}
	return maxNum

}
package main

import (
	"fmt"
	"math"
	"strconv"
	"net/http"

	"btap/golang/file"
	"btap/golang/router"
	"btap/golang/algorithm"
)

func main() {
	//tính toán + - * \
	fmt.Printf("%.1f\n", math.Round(algorithm.Calculate("div", 7, 1.2)))
	fmt.Println("===========================")
	fmt.Println("")

	//thuật toán sắp xếp
	var array = []int{1, 4, 3, 9, 2, 7, 4, 5}
	var array1 = []int{1, 4, 3, 9, 2, 7, 4, 5}
	var array2 = []int{1, 4, 3, 9, 2, 7, 4, 5}
	fmt.Println(algorithm.BubbleSort(array))
	fmt.Println(algorithm.QuickSort(array1, 0, len(array1)-1))
	fmt.Println(algorithm.MergeSort(array2))
	fmt.Println("===========================")
	fmt.Println("")

	//dãy Fibonacci
	for i := 0; i <= 9; i++ {
		fmt.Print(strconv.Itoa(algorithm.Fibonacci(i)) + " ")
	}
	fmt.Println("===========================")
	fmt.Println("")

	//thao tac voi file json
	file.ParseJSON()
	fmt.Println("===========================")
	fmt.Println("")

	//thao tac voi file text
	fmt.Println("Thao tác với file text")
	answer := file.FindMinMaxAndCalcAverage()
	fmt.Println("Giá trị nhỏ nhất là: ", answer[0])
	fmt.Println("Giá trị lớn nhất là: ", answer[1])
	fmt.Println("Giá trị trung bình là: ", answer[2]/answer[3])
	fmt.Println("===========================")
	fmt.Println("")

	//kiểm tra số nguyên tố file text\
	checkPrime := file.CheckPrime()
	for key, value := range checkPrime {
		fmt.Printf("Giá trị: %d", key)
		fmt.Printf("\t\tKiểm tra có phải số nguyên tố không: %t\n", value)
	}
	fmt.Println("===========================")
	fmt.Println("")

	//kiểm tra giá trị có tồn tại trong file hay không
	file.FindNumberInFile([]int{1, 2, 3, 4, 5, 6})
	fmt.Println("===========================")
	fmt.Println("")

	//ghi ra file

	//ghi bằng os
	file.WriteFile()

	//ghi bằng io/ioutil
	file.WriteFileByIO()

	//net/http
	router.CalculateRouter()
	http.ListenAndServe(":8080", nil)
}

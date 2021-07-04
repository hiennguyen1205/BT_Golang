package file

import (
	"os"
	"fmt"
	"math"
	"sort"
	"bufio"
	"strconv"
	"io/ioutil"
)

//hàm đọc file text, trả về các số trong file và tổng của các sô đó để
//phục vụ cho việc tính toán ở hàm sau.
func readFIleText() (result []int, sum int) {
	data, err := os.Open("number.txt")
	if err != nil {
		fmt.Println("error: ", err)
	}
	defer data.Close()
	
	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanWords)
	sum = 0
	for scanner.Scan() {
		element, _ := strconv.Atoi(scanner.Text())
		result = append(result, element)
		sum += element
	}
	return result, sum
}

//Hàm tìm giá trị lớn nhất, nhỏ nhất, trung bình của các giá trị trong file
//retrun 1 mảng với 4 phần tử lầ lượt là: min, max, tổng giá trị các phân tử và tổng số phần tử
func FindMinMaxAndCalcAverage() (result [4]int) {
	answer, sum := readFIleText()
	sort.Slice(answer, func(i, j int) bool { return answer[i] < answer[j] })
	// fmt.Println(answer)
	result[0] = answer[0]
	result[1] = answer[len(answer)-1]
	result[2] = sum
	result[3] = len(answer)
	return result
}

//hàm kiểm tra giá trị có phải số nguyên tố không
func isPrime(number int) bool {
	if number == 2 {
		return true
	}
	if number < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt((float64)(number))); i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

//hàm kiểm tra các giá trị file đầu vào có phải số nguyên tố không
//return lại 1 map với key là các giá trị, value là kiểu bool kiểm tra đúng sai
func CheckPrime() (result map[int]bool) {
	result = make(map[int]bool)
	listNumberInFile, _ := readFIleText()
	for _, value := range listNumberInFile {
		result[value] = isPrime(value)
	}
	return result
}

//hàm kiểm tra một số giá trị có tồn tại trong file hay không
//đầu vào là các giá trị cần tìm, output ra một map với key là 
//các giá trị cần tìm, va lue là true nếu tìm được và ngược lại
func FindNumberInFile(array []int) (result map[int]bool) {
	result = make(map[int]bool)
	checkExists := make([]int,FindMinMaxAndCalcAverage()[2])
	listNumber,_ := readFIleText()
	for _, number := range listNumber {
		checkExists[number]++
	}
	for _, number := range array{
		if checkExists[number] !=0 {
			result[number] = true
		} else {
			result[number] = false
		}
	}
	for key, value := range result {
		fmt.Println("key: ",key)
		fmt.Println("value: ",value)
	}
	return result
}


//Ghi file
//Bước 1: tạo file
//Bước 2 : Ghi :v
//Bước 3: đóng file
func WriteFile(){
	//Bước 1: create file
	file, err := os.Create("write_by_string.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	//Bước 2: ghi file, ở đây dùng WriteString thì sẽ ghi string vào file, return số lượng byte đã ghi
	lengthString, err := file.WriteString("Hello World 0123456789")
	 if err != nil {
        fmt.Println(err)
		//lỗi thì đóng file
        file.Close()
        return
    }
	fmt.Println(lengthString, " ký tự được ghi thành công.")

	//Bước 3: đóng fileText
	err = file.Close()
	if err != nil{
		fmt.Println(err)
        return
	}
}

//ghi file sử dụng io/ioutil
func WriteFileByIO(){
	//Bước 1: tạo file
	file, err := os.Create("write_by_ioutil.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	//Bước 2: ghi file
	message := []byte("Heloo cac cau, hello cc")
	err = ioutil.WriteFile("write_by_ioutil.txt", message, 0644)
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}
	
	//Bước 3: đóng fileText
	err = file.Close()
	if err != nil{
		fmt.Println(err)
        return
	}

}
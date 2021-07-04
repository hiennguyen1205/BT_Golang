package controller

import (
	"fmt"
	"strconv"
	"net/http"
	
	"btap/golang/algorithm"
)


func Handler(w http.ResponseWriter, r *http.Request) {

	//lấy tham số từ url
	calculation := r.URL.Query().Get("calculation")
	operator1, _ := strconv.ParseFloat(r.URL.Query().Get("operator1"), 64)
	operator2, _ := strconv.ParseFloat(r.URL.Query().Get("operator2"), 64)

	//hiển thị kết quả
	fmt.Fprintln(w, "Phép tính: ", operator1,calculation,operator2)
	fmt.Fprintln(w,"Kết quả: ", algorithm.Calculate(calculation, operator1, operator2))

}
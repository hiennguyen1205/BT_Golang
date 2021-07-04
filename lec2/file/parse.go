package file

import (
	"fmt"
	"io/ioutil"
	"encoding/json"

	"btap/golang/models"
)

func ParseJSON() {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println(err)
	}
	//khoi tao bien chua data
	var listUser []models.Person
	err = json.Unmarshal(data, &listUser)
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println("People: ", listUser)
	fmt.Println("")
	fmt.Println("===========================")
	var content []byte
	content, err = json.Marshal(listUser)
	if err != nil {
		return
	}
	fmt.Println((string(content)))
}
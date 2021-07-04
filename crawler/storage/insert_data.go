package storage

import (
	"btap/crawler/connect"
	"btap/crawler/crawl"
)

const (
	link = "https://www.imdb.com/chart/top/?ref_=nv_mv_250"
)

func InsertData() {
	//kết nối với database
	db := connect.Connect()

	//xóa bảng trong database nếu tồn tại
	_, errorDrop := db.Exec("DROP TABLE IF EXISTS Movies")
	if errorDrop != nil {
		panic(errorDrop)
	}
	
	//tạo bảng trong database
	_, error := db.Exec(
		"CREATE TABLE IF NOT EXISTS Movies (" +
			"id INT NOT NULL AUTO_INCREMENT PRIMARY KEY," +
			"title NVARCHAR(100) NOT NULL," +
			"rating FLOAT CHECK(rating>=0 AND rating <=10)," +
			"url VARCHAR(255) NOT NULL)")
	if error != nil {
		panic(error)
	}

	//crawl dữ liệu
	movies := crawl.Crawl(link)

	//insert dữ liệu vào bảng
	for _, movie := range movies {
		_, err := db.Exec("INSERT INTO Movies (title,rating,url) VALUES (?,?,?)", movie.Title, movie.Rating, movie.Url)
		if err != nil {
			panic(err)
		}
	}

}

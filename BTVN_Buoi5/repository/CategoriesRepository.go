package repository

import (
	"errors"
	"fmt"

	"github.com/abcxyz/model"
)

type CategoryRepo struct {
	autoID int64
	categorys  map[int64]*model.Category
}

var Category CategoryRepo

func init() { // hàm luôn chạy khi import package
	Category = CategoryRepo{autoID:0}
	Category.categorys = make(map[int64]*model.Category)
	Category.InitData("sql:Connect vào đây này")
}

func (r *CategoryRepo) autoId() int64 {
	r.autoID += 1
	return r.autoID
}
func (r *CategoryRepo) CreateCategory(category *model.Category) int64 {
	category.Id = r.autoId()
	r.categorys[category.Id] = category
	return category.Id
}
func (r *CategoryRepo) InitData(msg string) {
	fmt.Println(msg)
	r.CreateCategory(&model.Category{
      Name: "Women",
	})

	r.CreateCategory(&model.Category{
      Name: "Men",
	})
	r.CreateCategory(&model.Category{
      Name: "Kids",
		
	})
}
func (r *CategoryRepo) GetAllCategory() map[int64]*model.Category {
	return r.categorys
}

func (r *CategoryRepo) UpdateCategoryPatch(book *model.Category) error {
	if _, ok := r.categorys[book.Id]; ok {
		r.categorys[book.Id] = book
		return nil //tìm được
	} else {
		return errors.New("category not found")
	}
}

func (r *CategoryRepo) Upsert(category *model.Category) int64 {
	if _, ok := r.categorys[category.Id]; ok {
		r.categorys[category.Id] = category //tìm thấy thì update
		return category.Id
	} else { //không tìm thấy thì tạo mới
		return r.CreateCategory(category)
	}
}


func (r *CategoryRepo) DeleteCategoryById(Id int64) error {
	if _, ok := r.categorys[Id]; ok {
		delete(r.categorys, Id)
		return nil
	} else {
		return errors.New("category not found")
	}
}
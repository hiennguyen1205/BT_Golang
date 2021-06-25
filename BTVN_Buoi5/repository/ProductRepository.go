package repository

import (
	"errors"
	"fmt"

	"github.com/abcxyz/model"
)

type ProductRepo struct {
	autoID   int64
	products map[int64]*model.Product
}

var Products ProductRepo

func init() { // hàm luôn chạy khi import package
	Products = ProductRepo{autoID: 0}
	Products.products = make(map[int64]*model.Product)
	Products.InitData("sql:Connect vào đây này")
}

func (r *ProductRepo) autoId() int64 {
	r.autoID += 1
	return r.autoID
}
func (r *ProductRepo) CreateProduct(product *model.Product) int64 {
	product.Id = r.autoId()
	r.products[product.Id] = product
	return product.Id
}
func (r *ProductRepo) InitData(msg string) {
	fmt.Println(msg)
	r.CreateProduct(&model.Product{
		CategoryId: 2,
		Image:      []string{"https://st.focusedcollection.com/14026668/i/140/focused_173588902-stock-photo-portrait-couple-having-fun-sofa.jpg"},
		Name:       "Herschel supply co 25l",
		Price:      75,
		Sale:       0,
		CreatedAt:  1614362898000,
		ModifiedAt: 1615410795000,
	})

	r.CreateProduct(&model.Product{
		CategoryId: 1,
		Image:      []string{"https://st.focusedcollection.com/14026668/i/140/focused_173588902-stock-photo-portrait-couple-having-fun-sofa.jpg"},
		Name:       "Denim jacket blue",
		Price:      92.5,
		Sale:       10,
		CreatedAt:  1610281342000,
		ModifiedAt: 1619283693000,
	})
	r.CreateProduct(&model.Product{
		CategoryId: 3,
		Image:      []string{"https://st.focusedcollection.com/14026668/i/140/focused_173588902-stock-photo-portrait-couple-having-fun-sofa.jpg"},
		Name:       "Coach slim easton black",
		Price:      192.5,
		Sale:       20,
		CreatedAt:  1615745962000,
		ModifiedAt: 1615976362000,
	})
}
func (r *ProductRepo) GetAllProducts() map[int64]*model.Product {
	// listProducts := r.products
	// var result []*model.Product
	// for _, value := range listProducts {
	// 	result = append(result, value)
	// }
	// return result
	return r.products
}

func (r *ProductRepo) UpdateProductPatch(product *model.Product) error {
	if _, ok := r.products[product.Id]; ok {
		r.products[product.Id] = product
		return nil //tìm được
	} else {
		return errors.New("product not found")
	}
}

func (r *ProductRepo) Upsert(product *model.Product) int64 {
	if _, ok := r.products[product.Id]; ok {
		r.products[product.Id] = product //tìm thấy thì update
		return product.Id
	} else { //không tìm thấy thì tạo mới
		return r.CreateProduct(product)
	}
}

func (r *ProductRepo) DeleteProductById(Id int64) error {
	if _, ok := r.products[Id]; ok {
		delete(r.products, Id)
		return nil
	} else {
		return errors.New("product not found")
	}
}

func (r *ProductRepo) FindProductById(Id int64) (*model.Product, error) {
	if book, ok := r.products[Id]; ok {
		return book, nil //tìm được
	} else {
		return nil, errors.New("book not found")
	}
}

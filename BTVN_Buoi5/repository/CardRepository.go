package repository

import (
	"errors"
	"fmt"

	"github.com/abcxyz/model"
)

type CartRepo struct {
	autoID int64
	carts  map[int64]*model.Cart
}

var Carts CartRepo

func init() { // hàm luôn chạy khi import package
	Carts = CartRepo{autoID: 0}
	Carts.carts = make(map[int64]*model.Cart)
	Carts.InitData("sql:Connect vào đây này")
}

func (r *CartRepo) autoId() int64 {
	r.autoID += 1
	return r.autoID
}
func (r *CartRepo) CreateCart(user *model.Cart) int64 {
	user.Id = r.autoId()
	r.carts[user.Id] = user
	return user.Id
}
func (r *CartRepo) InitData(msg string) {
	fmt.Println(msg)
	r.CreateCart(&model.Cart{
		CategoryId: 2,
		Image:      "/uploads/images/item-02.jpg",
		Name:       "Herschel supply co 25l",
		Price:      75,
		Quantity:   2,
		CreatedAt:  1614362898000,
		ModifiedAt: 1615410795000,
	})

	r.CreateCart(&model.Cart{
		CategoryId: 2,
		Image:      "/uploads/images/item-02.jpg",
		Name:       "Herschel supply co 25l",
		Price:      75,
		Quantity:   2,
		CreatedAt:  1614362898000,
		ModifiedAt: 1615410795000,
	})
	r.CreateCart(&model.Cart{
		CategoryId: 2,
		Image:      "/uploads/images/item-02.jpg",
		Name:       "Herschel supply co 25l",
		Price:      75,
		Quantity:   2,
		CreatedAt:  1614362898000,
		ModifiedAt: 1615410795000,
	})
}
func (r *CartRepo) GetAllCarts() map[int64]*model.Cart {
	return r.carts
}

func (r *CartRepo) UpdateCartPatch(cart *model.Cart) error {
	if _, ok := r.carts[cart.Id]; ok {
		r.carts[cart.Id] = cart
		return nil //tìm được
	} else {
		return errors.New("user not found")
	}
}

func (r *CartRepo) Upsert(user *model.Cart) int64 {
	if _, ok := r.carts[user.Id]; ok {
		r.carts[user.Id] = user //tìm thấy thì update
		return user.Id
	} else { //không tìm thấy thì tạo mới
		return r.CreateCart(user)
	}
}

func (r *CartRepo) DeleteCartById(Id int64) error {
	if _, ok := r.carts[Id]; ok {
		delete(r.carts, Id)
		return nil
	} else {
		return errors.New("user not found")
	}
}

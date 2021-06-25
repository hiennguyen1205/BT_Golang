package repository

import (
	"errors"
	"fmt"

	"github.com/abcxyz/model"
)

type UserRepo struct {
	autoID int64
	users  map[int64]*model.User
}

var Users UserRepo

func init() { // hàm luôn chạy khi import package
	Users = UserRepo{autoID: 0}
	Users.users = make(map[int64]*model.User)
	Users.InitData("sql:Connect vào đây này")
}

func (r *UserRepo) autoId() int64 {
	r.autoID += 1
	return r.autoID
}
func (r *UserRepo) CreateUser(user *model.User) int64 {
	user.Id = r.autoId()
	r.users[user.Id] = user
	return user.Id
}
func (r *UserRepo) InitData(msg string) {
	fmt.Println(msg)
	r.CreateUser(&model.User{
		FirstName:  "Administrator",
		LastName:   "",
		Username:   "admin",
		Email:      "admin@gmail.com",
		Password:   "admin",
		Avatar:     "https://robohash.org/eaquequasincidunt.png?size=50x50&set=set1",
		Gender:     "Genderfluid",
		Phone:      "933-658-1213",
		Birthday:   "1994-03-23",
		Status:     true,
		CreatedAt:  1609483221000,
		ModifiedAt: 1609483221000,
	})

	r.CreateUser(&model.User{
		FirstName:  "Client 1",
		LastName:   "",
		Username:   "client1",
		Email:      "client1@gmail.com",
		Password:   "client",
		Avatar:     "https://robohash.org/accusantiumminimamagni.png?size=50x50&set=set1",
		Gender:     "Male",
		Phone:      "510-449-7332",
		Birthday:   "2002-03-11",
		Status:     false,
		CreatedAt:  1617440961000,
		ModifiedAt: 1618301961000,
	})
	r.CreateUser(&model.User{
		FirstName:  "Client 2",
		LastName:   "",
		Username:   "client2",
		Email:      "client2@gmail.com",
		Password:   "kjU6qK1Bm",
		Avatar:     "https://robohash.org/voluptatemdebitiset.png?size=50x50&set=set1",
		Gender:     "Female",
		Phone:      "676-983-4977",
		Birthday:   "1997-09-29",
		Status:     false,
		CreatedAt:  1615745961000,
		ModifiedAt: 1615976361000,
	})
}
func (r *UserRepo) GetAllUsers() map[int64]*model.User {
	return r.users
}

func (r *UserRepo) UpdateUserPatch(book *model.User) error {
	if _, ok := r.users[book.Id]; ok {
		r.users[book.Id] = book
		return nil //tìm được
	} else {
		return errors.New("user not found")
	}
}

func (r *UserRepo) Upsert(user *model.User) int64 {
	if _, ok := r.users[user.Id]; ok {
		r.users[user.Id] = user //tìm thấy thì update
		return user.Id
	} else { //không tìm thấy thì tạo mới
		return r.CreateUser(user)
	}
}

func (r *UserRepo) DeleteUserById(Id int64) error {
	if _, ok := r.users[Id]; ok {
		delete(r.users, Id)
		return nil
	} else {
		return errors.New("user not found")
	}
}

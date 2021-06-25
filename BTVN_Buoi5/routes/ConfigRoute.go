package routes

import (
	"github.com/abcxyz/controller"
	"github.com/gofiber/fiber/v2"
)

func ConfigUserRouter(router *fiber.Router) {
	//return all user
	(*router).Get("", controller.GetAllUsers)
	//create user :insert 1 bản ghi
	(*router).Post("", controller.CreateUser)
	//patch cập nhật 1 bản ghi (lỗi nếu không tìm đc id)
	(*router).Patch("", controller.UpdateUserPatch)
	//put cập nhật 1 bản ghi (không tìm được thì tạo mới) 
	(*router).Put("", controller.UpSertUser)
	//delete
	(*router).Delete("", controller.DeleteUserById)
}

func ConfigProductRouter(router *fiber.Router) {
	//return all product
	(*router).Get("", controller.GetAllProducts)
	//create product :insert 1 bản ghi
	(*router).Post("", controller.CreateProduct)
	//patch cập nhật 1 bản ghi (lỗi nếu không tìm đc id)
	(*router).Patch("", controller.UpdateProductPatch)
	//put cập nhật 1 bản ghi (không tìm được thì tạo mới) 
	(*router).Put("", controller.UpSertProduct)
	//delete
	(*router).Delete("/:id", controller.DeleteProductById)
}

func ConfigCategoryRouter(router *fiber.Router) {
	//return all Category
	(*router).Get("", controller.GetAllCategories)
	//create Category :insert 1 bản ghi
	(*router).Post("", controller.CreateCategory)
	//patch cập nhật 1 bản ghi (lỗi nếu không tìm đc id)
	(*router).Patch("", controller.UpdateCategoryPatch)
	//put cập nhật 1 bản ghi (không tìm được thì tạo mới) 
	(*router).Put("", controller.UpSertCategory)
	//delete
	(*router).Delete("", controller.DeleteCategoryById)
}

func ConfigCartRouter(router *fiber.Router) {
	//return all Cart
	(*router).Get("", controller.GetAllCarts)
	//create Cart :insert 1 bản ghi
	(*router).Post("", controller.CreateCart)
	//patch cập nhật 1 bản ghi (lỗi nếu không tìm đc id)
	(*router).Patch("", controller.UpdateCartPatch)
	//put cập nhật 1 bản ghi (không tìm được thì tạo mới) 
	(*router).Put("", controller.UpSertCart)
	//delete
	(*router).Delete("", controller.DeleteCartById)
}

func ConfigReviewRouter(router *fiber.Router) {
	//return all Cart
	(*router).Get("", controller.GetAllReviews)
	//create Cart :insert 1 bản ghi
	(*router).Post("", controller.CreateReview)
	//patch cập nhật 1 bản ghi (lỗi nếu không tìm đc id)
	// (*router).Patch("", controller.UpdateCartPatch)
	// //put cập nhật 1 bản ghi (không tìm được thì tạo mới) 
	(*router).Put("", controller.UpsertReview)
	//delete
	(*router).Delete("/:id", controller.DeleteReviewById)
}
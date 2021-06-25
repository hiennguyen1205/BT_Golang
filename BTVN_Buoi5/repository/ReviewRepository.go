package repository

import (
	"errors"
	"fmt"

	"github.com/abcxyz/model"
)

type ReviewRepo struct {
	reviews map[int64]*model.Review
	autoID  int64
}

var Reviews ReviewRepo //Khai báo biến toàn cục, global variable

func init() { //func init luôn chạy đầu tiên khi chúng ta import package
	Reviews = ReviewRepo{autoID: 0}
	Reviews.reviews = make(map[int64]*model.Review)
	Reviews.InitData("sql:45312")
}

func (r *ReviewRepo) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}

func (r *ReviewRepo) CreateNewReview(review *model.Review) int64 {
	nextID := r.getAutoID()
	review.Id = nextID
	r.reviews[nextID] = review
	return nextID
}

func (r *ReviewRepo) InitData(connection string) {
	fmt.Println("Connect to ", connection)
}

func (r *ReviewRepo) GetAllReviews() map[int64]*model.Review {
	return r.reviews
}

func (r *ReviewRepo) Upsert(review *model.Review) int64 {
	if _, ok := r.reviews[review.Id]; ok {
		r.reviews[review.Id] = review
		return review.Id
	} else {
		return r.CreateNewReview(review)
	}
}

func (r *ReviewRepo) AverageRating(productId int64) (result float32) {
	listAllReview := r.GetAllReviews()
	countReviewsOfProduct := 0
	sumRating := 0
	for _, review := range listAllReview {
		if review.ProductId == productId {
			countReviewsOfProduct++
			sumRating += review.Rating
		}
	}
	if(countReviewsOfProduct == 0) {
		return 0;
	}
	result = (float32)(sumRating / countReviewsOfProduct)
	return result
}

func (r *ReviewRepo) DeleteReviewById(Id int64) error {
	if _, ok := r.reviews[Id]; ok {
		delete(r.reviews, Id)
		return nil
	} else {
		return errors.New("product not found")
	}
}
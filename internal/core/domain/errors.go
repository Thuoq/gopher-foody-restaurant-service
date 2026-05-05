package domain

import "errors"

var (
	ErrRestaurantNotFound = errors.New("restaurant not found")
	ErrUnauthorized       = errors.New("unauthorized access to this resource")
	ErrFoodNotFound       = errors.New("food not found")
	ErrCategoryNotFound   = errors.New("category not found")
	ErrDuplicateResource  = errors.New("resource already exists")
	ErrInvalidInput       = errors.New("invalid input data")
)

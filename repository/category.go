package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"
	"errors"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetCategoriesByUserId(ctx context.Context, id int) ([]entity.Category, error)
	StoreCategory(ctx context.Context, category *entity.Category) (categoryId int, err error)
	StoreManyCategory(ctx context.Context, categories []entity.Category) error
	GetCategoryByID(ctx context.Context, id int) (entity.Category, error)
	UpdateCategory(ctx context.Context, category *entity.Category) error
	DeleteCategory(ctx context.Context, id int) error
	NextCategory(id int) (entity.Category, error)
	PreviousCategory(id int) (entity.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) GetCategoriesByUserId(ctx context.Context, id int) ([]entity.Category, error) {
	var result []entity.Category
	if err := r.db.WithContext(ctx).Table("categories").Where("user_id = ?", id).Find(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []entity.Category{}, nil
		} else {
			return nil, err
		}
	}	
	return result, nil
}

func (r *categoryRepository) StoreCategory(ctx context.Context, category *entity.Category) (categoryId int, err error) {
	if err := r.db.WithContext(ctx).Create(&category).Error; err != nil {
		return 0, err
	}
	return category.ID, nil
}

func (r *categoryRepository) StoreManyCategory(ctx context.Context, categories []entity.Category) error {
	if err := r.db.WithContext(ctx).Create(&categories).Error; err != nil {
		return err
	}
	return nil
}

func (r *categoryRepository) GetCategoryByID(ctx context.Context, id int) (entity.Category, error) {
	var result entity.Category
	if err := r.db.WithContext(ctx).Table("categories").Where("id = ?", id).Find(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Category{}, nil
		} else {
			return entity.Category{}, err
		}
	}
	return result, nil
}

func (r *categoryRepository) UpdateCategory(ctx context.Context, category *entity.Category) error {
	if err := r.db.WithContext(ctx).Table("categories").Where("id = ?", category.ID).Updates(category).Error; err != nil {
		return err
	}
	return nil
}

func (r *categoryRepository) DeleteCategory(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).Delete(&entity.Category{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *categoryRepository) NextCategory(id int) (entity.Category, error) {
	var result entity.Category
	if err := r.db.Table("categories").Where("id > ? ", id).Order("id ASC").First(&result).Error; err != nil {
		return entity.Category{}, err
	}
	return result, nil
}

func (r *categoryRepository) PreviousCategory(id int) (entity.Category, error) {
	var result entity.Category
	if err := r.db.Table("categories").Where("id < ? ", id).Order("id DESC").First(&result).Error; err != nil {
		return entity.Category{}, err
	}
	return result, nil
}
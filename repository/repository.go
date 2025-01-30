package repository

import "gorm.io/gorm"

type ProductRepository struct {
	// Gorm Database Connection Used for Product Operations
	RepoDataSetup *gorm.DB
}

func CreateProductRepository(database *gorm.DB) *ProductRepository {
	// Create, Return Instance of ProductRepository with Database Connection
	// Initialize, Return ProductRepository
	return &ProductRepository{RepoDataSetup: database}
}
package repository

import "gorm.io/gorm"

type RepositoryFactory struct {
    db *gorm.DB
}

func NewRepositoryFactory(db *gorm.DB) *RepositoryFactory {
    return &RepositoryFactory{db: db}
}

func (f *RepositoryFactory) GetDB() *gorm.DB {
    return f.db
}

func (f *RepositoryFactory) GetUserRepository() *UserRepository {
    return NewUserRepository(f.db)
}

func (f *RepositoryFactory) GetOrderRepository() *OrderRepository {
    return NewOrderRepository(f.db)
}

func (f *RepositoryFactory) GetCartRepository() *CartRepository {
    return NewCartRepository(f.db)
}

func (f *RepositoryFactory) GetProductRepository() *ProductRepository {
    return NewProductRepository(f.db)
}

func (f *RepositoryFactory) GetAdvertisementRepository() *AdvertisementRepository {
    return NewAdvertisementRepository(f.db)
}

func (f *RepositoryFactory) GetReviewRepository() *ReviewRepository {
    return NewReviewRepository(f.db)
} 
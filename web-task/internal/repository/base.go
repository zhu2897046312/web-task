package repository

import "gorm.io/gorm"

type Repository interface {
    Create(interface{}) error
    Update(interface{}) error
    Delete(interface{}) error
}

type BaseRepository struct {
    db *gorm.DB
}

func NewBaseRepository(db *gorm.DB) *BaseRepository {
    return &BaseRepository{db: db}
}

func (r *BaseRepository) Create(v interface{}) error {
    return r.db.Create(v).Error
}

func (r *BaseRepository) Update(v interface{}) error {
    return r.db.Save(v).Error
}

func (r *BaseRepository) Delete(v interface{}) error {
    return r.db.Delete(v).Error
} 
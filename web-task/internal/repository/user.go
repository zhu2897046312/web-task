package repository

import (
	"errors"
	"time"
	"web-task/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	*BaseRepository
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateProfile(user *models.User) error {
	return r.db.Model(user).Select("nickname", "avatar").Updates(user).Error
}

func (r *UserRepository) UpdatePassword(userID uint, hashedPassword string) error {
	return r.db.Model(&models.User{}).Where("id = ?", userID).
		Update("password", hashedPassword).Error
}

func (r *UserRepository) UpdateEmailVerification(userID uint, verified bool) error {
	return r.db.Model(&models.User{}).Where("id = ?", userID).
		Update("email_verified", verified).Error
}

func (r *UserRepository) SetVerifyToken(userID uint, token string, expiry time.Time) error {
	return r.db.Model(&models.User{}).Where("id = ?", userID).
		Updates(map[string]interface{}{
			"verify_token": token,
			"token_expiry": expiry,
		}).Error
}

func (r *UserRepository) ClearVerifyToken(userID uint) error {
	return r.db.Model(&models.User{}).Where("id = ?", userID).
		Updates(map[string]interface{}{
			"verify_token": "",
			"token_expiry": nil,
		}).Error
}

func (r *UserRepository) VerifyEmailByToken(token string) error {
	result := r.db.Model(&models.User{}).
		Where("verify_token = ? AND token_expiry > ? AND email_verified = ?", 
			token, time.Now(), false).
		Updates(map[string]interface{}{
			"email_verified": true,
			"verify_token": "",
			"token_expiry": nil,
		})
	
	if result.Error != nil {
		return result.Error
	}
	
	if result.RowsAffected == 0 {
		return errors.New("invalid or expired token")
	}
	
	return nil
}

func (r *UserRepository) FindByVerifyToken(token string) (*models.User, error) {
	var user models.User
	err := r.db.Where("verify_token = ? AND token_expiry > ?", token, time.Now()).
		First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) SoftDelete(userID uint) error {
	return r.db.Delete(&models.User{}, userID).Error
}

func (r *UserRepository) IsEmailVerified(userID uint) (bool, error) {
	var user models.User
	err := r.db.Select("email_verified").First(&user, userID).Error
	if err != nil {
		return false, err
	}
	return user.EmailVerified, nil
}

func (r *UserRepository) UpdateRole(userID uint, role string) error {
	return r.db.Model(&models.User{}).Where("id = ?", userID).
		Update("role", role).Error
}

func (r *UserRepository) UpdateVerifyToken(userID uint, token string, expiry time.Time) error {
	return r.db.Model(&models.User{}).Where("id = ?", userID).
		Updates(map[string]interface{}{
			"verify_token": token,
			"token_expiry": expiry,
		}).Error
}

func (r *UserRepository) VerifyEmail(token string) error {
	return r.db.Model(&models.User{}).
		Where("verify_token = ? AND token_expiry > ? AND email_verified = ?", 
			token, time.Now(), false).
		Updates(map[string]interface{}{
			"email_verified": true,
			"verify_token": "",
			"token_expiry": nil,
		}).Error
}

func (r *UserRepository) GetByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) ListAddresses(userID uint) ([]models.Address, error) {
	var addresses []models.Address
	err := r.db.Where("user_id = ?", userID).Find(&addresses).Error
	return addresses, err
}

func (r *UserRepository) CreateAddress(address *models.Address) error {
	return r.db.Create(address).Error
}

func (r *UserRepository) GetAddressByID(id uint) (*models.Address, error) {
	var address models.Address
	err := r.db.First(&address, id).Error
	if err != nil {
		return nil, err
	}
	return &address, nil
}

func (r *UserRepository) UpdateAddress(address *models.Address) error {
	return r.db.Save(address).Error
}

func (r *UserRepository) DeleteAddress(id uint) error {
	return r.db.Delete(&models.Address{}, id).Error
}

func (r *UserRepository) UnsetDefaultAddresses(userID uint) error {
	return r.db.Model(&models.Address{}).
		Where("user_id = ?", userID).
		Update("is_default", false).Error
} 
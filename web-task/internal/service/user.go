package service

import (
    "errors"
    "web-task/internal/models"
    "golang.org/x/crypto/bcrypt"
)

type UserService struct {
    *Service
}

func NewUserService(base *Service) *UserService {
    return &UserService{Service: base}
}

func (s *UserService) Register(user *models.User) error {
    // 检查用户名是否已存在
    if _, err := s.repoFactory.GetUserRepository().GetByUsername(user.Username); err == nil {
        return errors.New("username already exists")
    }

    // 检查邮箱是否已存在
    if _, err := s.repoFactory.GetUserRepository().GetByEmail(user.Email); err == nil {
        return errors.New("email already exists")
    }

    // 加密密码
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(hashedPassword)

    return s.repoFactory.GetUserRepository().Create(user)
}

func (s *UserService) Login(username, password string) (*models.User, error) {
    user, err := s.repoFactory.GetUserRepository().GetByUsername(username)
    if err != nil {
        return nil, errors.New("invalid username or password")
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return nil, errors.New("invalid username or password")
    }

    return user, nil
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
    return s.repoFactory.GetUserRepository().GetByID(id)
}

func (s *UserService) UpdateUser(user *models.User) error {
    // 如果更新了密码，需要加密
    if user.Password != "" {
        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
        if err != nil {
            return err
        }
        user.Password = string(hashedPassword)
    }
    return s.repoFactory.GetUserRepository().Update(user)
}

func (s *UserService) AddAddress(address *models.Address) error {
    // 检查用户是否存在
    if _, err := s.GetUserByID(address.UserID); err != nil {
        return errors.New("user not found")
    }
    
    // 如果是默认地址，先将其他地址设置为非默认
    if address.IsDefault {
        addresses, err := s.ListAddresses(address.UserID)
        if err != nil {
            return err
        }
        for _, addr := range addresses {
            if addr.IsDefault {
                addr.IsDefault = false
                if err := s.repoFactory.GetUserRepository().Update(&addr); err != nil {
                    return err
                }
            }
        }
    }
    
    return s.repoFactory.GetUserRepository().Create(address)
}

func (s *UserService) ListAddresses(userID uint) ([]models.Address, error) {
    return s.repoFactory.GetUserRepository().ListAddresses(userID)
} 
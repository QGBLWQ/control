package model

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User the user model
type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// TableName for gorm
func (User) TableName() string {
	return "users"
}

// EncryptPassword encrypts the password and stores the hash
func (u *User) EncryptPassword() error {
	if u.Password == "" {
		return errors.New("password cannot be empty")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}

// GetFirstByID retrieves a user by their ID
func (u *User) GetFirstByID(id string) error {
	err := DB().Where("id=?", id).First(u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("user not found")
	}
	return err
}

// GetFirstByEmail retrieves a user by their email
func (u *User) GetFirstByEmail(email string) error {
	err := DB().Where("email=?", email).First(u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("user not found")
	}
	return err
}

// Create inserts a new user into the database
func (u *User) Create() error {
	if err := DB().Create(u).Error; err != nil {
		return err
	}
	return nil
}

// Signup registers a new user
func (u *User) Signup() error {
	var existingUser User
	if err := existingUser.GetFirstByEmail(u.Email); err == nil {
		return errors.New("user already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if err := u.EncryptPassword(); err != nil {
		return err
	}

	return u.Create()
}

// Login checks if the provided password matches the stored hash
func (u *User) Login(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

// LoginByEmailAndPassword logs a user in using their email and password
func LoginByEmailAndPassword(email, password string) (*User, error) {
	var user User
	if err := user.GetFirstByEmail(email); err != nil {
		return nil, err
	}

	if err := user.Login(password); err != nil {
		return nil, err
	}

	return &user, nil
}

// GetAllUsers retrieves all users (admin use only)
func GetAllUsers(users *[]User) error {
	if err := DB().Find(users).Error; err != nil {
		return err
	}
	return nil
}

// DeleteUserByID deletes a user by their ID (admin use only)
func DeleteUserByID(id string) error {
	var user User
	if err := DB().Where("id = ?", id).First(&user).Error; err != nil {
		return err
	}
	return DB().Delete(&user).Error
}

// AdminCreateUser creates a new user as an admin
func AdminCreateUser(user *User) error {
	var existingUser User
	if err := DB().Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return errors.New("email already taken")
	}

	if err := user.EncryptPassword(); err != nil {
		return err
	}

	return DB().Create(user).Error
}

// AdminUpdateUser updates a user's information (admin use only)
func AdminUpdateUser(user *User) error {
	// Check if the password needs to be encrypted
	if user.Password != "" {
		if err := user.EncryptPassword(); err != nil {
			return err
		}
	}

	return DB().Save(user).Error
}
package user_management_model

import (
    "errors"
    "fmt"
    "time"

    "github.com/bxcodec/faker/v4"
    "gorm.io/gorm"
)

// User represents the users table in the database.
type User struct {
    ID        uint           `gorm:"primaryKey" json:"id"`
    Name      string         `gorm:"type:varchar(100)" json:"name"`
    Email     string         `gorm:"uniqueIndex;type:varchar(100)" json:"email"`
    Password  string         `gorm:"type:varchar(255)" json:"-"`
    CreatedAt int64          `json:"created_at"`
    UpdatedAt int64          `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
    Roleable  []Roleable     `gorm:"polymorphic:Roleable;"`
}

// Implement RoleableEntity for User
func (u User) GetID() uint {
    return u.ID
}

func (u User) GetType() string {
    return "users"
}

// SetTimestamps updates CreatedAt and UpdatedAt fields for the user.
func (u *User) SetTimestamps() {
    now := time.Now().Unix()
    if u.CreatedAt == 0 {
        u.CreatedAt = now
    }
    u.UpdatedAt = now
}

// CreateUser inserts a new user into the database.
func CreateUser(db *gorm.DB, user *User) error {
    user.SetTimestamps()
    return db.Create(&user).Error
}

// GetAllUsers retrieves all users from the database.
func GetAllUsers(db *gorm.DB) ([]User, error) {
    var users []User
    err := db.Find(&users).Error
    return users, err
}

// GetUser retrieves a user by ID from the database.
func GetUser(db *gorm.DB, id string) (User, error) {
    var user User
    err := db.First(&user, id).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return User{}, fmt.Errorf("user with ID %s not found", id)
    }
    return user, err
}

// UpdateUser updates an existing user's information in the database.
func UpdateUser(db *gorm.DB, user *User) error {
    user.SetTimestamps()
    result := db.Model(&User{}).Where("id = ?", user.ID).Updates(user)

    if result.RowsAffected == 0 && result.Error == nil {
        return fmt.Errorf("no user found with ID: %d", user.ID)
    }
    return result.Error
}

// DeleteUser deletes a user by ID from the database.
func DeleteUser(db *gorm.DB, id string) error {
    result := db.Delete(&User{}, id)
    if result.RowsAffected == 0 && result.Error == nil {
        return fmt.Errorf("no user found with ID: %s", id)
    }
    return result.Error
}

// GenerateFakeUsers generates a specified number of fake users for testing.
func GenerateFakeUsers(db *gorm.DB, count int) error {
    for i := 0; i < count; i++ {
        user := User{
            Name:     faker.Name(),
            Email:    faker.Email(),
            Password: faker.Password(),
        }
        user.SetTimestamps()
        if err := db.Create(&user).Error; err != nil {
            return err
        }
    }
    return nil
}

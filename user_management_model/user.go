package user_management_model 

import (
    "time"
    "fmt"

    "github.com/bxcodec/faker/v4"
    "gorm.io/gorm"
)

// User represents the users table in the database
type User struct {
    ID        uint            `gorm:"primaryKey" json:"id"`
    Name      string          `gorm:"type:varchar(100)" json:"name"`
    Email     string          `gorm:"uniqueIndex;type:varchar(100)" json:"email"`
    Password  string          `gorm:"type:varchar(255)" json:"-"`
    CreatedAt int64           `json:"created_at"`
    UpdatedAt int64           `json:"updated_at"`
    DeletedAt gorm.DeletedAt  `gorm:"index" json:"-"`
    Roleable []Roleable       `gorm:"polymorphic:Roleable;"`
}

// Implement RoleableEntity for User
func (u User) GetID() uint {
    return u.ID
}

func (u User) GetType() string {
    return "users"
}

// CreateUser creates a new user
func CreateUser(db *gorm.DB, user User) (err error) {

    // Set timestamps
    user.CreatedAt = time.Now().Unix()
    user.UpdatedAt = time.Now().Unix()

    err = db.Create(&user).Error

    return
}

// GetAllUsers retrieves all users
func GetAllUsers(db *gorm.DB) (err error) {

    var users []User;

    err = db.Find(&users).Error

    return 
}

// GetUser retrieves a user by ID
func GetUser(db *gorm.DB, id int, user User) (err error) {

    err = db.First(&user, id).Error

    return
}

// UpdateUser updates an existing user by ID
func UpdateUser(db *gorm.DB, user User) (err error) {

    // Update timestamps
    user.UpdatedAt = time.Now().Unix()

    result := db.Model(&User{}).Where("id = ?", user.ID).Updates(user)

    err = result.Error

    if result.RowsAffected == 0 && err == nil {
      err = fmt.Errorf("no user found with ID: %d", user.ID)
    }

    return 
}

// DeleteUser deletes a user by ID
func DeleteUser(db *gorm.DB, id int) (err error) {
    err = db.Delete(&User{}, id).Error;

    return
}

func GenerateFakeUsers(db *gorm.DB, count int) error {

    for i := 0; i < count; i++ {
        user := User{
            Name:      faker.Name(),
            Email:     faker.Email(),
            Password:  faker.Password(),
            CreatedAt: time.Now().Unix(),
            UpdatedAt: time.Now().Unix(),
        }

        // Create user
        if err := db.Create(&user).Error; err != nil {
            return err
        }

    }

    return nil
}

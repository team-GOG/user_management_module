package user_management_model

import (
    "errors"
    "fmt"
    "time"

    "github.com/bxcodec/faker/v4"
    "gorm.io/gorm"
)

// Admin represents the admins table in the database.
type Admin struct {
    ID        uint           `gorm:"primaryKey" json:"id"`
    Name      string         `gorm:"not null" json:"name"`
    Email     string         `gorm:"unique;not null" json:"email"`
    Password  string         `gorm:"type:varchar(255)" json:"-"`
    CreatedAt int64          `json:"created_at"`
    UpdatedAt int64          `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
    Roleable  []Roleable     `gorm:"polymorphic:Roleable;"`
}

// Implement RoleableEntity for Admin
func (a Admin) GetID() uint {
    return a.ID
}

func (a Admin) GetType() string {
    return "admins"
}

// SetTimestamps updates CreatedAt and UpdatedAt fields for the admin.
func (a *Admin) SetTimestamps() {
    now := time.Now().Unix()
    if a.CreatedAt == 0 {
        a.CreatedAt = now
    }
    a.UpdatedAt = now
}

// CreateAdmin inserts a new admin into the database.
func CreateAdmin(db *gorm.DB, admin *Admin) error {
    admin.SetTimestamps()
    return db.Create(&admin).Error
}

// GetAllAdmins retrieves all admins from the database.
func GetAllAdmins(db *gorm.DB) ([]Admin, error) {
    var admins []Admin
    err := db.Find(&admins).Error
    return admins, err
}

// GetAdmin retrieves an admin by ID from the database.
func GetAdmin(db *gorm.DB, id string) (Admin, error) {
    var admin Admin
    err := db.First(&admin, id).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return Admin{}, fmt.Errorf("admin with ID %s not found", id)
    }
    return admin, err
}

// UpdateAdmin updates an existing admin's information in the database.
func UpdateAdmin(db *gorm.DB, admin *Admin) error {
    admin.SetTimestamps()
    result := db.Model(&Admin{}).Where("id = ?", admin.ID).Updates(admin)

    if result.RowsAffected == 0 && result.Error == nil {
        return fmt.Errorf("no admin found with ID: %d", admin.ID)
    }
    return result.Error
}

// DeleteAdmin deletes an admin by ID from the database.
func DeleteAdmin(db *gorm.DB, id string) error {
    result := db.Delete(&Admin{}, id)
    if result.RowsAffected == 0 && result.Error == nil {
        return fmt.Errorf("no admin found with ID: %s", id)
    }
    return result.Error
}

// GenerateFakeAdmins generates a specified number of fake admins for testing.
func GenerateFakeAdmins(db *gorm.DB, count int) error {
    for i := 0; i < count; i++ {
        admin := Admin{
            Name:     faker.Name(),
            Email:    faker.Email(),
            Password: faker.Password(),
        }
        admin.SetTimestamps()
        if err := db.Create(&admin).Error; err != nil {
            return err
        }
    }
    return nil
}

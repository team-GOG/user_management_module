package user_management_model

import (
    "time"
    "fmt"

    "github.com/bxcodec/faker/v4"
    "gorm.io/gorm"
)

type Admin struct {
    ID   uint                 `gorm:"primaryKey"`
    Name string               `gorm:"not null"`
    Email string              `gorm:"unique;not null"`
    Password  string          `gorm:"type:varchar(255)" json:"-"`
    CreatedAt int64           `json:"created_at"`
    UpdatedAt int64           `json:"updated_at"`
    DeletedAt gorm.DeletedAt  `gorm:"index" json:"-"`
    Roleable []Roleable       `gorm:"polymorphic:Roleable;"`
}

// Implement RoleableEntity for Admin
func (a Admin) GetID() uint {
    return a.ID
}

func (a Admin) GetType() string {
    return "admins"
}



// CreateAdmin creates a new admin
func CreateAdmin(db *gorm.DB, admin *Admin) (err error) {

    // Set timestamps
    admin.CreatedAt = time.Now().Unix()
    admin.UpdatedAt = time.Now().Unix()

    err = db.Create(&admin).Error

    return
}

// GetAllAdmins retrieves all admins
func GetAllAdmins(db *gorm.DB, admins *[]Admin) (err error) {
    err = db.Find(&admins).Error

    return 
}

// GetAdmin retrieves a admin by ID
func GetAdmin(db *gorm.DB, id string, admin *Admin) (err error) {

    err = db.First(&admin, id).Error

    return
}

// UpdateAdmin updates an existing admin by ID
func UpdateAdmin(db *gorm.DB, admin *Admin) (err error) {

    // Update timestamps
    admin.UpdatedAt = time.Now().Unix()

    result := db.Model(&Admin{}).Where("id = ?", admin.ID).Updates(admin)

    err = result.Error

    if result.RowsAffected == 0 && err == nil {
      err = fmt.Errorf("no admin found with ID: %d", admin.ID)
    }

    return 
}

// DeleteAdmin deletes a admin by ID
func DeleteAdmin(db *gorm.DB, id string) (err error) {
    err = db.Delete(&Admin{}, id).Error;

    return
}

func GenerateFakeAdmins(db *gorm.DB, count int) error {

    for i := 0; i < count; i++ {

        admin := Admin{
            Name:      faker.Name(),
            Email:     faker.Email(),
            Password:  faker.Password(),
            CreatedAt: time.Now().Unix(),
            UpdatedAt: time.Now().Unix(),
        }

        // Create admin
        if err := db.Create(&admin).Error; err != nil {
            return err
        }

    }

    return nil
}

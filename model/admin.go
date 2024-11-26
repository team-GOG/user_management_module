package user_management_model

import (
    "time"

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

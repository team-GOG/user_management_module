package user_management_model

import "gorm.io/gorm"

type Permission struct {
    ID          uint   `gorm:"primaryKey"`
    Name        string `gorm:"unique;not null"`
}

// CreatePermissions inserts an array of permissions into the database
func CreatePermissions(db *gorm.DB, permissions []Permission) error {
    // Use GORM's Create method for batch insertion
    if err := db.Create(&permissions).Error; err != nil {
        return err
    }
    return nil
}

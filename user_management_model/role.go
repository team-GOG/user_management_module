package user_management_model

import (
    "errors"
    "fmt"
    "time"

    "github.com/bxcodec/faker/v4"
    "gorm.io/gorm"
)

// Role represents the roles table in the database.
type Role struct {
    ID          uint           `gorm:"primaryKey" json:"id"`
    Name        string         `gorm:"unique;not null" json:"name"`
    Permissions []Permission   `gorm:"many2many:role_permissions;" json:"permissions"`
    CreatedAt   int64          `json:"created_at"`
    UpdatedAt   int64          `json:"updated_at"`
    DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// Implement RoleableEntity for Role
func (r Role) GetID() uint {
    return r.ID
}

func (r Role) GetType() string {
    return "roles"
}

// SetTimestamps updates CreatedAt and UpdatedAt fields for the role.
func (r *Role) SetTimestamps() {
    now := time.Now().Unix()
    if r.CreatedAt == 0 {
        r.CreatedAt = now
    }
    r.UpdatedAt = now
}

// CreateRole inserts a new role into the database.
func CreateRole(db *gorm.DB, role *Role) error {
    role.SetTimestamps()
    return db.Create(&role).Error
}

// GetAllRoles retrieves all roles from the database.
func GetAllRoles(db *gorm.DB) ([]Role, error) {
    var roles []Role
    err := db.Find(&roles).Error
    return roles, err
}

// GetRole retrieves a role by ID from the database.
func GetRole(db *gorm.DB, id string) (Role, error) {
    var role Role
    err := db.First(&role, id).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return Role{}, fmt.Errorf("role with ID %s not found", id)
    }
    return role, err
}

// UpdateRole updates an existing role's information in the database.
func UpdateRole(db *gorm.DB, role *Role) error {
    role.SetTimestamps()
    result := db.Model(&Role{}).Where("id = ?", role.ID).Updates(role)

    if result.RowsAffected == 0 && result.Error == nil {
        return fmt.Errorf("no role found with ID: %d", role.ID)
    }
    return result.Error
}

// DeleteRole deletes a role by ID from the database.
func DeleteRole(db *gorm.DB, id string) error {
    result := db.Delete(&Role{}, id)
    if result.RowsAffected == 0 && result.Error == nil {
        return fmt.Errorf("no role found with ID: %s", id)
    }
    return result.Error
}

// GenerateFakeRoles generates a specified number of fake roles for testing.
func GenerateFakeRoles(db *gorm.DB, count int) error {
    for i := 0; i < count; i++ {
        role := Role{
            Name: faker.Name(),
        }
        role.SetTimestamps()
        if err := db.Create(&role).Error; err != nil {
            return err
        }
    }
    return nil
}

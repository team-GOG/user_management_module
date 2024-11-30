package user_management_model

import (
    "time"
    "fmt"

    "github.com/bxcodec/faker/v4"
    "gorm.io/gorm"
)

type Role struct {
    ID          uint   `gorm:"primaryKey"`
    Name        string `gorm:"unique;not null"`
    Permissions []Permission `gorm:"many2many:role_permissions;"`
    CreatedAt int64           `json:"created_at"`
    UpdatedAt int64           `json:"updated_at"`
    DeletedAt gorm.DeletedAt  `gorm:"index" json:"-"`
}


// Implement RoleableEntity for Role
func (r Role) GetID() uint {
    return r.ID
}

func (r Role) GetType() string {
    return "roles"
}

// CreateRole creates a new role
func CreateRole(db *gorm.DB, role *Role) (err error) {

    // Set timestamps
    role.CreatedAt = time.Now().Unix()
    role.UpdatedAt = time.Now().Unix()

    err = db.Create(&role).Error

    return
}

// GetAllRoles retrieves all roles
func GetAllRoles(db *gorm.DB, roles *[]Role) (err error) {

    err = db.Find(&roles).Error

    return 

}

// GetRole retrieves a role by ID
func GetRole(db *gorm.DB, id string) (role Role, err error) {

    err = db.First(&role, id).Error

    return
}

// UpdateRole updates an existing role by ID
func UpdateRole(db *gorm.DB, role *Role) (err error) {
    result := db.Model(&Role{}).Where("id = ?", role.ID).Updates(role)

    err = result.Error

    if result.RowsAffected == 0 && err == nil {
      err = fmt.Errorf("no role found with ID: %d", role.ID)
    }

    return 
}

// DeleteRole deletes a role by ID
func DeleteRole(db *gorm.DB, id string) (err error) {
    err = db.Delete(&Role{}, id).Error;

    return
}

func GenerateFakeRoles(db *gorm.DB, count int) error {

    for i := 0; i < count; i++ {
        role := Role{
            Name:      faker.Name(),
            CreatedAt: time.Now().Unix(),
            UpdatedAt: time.Now().Unix(),
        }

        // Create role
        if err := db.Create(&role).Error; err != nil {
            return err
        }

    }

    return nil
}

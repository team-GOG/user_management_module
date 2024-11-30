package user_management_model

import (
    "errors"
    "gorm.io/gorm"
    "time"
)

var DB *gorm.DB

func GetDB() *gorm.DB {
	return DB
}

// CreateSuperAdmin creates a superadmin user, assigns a role, and grants all permissions
func CreateSuperAdmin(db *gorm.DB, superAdmin Admin) error {
    return db.Transaction(func(tx *gorm.DB) error {
        // Check if superadmin already exists
        var existingAdmin Admin
        if err := tx.Where("email = ?", superAdmin.Email).First(&existingAdmin).Error; err == nil {
            return errors.New("superadmin with this email already exists")
        }

        // Create superadmin user
        superAdmin.CreatedAt = time.Now().Unix()
        superAdmin.UpdatedAt = time.Now().Unix()

        if err := tx.Create(&superAdmin).Error; err != nil {
            return err
        }

		
        // Check if superadmin role already exists
		var superAdminRole Role

        if err := tx.Where("name = ?", "SuperAdmin").First(&superAdminRole).Error; err != nil {
			// Create a role for superadmin
			superAdminRole = Role{
				Name:      "SuperAdmin",
				CreatedAt: time.Now().Unix(),
				UpdatedAt: time.Now().Unix(),
			}
		}


        if err := tx.Create(&superAdminRole).Error; err != nil {
            return err
        }

		permissions, err := GetAllPermissions(tx)

		if err != nil {
			return err
		}

        // Associate all permissions to the role
        if err := tx.Model(&superAdminRole).Association("Permissions").Append(permissions); err != nil {
            return err
        }

        // Associate the role with the superadmin
        if err := tx.Model(&superAdmin).Association("Roleable").Append(superAdminRole); err != nil {
            return err
        }

        return nil
    })
}


// CreateAllNeededPermissions creates all needed permissions
func CreateAllNeededPermissions(db *gorm.DB) error {
    return db.Transaction(func(tx *gorm.DB) error {

        permissions := []Permission{
            {Name: "CreateUser"},
            {Name: "UpdateUser"},
            {Name: "DeleteUser"},
            {Name: "CreateAdmin"},
            {Name: "UpdateAdmin"},
            {Name: "DeleteAdmin"},
            {Name: "CreateRole"},
            {Name: "UpdateRole"},
            {Name: "DeleteRole"},
            {Name: "CreatePermission"},
            {Name: "UpdatePermission"},
            {Name: "DeletePermission"},
        }

		CreatePermissions(tx, permissions)

        return nil
    })
}





package user_management_model

type RolePermission struct {
    RoleID       uint `gorm:"primaryKey"`
    PermissionID uint `gorm:"primaryKey"`
}

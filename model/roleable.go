package user_management_model

type Roleable struct {
    ID           uint   `gorm:"primaryKey"`
    RoleID       uint   `gorm:"not null"`
    Role         Role   `gorm:"foreignKey:RoleID"`
    RoleableID   uint   `gorm:"not null"`  // ID of the associated entity
    RoleableType string `gorm:"not null"`  // Type of the associated entity (e.g., "users", "admins")
}

// Define an interface for Roleable entities
type RoleableEntity interface {
    GetID() uint
    GetType() string
}

// Add a method to assign a role to any Roleable entity
func AssignRole(db *gorm.DB, roleID uint, entity RoleableEntity) error {
    roleable := Roleable{
        RoleID:       roleID,
        RoleableID:   entity.GetID(),
        RoleableType: entity.GetType(),
    }
    return db.Create(&roleable).Error
}

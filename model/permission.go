package user_management_model


type Permission struct {
    ID          uint   `gorm:"primaryKey"`
    Name        string `gorm:"unique;not null"`
}


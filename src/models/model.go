package models

// import (
// 	"github.com/jinzhu/gorm"
// 	uuid "github.com/satori/go.uuid"
// )

type Model struct {
	Id uint `json:"id"`
}

// type Base struct {
// 	ID uuid.UUID `gorm:"type:char(36);primary_key"`
// 	// add more common columns like CreatedAt
// 	// CreatedAt *time.Time
// 	// ...
// }

// This functions are called before creating Base
// func (base *Base) BeforeCreate(scope *gorm.Scope) error {
// 	return scope.SetColumn("ID", uuid.NewV4())
// }

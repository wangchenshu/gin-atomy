package model

// Host
type Host struct {
	Address  string `json:"address"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DB       string `json:"db"`
}

// Products
type Products struct {
	ID    int    `json:"id" gorm:"column:id"`
	Name  string `json:"name" gorm:"type:varchar(255)"`
	Price string `json:"price" gorm:"type:varchar(100)"`
	Point string `json:"point" gorm:"type:varchar(100)"`
}

package domain

type BaseUser struct {
	Name     string `gorm:"index;not null; type:varchar(60)" json:"name"`
	Username string `gorm:"uniqueIndex;not null;type:varchar(60)" json:"username" validate:"required"`
	Intro    string `gorm:"not null; type:varchar(512)" json:"intro"`
	Avatar   string `gorm:"type:varchar(1024)" json:"avatar"`
}

type Avatar struct {
	Avatar string `json:"avatar"`
}

package schema

import "time"

type Author struct {
	Id          int       `gorm:"column:id;type:int;primaryKey;autoIncrement" json:"id"`                                                       // 作者ID
	Name        string    `gorm:"column:name;type:varchar(100);not null" json:"name"`                                                          // 作者姓名
	CountryId   int       `gorm:"column:country_id;type:int;not null" json:"country_id"`                                                       // 国籍（关联国家表）
	Description string    `gorm:"column:description;type:varchar(500);default:null" json:"description,omitempty"`                              // 作者简介
	CreatedDate time.Time `gorm:"column:created_date;type:datetime;default:current_timestamp" json:"created_date"`                             // 创建时间
	UpdatedDate time.Time `gorm:"column:updated_date;type:datetime;default:current_timestamp on update current_timestamp" json:"updated_date"` // 更新时间
	Country     Country   `gorm:"foreignKey:CountryId" json:"country,omitempty"`                                                               // 关联国家信息
}

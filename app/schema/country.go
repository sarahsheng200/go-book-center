package schema

import "time"

type Country struct {
	Id          int       `gorm:"column:id;type:int;primaryKey;autoIncrement" json:"id"`                                                       // 国家ID
	Name        string    `gorm:"column:name;type:varchar(100);not null;uniqueIndex:uk_name" json:"name"`                                      // 国家名称
	Code        string    `gorm:"column:code;type:varchar(10);not null;uniqueIndex:uk_code" json:"code"`                                       // 国家代码（如CN、US）
	Description string    `gorm:"column:description;type:varchar(500);default:null" json:"description,omitempty"`                              // 描述
	CreatedDate time.Time `gorm:"column:created_date;type:datetime;default:current_timestamp" json:"created_date"`                             // 创建时间
	UpdatedDate time.Time `gorm:"column:updated_date;type:datetime;default:current_timestamp on update current_timestamp" json:"updated_date"` // 更新时间
}

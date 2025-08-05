package schema

import (
	"time"
)

type Category struct {
	Id          int64     `gorm:"column:id;type:bigint;primaryKey;autoIncrement" json:"id"`                                                  // 分类ID
	Name        string    `gorm:"column:name;type:varchar(50);not null" json:"name"`                                                         // 分类名
	ParentId    *int64    `gorm:"column:parent_id;type:bigint;default:null" json:"parent_id,omitempty"`                                      // 上级分类ID（NULL为顶级）
	Remark      string    `gorm:"column:remark;type:varchar(500);default:null" json:"remark,omitempty"`                                      // 备注
	CreatedTime time.Time `gorm:"column:created_time;type:datetime;default:current_timestamp" json:"created_time"`                           // 创建时间
	CreatedId   int       `gorm:"column:created_id;type:int;not null" json:"created_id"`                                                     // 创建人
	UpdateId    *int      `gorm:"column:update_id;type:int;default:null" json:"update_id,omitempty"`                                         // 更新人
	UpdateTime  time.Time `gorm:"column:update_time;type:datetime;default:current_timestamp on update current_timestamp" json:"update_time"` // 更新时间

	_ struct{} `gorm:"uniqueIndex:uk_name_parent:name,parent_id"` // 同一父级分类名唯一
}

type Categories []Category

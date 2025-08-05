package schema

import "time"

// BookNFC 书籍NFC表结构体（实体书籍唯一标识）
type BookNFC struct {
	Id          int       `gorm:"column:id;type:int;primaryKey;autoIncrement" json:"id"`                                                       // NFC记录ID
	BookId      int       `gorm:"column:book_id;type:int;not null" json:"book_id"`                                                             // 关联书籍
	NfcNumber   string    `gorm:"column:nfc_number;type:varchar(45);not null;uniqueIndex:uk_nfc" json:"nfc_number"`                            // NFC唯一编号
	Status      int8      `gorm:"column:status;type:tinyint;default:0" json:"status"`                                                          // 状态（0-库存中，1-已售出，2-损坏）
	IsDeleted   int8      `gorm:"column:is_deleted;type:tinyint;default:0" json:"is_deleted"`                                                  // 是否删除
	CreatedDate time.Time `gorm:"column:created_date;type:datetime;default:current_timestamp" json:"created_date"`                             // 创建时间
	UpdatedDate time.Time `gorm:"column:updated_date;type:datetime;default:current_timestamp on update current_timestamp" json:"updated_date"` // 更新时间

	Book Book `gorm:"foreignKey:BookId" json:"book,omitempty"` // 关联书籍信息
}

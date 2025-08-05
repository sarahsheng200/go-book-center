package schema

import "time"

// UserBook 对应数据库中的user_book表（用户与书籍的关联：收藏/购买/借阅）
type UserBook struct {
	Id          int       `gorm:"column:id;type:int;primaryKey;autoIncrement" json:"id"`                           // 关联ID
	UserId      int       `gorm:"column:user_id;type:int;not null" json:"user_id"`                                 // 用户ID
	BookId      int       `gorm:"column:book_id;type:int;not null" json:"book_id"`                                 // 书籍ID
	BookNfcId   *int      `gorm:"column:book_nfc_id;type:int;default:null" json:"book_nfc_id,omitempty"`           // 书籍NFC ID（购买/借阅时必填）
	Status      int8      `gorm:"column:status;type:tinyint;default:0" json:"status"`                              // 关联类型（0-收藏，1-购买，2-借阅）
	CreatedDate time.Time `gorm:"column:created_date;type:datetime;default:current_timestamp" json:"created_date"` // 关联时间
	_           struct{}  `gorm:"uniqueIndex:uk_user_book_status:user_id,book_nfc_id,status"`                      // 唯一约束：同一用户对同一NFC书籍的同一操作唯一
	User        User      `gorm:"foreignKey:UserId" json:"user,omitempty"`                                         // 关联用户信息
	Book        Book      `gorm:"foreignKey:BookId" json:"book,omitempty"`                                         // 关联书籍信息
	BookNFC     *BookNFC  `gorm:"foreignKey:BookNfcId" json:"book_nfc,omitempty"`                                  // 关联NFC信息
}

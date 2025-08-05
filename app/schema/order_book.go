package schema

import "time"

// OrderBook 订单明细表结构体（关联订单与书籍NFC）
type OrderBook struct {
	Id           int       `gorm:"column:id;type:int;primaryKey;autoIncrement" json:"id"`                                                       // 明细ID
	OrderId      int       `gorm:"column:order_id;type:int;not null" json:"order_id"`                                                           // 关联订单
	BookNfcId    int       `gorm:"column:book_nfc_id;type:int;not null" json:"book_nfc_id"`                                                     // 关联实体书籍（NFC）
	UnitPrice    float64   `gorm:"column:unit_price;type:decimal(10,2);default:0.00" json:"unit_price"`                                         // 单本售价
	BookName     string    `gorm:"column:book_name;type:varchar(100);not null" json:"book_name"`                                                // 下单时的书籍名称
	BookCategory string    `gorm:"column:book_category;type:varchar(50);not null" json:"book_category"`                                         // 下单时的书籍分类
	Status       int8      `gorm:"column:status;type:tinyint;default:0" json:"status"`                                                          // 状态（0-正常，4-退款中，5-已退款）
	CreatedDate  time.Time `gorm:"column:created_date;type:datetime;default:current_timestamp" json:"created_date"`                             // 创建时间
	UpdatedDate  time.Time `gorm:"column:updated_date;type:datetime;default:current_timestamp on update current_timestamp" json:"updated_date"` // 更新时间
	Order        Order     `gorm:"foreignKey:OrderId" json:"order,omitempty"`                                                                   // 关联订单信息
	BookNFC      BookNFC   `gorm:"foreignKey:BookNfcId" json:"book_nfc,omitempty"`                                                              // 关联实体书籍信息
}

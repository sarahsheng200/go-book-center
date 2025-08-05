package schema

import "time"

// Order 订单表结构体
type Order struct {
	Id          int         `gorm:"column:id;type:int;primaryKey;autoIncrement" json:"id"`                                                       // 订单ID
	OrderNumber string      `gorm:"column:order_number;type:varchar(45);not null;uniqueIndex:uk_order_number" json:"order_number"`               // 订单编号
	BuyerId     int         `gorm:"column:buyer_id;type:int;not null" json:"buyer_id"`                                                           // 买家ID
	SellerId    int         `gorm:"column:seller_id;type:int;not null" json:"seller_id"`                                                         // 卖家ID（默认平台管理员）
	TotalPrice  float64     `gorm:"column:total_price;type:decimal(10,2);default:0.00" json:"total_price"`                                       // 订单总金额
	Status      int8        `gorm:"column:status;type:tinyint;default:0" json:"status"`                                                          // 状态（0-未付款，1-已付款，2-已取消，3-已完成）
	CreatedDate time.Time   `gorm:"column:created_date;type:datetime;default:current_timestamp" json:"created_date"`                             // 创建时间
	CreatorId   *int        `gorm:"column:creator_id;type:int;default:null" json:"creator_id,omitempty"`                                         // 创建人
	UpdatedDate time.Time   `gorm:"column:updated_date;type:datetime;default:current_timestamp on update current_timestamp" json:"updated_date"` // 更新时间
	UpdatorId   *int        `gorm:"column:updator_id;type:int;default:null" json:"updator_id,omitempty"`                                         // 更新人
	Buyer       User        `gorm:"foreignKey:BuyerId" json:"buyer,omitempty"`                                                                   // 关联买家信息
	Seller      User        `gorm:"foreignKey:SellerId" json:"seller,omitempty"`                                                                 // 关联卖家信息
	Items       []OrderBook `gorm:"foreignKey:OrderId" json:"items,omitempty"`                                                                   // 订单明细
}

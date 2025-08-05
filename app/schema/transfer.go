package schema

import "time"

// Transfer 资金交易表结构体（用户充值提现记录）
type Transfer struct {
	Id             int       `gorm:"column:id;type:int;primaryKey;autoIncrement" json:"id"`                                                       // 交易ID
	TransferNumber string    `gorm:"column:transfer_number;type:varchar(45);not null;uniqueIndex:uk_transfer_number" json:"transfer_number"`      // 交易编号
	UserId         int       `gorm:"column:user_id;type:int;not null" json:"user_id"`                                                             // 关联用户
	OrderId        *int      `gorm:"column:order_id;type:int;default:null" json:"order_id,omitempty"`                                             // 关联订单（充值/提现可为NULL）
	Amount         float64   `gorm:"column:amount;type:decimal(10,2);default:0.00" json:"amount"`                                                 // 交易金额（正数：收入/充值；负数：支出/消费）
	Type           int8      `gorm:"column:type;type:tinyint;default:0" json:"type"`                                                              // 类型（0-充值，1-消费，2-提现，3-退款）
	CreatorId      *int      `gorm:"column:creator_id;type:int;default:null" json:"creator_id,omitempty"`                                         // 操作人
	UpdatorId      *int      `gorm:"column:updator_id;type:int;default:null" json:"updator_id,omitempty"`                                         // 更新人
	CreatedDate    time.Time `gorm:"column:created_date;type:datetime;default:current_timestamp" json:"created_date"`                             // 创建时间
	UpdatedDate    time.Time `gorm:"column:updated_date;type:datetime;default:current_timestamp on update current_timestamp" json:"updated_date"` // 更新时间
	_              struct{}  `gorm:"index:idx_user_type_time:user_id,type,created_date"`                                                          // 复合索引：用户+类型+时间
	User           User      `gorm:"foreignKey:UserId" json:"user,omitempty"`                                                                     // 关联用户信息
	Order          *Order    `gorm:"foreignKey:OrderId" json:"order,omitempty"`                                                                   // 关联订单信息
}

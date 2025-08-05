package schema

import (
	"time"
)

type User struct {
	Id             int       `gorm:"column:id;type:int;primaryKey;autoIncrement" json:"id"`                                                       // 用户ID
	AccountNumber  string    `gorm:"column:account_number;type:varchar(45);not null;uniqueIndex:uk_account" json:"account_number"`                // 账号
	Name           string    `gorm:"column:name;type:varchar(100);not null" json:"name"`                                                          // 姓名
	Password       string    `gorm:"column:password;type:varchar(100);not null" json:"-"`                                                         // 密码（加密存储）
	Email          string    `gorm:"column:email;type:varchar(45);not null;uniqueIndex:uk_email" json:"email"`                                    // 邮箱
	AccountBalance float64   `gorm:"column:account_balance;type:decimal(10,2);default:0.00" json:"account_balance"`                               // 账户余额
	IsAdmin        int8      `gorm:"column:is_admin;type:tinyint;default:0" json:"is_admin"`                                                      // 是否管理员（0-否，1-是）
	IsDeleted      int8      `gorm:"column:is_deleted;type:tinyint;default:0" json:"is_deleted"`                                                  // 是否删除（0-否，1-是）
	CreatedDate    time.Time `gorm:"column:created_date;type:datetime;default:current_timestamp" json:"created_date"`                             // 创建时间
	UpdatedDate    time.Time `gorm:"column:updated_date;type:datetime;default:current_timestamp on update current_timestamp" json:"updated_date"` // 更新时间

}

type UserAuth struct {
	Id      int
	Name    string
	IsAdmin int
}

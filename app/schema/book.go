package schema

//order, category,transfer
import (
	"time"
)

type Book struct {
	Id          int       `gorm:"column:id;type:int;primaryKey;autoIncrement" json:"id"`                                                       // 书籍ID
	Name        string    `gorm:"column:name;type:varchar(100);not null" json:"name,omitempty"`                                                // 书名
	CategoryId  int64     `gorm:"column:category_id;type:bigint;not null" json:"category_id,omitempty"`                                        // 所属分类
	AuthorId    int       `gorm:"column:author_id;type:int;not null" json:"author_id,omitempty"`                                               // 作者ID
	SellPrice   float64   `gorm:"column:sell_price;type:decimal(10,2);default:0.00" json:"sell_price,omitempty"`                               // 售价
	StockCount  int       `gorm:"column:stock_count;type:int;default:0" json:"stock_count,omitempty"`                                          // 库存数量
	Description string    `gorm:"column:description;type:varchar(500);default:null" json:"description,omitempty"`                              // 书籍简介
	IsDeleted   int8      `gorm:"column:is_deleted;type:tinyint;default:0" json:"is_deleted,omitempty"`                                        // 是否删除（0-否，1-是）
	CreatedDate time.Time `gorm:"column:created_date;type:datetime;default:current_timestamp" json:"created_date"`                             // 创建时间
	UpdatedDate time.Time `gorm:"column:updated_date;type:datetime;default:current_timestamp on update current_timestamp" json:"updated_date"` // 更新时间

	Category      Category   `gorm:"foreignKey:CategoryId" json:"category,omitempty"` // 关联分类信息
	Author        Author     `gorm:"foreignKey:AuthorId" json:"author,omitempty"`     // 关联作者信息
	CategoryChain []Category `gorm:"foreignKey:ParentId" json:"category_chain,omitempty"`
}

type BookList []Book

type UpdateBook struct {
	Name        string  `gorm:"column:name;type:varchar(100);not null" json:"name"`
	CategoryId  int64   `gorm:"column:category_id;type:bigint;not null" json:"category_id"`
	AuthorId    int     `gorm:"column:author_id;type:int;not null" json:"author_id"`
	SellPrice   float64 `gorm:"column:sell_price;type:decimal(10,2)" json:"sell_price"`
	StockCount  int     `gorm:"column:stock_count;type:int;default:0" json:"stock_count"`
	Description string  `gorm:"column:description;type:varchar(500)" json:"description"`
}

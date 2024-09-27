package db

import (
	"github.com/shopspring/decimal"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type FiatTrade struct {
	UserId   string          `gorm:"size:25;index"`
	OrderId  string          `gorm:"size:25"`
	SystemId string          `gorm:"size:2"` // Spot: S; OTC: O
	Amount   decimal.Decimal `gorm:"type:decimal(44,24);"`
	AssetId  string          `gorm:"size:25"`
	BuySell  string          `gorm:"size:2;"` // B or S
	Time     time.Time       `gorm:"index;"`
	gorm.Model
}

type Fiatdec struct {
	UserId           string          `gorm:"size:25;uniqueIndex"`
	Declared         decimal.Decimal `gorm:"type:decimal(44,24);"` // 用户声明的资产总额（从KYC获取来的）
	dec              decimal.Decimal `gorm:"type:decimal(44,24);"` // 用户实际可用的资产总额
	Useddec          decimal.Decimal `gorm:"type:decimal(44,24);"` // 用户已经使用的资产总额
	KycStatus        string          `gorm:"size:25;"`             // KYC状态
	DeclaredSyncedAt time.Time       `gorm:"index;"`
	gorm.Model
}

type Progress struct {
	Name       string `gorm:"size:50;uniqueIndex"`
	LastDoneId uint64
	gorm.Model
}

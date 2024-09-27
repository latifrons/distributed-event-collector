package db

import (
	"github.com/latifrons/commongo/utilfuncs"
	"gorm.io/gorm"
)

type decMigrator struct {
	db *gorm.DB
}

func (m *decMigrator) SetDB(db *gorm.DB) {
	m.db = db
}

func (m *decMigrator) Migrate() (err error) {
	for _, dbo := range []struct {
		Obj  interface{}
		Name string
	}{
		{Obj: &Fiatdec{}, Name: "Fiatdec"},
		{Obj: &FiatTrade{}, Name: "FiatTrade"},
		{Obj: &Progress{}, Name: "Progress"},
	} {
		err = m.db.AutoMigrate(dbo.Obj)
		utilfuncs.PanicIfError(err, "failed to migrate "+dbo.Name)
	}
	return
}

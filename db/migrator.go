package db

import (
	"github.com/latifrons/commongo/utilfuncs"
	"gorm.io/gorm"
)

type templatexxMigrator struct {
	db *gorm.DB
}

func (m *templatexxMigrator) SetDB(db *gorm.DB) {
	m.db = db
}

func (m *templatexxMigrator) Migrate() (err error) {
	for _, dbo := range []struct {
		Obj  interface{}
		Name string
	}{
		{Obj: &Fiattemplatexx{}, Name: "Fiattemplatexx"},
		{Obj: &FiatTrade{}, Name: "FiatTrade"},
		{Obj: &Progress{}, Name: "Progress"},
	} {
		err = m.db.AutoMigrate(dbo.Obj)
		utilfuncs.PanicIfError(err, "failed to migrate "+dbo.Name)
	}
	return
}

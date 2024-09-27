package db

import (
	"database/sql"
	"fmt"
	"github.com/latifrons/distributed-event-collector/consts"
	"github.com/latifrons/latigo/berror"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

func WithStackErr(err error) error {
	if err == nil {
		return err
	}
	return errors.Wrap(err, "DB error")
}

func CheckNonExistError(errx error) (exists bool, err error) {

	if errx == nil {
		exists = true
		return
	}
	if errx == gorm.ErrRecordNotFound {
		// clear the error
		exists = false
		err = nil
		return
	}
	err = errx
	return
}

func CheckSaveOk(result *gorm.DB, targetName string, expectedRowsAffected int) (err error) {
	if result.Error != nil {
		err = result.Error
		return
	}
	if result.RowsAffected != int64(expectedRowsAffected) {
		err = berror.NewBusinessFail(nil, consts.ErrCAS, fmt.Sprintf("rows not affected: %s. expected %d, actual %d", targetName, expectedRowsAffected, result.RowsAffected))
		return
	}
	return
}

//func SqlNullTimeToInt64Default(value sql.NullTime) int64 {
//	if value.Valid {
//		return value.Time.Unix()
//	}
//	return 0
//}

func SqlNullTimeToMillisecondInt64Default(value sql.NullTime) uint64 {
	if value.Valid {
		return uint64(value.Time.UnixMilli())
	}
	return 0
}

func TimeToInt64(t *time.Time) *int64 {
	if t == nil {
		return nil
	}
	v := t.Unix()
	return &v
}

func SqlNullStringToStringDefault(value sql.NullString) string {
	if value.Valid {
		return value.String
	}
	return ""
}

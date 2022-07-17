package main

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"os"
	"regexp"
	"streamtodb/domain/service"
	"streamtodb/infra/input"
	"streamtodb/infra/persistence"
	"streamtodb/interfaces"
	"testing"
)

const bufSize = 1024 * 1024

var (
	db      *gorm.DB
	sqlMock sqlmock.Sqlmock
)

func initDbMock(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf(err.Error())
	}
	gormDb, err := gorm.Open("postgres", db)
	if err != nil {
		t.Errorf(err.Error())
	}
	gormDb.LogMode(true)
	return gormDb, mock
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestCreatePort(t *testing.T) {
	db, sqlMock = initDbMock(t)
	handler := interfaces.NewPortHandler(service.NewPortService(persistence.NewPortRepository(db)))
	producer := input.NewProducer()
	go producer.Start(dataStream("testdata/ports.json"))

	sqlCreate("AEAJM")
	sqlCreate("ZWUTA")
	connect(producer, handler)
	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdatePort(t *testing.T) {
	db, sqlMock = initDbMock(t)
	handler := interfaces.NewPortHandler(service.NewPortService(persistence.NewPortRepository(db)))
	producer := input.NewProducer()
	go producer.Start(dataStream("testdata/ports.json"))

	sqlUpdate("AEAJM")
	sqlUpdate("ZWUTA")
	connect(producer, handler)
	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func sqlCreate(codename string) {
	sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "ports" WHERE (codename = $1) LIMIT 1`)).
		WithArgs(codename).WillReturnError(gorm.ErrRecordNotFound)
	sqlMock.ExpectBegin()
	sqlString := regexp.QuoteMeta(`INSERT INTO "ports" ("codename","name","city","country","alias","regions","coordinates","province","timezone","unlocs","code","created_at","updated_at") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13) RETURNING "ports"."id"`)
	sqlMock.ExpectQuery(sqlString).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
		sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
		sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
		sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
	).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("2fa20b76-cff5-4f69-bc13-a801e5d09bbb"))
	sqlMock.ExpectCommit()
}

func sqlUpdate(codename string) {
	sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "ports" WHERE (codename = $1) LIMIT 1`)).
		WithArgs(codename).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("2fa20b76-cff5-4f69-bc13-a801e5d09bbb"))
	sqlMock.ExpectBegin()
	sqlString := regexp.QuoteMeta(`UPDATE "ports" SET "alias" = $1, "city" = $2, "code" = $3, "codename" = $4, "coordinates" = $5, "country" = $6, "id" = $7, "name" = $8, "province" = $9, "regions" = $10, "timezone" = $11, "unlocs" = $12, "updated_at" = $13  WHERE "ports"."id" = $14`)
	sqlMock.ExpectExec(sqlString).WithArgs(
		sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
		sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
		sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
		sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
	).WillReturnResult(sqlmock.NewResult(0, 1))
	sqlMock.ExpectCommit()
}

package repository_test

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"shrt-server/repository"
	"shrt-server/types/entity"
	"testing"
)

func dbMock(t *testing.T) (*sql.DB, *gorm.DB, sqlmock.Sqlmock) {
	sqldb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	gormdb, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqldb,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		t.Fatal(err)
	}
	return sqldb, gormdb, mock
}

func Test_shrtRepository_Create_Success(t *testing.T) {
	sqlDB, db, mock := dbMock(t)
	var anyTime = sqlmock.AnyArg()
	defer sqlDB.Close()

	shrtRepo := repository.NewShrtRepository(db)

	shrt := &entity.Shrt{
		LongURL: "https://google.com",
		Slug:    "slug456",
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `Shrts`").
		WithArgs(shrt.LongURL, shrt.Slug, 0, anyTime, anyTime).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := shrtRepo.Create(shrt)

	if err != nil {
		t.Errorf("Create() error = %v", err.Error())
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Expectations were not met: %s", err)
	}
}

func Test_shrtRepository_Create_Error(t *testing.T) {
	sqlDB, db, mock := dbMock(t)
	var anyTime = sqlmock.AnyArg()
	defer sqlDB.Close()

	shrtRepo := repository.NewShrtRepository(db)

	shrt := &entity.Shrt{
		LongURL: "https://google.com",
		Slug:    "slug456",
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `Shrts`").
		WithArgs(shrt.LongURL, shrt.Slug, 0, anyTime, anyTime).
		WillReturnResult(sqlmock.NewResult(1, 1)) // Simulating a successful query
	mock.ExpectCommit().
		WillReturnError(sql.ErrConnDone) // Simulating an error during commit

	err := shrtRepo.Create(shrt)

	if err == nil {
		t.Error("Expected an error, but got nil")
	} else if err != sql.ErrConnDone {
		t.Errorf("Expected connection error, got %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Expectations were not met: %s", err)
	}
}

func Test_shrtRepository_FindBySlug(t *testing.T) {

}

func Test_shrtRepository_UpdateVisit(t *testing.T) {

}

package account

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	gorm "github.com/jinzhu/gorm"
)

func Test_postgresRepository_Create(t *testing.T) {
	createQuery := regexp.QuoteMeta(`INSERT INTO "accounts" ("document_number") VALUES ($1) RETURNING "accounts"."id"`)
	document := "44639467828"
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	mockDB, err := gorm.Open("postgres", db)
	if err != nil {
		t.Fatalf("failed to connect database err: %s", err)
	}
	defer mockDB.Close()

	// Create OK
	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	mock.ExpectBegin()
	mock.ExpectQuery(createQuery).WithArgs(document).WillReturnRows(rows)
	mock.ExpectCommit()

	// Create Error
	mock.ExpectBegin()
	mock.ExpectQuery(createQuery).WithArgs(document).WillReturnError(fmt.Errorf("some error"))

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		a *Account
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Create OK",
			fields: fields{
				db: mockDB,
			},
			args: args{
				&Account{
					DocumentNumber: document,
				},
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "Create Error",
			fields: fields{
				db: mockDB,
			},
			args: args{
				&Account{
					DocumentNumber: document,
				},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &postgresRepository{
				db: tt.fields.db,
			}
			got, err := r.Create(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("postgresRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("postgresRepository.Create() = %v, want %v", got, tt.want)
			}
		})
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_postgresRepository_Get(t *testing.T) {
	query := regexp.QuoteMeta(`SELECT * FROM "accounts"  WHERE ("accounts"."id" = 1) ORDER BY "accounts"."id" ASC LIMIT 1`)
	id := 1
	document := "44639467828"
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	mockDB, err := gorm.Open("postgres", db)
	if err != nil {
		t.Fatalf("failed to connect database err: %s", err)
	}
	defer mockDB.Close()

	// Select OK
	rows := sqlmock.NewRows([]string{"id", "document_number"}).AddRow(id, document)
	mock.ExpectQuery(query).WillReturnRows(rows)

	// Select Error
	mock.ExpectQuery(query).WillReturnError(fmt.Errorf("some error"))

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Account
		wantErr bool
	}{
		{
			name: "Select OK",
			fields: fields{
				db: mockDB,
			},
			args: args{
				id: id,
			},
			want: &Account{
				ID:             id,
				DocumentNumber: document,
			},
			wantErr: false,
		},
		{
			name: "Select Error",
			fields: fields{
				db: mockDB,
			},
			args: args{
				id: id,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &postgresRepository{
				db: tt.fields.db,
			}
			got, err := r.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("postgresRepository.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postgresRepository.Get() = %v, want %v", got, tt.want)
			}
		})
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

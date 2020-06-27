package transaction

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/julioc98/pismo/internal/app/account"
)

func Test_postgresRepository_Create(t *testing.T) {
	createQuery := regexp.QuoteMeta(`INSERT INTO "transactions" ("account_id","operation_id","amount","created_at") VALUES ($1,$2,$3,$4) RETURNING "transactions"."id"`)
	accountID := 1
	operationID := 1
	amount := -10.01
	createdAt := time.Now()
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
	mock.ExpectQuery(createQuery).WithArgs(accountID, operationID, amount, createdAt).WillReturnRows(rows)
	mock.ExpectCommit()

	// Create Error
	mock.ExpectBegin()
	mock.ExpectQuery(createQuery).WithArgs(accountID, operationID, amount, createdAt).WillReturnError(fmt.Errorf("some error"))

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		a *Transaction
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
				&Transaction{
					AccountID:   accountID,
					OperationID: operationID,
					Amount:      amount,
					CreatedAt:   createdAt,
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
				&Transaction{
					AccountID:   accountID,
					OperationID: operationID,
					Amount:      amount,
					CreatedAt:   createdAt,
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
}

func Test_postgresRepository_Get(t *testing.T) {
	mainQuery := regexp.QuoteMeta(`SELECT * FROM "transactions" WHERE ("transactions"."id" = 1) ORDER BY "transactions"."id" ASC LIMIT 1`)
	accountQuery := regexp.QuoteMeta(`SELECT * FROM "accounts" WHERE ("id" IN ($1)) ORDER BY "accounts"."id" ASC`)
	operationQuery := regexp.QuoteMeta(`SELECT * FROM "operations" WHERE ("id" IN ($1)) ORDER BY "operations"."id" ASC`)
	id := 1
	accountID := 1
	operationID := 1
	amount := -10.01
	createdAt := time.Now()
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
	rows := sqlmock.NewRows([]string{"id", "account_id", "operation_id", "amount", "created_at"}).AddRow(id, accountID, operationID, amount, createdAt)
	mock.ExpectQuery(mainQuery).WillReturnRows(rows)

	accountRows := sqlmock.NewRows([]string{"id"}).AddRow(accountID)
	mock.ExpectQuery(accountQuery).WillReturnRows(accountRows)

	operationRows := sqlmock.NewRows([]string{"id"}).AddRow(operationID)
	mock.ExpectQuery(operationQuery).WillReturnRows(operationRows)

	// Select Error
	mock.ExpectQuery(mainQuery).WillReturnError(fmt.Errorf("some error"))

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
		want    *Transaction
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
			want: &Transaction{
				ID:        id,
				AccountID: accountID,
				Account: account.Account{
					ID: accountID,
				},
				OperationID: operationID,
				Operation: Operation{
					ID: operationID,
				},
				Amount:    amount,
				CreatedAt: createdAt,
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
}

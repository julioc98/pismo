package transaction

import (
	"errors"
	"testing"
)

type repositoryMoc struct {
	transaction *Transaction
	err         error
}

func (r repositoryMoc) Create(transaction *Transaction) (int, error) {
	return r.transaction.ID, r.err
}

func (r repositoryMoc) Get(id int) (*Transaction, error) {
	return r.transaction, r.err
}

func Test_service_Create(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		t *Transaction
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Create Transaction With Success",
			fields: fields{
				repo: repositoryMoc{
					transaction: &Transaction{
						ID: 1,
					},
					err: nil,
				},
			},
			args: args{
				t: &Transaction{},
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "Create Transaction With Error",
			fields: fields{
				repo: repositoryMoc{
					transaction: &Transaction{
						ID: 0,
					},
					err: errors.New("mock err"),
				},
			},
			args: args{
				t: &Transaction{},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service{
				repo: tt.fields.repo,
			}
			got, err := s.Create(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("service.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

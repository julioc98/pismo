package account

import (
	"errors"
	"reflect"
	"testing"
)

type repositoryMoc struct {
	account *Account
	err     error
}

func (r repositoryMoc) Create(account *Account) (int, error) {
	return r.account.ID, r.err
}

func (r repositoryMoc) Get(id int) (*Account, error) {
	return r.account, r.err
}

func Test_service_Create(t *testing.T) {
	type fields struct {
		repo Repository
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
			name: "Create Account With Success",
			fields: fields{
				repo: repositoryMoc{
					account: &Account{
						ID: 1,
					},
					err: nil,
				},
			},
			args:    args{},
			want:    1,
			wantErr: false,
		},
		{
			name: "Create Account With Error",
			fields: fields{
				repo: repositoryMoc{
					account: &Account{},
					err:     errors.New("mock err"),
				},
			},
			args:    args{},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service{
				repo: tt.fields.repo,
			}
			got, err := s.Create(tt.args.a)
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

func Test_service_Get(t *testing.T) {
	type fields struct {
		repo Repository
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
			name: "Get Account With Success",
			fields: fields{
				repo: repositoryMoc{
					account: &Account{
						ID:             1,
						DocumentNumber: "123456789",
					},
					err: nil,
				},
			},
			args: args{},
			want: &Account{
				ID:             1,
				DocumentNumber: "123456789",
			},
			wantErr: false,
		},
		{
			name: "Get Account With Error",
			fields: fields{
				repo: repositoryMoc{
					account: nil,
					err:     errors.New("mock err"),
				},
			},
			args:    args{},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service{
				repo: tt.fields.repo,
			}
			got, err := s.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

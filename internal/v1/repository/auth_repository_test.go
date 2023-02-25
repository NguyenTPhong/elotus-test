package repository

import (
	"elotus/internal/v1/repository/model"
	db2 "elotus/test/db"
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestAuthRepositoryImpl_CreateUser(t *testing.T) {

	type args struct {
		user *model.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Create user",
			args: args{
				user: &model.User{
					Username: "test_user",
					Password: "abc",
					Salt:     "123",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, cancelFunc := db2.GetTestDB()
			defer cancelFunc()
			u := &AuthRepositoryImpl{
				database: db,
			}
			if err := u.CreateUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAuthRepositoryImpl_FindFirstUser(t *testing.T) {

	type args struct {
		filter UserFilter
	}
	tests := []struct {
		name         string
		args         args
		mockDatabase func(db *gorm.DB)
		want         *model.User
		wantErr      bool
	}{
		// TODO: Add test cases.
		{
			name: "Find user by username",
			args: args{
				filter: UserFilter{
					Username: "test_user",
				},
			},
			mockDatabase: func(db *gorm.DB) {
				db.Create(&model.User{
					Username: "test_user",
					Password: "abc",
					Salt:     "123",
				})
			},
			want: &model.User{
				Username: "test_user",
				Password: "abc",
				Salt:     "123",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, cancelFunc := db2.GetTestDB()
			defer cancelFunc()
			u := &AuthRepositoryImpl{
				database: db,
			}

			tt.mockDatabase(db)

			got, err := u.FindFirstUser(tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindFirstUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Username, tt.want.Username) {
				t.Errorf("FindFirstUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

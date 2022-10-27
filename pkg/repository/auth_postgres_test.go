package repository

import (
	"github.com/kirill0909/resource-keeper-api/models"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"testing"
)

func TestRepository_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub DB connection", err.Error())
	}
	defer db.Close()

	repo := NewAuthPostgres(db)

	testTable := []struct {
		name    string
		mock    func()
		input   models.User
		want    int
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO users").
					WithArgs("John Down", "john@gmail.com", "JohnPass").WillReturnRows(rows)
			},
			input: models.User{
				Name:     "John Down",
				Email:    "john@gmail.com",
				Password: "JohnPass",
			},
			want: 1,
		},
		{
			name: "Empty field",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery("INSERT INTO users").
					WithArgs("John Down", "john@gmail.com", "").WillReturnRows(rows)
			},
			input: models.User{
				Name:     "John Down",
				Email:    "john@gmail.com",
				Password: "",
			},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mock()

			got, err := repo.CreateUser(testCase.input)
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestRepository_GetUser(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub DB connection", err.Error())
	}
	defer db.Close()

	repo := NewAuthPostgres(db)

	type args = struct {
		email    string
		password string
	}

	testTable := []struct {
		name    string
		mock    func()
		input   args
		want    models.User
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "email", "password"}).
					AddRow(1, "John Down", "john@gmail.com", "JohnPass")
				mock.ExpectQuery("SELECT (.+) FROM users").
					WithArgs("john@gmail.com", "JohnPass").WillReturnRows(rows)
			},
			input: args{"john@gmail.com", "JohnPass"},
			want:  models.User{Id: 1, Name: "John Down", Email: "john@gmail.com", Password: "JohnPass"},
		},
		{
			name: "Not Found",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "email", "password"})
				mock.ExpectQuery("SELECT (.+) FROM users").
					WithArgs("not", "found").WillReturnRows(rows)
			},
			input:   args{"not", "found"},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mock()

			got, err := repo.GetUser(testCase.input.email, testCase.input.password)
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.want, got)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}

}

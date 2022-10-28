package repository

import (
	"github.com/kirill0909/resource-keeper-api/models"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"testing"
)

func TestUserResourcePostgres_CreateResource(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub DB connection", err.Error())
	}
	defer db.Close()

	repo := NewUserResourcePostgres(db)

	testTable := []struct {
		name    string
		mock    func()
		input   models.UserResource
		want    int
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO users_resources").
					WithArgs(1, "rname", "rlogin", "rpass").WillReturnRows(rows)
			},
			input: models.UserResource{
				UID:              1,
				ResourceName:     "rname",
				ResourceLogin:    "rlogin",
				ResourcePassword: "rpass",
			},
			want: 1,
		},
		{
			name: "Empty Value",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery("INSERT INTO users_resources").
					WithArgs(1, "rname", "rlogin", "").WillReturnRows(rows)
			},
			input: models.UserResource{
				UID:              1,
				ResourceName:     "rname",
				ResourceLogin:    "rlogin",
				ResourcePassword: "",
			},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mock()

			got, err := repo.CreateResource(testCase.input)
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

func TestUserResourcePostgres_GetAllResources(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub DB connection", err.Error())
	}
	defer db.Close()

	repo := NewUserResourcePostgres(db)

	testTable := []struct {
		name    string
		mock    func()
		userId  int
		want    []models.UserResource
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "user_id", "resource_name", "resource_login_enc", "resource_password_enc"}).
					AddRow(1, 1, "rname", "rlogin", "rpass").
					AddRow(2, 1, "rname", "rlogin", "rpass").
					AddRow(3, 1, "rname", "rlogin", "rpass")
				mock.ExpectQuery("SELECT (.+) FROM users_resources").WithArgs(1).WillReturnRows(rows)
			},
			want: []models.UserResource{
				{Id: 1, UID: 1, ResourceName: "rname", ResourceLogin: "rlogin", ResourcePassword: "rpass"},
				{Id: 2, UID: 1, ResourceName: "rname", ResourceLogin: "rlogin", ResourcePassword: "rpass"},
				{Id: 3, UID: 1, ResourceName: "rname", ResourceLogin: "rlogin", ResourcePassword: "rpass"},
			},
			userId: 1,
		},
		{
			name: "User id not found",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "user_id", "resource_name", "resource_login_enc", "resource_password_enc"})
				mock.ExpectQuery("SELECT (.+) FROM users_resources").WithArgs(2).WillReturnRows(rows)
			},
			want:   nil,
			userId: 2,
		},
		{
			name: "User id is negative number",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "user_id", "resource_name", "resource_login_enc", "resource_password_enc"})
				mock.ExpectQuery("SELECT (.+) FROM users_resources").WithArgs(-1).WillReturnRows(rows)
			},
			want:   nil,
			userId: -1,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mock()

			got, err := repo.GetAllResources(testCase.userId)
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

func TestUserResourcePostgres_GetById(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub DB connection", err.Error())
	}
	defer db.Close()

	repo := NewUserResourcePostgres(db)

	testTable := []struct {
		name       string
		mock       func()
		userId     int
		resourceId int
		want       models.UserResource
		wantErr    bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "user_id", "resource_name", "resource_login_enc", "resource_password_enc"}).
					AddRow(1, 1, "rname", "rlogin", "rpass")
				mock.ExpectQuery("SELECT (.+) FROM users_resources").WithArgs(1, 1).WillReturnRows(rows)
			},
			userId:     1,
			resourceId: 1,
			want:       models.UserResource{Id: 1, UID: 1, ResourceName: "rname", ResourceLogin: "rlogin", ResourcePassword: "rpass"},
		},
		{
			name: "No rows in resutl set",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "user_id", "resource_name", "resource_login_enc", "resource_password_enc"})
				mock.ExpectQuery("SELECT (.+) FROM users_resources").WithArgs(1, 1).WillReturnRows(rows)
			},
			userId:     1,
			resourceId: 1,
			wantErr:    true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mock()

			got, err := repo.GetById(testCase.userId, testCase.resourceId)
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

func TestUserResourcePostgres_DeleteResource(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub DB connection", err.Error())
	}
	defer db.Close()

	repo := NewUserResourcePostgres(db)

	testTable := []struct {
		name       string
		mock       func()
		userId     int
		resourceId int
		want       int
		wantErr    bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("DELETE FROM users_resources (.+)").WithArgs(1, 1).WillReturnRows(rows)
			},
			userId:     1,
			resourceId: 1,
			want:       1,
		},
		{
			name: "No found",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery("DELETE FROM users_resources").WithArgs(1, 1).WillReturnRows(rows)
			},
			userId:     1,
			resourceId: 1,
			wantErr:    true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mock()

			got, err := repo.DeleteResource(testCase.userId, testCase.resourceId)
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

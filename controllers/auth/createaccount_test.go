package controllers

import (
	"context"
	"fmt"
	"testing"

	dm "github.com/samverrall/spacecms-api/datastore/mocks"
	"github.com/samverrall/spacecms-api/gen/auth"
	"github.com/samverrall/spacecms-api/hasher/argon2id"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/mock"
)

func Test_authservice_CreateAccount(t *testing.T) {
	const (
		failedToInsertErr     = "failed@test.com"
		checkExistingUserFail = "existing@fail.com"
		successInsert         = "success@test.com"
	)

	var getGetUserByEmailMock = func(m *dm.DBInterface, t *testing.T, email string) {
		// GetUserByEmail(ctx context.Context, email string) (*auth.User, error)
		switch email {
		case checkExistingUserFail:
			m.On("GetUserByEmail", mock.AnythingOfType("*context.emptyCtx"), email).
				Return(nil, fmt.Errorf("Failed to check if user exists, unit test error."))
		default:
			m.On("GetUserByEmail", mock.AnythingOfType("*context.emptyCtx"), email).
				Return(nil, nil)
		}
	}

	var getCreateUserMock = func(m *dm.DBInterface, t *testing.T, email string) {
		switch email {
		case failedToInsertErr:
			m.On("CreateUser", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string"), email, mock.AnythingOfType("string"), mock.AnythingOfType("string")).
				Return(fmt.Errorf("failed to create user, unit test error"))
		default:
			m.On("CreateUser", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string"), email, mock.AnythingOfType("string"), mock.AnythingOfType("string")).
				Return(nil)
		}
	}

	var getDatastoreMock = func(t *testing.T, email string) *dm.DBInterface {
		var mock dm.DBInterface
		getGetUserByEmailMock(&mock, t, email)
		getCreateUserMock(&mock, t, email)
		return &mock
	}

	ctx := context.Background()
	tests := []struct {
		name    string
		payload *auth.User
		wantErr bool
	}{
		{
			name: "Check if user exists fail",
			payload: &auth.User{
				Email: checkExistingUserFail,
			},
			wantErr: true,
		},
		{
			name: "Insert user fail",
			payload: &auth.User{
				Email: failedToInsertErr,
			},
			wantErr: true,
		},
		{
			name: "Successfully inserted user",
			payload: &auth.User{
				Email: successInsert,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &authservice{
				logger: &logrus.Logger{},
				dbi:    getDatastoreMock(t, tt.payload.Email),
				hasher: argon2id.New(),
			}
			if err := s.CreateAccount(ctx, tt.payload); (err != nil) != tt.wantErr {
				t.Errorf("authservice.CreateAccount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

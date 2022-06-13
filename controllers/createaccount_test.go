package controllers

import (
	"context"
	"fmt"
	"testing"

	dm "github.com/samverrall/invoice-app/datastore/mocks"
	invoice "github.com/samverrall/invoice-app/gen/invoice"
	"github.com/samverrall/invoice-app/hasher/argon2id"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/mock"
)

func Test_invoicesrvc_CreateAccount(t *testing.T) {
	const (
		failedToInsertErr = "failed@test.com"
		successInsert     = "success@test.com"
	)

	var getGetUserByEmailMock = func(m *dm.DBInterface, t *testing.T, email string) {
		// GetUserByEmail(ctx context.Context, email string) (*invoice.User, error)
		switch email {
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
		payload *invoice.User
		wantErr bool
	}{
		{
			name: "Insert user fail",
			payload: &invoice.User{
				Email: failedToInsertErr,
			},
			wantErr: true,
		},
		{
			name: "Successfuly inserted user",
			payload: &invoice.User{
				Email: successInsert,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &invoicesrvc{
				logger: &logrus.Logger{},
				dbi:    getDatastoreMock(t, tt.payload.Email),
				hasher: argon2id.New(),
			}
			if err := s.CreateAccount(ctx, tt.payload); (err != nil) != tt.wantErr {
				t.Errorf("invoicesrvc.CreateAccount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

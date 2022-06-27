package controllers

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	dm "github.com/samverrall/invoice-api-service/datastore/mocks"
	"github.com/samverrall/invoice-api-service/gen/invoice"
	"github.com/samverrall/invoice-api-service/hasher"
	hm "github.com/samverrall/invoice-api-service/hasher/mocks"
	"github.com/samverrall/invoice-api-service/tokens/jwttoken"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/mock"
)

func Test_invoicesrvc_GrantToken(t *testing.T) {
	const (
		failedToInsertErr     = "failed@test.com"
		checkExistingUserFail = "existing@fail.com"
		successInsert         = "success@test.com"
		compareHashFail       = "compare@fail.com"
		compareHashFalse      = "false@compare.com"
	)

	var getGetUserByEmailMock = func(m *dm.DBInterface, t *testing.T, email string) {
		// GetUserByEmail(ctx context.Context, email string) (*invoice.User, error)
		switch email {
		case checkExistingUserFail:
			m.On("GetUserByEmail", mock.AnythingOfType("*context.emptyCtx"), email).
				Return(nil, fmt.Errorf("Failed to check if user exists, unit test error."))
		default:
			m.On("GetUserByEmail", mock.AnythingOfType("*context.emptyCtx"), email).
				Return(&invoice.User{
					Email:    email,
					Name:     "Test",
					Password: "hash",
				}, nil)
		}
	}

	var getDatastoreMock = func(t *testing.T, email string) *dm.DBInterface {
		var mock dm.DBInterface
		getGetUserByEmailMock(&mock, t, email)
		return &mock
	}

	// Compare(hashed string, pwd string) (bool, error)
	getCompareMocks := func(t *testing.T, m *hm.Hasher, email string) {
		switch email {
		case compareHashFail:
			m.On("Compare", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
				Return(false, fmt.Errorf("failed to generate hash unit test fail")).
				Once()
		case compareHashFalse:
			m.On("Compare", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
				Return(false, nil).
				Once()
		default:
			m.On("Compare", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
				Return(true, nil).
				Once()
		}
	}

	var getHasherMocks = func(t *testing.T, email string) hasher.Hasher {
		m := hm.Hasher{}
		getCompareMocks(t, &m, email)
		return &m
	}

	ctx := context.Background()
	tests := []struct {
		name    string
		want    *invoice.Token
		p       *invoice.GrantTokenPayload
		wantErr bool
	}{
		{
			name: "Check if user exists fail",
			p: &invoice.GrantTokenPayload{
				Email: checkExistingUserFail,
			},
			wantErr: true,
		},
		{
			name: "Fail to compare password hash",
			p: &invoice.GrantTokenPayload{
				Email: compareHashFail,
			},
			wantErr: true,
		},
		{
			name: "Password hash compare does not match returns an error",
			p: &invoice.GrantTokenPayload{
				Email: compareHashFalse,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &invoicesrvc{
				logger:  &logrus.Logger{},
				dbi:     getDatastoreMock(t, tt.p.Email),
				hasher:  getHasherMocks(t, tt.p.Email),
				tokener: jwttoken.New(15),
			}
			got, err := s.GrantToken(ctx, tt.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("invoicesrvc.GrantToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invoicesrvc.GrantToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

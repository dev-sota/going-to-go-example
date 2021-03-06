package user

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ispec-inc/going-to-go-example/pkg/apperror"
	"github.com/ispec-inc/going-to-go-example/pkg/domain/mock"
	"github.com/ispec-inc/going-to-go-example/pkg/domain/model"
	"github.com/ispec-inc/going-to-go-example/pkg/password"
	"github.com/stretchr/testify/assert"
)

func TestUserUsecase_Find(t *testing.T) {
	cases := map[string]struct {
		inp     FindInput
		out     FindOutput
		errCode apperror.Code
	}{
		"success": {
			inp: FindInput{
				ID: int64(1),
			},
			out: FindOutput{
				User: model.User{
					ID:       int64(1),
					Email:    "test@example.com",
					Password: "hashed_password",
					Name:     "test-user",
					Age:      int(25),
				},
			},
			errCode: apperror.CodeNoError,
		},
		"not found": {
			inp: FindInput{
				ID: int64(1),
			},
			out:     FindOutput{},
			errCode: apperror.CodeNotFound,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			um := mock.NewMockUser(ctrl)
			aerr := apperror.NewTestError(c.errCode)
			um.EXPECT().Find(c.inp.ID).Return(c.out.User, aerr)

			u := Usecase{user: um}
			out, aerr := u.Find(c.inp)

			assert.Equal(t, c.out, out)
			apperror.AssertError(t, c.errCode, aerr)
		})
	}
}

func TestUserUsecase_Add_Success(t *testing.T) {
	cases := map[string]struct {
		inp             AddInput
		out             AddOutput
		errCode         apperror.Code
		expectedErrCode apperror.Code
	}{
		"success": {
			inp: AddInput{
				User: model.User{
					Email:    "test@example.com",
					Password: "raw_password",
					Name:     "test-user",
					Age:      int(25),
				},
			},
			out: AddOutput{
				User: model.User{
					ID:       int64(1),
					Email:    "test@example.com",
					Password: "hashed_password",
					Name:     "test-user",
					Age:      int(25),
				},
			},
			errCode:         apperror.CodeNoError,
			expectedErrCode: apperror.CodeNotFound,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			um := mock.NewMockUser(ctrl)
			aerr := apperror.NewTestError(c.errCode)
			exerr := apperror.NewTestError(c.expectedErrCode)
			um.EXPECT().FindByEmail(c.inp.User.Email).Return(model.User{}, exerr)
			um.EXPECT().Create(gomock.Any()).SetArg(0, c.inp.User).Return(aerr)
			um.EXPECT().Find(c.inp.User.ID).Return(c.out.User, aerr)

			u := Usecase{user: um}
			out, aerr := u.Add(c.inp)

			assert.Equal(t, c.out, out)
			apperror.AssertError(t, c.errCode, aerr)
		})
	}
}

func TestUserUsecase_Login_Success(t *testing.T) {
	raw_password := "raw_password"
	hashed_password, err := password.Encrypt(raw_password)
	if err != nil {
		t.Fatal(err)
	}

	cases := map[string]struct {
		inp     LoginInput
		res     FindOutput
		errCode apperror.Code
	}{
		"success": {
			inp: LoginInput{
				Email:    "test@example.com",
				Password: raw_password,
			},
			res: FindOutput{
				User: model.User{
					ID:       int64(1),
					Email:    "test@example.com",
					Password: hashed_password,
					Name:     "test-user",
					Age:      int(25),
				},
			},
			errCode: apperror.CodeNoError,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			um := mock.NewMockUser(ctrl)
			aerr := apperror.NewTestError(c.errCode)
			um.EXPECT().FindByEmail(c.inp.Email).Return(c.res.User, aerr)

			u := Usecase{user: um}
			_, aerr = u.Login(c.inp)

			apperror.AssertError(t, c.errCode, aerr)
		})
	}
}

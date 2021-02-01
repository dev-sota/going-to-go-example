package user

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ispec-inc/going-to-go-example/pkg/apperror"
	"github.com/ispec-inc/going-to-go-example/pkg/domain/mock"
	"github.com/ispec-inc/going-to-go-example/pkg/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestUserUsecase_FindUser(t *testing.T) {
	cases := map[string]struct {
		inp     FindUserInput
		out     FindUserOutput
		errCode apperror.Code
	}{
		"success": {
			inp: FindUserInput{
				ID: int64(1),
			},
			out: FindUserOutput{
				User: model.User{
					ID:   int64(1),
					Name: "dev-sota",
					Age:  int(25),
				},
			},
			errCode: apperror.CodeNoError,
		},
		"not found": {
			inp: FindUserInput{
				ID: int64(1),
			},
			out:     FindUserOutput{},
			errCode: apperror.CodeNotFound,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			um := mock.NewMockUser(ctrl)

			aerr := apperror.NewTestError(c.errCode)
			um.EXPECT().Find(c.inp.ID).Return(c.out.User, aerr)

			u := Usecase{user: um}
			out, aerr := u.FindUser(c.inp)

			assert.Equal(t, c.out, out)
			apperror.AssertError(t, c.errCode, aerr)

			ctrl.Finish()
		})
	}
}

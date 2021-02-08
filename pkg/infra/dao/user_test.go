package dao

import (
	"testing"

	"github.com/ispec-inc/going-to-go-example/pkg/domain/model"
	"github.com/ispec-inc/going-to-go-example/pkg/infra/entity"
	"github.com/stretchr/testify/assert"
)

func TestUserDao_Find(t *testing.T) {
	t.Helper()
	d := NewUser(db)

	cases := []struct {
		name string
		id   int64
		want int64
		err  bool
	}{
		{
			name: "Found",
			id:   int64(1),
			want: int64(1),
			err:  false,
		},
		{
			name: "NotFound",
			id:   int64(2),
			want: int64(0),
			err:  true,
		},
	}
	for i := range cases {
		tc := cases[i]
		t.Run(tc.name, func(t *testing.T) {
			if err := prepareTestData("./testdata/user/find.sql"); err != nil {
				t.Error(err)
			}

			opt, aerr := d.Find(tc.id)

			assert.Exactly(t, tc.want, opt.ID)
			if tc.err {
				assert.Error(t, aerr)
			} else {
				assert.NoError(t, aerr)
			}
		})
	}
}

func TestUserDao_FindByEmail(t *testing.T) {
	t.Helper()
	d := NewUser(db)

	cases := []struct {
		name  string
		email string
		want  int64
		err   bool
	}{
		{
			name:  "Found",
			email: "test@example.com",
			want:  int64(1),
			err:   false,
		},
		{
			name:  "NotFound",
			email: "notfound@example.com",
			want:  int64(0),
			err:   true,
		},
	}
	for i := range cases {
		tc := cases[i]
		t.Run(tc.name, func(t *testing.T) {
			if err := prepareTestData("./testdata/user/find_by_email.sql"); err != nil {
				t.Error(err)
			}

			opt, aerr := d.FindByEmail(tc.email)

			assert.Exactly(t, tc.want, opt.ID)
			if tc.err {
				assert.Error(t, aerr)
			} else {
				assert.NoError(t, aerr)
			}
		})
	}
}

func TestUserDao_Create(t *testing.T) {
	t.Helper()
	d := NewUser(db)

	cases := []struct {
		name       string
		model      model.User
		createdCnt int
		err        bool
	}{
		{
			name: "Created",
			model: model.User{
				Email:    "test2@example.com",
				Password: "hashed_password",
				Name:     "test_user2",
				Age:      25,
			},
			createdCnt: 1,
			err:        false,
		},
		{
			name: "AlreadyExist",
			model: model.User{
				Email:    "test@example.com",
				Password: "hashed_password",
				Name:     "test_user",
				Age:      25,
			},
			createdCnt: 0,
			err:        true,
		},
	}
	for i := range cases {
		tc := cases[i]
		t.Run(tc.name, func(t *testing.T) {
			if err := prepareTestData("./testdata/user/create.sql"); err != nil {
				t.Error(err)
			}

			var before, after []entity.User

			d.db.Find(&before)

			aerr := d.Create(&tc.model)

			d.db.Find(&after)

			assert.Exactly(t, tc.createdCnt, len(after)-len(before))
			if tc.err {
				assert.Error(t, aerr)
			} else {
				assert.NoError(t, aerr)
			}
		})
	}
}

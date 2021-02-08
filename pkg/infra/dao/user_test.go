package dao

import (
	"testing"

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

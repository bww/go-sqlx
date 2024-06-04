package sqlx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testArgs struct {
	Query  string
	N      int
	Expect string
}

func TestArgs(t *testing.T) {
	tests := []testArgs{
		{
			Query:  "created_at > $N",
			N:      0,
			Expect: "created_at > $1",
		},
		{
			Query:  "created_at > $1",
			N:      0,
			Expect: "created_at > $1",
		},
		{
			Query:  "created_at > $_$@ $$$ $?",
			N:      0,
			Expect: "created_at > $_$@ $$$ $?",
		},
		{
			Query:  "created_at > $1 $N $N $N",
			N:      1,
			Expect: "created_at > $1 $2 $3 $4",
		},
	}
	for _, e := range tests {
		r := Args(e.N, e.Query)
		assert.Equal(t, r, e.Expect)
	}
}

func TestWhere(t *testing.T) {
	tests := []testArgs{
		{
			Query:  "created_at > $N",
			N:      0,
			Expect: " WHERE created_at > $1",
		},
		{
			Query:  "created_at > $N",
			N:      1,
			Expect: " AND created_at > $2",
		},
		{
			Query:  "created_at > $1",
			N:      0,
			Expect: " WHERE created_at > $1",
		},
		{
			Query:  "created_at > $1",
			N:      1,
			Expect: " AND created_at > $1",
		},
		{
			Query:  "created_at > $_$@ $$$ $?",
			N:      0,
			Expect: " WHERE created_at > $_$@ $$$ $?",
		},
		{
			Query:  "created_at > $1 $N $N $N",
			N:      1,
			Expect: " AND created_at > $1 $2 $3 $4",
		},
	}
	for _, e := range tests {
		r := Where(e.N, e.Query)
		assert.Equal(t, r, e.Expect)
	}
}

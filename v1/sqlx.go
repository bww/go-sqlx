package sqlx

import (
	"fmt"
	"strconv"
	"strings"
)

func List(s, n int) string {
	b := strings.Builder{}
	for i := 0; i < n; i++ {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(fmt.Sprintf("$%d", s+i+1))
	}
	return b.String()
}

func Wheref(n int, f string, a ...interface{}) string {
	return Where(n, fmt.Sprintf(f, a...))
}

func Where(n int, f string) string {
	var q string
	if n > 0 {
		q += " AND "
	} else {
		q += " WHERE "
	}
	return q + Args(n, f)
}

func Argsf(n int, f string, a ...interface{}) string {
	return Args(n, fmt.Sprintf(f, a...))
}

func Args(n int, s string) string {
	var r strings.Builder
	var e int
	for _, c := range s {
		if c == '$' {
			r.WriteRune(c)
			e++
		} else if e%2 == 1 && c == 'N' {
			n++
			r.WriteString(strconv.FormatInt(int64(n), 10))
			e = 0
		} else {
			r.WriteRune(c)
			e = 0
		}
	}
	return r.String()
}

// Generalize converts any slice into a slice of interface{}. This
// is useful when copying a slice of values of an arbitrary type
// into query arguments.
func Generalize[T any](s []T) []interface{} {
	g := make([]any, len(s))
	for i, e := range s {
		g[i] = e
	}
	return g
}

// Append arguments of any type to a slice of type []interface{}. This
// functions as a specialized version of the built-in append() operation
// which is suitable for appending to a slice of empty interfaces.
func Append[T any](a []interface{}, s ...T) []interface{} {
	for _, e := range s {
		a = append(a, e)
	}
	return a
}

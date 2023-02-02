package pgdialect

import (
	"testing"

	"github.com/uptrace/bun/schema"
)

type tag struct {
	ID    int32
	Label string
}

func ptr[T any](v T) *T {
	return &v
}

func TestArrayAppend(t *testing.T) {
	tcases := []struct {
		input interface{}
		out   string
	}{
		{
			input: []byte{1, 2},
			out:   `'{1,2}'`,
		},
		{
			input: []*byte{ptr(byte(1)), ptr(byte(2))},
			out:   `'{1,2}'`,
		},
		{
			input: []int{1, 2},
			out:   `'{1,2}'`,
		},
		{
			input: []*int{ptr(1), ptr(2)},
			out:   `'{1,2}'`,
		},
		{
			input: []string{"foo", "bar"},
			out:   `'{"foo","bar"}'`,
		},
		{
			input: []*string{ptr("foo"), ptr("bar")},
			out:   `'{"foo","bar"}'`,
		},
		{
			input: []tag{{1, "foo"}, {2, "bar"}},
			out:   `'[{"ID":1,"Label":"foo"},{"ID":2,"Label":"bar"}]'`,
		},
		{
			input: &[]tag{{1, "foo"}, {2, "bar"}},
			out:   `'[{"ID":1,"Label":"foo"},{"ID":2,"Label":"bar"}]'`,
		},
		{
			input: &[]*tag{{1, "foo"}, {2, "bar"}},
			out:   `'[{"ID":1,"Label":"foo"},{"ID":2,"Label":"bar"}]'`,
		},
		{
			input: ptr(&[]*tag{{1, "foo"}, {2, "bar"}}),
			out:   `'[{"ID":1,"Label":"foo"},{"ID":2,"Label":"bar"}]'`,
		},
		{
			input: []**tag{ptr(ptr(tag{1, "foo"})), ptr(ptr(tag{2, "bar"}))},
			out:   `'[{"ID":1,"Label":"foo"},{"ID":2,"Label":"bar"}]'`,
		},
		{
			input: &[]**tag{ptr(ptr(tag{1, "foo"})), ptr(ptr(tag{2, "bar"}))},
			out:   `'[{"ID":1,"Label":"foo"},{"ID":2,"Label":"bar"}]'`,
		},
		{
			input: ptr(&[]**tag{ptr(ptr(tag{1, "foo"})), ptr(ptr(tag{2, "bar"}))}),
			out:   `'[{"ID":1,"Label":"foo"},{"ID":2,"Label":"bar"}]'`,
		},
		{
			input: ptr(ptr(&[]**tag{ptr(ptr(tag{1, "foo"})), ptr(ptr(tag{2, "bar"}))})),
			out:   `'[{"ID":1,"Label":"foo"},{"ID":2,"Label":"bar"}]'`,
		},
	}

	for _, tcase := range tcases {
		out, err := Array(tcase.input).AppendQuery(schema.NewFormatter(New()), []byte{})
		if err != nil {
			t.Fatal(err)
		}

		if string(out) != tcase.out {
			t.Errorf("expected output to be %s, was %s", tcase.out, string(out))
		}
	}
}

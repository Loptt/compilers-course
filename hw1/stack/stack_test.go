package stack

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestPush(t *testing.T) {
	tests := []struct {
		in   func() *Stack
		add  []interface{}
		want func() *Stack
	}{
		{
			in: func() *Stack {
				var s Stack
				return &s
			},
			add: []interface{}{1},
			want: func() *Stack {
				var s Stack
				s.head = &Node{1, nil}
				return &s
			},
		},
		{
			in: func() *Stack {
				var s Stack
				return &s
			},
			add: []interface{}{3,1,2},
			want: func() *Stack {
				var s Stack
				s.head = &Node{2, &Node{1, &Node{3, nil}}}
				return &s
			},
		},
		{
			in: func() *Stack {
				var s Stack
				s.head = &Node{"hey", nil}
				return &s
			},
			add: []interface{}{"hola", "que tal"},
			want: func() *Stack {
				var s Stack
				s.head = &Node{"que tal", &Node{"hola", &Node{"hey", nil}}}
				return &s
			},
		},
	}

	for _, test := range tests {
		s := *test.in()

		for _, add := range test.add {
			s.Push(add)
		}

		got := s
		want := *test.want()
		if !cmp.Equal(s, want) {
			t.Errorf("Push(%v) = %v, expected %v", test.add, got, want)
		}
	}
}

func TestPop(t *testing.T) {
	tests := []struct {
		in   func() *Stack
		want func() *Stack
	}{
		{
			in: func() *Stack {
				var s Stack
				s.Push(3)
				s.Push(2)
				s.Push(1)
				return &s
			},
			want: func() *Stack {
				var s Stack
				s.Push(3)
				s.Push(2)
				return &s
			},
		},
		{
			in: func() *Stack {
				var s Stack
				s.Push("hola")
				s.Push("que")
				return &s
			},
			want: func() *Stack {
				var s Stack
				s.Push("hola")
				return &s
			},
		},
		{
			in: func() *Stack {
				var s Stack
				s.Push(23.55)
				return &s
			},
			want: func() *Stack {
				var s Stack
				return &s
			},
		},
	}

	for _, test := range tests {
		s := *test.in()
		s.Pop()
		got := s
		want := *test.want()
		if !cmp.Equal(s, want) {
			t.Errorf("Pop() = %v, expected %v", got, want)
		}
	}
}

func TestTop(t *testing.T) {
	tests := []struct {
		in   func() *Stack
		want interface{}
	}{
		{
			in: func() *Stack {
				var s Stack
				s.Push(2)
				s.Push(11)
				return &s
			},
			want: 11,
		},
		{
			in: func() *Stack {
				var s Stack
				s.Push("ey")
				s.Push("ho")
				s.Push("yo")
				return &s
			},
			want: "yo",
		},
	}

	for _, test := range tests {
		in := *test.in()
		got := in.Top()

		if got != test.want {
			t.Errorf("Top() = %v, expected %v", got, test.want)
		}
	}
}
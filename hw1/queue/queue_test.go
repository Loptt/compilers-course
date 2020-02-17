package queue

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestPush(t *testing.T) {
	tests := []struct {
		in   func() *Queue
		add  []interface{}
		want func() *Queue
	}{
		{
			in: func() *Queue {
				n := &Node{1, nil}
				return &Queue{head: n, tail: n}
			},
			add: []interface{}{2},
			want: func() *Queue {
				t := &Node{2, nil}
				n := &Node{1, t}
				return &Queue{head: n, tail: t}
			},
		},
		{
			in: func() *Queue {
				return &Queue{head: nil, tail: nil}
			},
			add: []interface{}{3,4,5},
			want: func() *Queue {
				t := &Node{5, nil}
				n := &Node{3, &Node{4, t}}
				return &Queue{head: n, tail: t}
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
		in   func() *Queue
		want func() *Queue
	}{
		{
			in: func() *Queue {
				n := &Node{1, nil}
				return &Queue{head: n, tail: n}
			},
			want: func() *Queue {
				return &Queue{head: nil, tail: nil}
			},
		},
		{
			in: func() *Queue {
				t := &Node{5, nil}
				n := &Node{3, &Node{4, t}}
				return &Queue{head: n, tail: t}
			},
			want: func() *Queue {
				t := &Node{5, nil}
				n := &Node{4, t}
				return &Queue{head: n, tail: t}
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

func TestFront(t *testing.T) {
	tests := []struct {
		in   func() *Queue
		want interface{}
	}{
		{
			in: func() *Queue {
				var q Queue
				q.Push(2)
				q.Push(11)
				return &q
			},
			want: 2,
		},
		{
			in: func() *Queue {
				var q Queue
				q.Push("ey")
				q.Push("ho")
				q.Push("yo")
				return &q
			},
			want: "ey",
		},
	}

	for _, test := range tests {
		in := *test.in()
		got := in.Front()

		if got != test.want {
			t.Errorf("Top() = %v, expected %v", got, test.want)
		}
	}
}
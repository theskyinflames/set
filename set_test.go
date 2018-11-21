package set

import (
	"reflect"
	"sync"
	"testing"
)

var (
	key1 func() string = func() string {
		return "key1"
	}

	key2 func() string = func() string {
		return "key2"
	}
)

func load() (s *Set) {

	s = &Set{
		RWMutex: sync.RWMutex{},
		set:     sync.Map{},
	}

	s.Add(key1, 1)
	s.Add(key2, 2)
	return
}

func TestSet_Add(t *testing.T) {

	s := &Set{
		RWMutex: sync.RWMutex{},
		set:     sync.Map{},
	}

	type args struct {
		key  Key
		item interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "given an empty set, when a new element is added then it results ok",
			args: args{
				key:  key1,
				item: 1,
			},
			wantErr: false,
		},
		{
			name: "given a set, when a already existing element is added then it results error",
			args: args{
				key:  key1,
				item: 1,
			},
			wantErr: true,
		},
		{
			name: "given a set, when a new element is added then it results error",
			args: args{
				key:  key2,
				item: 2,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := s.Add(tt.args.key, tt.args.item); (err != nil) != tt.wantErr {
				t.Errorf("Set.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSet_Pop(t *testing.T) {

	s := load()

	type args struct {
		key Key
	}
	tests := []struct {
		name     string
		args     args
		wantItem interface{}
		wantErr  bool
	}{
		{
			name: "given a set, when it tried to extract an existing element ",
			args: args{
				key: key1,
			},
			wantItem: 1,
			wantErr:  false,
		},
		{
			name: "given a set, when it tried to extract a non existing element ",
			args: args{
				key: func() string {
					return "key4"
				},
			},
			wantItem: 4,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotItem, err := s.Pop(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Set.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotItem != nil {
				if !reflect.DeepEqual(gotItem, tt.wantItem) {
					t.Errorf("Set.Pop() = %v, want %v", gotItem, tt.wantItem)
				}
			}
		})
	}
}

func TestSet_Peek(t *testing.T) {

	s := load()

	type args struct {
		key Key
	}
	tests := []struct {
		name     string
		args     args
		wantItem interface{}
		wantErr  bool
	}{
		{
			name: "given a set, when it tried to peek an existing element ",
			args: args{
				key: key1,
			},
			wantItem: 1,
			wantErr:  false,
		},
		{
			name: "given a set, when it tried to peek a non existing element ",
			args: args{
				key: func() string {
					return "key4"
				},
			},
			wantItem: 1,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotItem, err := s.Peek(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Set.Peek() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotItem != nil {
				if !reflect.DeepEqual(gotItem, tt.wantItem) {
					t.Errorf("Set.Peek() = %v, want %v", gotItem, tt.wantItem)
				}
			}
		})
	}
}

func TestSet_CountOfNonEmptySet(t *testing.T) {

	s := load()

	tests := []struct {
		name      string
		wantCount int
	}{
		{
			name:      "given a non empty set, when it's count, then it result in non zero value",
			wantCount: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := s.Count(); gotCount != tt.wantCount {
				t.Errorf("Set.Count() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func TestSet_CountOfEmptySet(t *testing.T) {

	s := &Set{
		RWMutex: sync.RWMutex{},
		set:     sync.Map{},
	}

	tests := []struct {
		name      string
		wantCount int
	}{
		{
			name:      "given a empty set, when it's count, then it result in zero value",
			wantCount: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := s.Count(); gotCount != tt.wantCount {
				t.Errorf("Set.Count() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func TestSet_Clear(t *testing.T) {

	s := load()

	tests := []struct {
		name string
	}{
		{
			name: "given a non empty set, when once it's cleared, then its count function returns zero value",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s.Clear()
			count := s.Count()
			if count != 0 {
				t.Errorf("Set.Clear has failed")
			}
		})
	}
}

func TestSet_Iterator(t *testing.T) {
	s := load()
	tests := []struct {
		name string
	}{
		{
			name: "given a non empty set, when retrieving an iterator, all of its items are retrieved",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := 0
			iterator := s.Iterator()
			for {
				item, end := iterator()
				t.Logf("retrieved item : %v, ok: %t", item, end)
				if end {
					if c != 1 {
						t.Errorf("iterated wrong number of items, got %d, wanted 2", c)
					}
					break
				}
				c++
			}
		})
	}
}

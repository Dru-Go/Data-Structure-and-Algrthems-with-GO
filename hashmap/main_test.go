package main

import (
	"reflect"
	"testing"
)

func TestHashTable_Search(t *testing.T) {
	ht := NewHashTable(10)
	ht.Insert("dog", 1)
	ht.Insert("cat", 2)
	type fields struct {
		size  int
		table []*Node
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		value  int
		found  bool
	}{
		{
			name: "Test search with non existing query",
			fields: fields{
				size:  10,
				table: ht.table,
			},
			args: args{
				key: "bare",
			},
			value: 0,
			found: false,
		},
		{
			name: "Test search with an existing query",
			fields: fields{
				size:  10,
				table: ht.table,
			},
			args: args{
				key: "dog",
			},
			value: 1,
			found: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HashTable{
				size:  tt.fields.size,
				table: tt.fields.table,
			}
			got, got1 := h.Search(tt.args.key)
			if got != tt.value {
				t.Errorf("HashTable.Search() got = %v, want %v", got, tt.value)
			}
			if got1 != tt.found {
				t.Errorf("HashTable.Search() got1 = %v, want %v", got1, tt.found)
			}
		})
	}
}

func TestHashTable_SearchRegex(t *testing.T) {
	ht := NewHashTable(10)
	ht.Insert("dog", 1)
	ht.Insert("cat", 2)
	type fields struct {
		size  int
		table []*Node
	}
	type args struct {
		pattern string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		value  []int
		found  bool
	}{
		{
			name: "Test search with an existing query",
			fields: fields{
				size:  10,
				table: ht.table,
			},
			args: args{
				pattern: "dog",
			},
			value: []int{1},
			found: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HashTable{
				size:  tt.fields.size,
				table: tt.fields.table,
			}
			got, got1 := h.SearchRegex(tt.args.pattern)
			if !reflect.DeepEqual(got, tt.value) {
				t.Errorf("HashTable.SearchRegex() got = %v, want %v", got, tt.value)
			}
			if got1 != tt.found {
				t.Errorf("HashTable.SearchRegex() got1 = %v, want %v", got1, tt.found)
			}
		})
	}
}

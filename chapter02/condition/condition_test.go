package main

import (
	"reflect"
	"testing"
)

func TestNewQueue(t *testing.T) {
	capacity := 10

	type args struct {
		capacity int
	}
	tests := []struct {
		name string
		args args
		want *Queue
	}{
		// TODO: Add test cases.
		{
			name: "Test 1",
			args: args{capacity: capacity},
			want: &Queue{elements: make([]int, capacity),
				front: 0,  // Read from elements[front]
				rear:  -1, // Write to elements[rear]
				len:   0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewQueue(tt.args.capacity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

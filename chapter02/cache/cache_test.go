package main

import (
	"reflect"
	"testing"
)

func Test_retrieveData(t *testing.T) {

	type args struct {
		ID string
	}
	tests := []struct {
		name  string
		args  args
		want  *Data
		want1 bool
	}{
		{
			name:  "Test a",
			args:  args{ID: "a"},
			want:  &Data{ID: "a", Payload: "payload"},
			want1: true,
		},
		{
			name:  "Test b",
			args:  args{ID: "b"},
			want:  &Data{ID: "b", Payload: "payload"},
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := retrieveData(tt.args.ID)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("retrieveData() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("retrieveData() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

package main

import (
	"reflect"
	"testing"
)

func TestFizzBuzzCollection(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want map[int]string
	}{
		{
			"If given size 5, counts to 5 returning fizz/buzz value for each",
			args{5},
			map[int]string{1: "1", 2: "2", 3: "fizz", 4: "4", 5: "buzz"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FizzBuzzCollection(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FizzBuzzCollection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFizzBuzzItem(t *testing.T) {
	type args struct {
		item int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Given 4, returns 4", args{4}, "4"},
		{"Given 99, returns fizz", args{99}, "fizz"},
		{"Given 245, returns buzz", args{245}, "buzz"},
		{"Given 15, returns fizz-buzz", args{15}, "fizz-buzz"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FizzBuzzItem(tt.args.item); got != tt.want {
				t.Errorf("FizzBuzzItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

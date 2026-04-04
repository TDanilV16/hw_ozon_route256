package main

import (
	"slices"
	"testing"
)

func TestDistributeApplesInCrates(t *testing.T) {
	type args struct {
		degree int
		base   int
	}

	tests := []struct {
		name           string
		args           args
		want           [][]int
		failsWithError bool
	}{
		{
			name:           "Positive base, positive degree, total apples 4",
			args:           args{base: 2, degree: 2},
			want:           [][]int{{4, 3, 2, 1}},
			failsWithError: false,
		},
		{
			name:           "Positive base, negative degree, total apples 16",
			args:           args{base: 2, degree: 4},
			want:           [][]int{{16, 15, 14, 13, 12, 11, 10, 9, 8, 7}, {6, 5, 4, 3, 2, 1}},
			failsWithError: false,
		},
		{
			name:           "Negative base, positive degree, fails with error",
			args:           args{base: -1, degree: 2},
			want:           nil,
			failsWithError: true,
		},
		{
			name:           "Positive base, negative degree, fails with error",
			args:           args{base: -2, degree: 2},
			want:           nil,
			failsWithError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DistributeApplesInCrates(tt.args.degree, tt.args.base)
			if (err != nil) != tt.failsWithError {
				t.Errorf("DistributeApplesInCrates() error = %v, wantErr %v", err, tt.failsWithError)
				return
			}
			if !slices.EqualFunc(got, tt.want, slices.Equal) {
				t.Errorf("DistributeApplesInCrates() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRemove(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	got := remove(input, 3)
	if !reflect.DeepEqual(got, []int{1, 2, 4, 5, 6}) {
		t.Fatal()
	}
}

func TestParse(t *testing.T) {
	wantBegin := newSuduko()
	wantBegin.board[0][0] = 1
	wantBegin.board[3][3] = 2
	wantBegin.board[6][6] = 3

	wantGoal := newSuduko()
	wantGoal.board[1][1] = 1
	wantGoal.board[4][4] = 2
	wantGoal.board[7][7] = 3

	gotBegin, gotGoal, err := parse("./testdata/testRead")
	if err != nil {
		t.Fatal("unexpected error")
	}

	if !reflect.DeepEqual(gotBegin, wantBegin) {
		gotBegin.dump()
		wantBegin.dump()
		t.Fatal("begin sudukos do not match")
	}
	if !reflect.DeepEqual(gotGoal, wantGoal) {
		gotGoal.dump()
		wantGoal.dump()
		t.Fatal("goal sudukos do not match")

	}
}
func TestCheckSubGrid(t *testing.T) {

	tests := []struct {
		mx, my, m []int
		want      bool
	}{
		{
			mx:   []int{0, 1, 2},
			my:   []int{0, 1, 2},
			m:    []int{1, 2, 3},
			want: true,
		},
		{
			mx:   []int{0, 1, 2},
			my:   []int{0, 1, 2},
			m:    []int{1, 1, 3},
			want: false,
		},
		{
			mx:   []int{8, 8, 8},
			my:   []int{0, 1, 2},
			m:    []int{1, 1, 3},
			want: false,
		},
		{
			mx:   []int{8, 8, 8},
			my:   []int{7, 6, 8},
			m:    []int{8, 7, 6},
			want: true,
		},
		{
			mx:   []int{8, 8, 8},
			my:   []int{7, 6, 8},
			m:    []int{6, 7, 6},
			want: false,
		},
	}

	for _, tc := range tests {
		s := newSuduko()
		for i := range tc.m {
			s.board[tc.my[i]][tc.mx[i]] = tc.m[i]
		}

		if tc.want != s.checkSubGrid(tc.my[0], tc.mx[0]) {
			s.dump()
			t.Fatal()
		}
	}
}

func TestUnsolvable(t *testing.T) {
	tests := []struct {
		inputPath string
	}{
		{inputPath: "./testdata/test_03"},
	}

	for _, tc := range tests {
		got, want, err := parse(tc.inputPath)
		if err != nil {
			t.Fatal("unexpected error")
		}

		if got.solve(0) {
			t.Fatal("got a solution to unsolvable puzzle")
		}

		if !reflect.DeepEqual(got.board, want.board) {
			fmt.Println("got:")
			got.dump()
			fmt.Println("want:")
			want.dump()
			t.Fatal("boards not same")
		}
	}
}
func TestSolvable(t *testing.T) {
	tests := []struct {
		inputPath string
	}{
		{inputPath: "./testdata/test_01"},
		{inputPath: "./testdata/test_02"},
	}

	for _, tc := range tests {
		got, want, err := parse(tc.inputPath)
		if err != nil {
			t.Fatal("unexpected error")
		}

		if !got.solve(0) {
			got.dump()
			t.Fatal("unable to solve")
		}

		if !reflect.DeepEqual(got.board, want.board) {
			fmt.Println("got:")
			got.dump()
			fmt.Println("want:")
			want.dump()
			t.Fatal("boards not same")
		}
	}
}

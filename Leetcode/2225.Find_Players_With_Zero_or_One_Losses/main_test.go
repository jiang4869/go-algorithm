package main

import (
	"fmt"
	"testing"
)

func Test_Case1(t *testing.T) {
	matches := [][]int{
		{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9},
	}

	winners := findWinners(matches)
	fmt.Println(winners)

}

package jump

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	testCases := []struct {
		key     uint64
		buckets int32
		want    int32
	}{
		{0, 0, 0},
		{1, 1, 0},
		{42, 100, 43},
		{1337, 2000, 899},
		{128, 2048, 1995},
		{0x27BB2EE687B0B0FD, 5024, 592},
		{0, -1, 0},
		{2048, -2048, 0},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf(
			"input_key_%d_buckets_%d", tc.key, tc.buckets),
			func(t *testing.T) {
				got := Hash(tc.key, tc.buckets)
				if got != tc.want {
					t.Errorf("fail got %v want %v", got, tc.want)
				}
			})
	}
}

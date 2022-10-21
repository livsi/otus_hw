package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	tests := []struct {
		name            string
		source          string
		dest            string
		expectedContent string
		offset          int64
		limit           int64
	}{
		{
			name:            "full copy",
			source:          "testdata/input.txt",
			dest:            "out.txt",
			expectedContent: "testdata/input.txt",
			offset:          0,
			limit:           0,
		},
		{
			name:            "out_offset0_limit10",
			source:          "testdata/input.txt",
			dest:            "out.txt",
			expectedContent: "testdata/out_offset0_limit10.txt",
			offset:          0,
			limit:           10,
		},
		{
			name:            "out_offset0_limit1000",
			source:          "testdata/input.txt",
			dest:            "out.txt",
			expectedContent: "testdata/out_offset0_limit1000.txt",
			offset:          0,
			limit:           1000,
		},
		{
			name:            "out_offset0_limit10000",
			source:          "testdata/input.txt",
			dest:            "out.txt",
			expectedContent: "testdata/out_offset0_limit10000.txt",
			offset:          0,
			limit:           10000,
		},
		{
			name:            "out_offset100_limit1000",
			source:          "testdata/input.txt",
			dest:            "out.txt",
			expectedContent: "testdata/out_offset100_limit1000.txt",
			offset:          100,
			limit:           1000,
		},
		{
			name:            "out_offset6000_limit1000",
			source:          "testdata/input.txt",
			dest:            "out.txt",
			expectedContent: "testdata/out_offset6000_limit1000.txt",
			offset:          6000,
			limit:           1000,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := Copy(tc.source, tc.dest, tc.offset, tc.limit)
			defer func() {
				err := os.Remove(tc.dest)
				if err != nil {
					fmt.Printf("error on remove file: %v", err)
				}
			}()
			require.Nil(t, err)

			expectedContent, _ := os.ReadFile(tc.expectedContent)
			dest, _ := os.ReadFile(tc.dest)

			require.Equal(t, expectedContent, dest)
		})
	}
}

func TestErrors(t *testing.T) {
	tests := []struct {
		name    string
		source  string
		dest    string
		offset  int64
		limit   int64
		errorIs error
	}{
		{
			name:    "source and dest empty",
			source:  "",
			dest:    "",
			offset:  0,
			limit:   0,
			errorIs: ErrInvalidParams,
		},
		{
			name:    "source not exist, dest empty",
			source:  "testdata/dummySource",
			dest:    "",
			offset:  0,
			limit:   0,
			errorIs: ErrInvalidParams,
		},
		{
			name:   "source not exist, dest exist",
			source: "testdata/dummySource",
			dest:   "dummySource.txt",
			offset: 0,
			limit:  0,
		},
		{
			name:    "source exist, dest empty",
			source:  "testdata/dummySource",
			dest:    "",
			offset:  0,
			limit:   0,
			errorIs: ErrInvalidParams,
		},
		{
			name:    "source exist, dest not empty, limit <0",
			source:  "testdata/input.txt",
			dest:    "out.txt",
			offset:  0,
			limit:   -10,
			errorIs: ErrInvalidParams,
		},
		{
			name:    "source exist, dest not empty, offset <0",
			source:  "testdata/input.txt",
			dest:    "out.txt",
			offset:  -10,
			limit:   10,
			errorIs: ErrInvalidParams,
		},
		{
			name:    "source exist, dest not empty, offset exceeds file size",
			source:  "testdata/out_offset0_limit0.txt",
			dest:    "out.txt",
			offset:  10000000,
			limit:   0,
			errorIs: ErrOffsetExceedsFileSize,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := Copy(tc.source, tc.dest, tc.offset, tc.limit)
			require.Error(t, err)
			if tc.errorIs != nil {
				require.ErrorIs(t, err, tc.errorIs)
			}
		})
	}
}

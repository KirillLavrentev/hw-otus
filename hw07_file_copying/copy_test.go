package main

import (
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopyErrors(t *testing.T) {
	t.Run("offset more than filesize", func(t *testing.T) {
		err := Copy("testdata/input.txt", "out.txt", 99999999999999, 0)
		require.EqualError(t, err, ErrOffsetExceedsFileSize.Error())
	})
	t.Run("file does not exist", func(t *testing.T) {
		err := Copy("no_file.no", "out.txt", 0, 0)
		require.EqualError(t, err, "stat no_file.no: no such file or directory")
	})

	t.Run("check whether file is a directory", func(t *testing.T) {
		err := Copy("./testdata", "out.txt", 0, 0)
		require.EqualError(t, err, ErrUnsupportedFile.Error())
	})
	t.Run("check whether file is a regular file", func(t *testing.T) {
		err := Copy("/dev/null", "out.txt", 0, 0)
		require.EqualError(t, err, ErrUnsupportedFile.Error())
	})
}

func TestCopy(t *testing.T) {
	from := "testdata/input.txt"
	testCases := []struct {
		name   string
		toPath string
		offset int64
		limit  int64
		err    error
	}{
		{
			name:   "test offset=0 and limit=0",
			offset: 0,
			limit:  0,
			toPath: "out_offset0_limit0.txt",
		},
		{
			name:   "test offset=0 and limit=10",
			offset: 0,
			limit:  10,
			toPath: "out_offset0_limit10.txt",
		},
		{
			name:   "test offset=0 and limit=1000",
			offset: 0,
			limit:  1000,
			toPath: "out_offset0_limit1000.txt",
		},
		{
			name:   "test offset=0 and limit=10000",
			offset: 0,
			limit:  10000,
			toPath: "out_offset0_limit10000.txt",
		},
		{
			name:   "test offset=100 and limit=1000",
			offset: 100,
			limit:  1000,
			toPath: "out_offset100_limit1000.txt",
		},
		{
			name:   "test offset=6000 and limit=1000",
			offset: 6000,
			limit:  1000,
			toPath: "out_offset6000_limit1000.txt",
		},
		{
			name:   "test limit more than filesize",
			offset: 0,
			limit:  999999999999,
			toPath: "out_offset0_limit0.txt",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tmpDir, err := os.MkdirTemp("", "tmpTests")
			if err != nil {
				t.Fatal("Error in creation temp dir: ", err)
			}
			outFile := "out_offset" + strconv.Itoa(int(tc.offset)) + "_limit" + strconv.Itoa(int(tc.limit)) + ".txt"
			outPath := filepath.Join(tmpDir, outFile)
			defer os.RemoveAll(tmpDir)

			err = Copy(from, outPath, tc.offset, tc.limit)
			require.Nil(t, err)

			expected, err := os.ReadFile("testdata/" + tc.toPath)
			require.NoError(t, err)

			actual, err := os.ReadFile(outPath)
			require.NoError(t, err)

			require.Equal(t, expected, actual)
		})
	}
}

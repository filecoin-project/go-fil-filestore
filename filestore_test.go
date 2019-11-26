package filestore_test

import (
	"crypto/rand"
	"io/ioutil"
	"os"
	"testing"

	filestore "github.com/filecoin-project/go-fil-filestore"
	"github.com/stretchr/testify/require"
)

func randBytes(n int) []byte {
	arr := make([]byte, n)
	rand.Read(arr)
	return arr
}

func TestLocalFileStoreOpen(t *testing.T) {
	require.NoError(t, os.MkdirAll("_test/a/b/c/d", 0755))
	openOrFail := func(t *testing.T, filename string) *os.File {
		y, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		require.NoError(t, err)
		return y
	}
	f := openOrFail(t, "_test/a/b/c/d/existing.go")
	data := randBytes(100)
	_, err := f.Write(data)
	require.NoError(t, err)
	require.NoError(t, f.Close())

	fs := filestore.NewLocalFileStore("_test")
	file, err := fs.Open("a/b/c/d/existing.go")
	require.NoError(t, err)
	returnedData, err := ioutil.ReadAll(file)
	require.NoError(t, err)
	require.Equal(t, data, returnedData)

	/// ... add more tests for create/store/delete

}

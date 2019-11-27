package filestore

import "os"

type fd struct {
    *os.File
    filename string
}

func newFile(filename Path) (File, error) {
    var err error
    result := fd{filename: string(filename) }
    result.File, err = os.OpenFile(result.filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
    if err != nil {
        return nil, err
    }
    return &result, nil
}

func (f fd) Path() Path {
    return Path(f.filename)
}

func (f fd) Size() int64 {
    info, err := os.Stat(f.filename)
    if err != nil {
        return -1
    }
    return info.Size()
}

func (f fd) Close() error {
    return f.File.Close()
}

func (f fd) Read(p []byte) (n int, err error) {
    return f.File.Read(p)
}

func (f fd) Write(p []byte) (n int, err error) {
    return f.File.Write(p)
}

func (f fd) Seek(offset int64, whence int) (int64, error) {
    return f.File.Seek(offset, whence)
}



package storage

import "io"

// IStorage interface define methods for storage (bucket)
type IStorage interface {
	Download(key, path string) error
	Upload(key string, body io.ReadSeeker) error
}

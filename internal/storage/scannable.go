package storage

type Scannable interface {
	Scan(dest ...any) error
}

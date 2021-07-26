package contracts

type InMemoryCache interface {
	Set(key string, value []byte) error
	Get(key string) ([]byte, error)
	Exists(key string) (bool, error)
	Delete(key string) (int64, error)
}

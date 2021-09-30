package hashmap

type HashMap interface {
	Get(key string) (string, bool)
	Add(key, value string)
	Remove(key string) (string, bool)
}

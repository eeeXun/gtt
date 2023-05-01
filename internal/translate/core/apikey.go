package core

type APIKey struct {
	key string
}

func (k *APIKey) SetAPIKey(key string) {
	k.key = key
}

func (k *APIKey) GetAPIKey() string {
	return k.key
}

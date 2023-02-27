package cache

// CacheTable represent a table like table of MySQL
type CacheTable struct {
	items map[string]any
}

func (ct *CacheTable) Get(key string) (res any, exist bool) {
	res, exist = ct.items[key]
	return
}

// TODO-@geyan 这里需要考虑并发吗？ 我觉得需要
func (ct *CacheTable) Put(key string, value any) {
	ct.items[key] = value
}

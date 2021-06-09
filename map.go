package mapreduce

func Map(items []interface{}, f func(item interface{}) error) {
	for _, item := range items {
		f(item)
	}
}

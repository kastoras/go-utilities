package utilities

import "sync"

func GetSingleton[T any](instance *T) *T {
	var once sync.Once
	var instanceValue *T

	once.Do(func() {
		instanceValue = instance
	})

	return instanceValue
}

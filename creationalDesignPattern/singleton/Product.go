package singleton

import (
	"sync"
)

// use function init - not recommended
// use sync.Once
// use mutex locks

type Product struct {
	Name string
}

var instance *Product
var once sync.Once
var mu sync.Mutex

/* problem*/
// func GetInstance() *Product {
// 	if instance == nil {
// 		time.Sleep(time.Second)
// 		instance = new(Product)
// 		instance.Name = "new Product"
// 	}
// 	return instance
// }

/*
 sync.Once
*/
// func GetInstance() *Product {
// 	once.Do(func() {
// 		instance = new(Product)
// 	})
// 	return instance
// }

/*
 mutex locks
*/
func GetInstance() *Product {
	mu.Lock()
	defer mu.Unlock()
	if instance == nil {
		instance = new(Product)
	}
	return instance
}

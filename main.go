package main

import (
	"fmt"
	"sync"
)

// Customer havuzda depolanacak örnek veri yapısıdır
type Customer struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

func main() {
	// sync.Pool nesnesi oluşturulur
	pool := &sync.Pool{
		New: func() interface{} {
			// Yeni bir Customer nesnesi oluşturulur ve geri döndürülür
			return &Customer{}
		},
	}

	// Havuza bir Customer nesnesi eklenir
	customer := pool.Get().(*Customer)
	customer.ID = 1
	customer.FirstName = "Örnek İsim"
	customer.LastName = "Örnek Soyisim"
	customer.Address = "Örnek Address"

	// Havuzdaki veri kullanılır
	fmt.Println("Havuzdaki Müşteri:", customer)

	// Veri havuza geri döndürülür
	pool.Put(customer)

	// Havuzdan bir veri alınır
	customerDataFromPool := pool.Get().(*Customer)
	fmt.Println("Havuzdan Alınan Müşteri:", customerDataFromPool)

	// Havuzdaki veri kullanılır
	// ...

	// Veri havuza geri döndürülür
	pool.Put(customerDataFromPool)
}

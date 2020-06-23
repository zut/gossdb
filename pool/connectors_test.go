package pool

import (
	"sync"
	"testing"

	"github.com/zut/gossdb/conf"
)

func BenchmarkConnectors_NewClient10(b *testing.B) {
	pool := NewConnectors(&conf.Config{
		Host:        "127.0.0.1",
		Port:        8888,
		MaxWaitSize: 10000,
		PoolSize:    5,
		MaxPoolSize: 10,
		MinPoolSize: 10,
	})
	err := pool.Start()
	if err != nil {
		b.Fatal(err)
	}
	b.SetParallelism(10)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c, err := pool.NewClient()
			if err == nil {
				//_, _ = c.Info()
				c.Close()
			} else {
				//b.Error(err)
			}
		}
	})

	pool.Close()
}

func Test1(t *testing.T) {
	pool := NewConnectors(&conf.Config{
		Host:         "127.0.0.1",
		Port:         8888,
		MaxWaitSize:  10000,
		PoolSize:     20,
		MinPoolSize:  10,
		MaxPoolSize:  10,
		HealthSecond: 2,
	})
	err := pool.Start()
	if err != nil {
		panic(err)
	}

	c, err := pool.NewClient()
	if err == nil {
		_, _ = c.Get("a")
		c.Close()
	} else {
		t.Error(err)
	}

	pool.Close()
}
func Test1000(t *testing.T) {
	pool := NewConnectors(&conf.Config{
		Host:         "127.0.0.1",
		Port:         8888,
		MaxWaitSize:  10000,
		PoolSize:     10,
		MinPoolSize:  20,
		MaxPoolSize:  100,
		HealthSecond: 2,
	})
	err := pool.Start()
	if err != nil {
		panic(err)
	}
	var wait sync.WaitGroup
	for i := 0; i < 100; i++ {
		wait.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				c, err := pool.NewClient()
				if err == nil {
					_, _ = c.Get("a")
					c.Close()
				} else {
					t.Error(err)
				}
			}
			wait.Done()
		}()
	}
	wait.Wait()
	//time.Sleep(time.Second * 20)
	pool.Close()
}

func BenchmarkConnectors_NewClient100(b *testing.B) {
	pool := NewConnectors(&conf.Config{
		Host:        "127.0.0.1",
		Port:        8888,
		MaxWaitSize: 10000,
		PoolSize:    10,
		MinPoolSize: 100,
		MaxPoolSize: 100,
	})
	err := pool.Start()
	if err != nil {
		b.Fatal(err)
	}
	b.SetParallelism(100)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c, err := pool.NewClient()
			if err == nil {
				//_, _ = c.Info()
				c.Close()
			} else {
				//b.Error(err)
			}
		}
	})

	pool.Close()
}
func BenchmarkConnectors_NewClient1000(b *testing.B) {
	pool := NewConnectors(&conf.Config{
		Host:        "127.0.0.1",
		Port:        8888,
		MaxWaitSize: 10000,
		PoolSize:    20,
		MinPoolSize: 100,
		MaxPoolSize: 100,
		//Password:    "vdsfsfafapaddssrd#@Ddfasfdsfedssdfsdfsd",
	})
	err := pool.Start()
	if err != nil {
		b.Fatal(err)
	}
	b.SetParallelism(1000)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c, err := pool.NewClient()
			if err == nil {
				//_, _ = c.Info()
				c.Close()
			} else {
				//b.Error(err)
			}
		}
	})

	pool.Close()
}
func BenchmarkConnectors_NewClient5000(b *testing.B) {
	pool := NewConnectors(&conf.Config{
		Host:        "127.0.0.1",
		Port:        8888,
		MaxWaitSize: 100000,
		PoolSize:    20,
		MaxPoolSize: 500,
		MinPoolSize: 500,
	})
	err := pool.Start()
	if err != nil {
		b.Fatal(err)
	}
	//pool.SetNewClientRetryCount(4)
	b.SetParallelism(5000)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c, err := pool.NewClient()
			if err == nil {
				//_, _ = c.Info()
				c.Close()
			} else {
				//b.Error(err)
			}
		}
	})

	pool.Close()
}

//func TestCheck(t *testing.T) {
//	pool := NewConnectors(&conf.Config{
//		Host:         "127.0.0.1",
//		Port:         8888,
//		MaxWaitSize:  10000,
//		PoolSize:     10,
//		MinPoolSize:  10,
//		MaxPoolSize:  10,
//		HealthSecond: 2,
//	})
//	err := pool.Start()
//	if err != nil {
//		panic(err)
//	}
//	defer pool.Close()
//	for {
//		c, err := pool.NewClient()
//		if err == nil {
//			if v, err := c.Get("a"); err == nil {
//				t.Log(v.String())
//			} else {
//				t.Error(err)
//			}
//			c.Close()
//		} else {
//			t.Error(err)
//		}
//	}
//
//}
func TestAutoClose1(t *testing.T) {
	pool := NewConnectors(&conf.Config{
		Host:         "127.0.0.1",
		Port:         8888,
		MaxWaitSize:  10000,
		PoolSize:     10,
		MinPoolSize:  10,
		MaxPoolSize:  10,
		HealthSecond: 2,
		AutoClose:    true,
	})
	err := pool.Start()
	if err != nil {
		panic(err)
	}
	defer pool.Close()
	//
	v, err := pool.GetClient().Get("a")
	t.Log(v, err)
}
func TestAutoClose2(t *testing.T) {
	pool := NewConnectors(&conf.Config{
		Host:         "127.0.0.1",
		Port:         8888,
		MaxWaitSize:  10000,
		PoolSize:     10,
		MinPoolSize:  10,
		MaxPoolSize:  10,
		HealthSecond: 2,
		AutoClose:    true,
	})
	//

	err := pool.Start()
	if err != nil {
		panic(err)
	}
	defer pool.Close()
	for i := 0; i < 100; i++ {
		c, err := pool.NewClient()
		if err != nil {
			panic(err)
		}
		v, err := c.Get("a")
		t.Log(v, err)
	}
}
func BenchmarkConnectors_NewClient5000a(b *testing.B) {

	b.SetParallelism(5000)
	a := "a"
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = a == ""
		}
	})

}
func BenchmarkConnectors_NewClient5000b(b *testing.B) {

	b.SetParallelism(5000)
	a := 0
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = a == 0
		}
	})

}

func TestAutoClose3(t *testing.T) {
	pool := NewConnectors(&conf.Config{
		Host:         "127.0.0.1",
		Port:         8888,
		MaxWaitSize:  10000,
		PoolSize:     10,
		MinPoolSize:  10,
		MaxPoolSize:  10,
		HealthSecond: 2,
		AutoClose:    true,
		Password:     "vdsfsfafapaddssrd#@Ddfasfdsfedssdfsdfsd",
	})
	//

	err := pool.Start()
	if err != nil {
		panic(err)
	}
	defer pool.Close()
	for i := 0; i < 100; i++ {
		c, err := pool.NewClient()
		if err != nil {
			panic(err)
		}
		if v, err := c.Get("a"); err != nil {
			t.Log(err)
		} else {
			t.Log(v)
		}
		c.Close()
	}
}
func BenchmarkConnectors_Set1k(b *testing.B) {
	pool := NewConnectors(&conf.Config{
		Host:        "127.0.0.1",
		Port:        8888,
		MaxWaitSize: 100000,
		PoolSize:    20,
		MaxPoolSize: 100,
		MinPoolSize: 100,
		AutoClose:   true,
	})
	err := pool.Start()
	if err != nil {
		b.Fatal(err)
	}
	//pool.SetNewClientRetryCount(4)
	b.SetParallelism(1000)
	k1 := make([]byte, 1024)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if err := pool.GetClient().Set("1k", k1); err != nil {
				b.Error(err)
			}
		}
	})
	pool.Close()
}

func BenchmarkConnectors_Get1k(b *testing.B) {
	pool := NewConnectors(&conf.Config{
		Host:        "127.0.0.1",
		Port:        8888,
		MaxWaitSize: 100000,
		PoolSize:    20,
		MaxPoolSize: 100,
		MinPoolSize: 100,
		AutoClose:   false,
	})
	err := pool.Start()
	if err != nil {
		b.Fatal(err)
	}
	//pool.SetNewClientRetryCount(4)
	b.SetParallelism(1000)
	//k1 := make([]byte, 1024)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if c, err := pool.NewClient(); err == nil {
				c.Close()
			}
		}
	})
	pool.Close()
}

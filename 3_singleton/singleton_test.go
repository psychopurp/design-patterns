package singleton

import (
	"fmt"
	"sync"
	"testing"
)

func TestSingleton(t *testing.T) {
	ins1 := GetInstance()
	fmt.Println(ins1)
	ins1.Version = "2.0"
	ins2 := GetInstance()
	fmt.Println(ins2)
	// ins2 := &Singleton{Version: "1.0.0"}
	fmt.Printf("%p  %p\n", ins1, ins2)
	if ins1 != ins2 {
		t.Fatal("instance is not equal")
	}
}

func TestParallelSingleton(t *testing.T) {
	wg := sync.WaitGroup{}
	const paralCount = 100000
	wg.Add(paralCount)
	instances := [paralCount]*Singleton{}
	for i := 0; i < paralCount; i++ {
		go func(index int) {
			instances[index] = GetInstance()
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i := 1; i < paralCount; i++ {
		if instances[i] != instances[i-1] {

			t.Fatal("nstance is not equal")
		}
	}

}

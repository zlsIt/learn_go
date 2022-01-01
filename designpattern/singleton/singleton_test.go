package singleton

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSingleton(t *testing.T) {
	equal := reflect.DeepEqual(getInstance(), getInstance())
	fmt.Println(equal)
}

func BenchmarkSingleton(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if getInstance() != getInstance() {
				b.Errorf("test fail")
			}
		}
	})
}

func TestLazyInstance(t *testing.T) {
	equal := reflect.DeepEqual(getLazyInstance(), getLazyInstance())
	t.Log(equal)
}

func BenchmarkLazyInstance(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if getLazyInstance() != getLazyInstance() {
				b.Errorf("test fail")
			}
		}
	})
}

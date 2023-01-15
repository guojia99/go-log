package glog

import (
	"log"
	"testing"
)

func BenchmarkGetCodeRunPackage(b *testing.B) {
	a := func() {
		func() {
			func() {
				func() {
					getCodeRunPackage(1)
				}()
			}()
		}()
	}
	for i := 0; i < b.N; i++ {
		a()
	}
}

func TestGetCodeRunPackage(t *testing.T) {
	//log.SetPrefix()
	log.SetFlags(log.Llongfile)
	log.Println(1)
}

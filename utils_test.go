/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/2/26 下午5:22.
 *  * Author: guojia(https://github.com/guojia99)
 */

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

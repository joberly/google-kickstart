package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type testTC struct {
	tc  testcase
	exp int
}

var testTCs = []testTC{
	testTC{
		tc: testcase{
			blocks: "ABAACCA",
			qs: []qstn{
				qstn{3, 6},
				qstn{4, 4},
				qstn{2, 5},
				qstn{6, 7},
				qstn{3, 7},
			},
		},
		exp: 3,
	},
}

func TestRunTestcase(t *testing.T) {
	for i, ttc := range testTCs {
		t.Run(fmt.Sprintf("TC %d", i), func(t *testing.T) {
			act := ttc.tc.run()
			if act != ttc.exp {
				t.Errorf("failed: actual %d expected %d", act, ttc.exp)
			}
		})
	}
}

func BenchmarkNewPrefixSums(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 1; i <= 100000; i *= 10 {
		sbytes := make([]byte, i)
		for j := 0; j < i; j++ {
			sbytes[j] = byte((r.Int() % 26) + 'A')
		}
		b.Run(fmt.Sprintf("Length %d", i), func(b *testing.B) {
			doBenchmarkForString(string(sbytes), b)
		})
	}
}

func doBenchmarkForString(s string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = newPrefixSums(s)
	}
}

func BenchmarkLargeTestcase(b *testing.B) {
	for n := 0; n < b.N; n++ {
		// Setup the test data for the testcase
		b.StopTimer()
		T := testcase{}
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		sz := 100000
		bs := make([]byte, sz)
		for i := 0; i < sz; i++ {
			bs[i] = byte((r.Int() % 26) + 'A')
		}
		T.blocks = string(bs)
		T.qs = make([]qstn, sz)
		for i := 0; i < sz; i++ {
			Qi := qstn{}
			Qi.left = (r.Int() % sz) + 1
			Qi.right = (r.Int() % (sz - Qi.left + 1)) + Qi.left
			T.qs[i] = Qi
		}

		// Run the testcase
		b.StartTimer()
		T.run()
	}
}

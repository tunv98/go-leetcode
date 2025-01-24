package random

import (
	"fmt"
	"testing"
)

type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		queue: make([]int, 0),
	}
}

func (r *RecentCounter) Ping(t int) int {
	r.queue = append(r.queue, t)
	for len(r.queue) > 0 && r.queue[0] < t-3000 {
		r.queue = r.queue[1:]
	}
	return len(r.queue)
}

func Test_Number_Of_recent_call(t *testing.T) {
	recentCounter := Constructor()
	recentCounter.Ping(1)
	recentCounter.Ping(100)
	recentCounter.Ping(3001)
	recentCounter.Ping(3002)
	fmt.Println(recentCounter.queue)
}

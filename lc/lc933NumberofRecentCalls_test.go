package lc

import "testing"

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
	// 移除所有小於t-3000的請求
	for len(r.queue) > 0 && r.queue[0] < t-3000 {
		r.queue = r.queue[1:]
	}
	return len(r.queue)
}

func TestRecentCounter(t *testing.T) {
	rc := Constructor()
	if got := rc.Ping(1); got != 1 {
		t.Errorf("rc.Ping(1) = %d; want 1", got)
	}
	if got := rc.Ping(100); got != 2 {
		t.Errorf("rc.Ping(100) = %d; want 2", got)
	}
	if got := rc.Ping(3001); got != 3 {
		t.Errorf("rc.Ping(3001) = %d; want 3", got)
	}
	if got := rc.Ping(3002); got != 3 {
		t.Errorf("rc.Ping(3002) = %d; want 3", got)
	}
}

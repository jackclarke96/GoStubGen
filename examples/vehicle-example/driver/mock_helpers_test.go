package driver

import (
	"sync"
	"time"
)

type QueuedItem[T any] struct {
	Fn    T
	Delay time.Duration
}

type methodConfig[T any] struct {
	enabled  bool
	mu       sync.Mutex
	queue    []QueuedItem[T]
	fallback interface{}
}

// Set a fallback function
func (m *methodConfig[T]) SetResponseFunc(f T) {
	m.fallback = f
}

// Set a static value as fallback
func (m *methodConfig[T]) SetStaticResponse(f T) {
	m.SetResponseFunc(f)
}

// Enqueue a function with delay
func (m *methodConfig[T]) EnqueueWithDelay(f T, d time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.queue = append(m.queue, QueuedItem[T]{Fn: f, Delay: d})
}

// Enqueue a static function with delay
func (m *methodConfig[T]) EnqueueStaticWithDelay(f T, d time.Duration) {
	m.EnqueueWithDelay(f, d)
}

// Enqueue multiple functions without delay
func (m *methodConfig[T]) SetResponseFuncQueue(fns []T) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, fn := range fns {
		m.queue = append(m.queue, QueuedItem[T]{Fn: fn, Delay: 0})
	}
}

// Enqueue a single function (shortcut)
func (m *methodConfig[T]) SetResponseFuncOnce(f T) {
	m.SetResponseFuncQueue([]T{f})
}

// Enqueue a function N times
func (m *methodConfig[T]) SetResponseFuncTimes(f T, times int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for range times {
		m.queue = append(m.queue, QueuedItem[T]{Fn: f, Delay: 0})
	}
}

// Clear queue
func (m *methodConfig[T]) ResetQueue() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.queue = nil
}

// Check queue length
func (m *methodConfig[T]) PeekQueueLength() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return len(m.queue)
}

// Get next response from queue or fallback
func (m *methodConfig[T]) NextResponse(defaultFunc T) T {
	m.mu.Lock()
	defer m.mu.Unlock()

	if len(m.queue) > 0 {
		item := m.queue[0]
		m.queue = m.queue[1:]
		if item.Delay > 0 {
			time.Sleep(item.Delay)
		}
		return item.Fn
	}

	if f, ok := m.fallback.(T); ok {
		return f
	}

	return defaultFunc
}

/*
func (m *mockSelfDriving) setApplyBrakesFunc(f func(float64) bool)
func (m *mockSelfDriving) setApplyBrakesResponse(output0 bool)
func (m *mockSelfDriving) enqueueApplyBrakesResponse(output0 bool)
func (m *mockSelfDriving) enqueueApplyBrakesResponseWithDelay(output0 bool, delay time.Duration)
func (m *mockSelfDriving) enqueueApplyBrakesResponseFunc(f func(float64) bool)
func (m *mockSelfDriving) enqueueApplyBrakesResponseFuncWithDelay(f func(float64) bool, d time.Duration)
func (m *mockSelfDriving) enableApplyBrakesMock()
func (m *mockSelfDriving) disableApplyBrakesMock()
*/

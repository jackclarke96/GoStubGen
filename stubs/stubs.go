package stubs

import (
	"sync"
	"time"
)

type QueuedItem[T any] struct {
	Fn    T
	Delay time.Duration
}

type MethodConfig[T any] struct {
	Enabled  bool
	mu       sync.Mutex
	queue    []QueuedItem[T]
	Fallback interface{}
}

// Set a Fallback function
func (m *MethodConfig[T]) SetResponseFunc(f T) {
	m.Fallback = f
}

// Set a static value as Fallback
func (m *MethodConfig[T]) SetStaticResponse(f T) {
	m.SetResponseFunc(f)
}

// Enqueue a function with delay
func (m *MethodConfig[T]) EnqueueWithDelay(f T, d time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.queue = append(m.queue, QueuedItem[T]{Fn: f, Delay: d})
}

// Enqueue a static function with delay
func (m *MethodConfig[T]) EnqueueStaticWithDelay(f T, d time.Duration) {
	m.EnqueueWithDelay(f, d)
}

// Enqueue multiple functions without delay
func (m *MethodConfig[T]) SetResponseFuncQueue(fns []T) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, fn := range fns {
		m.queue = append(m.queue, QueuedItem[T]{Fn: fn, Delay: 0})
	}
}

// Enqueue a single function (shortcut)
func (m *MethodConfig[T]) SetResponseFuncOnce(f T) {
	m.SetResponseFuncQueue([]T{f})
}

// Enqueue a function N times
func (m *MethodConfig[T]) SetResponseFuncTimes(f T, times int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i := 0; i < times; i++ {
		m.queue = append(m.queue, QueuedItem[T]{Fn: f, Delay: 0})
	}
}

// Clear queue
func (m *MethodConfig[T]) ResetQueue() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.queue = nil
}

// Check queue length
func (m *MethodConfig[T]) PeekQueueLength() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return len(m.queue)
}

// Get next response from queue or Fallback
func (m *MethodConfig[T]) NextResponse(defaultFunc T) T {
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

	if f, ok := m.Fallback.(T); ok {
		return f
	}

	return defaultFunc
}

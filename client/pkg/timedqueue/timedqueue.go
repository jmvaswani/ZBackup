package timedqueue

import (
	"log"
	"sync"
	"time"
)

// TimedQueue represents a timed queue where tasks are processed after a delay.
type TimedQueue struct {
	mu            sync.Mutex
	taskTimers    map[string]*time.Timer        // Map to store timers by task value.
	taskCallbacks map[string]func(string) error // Map to store the callback functions for each task.
	defaultDelay  time.Duration                 // Default delay for the tasks.
}

// NewTimedQueue creates a new instance of TimedQueue.
func NewTimedQueue(delay time.Duration) *TimedQueue {
	return &TimedQueue{
		taskTimers:    make(map[string]*time.Timer),
		taskCallbacks: make(map[string]func(string) error),
		defaultDelay:  delay,
	}
}

// AddTask adds a task with a value and a callback function.
// If a task with the same value exists, it resets the timer.
func (q *TimedQueue) AddTask(value string, callback func(string) error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	// If there's already a timer for this value, stop and reset it.
	if timer, exists := q.taskTimers[value]; exists {
		log.Printf("Stopping task with value %s ", value)
		timer.Stop()
	}

	// Set the callback function for this value.
	q.taskCallbacks[value] = callback

	// Create a new timer to execute the task after the specified delay.
	q.taskTimers[value] = time.AfterFunc(q.defaultDelay, func() {
		// Call the callback function when the timer expires.
		log.Printf("Calling task with value %s", value)
		err := callback(value)
		if err != nil {
			log.Printf("Error occured on task with value %s -> %s", value, err)
		}
		// Remove the task after execution.
		q.mu.Lock()
		delete(q.taskTimers, value)
		delete(q.taskCallbacks, value)
		q.mu.Unlock()
	})
}

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Task func()

type Deque struct {
	mu    sync.Mutex
	tasks []Task
}

func (d *Deque) pushFront(task Task) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.tasks = append([]Task{task}, d.tasks...)
}

func (d *Deque) PopFront() (Task, bool) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if len(d.tasks) == 0 {
		return nil, false
	}
	task := d.tasks[0]
	d.tasks = d.tasks[1:]
	return task, true
}

func (d *Deque) PopBack() (Task, bool) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if len(d.tasks) == 0 {
		return nil, false
	}
	task := d.tasks[len(d.tasks)-1]
	d.tasks = d.tasks[:len(d.tasks)-1]
	return task, true
}

type Worker struct {
	id    int
	deque *Deque
	pool  *WorkerPool
}

func (w *Worker) Start() {
	go func() {
		for {
			if task, ok := w.deque.PopFront(); ok {
				task()
			} else {
				if task, ok := w.pool.Steal(w.id); ok {
					task()
				} else {
					time.Sleep(10 * time.Millisecond)
				}
			}
		}
	}()
}

type WorkerPool struct {
	workers []*Worker
}

func NewWorkerPool(numWorkers int) *WorkerPool {
	pool := &WorkerPool{workers: make([]*Worker, numWorkers)}
	for i := 0; i < numWorkers; i++ {
		pool.workers[i] = &Worker{
			id:    i,
			deque: &Deque{},
			pool:  pool,
		}
	}
	return pool
}

func (p *WorkerPool) Steal(thiefID int) (Task, bool) {
	for i := 0; i < len(p.workers); i++ {
		if i != thiefID {
			if p.workers[i] == nil {
				fmt.Printf("workers %d is nil\n", i)
				return nil, false
			}
			if p.workers[i].deque == nil {
				fmt.Printf("deque is nil\n")
				return nil, false
			}
			if task, ok := p.workers[i].deque.PopBack(); ok {
				return task, true
			}
		}
	}
	return nil, false
}

func (p *WorkerPool) SubmitTask(task Task) {
	worker := p.workers[rand.Intn(len(p.workers))]
	worker.deque.pushFront(task)

}

func main() {
	pool := NewWorkerPool(300)
	for i := 0; i < 300; i++ {
		pool.workers[i].Start()
	}
	for i := 0; i < 10000; i++ {
		taskId := i
		pool.SubmitTask(func() {
			fmt.Printf("Task %d is being executed\n", taskId)
		})
	}

	time.Sleep(2 * time.Second)

}

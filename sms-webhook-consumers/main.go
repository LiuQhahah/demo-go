package main

import (
	"context"
	"fmt"
	"time"
)

type SMSWebhookConsumerService interface {
	Start()
	Stop()
}

type SMSWebhookConsumer struct {
	ctx context.Context
}

// Start begins the polling process.
func (qc *SMSWebhookConsumer) Start() {
	go qc.manageGoroutines()
}

// Stop signals the polling to stop.
func (qc *SMSWebhookConsumer) Stop() {
	qc.cancel()
}

// Increase/Decrease number of polling routines based on queue size
func (qc *SMSWebhookConsumer) manageGoroutines() {
	var currentRoutines int
	var cancels []context.CancelFunc = []context.CancelFunc{}
	for {
		select {
		case <-qc.ctx.Done():
			// Handle graceful shutdown of all goroutines
			return
		default:
			queueLength, err := qc.checkQueueLength() // Implement this function to check the queue length
			if err != nil {
				logger.Error(qc.ctx, "error checking queue length: ", err)
				time.Sleep(30 * time.Second)
				continue
			}
			// Calculate the desired number of routines based on queue length or other heuristics
			desiredRoutines := CalculateDesiredRoutines(queueLength, MIN_ROUTINES, MAX_ROUTINES)
			// Rebalance the desired number of routines
			desiredRoutines = RebalanceDesiredRoutines(currentRoutines, desiredRoutines)
			// Adjust the number of goroutines based on the desired count
			for i := currentRoutines; i < desiredRoutines; i++ {
				ctx, cancel := context.WithCancel(qc.ctx)
				cancels = append(cancels, cancel)
				ctx = context.WithValue(ctx, contextKeys.PollContextKey, len(cancels))
				logger.Info(qc.ctx, fmt.Sprintf("starting goroutine with context id: %d", len(cancels)))
				go qc.poll(ctx)
				currentRoutines++
			}
			for i := currentRoutines; i > desiredRoutines; i-- {
				// Signal one goroutine to stop through context
				cancels[len(cancels)-1]()
				cancels = cancels[:len(cancels)-1] // Remove the last cancel function
				logger.Info(qc.ctx, fmt.Sprintf("stopping goroutine with context id: %d", len(cancels)))
				currentRoutines--
			}
			time.Sleep(10 * time.Second)
		}
	}
}

// poll continuously polls the SQS queue for messages.
func (qc *SMSWebhookConsumer) poll(ctx context.Context) {
	errorBackoff := 5 * time.Second
	contextID := ctx.Value(contextKeys.PollContextKey).(int)
	logger.Info(ctx, fmt.Sprintf("polling started for queue with context id: %d", contextID))
	for {
		select {
		case <-qc.ctx.Done():
			logger.Info(qc.ctx, "polling stopped.")
			return
		case <-ctx.Done():
			logger.Info(qc.ctx, fmt.Sprintf("polling stopped, reduced by goroutines. Context id: %d", contextID))
			return
		default:
			// receives call the aws sdk function within it
			messages, err := qc.receiveMessages()

			// Process the received messages

			// Sleep for a short period to avoid throttling in case of empty receives
			tts := 2 * time.Second
			if len(messages) == 0 {
				tts = 5 * time.Second
			}
			time.Sleep(tts)
		}
	}
}
func CalculateDesiredRoutines(queueLength, minRoutines, maxRoutines int) int {
	desiredRoutines := queueLength / MAX_MSG_PER_ROUTINE
	if desiredRoutines < minRoutines {
		return minRoutines
	}
	if desiredRoutines > maxRoutines {
		return maxRoutines
	}
	return desiredRoutines
}

func RebalanceDesiredRoutines(currentRoutines int, desiredRoutines int) int {
	if desiredRoutines <= currentRoutines {
		return desiredRoutines
	}

	// gopsutil library is used to get cpu and memory utilisation

	cpu, err := CheckCPUUtilization()
	if err != nil {
		logger.Error(context.Background(), "error checking CPU utilization: ", err)
		return currentRoutines
	}

	mem, err := CheckMemoryUtilization()
	if err != nil {
		logger.Error(context.Background(), "error checking memory utilization: ", err)
		return currentRoutines
	}

	if cpu > MAX_CPU_PC || mem > MAX_MEM_PC {
		return max(currentRoutines-1, MIN_ROUTINES)
	}

	multiple := desiredRoutines/currentRoutines + 1

	if (cpu*float64(multiple) < 95) && (mem*float64(multiple) < 95) {
		return desiredRoutines
	}

	for ((cpu*float64(multiple) > 100) || (mem*float64(multiple) > 100)) && multiple > 1 {
		multiple -= 1
	}

	return currentRoutines * multiple
}

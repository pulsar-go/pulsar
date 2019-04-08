package queue

/*
	newJob(...).dispatch()
*/
type Job struct {
	Handler     func()
	ShouldQueue bool
}

// NewJob creates a new Job struct
func NewJob(handler func()) *Job {
	return &Job{
		Handler:     handler,
		ShouldQueue: false,
	}
}

// Queue tells the job to run on a goroutine or not.
func (j *Job) Queue(shouldQueue bool) *Job {
	j.ShouldQueue = shouldQueue
	return j
}

// Dispatch runs the handler in the queue or synchronously
func (j *Job) Dispatch() error {
	if j.ShouldQueue {
		return Dispatch(j.Handler)
	}
	j.Handler()
	return nil
}

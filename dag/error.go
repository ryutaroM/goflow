package dag

const (
	ErrorCycleDetectedNum = iota + 1
)

type ErrorCycleDetected struct{}

func (e *ErrorCycleDetected) Error() string {
	return "cycle detected in DAG"
}

func NewErrorCycleDetected() error {
	return &ErrorCycleDetected{}
}

func wrapError(code int) error {
	switch code {
	case ErrorCycleDetectedNum:
		return NewErrorCycleDetected()
	default:
		return nil
	}
}

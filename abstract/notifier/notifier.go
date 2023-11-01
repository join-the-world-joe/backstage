package notifier

const (
	CMD     = "CMD"
	Stop    = "Stop"
	Destroy = "Destroy"
)

type Notifier interface {
	Name() string
	Wait() <-chan string
	Emit(string) error
	Destroy()
}

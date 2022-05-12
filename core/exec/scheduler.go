package exec

var (
	queue []func() error
)

func Schedule(fn Processor) {
	// ctx := currentContext // TODO
	queue = append(queue, func() error {
		return fn(nil) // TODO: ctx
	})
}

func ProcessQueue() error {
	for i := 0; i < len(queue); i++ {
		entry := queue[i]
		if err := entry(); err != nil {
			return err
		}
	}
	return nil
}

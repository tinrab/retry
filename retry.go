package retry

import (
	"time"
)

// Do retries calling function f n-times.
// It returns an error if none of the tries succeeds.
func Do(n int, f func(attempt int) error) (err error) {
	for i := 0; i < n; i++ {
		err = f(i)
		if err == nil {
			return nil
		}
	}
	return err
}

// DoSleep retries calling function f n-times and sleeps for d after each call.
// It returns an error if none of the tries succeeds.
func DoSleep(n int, d time.Duration, f func(attempt int) error) (err error) {
	for i := 0; i < n; i++ {
		err = f(i)
		if err == nil {
			return nil
		}
		time.Sleep(d)
	}
	return err
}

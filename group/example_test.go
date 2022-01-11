package group

import (
	"errors"
	"fmt"
	"time"
)

func ExampleGroup_Add_basic() {
	var g Group
	{
		cancel := make(chan struct{})
		execute := func() error {
			select {
			case <-time.After(time.Second):
				fmt.Printf("The first actor had its time elapsed\n")
				return nil
			case <-cancel:
				fmt.Printf("The first actor was canceled\n")
				return nil
			}
		}
		interrupt := func(err error) {
			fmt.Printf("The first actor was interrupted with: %v\n", err)
			close(cancel)
		}
		g.Add(execute, interrupt)
	}
	{
		execute2 := func() error {
			fmt.Printf("The second actor is returning immediately\n")
			return errors.New("immediate teardown")
		}
		interrupt2 := func(err error) {
			// Note that this interrupt function is called, even though the
			// corresponding execute function has already returned.
			fmt.Printf("The second actor was interrupted with: %v\n", err)
		}
		g.Add(execute2, interrupt2)
	}
	fmt.Printf("The group was terminated with: %v\n", g.Run())
	// Output:
	// The second actor is returning immediately
	// The first actor was interrupted with: immediate teardown
	// The second actor was interrupted with: immediate teardown
	// The first actor was canceled
	// The group was terminated with: immediate teardown
}

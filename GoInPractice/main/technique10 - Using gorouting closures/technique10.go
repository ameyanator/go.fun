package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Outside a goroutine")

	go func() {
		// runtime.Gosched()
		fmt.Println("Inside a goroutine")
	}()
	fmt.Println("Outside again")
	runtime.Gosched()
}

/*
runtime.Gosched(). This is a
way to indicate to the Go runtime that you’re at a point where you could pause and
yield to the scheduler. If the scheduler has other tasks queued up (other goroutines),
it may then run one or more of them before coming back to this function.

There are other ways of yielding to the scheduler; perhaps the most common is to
call time.Sleep. But none gives you the explicit ability to tell the scheduler what to do
when you yield. At best, you can indicate to the scheduler only that the present gorou-
tine is at a point where it can or should pause. Most of the time, the outcome of yield-
ing to the scheduler is predictable. But keep in mind that other goroutines may also
hit points at which they pause, and in such cases, the scheduler may again continue
running your function.

For example, if you execute a goroutine that runs a database query, running
runtime.Gosched may not be enough to ensure that the other goroutine has com-
pleted its query. It may end up paused, waiting for the database, in which case the
scheduler may continue running your function. Thus, although calling the Go sched-
uler may guarantee that the scheduler has a chance to check for other goroutines,
you shouldn’t rely on it as a tool for ensuring that other goroutines have a chance to
complete.
*/

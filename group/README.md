# Find a way to manage goroutine lifecycles
## Background
It's common to manage a group of grountines which run synchronously and if one exit, others need to exit either. We have to write code like below: 

```go
go someMethod()  // cancel()
go http.Serve(listener,api) // listener.Close()
go cronJobs(cronCh) // Close(cronCh)
```

After start batch goroutine, We have to `select` and catch the result of synchronously func result.

## Explore
We can figure out a code craft for this situation. Here are the definitions of some key concepts.
1. Actor： we can recognize it a set of action of `go func(){}` and `cancelFunc()`
2. Execute: the execute function which run synchronously
3. Interrupt: cause the execute function return
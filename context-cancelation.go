/*
Simple excample how goroutine cancelation with Context works

This is excample how to prevent goroutine leaking
*/

func main() {
  fmt.Println("Hello")
  ctx, cancel := context.WithCancel(context.Background())

  go func(ct context.Context) {
    select {
    case <-ctx.Done():
      return
    case <-time.Tick(3 * time.Second):
      fmt.Println("Message from goroutine")
    }
  }(ctx)

  go func(ct context.Context) {
    select {
    case <-ctx.Done():
      return
    case <-time.Tick(3 * time.Second):
      fmt.Println("Message from goroutine 2")
    }
  }(ctx)

  fmt.Println("Canceling")
  cancel() // calling for a goroutine cancelation
  time.Sleep(30 * time.Second)
}

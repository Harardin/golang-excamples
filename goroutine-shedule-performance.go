package scrp

/*

Handles Goroutines shedule to save memmory

Allows to create goroutine only if we actually going to use them

*/

type Pool struct {
  Work chan func()
  Sem  chan struct{}
}

func New(size int) *Pool {
  return &Pool{
    Work: make(chan func()),
    Sem:  make(chan struct{}, size),
  }
}

func (p *Pool) Shedule(task func()) error {
  select {
  case p.Work <- task:
  case p.Sem <- struct{}{}:
    go p.worker(task)
  }
  return nil
}

func (p *Pool) worker(task func()) {
  defer func() {
    <-p.Sem
  }()
  for {
    task()
    task = <-p.Work
  }
}

package engine

import "sync"

type commandQueue struct {
	sync.Mutex
	cmds      []Command
	noCommand bool
	pushed    chan struct{}
}

type cmdExecutor struct {
	executor func()
}

type EventLoop struct {
	queue        *commandQueue
	awaitFinish  bool
	finishSignal chan struct{}
}

func (cq *commandQueue) push(cmd Command) {
	cq.Lock()
	defer cq.Unlock()
	cq.cmds = append(cq.cmds, cmd)
	if cq.noCommand {
		cq.noCommand = false
		cq.pushed <- struct{}{}
	}
}

func (cq *commandQueue) pull() Command {
	cq.Lock()
	defer cq.Unlock()
	if len(cq.cmds) == 0 {
		cq.noCommand = true
		cq.Unlock()
		<-cq.pushed
		cq.Lock()
	}
	cmd := cq.cmds[0]
	cq.cmds[0] = nil
	cq.cmds = cq.cmds[1:]
	return cmd
}

func (cq *commandQueue) length() int {
	cq.Lock()
	defer cq.Unlock()
	return len(cq.cmds)
}

func (ce *cmdExecutor) Execute(h Handler) {
	ce.executor()
}

func (el *EventLoop) Start() {
	el.queue = &commandQueue{
		pushed: make(chan struct{}),
	}
	el.finishSignal = make(chan struct{})
	go func() {
		for !(el.awaitFinish && el.queue.length() == 0) {
			cmd := el.queue.pull()
			cmd.Execute(el)
		}
		el.finishSignal <- struct{}{}
	}()
}

func (el *EventLoop) Post(cmd Command) {
	el.queue.push(cmd)
}

func (el *EventLoop) AwaitFinish() {
	finish := &cmdExecutor{func() { el.awaitFinish = true }}
	el.Post(finish)
	<-el.finishSignal
}

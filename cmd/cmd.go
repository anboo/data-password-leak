package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"
)

type C interface {
	Run(ctx context.Context) error
	Name() string
}

type Cmd struct {
	name string
	F    func(ctx context.Context) error
}

func (c Cmd) Run(ctx context.Context) error {
	return c.F(ctx)
}

func (c Cmd) Name() string {
	return c.name
}

func NewCmdFunc(name string, F func(ctx context.Context) error) C {
	return Cmd{name: name, F: F}
}

type App struct {
	ctx      context.Context
	commands []C
	lock     sync.RWMutex
}

func NewApp(ctx context.Context) *App {
	return &App{ctx: ctx, commands: []C{}, lock: sync.RWMutex{}}
}

func (a *App) RegisterCommand(cmd C) {
	a.lock.Lock()
	defer a.lock.Unlock()

	a.commands = append(a.commands, cmd)
}

func (a *App) Find(name string) (C, error) {
	a.lock.RLock()
	defer a.lock.RUnlock()

	for _, c := range a.commands {
		if c.Name() == name {
			return c, nil
		}
	}

	return nil, fmt.Errorf("command not found")
}

func (a *App) Execute() {
	if len(os.Args) < 2 {
		panic("Cannot fetch command from os.Args. Expected \"start\" or etc")
	}

	c, err := a.Find(strings.Trim(os.Args[1], ""))
	if err != nil {
		panic(err.Error())
	}

	err = c.Run(a.ctx)
	if err != nil {
		panic(err.Error())
	}
}

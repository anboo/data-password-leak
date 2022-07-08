package cmd

import (
	"context"
	"fmt"
	"log"
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
	ctx            context.Context
	commands       map[string]C
	lock           sync.RWMutex
	defaultCommand *C
}

func NewApp(ctx context.Context) *App {
	return &App{ctx: ctx, commands: map[string]C{}, lock: sync.RWMutex{}}
}

func (a *App) RegisterDefaultCommand(cmd C) {
	a.defaultCommand = &cmd
}

func (a *App) RegisterCommand(cmd C) {
	ex, _ := a.Find(cmd.Name())
	if ex != nil {
		log.Fatalf("Command %s already exists", cmd.Name())
	}

	a.lock.Lock()
	defer a.lock.Unlock()

	a.commands[cmd.Name()] = cmd
}

func (a *App) Find(name string) (C, error) {
	a.lock.RLock()
	defer a.lock.RUnlock()

	c, ok := a.commands[name]
	if !ok {
		return nil, fmt.Errorf("command %s not found", name)
	}

	return c, nil
}

func (a *App) Execute() {
	var c C
	var err error

	if len(os.Args) < 2 {
		if a.defaultCommand == nil {
			log.Fatalf("not found command")
		}
		c = *a.defaultCommand
	} else {
		c, err = a.Find(strings.Trim(os.Args[1], ""))
		if err != nil {
			panic(err.Error())
		}
	}

	err = c.Run(a.ctx)
	if err != nil {
		panic(err.Error())
	}
}

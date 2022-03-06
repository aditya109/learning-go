# SOLID PATTERNS

## References

1. [Reference to Swarvanu Sengupta's article](https://s8sg.medium.com/solid-principle-in-go-e1a624290346)

## Introduction - what, why and how ?

According to Robert C. Martin the symptoms of a rotting design or bad code are:
       1. Rigidity - *Code to be difficult to change even if the change is small.*
       2. Fragility - *Code to break whenever a new change is introduced in the system.*
       3. Immobility - *Code being not reusable.*
       4. Viscosity - *Hacking rather than finding a solution that would preserve the design when it comes to change.*

## S stands for Single responsibility principle SRP

> *"Do one thing and do it well."* - Mcllroy
>
> Meaning in lay man terms, *if you have a class or a function, it should undertake one primary responsibility of its knowledge, and not of any other.*

The single responsibility principle suggests that two separate aspects of a problem need to be handled by a different module. In other words, changes in a module should be originated from only one reason.

In OO language, if you have more than one responsibilities embedded into a single class, the internal logics becomes highly coupled, which makes the class less responsive towards changes. Similarly, if you have two separate classes say class `A` and class `B`, and if the consumer of class `A` needs to know about class `B`, then `A` and `B` are considered highly coupled. *SRP aims to maintain a good level of soupling that also maintains a good level of cohesion.*

A simple example could be the following package `journal`,

```go
// example/problem/journal/journal.go

package journal

type Journal struct {
	entries []string
	count   int
}

func (J *Journal) AddEntry(entry string) {
	J.count++
	J.entries = append(J.entries, entry)
}

func (J *Journal) DeleteEntry(position int) {
	J.entries = append(J.entries[:position], J.entries[position+1:]...)
}

func (J Journal) String() string {
	var result string
	for i := 0; i < J.count; i++ {
		result += fmt.Sprintf("Entry %d: %s", i+1, J.entries[i])
		result += fmt.Sprint("\n")
	}
	return result
}
```

Let's say till here the package `journal` is in accordance with SRP, but if we add more functions to it...

```go
// journal.go

package journal

type Journal struct {
	entries []string
	count   int
}

func (J *Journal) AddEntry(entry string) {
	// check above
}

func (J *Journal) DeleteEntry(position int) {
	// check above
}

func (J Journal) String() string {
	// check above
}

// Adding these new functions
func (J Journal) SaveToFile(filename string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND, 0666)
	check(err)
	defer file.Close()
	_, err = file.WriteString(J.String())
	check(err)
}

func (J Journal) LoadFileFromDisk(filename string) {
	// implementation pending
}

func (J Journal) LoadFileFromWeb(url string) {
	// implementation pending
}
```

At this point package `journal`has undertaken secondary responsibility as well as owning to its primary one, this is in violation of SRP, which later emerges as *GOD PROBLEM*.

A quick re-designing of code solves this problem:

```go
// example-1/solution/journal/journal.go

package journal

import (
	"fmt"
)

type Journal struct {
	entries []string
	count   int
}

func (J *Journal) AddEntry(entry string) {
	J.count++
	J.entries = append(J.entries, entry)
}

func (J *Journal) DeleteEntry(position int) {
	J.entries = append(J.entries[:position], J.entries[position+1:]...)
}

func (J Journal) String() string {
	var result string
	for i := 0; i < J.count; i++ {
		result += fmt.Sprintf("Entry %d: %s", i+1, J.entries[i])
		result += fmt.Sprint("\n")
	}
	return result
}
```

```go
// example-1/solution/persistancemanager/persistancemanager.go

package persistancemanager

import (
	"os"

	"github.com/aditya109/learning-go/design-patterns/solid-patterns/1-s/example-1/solution/journal"
)

// Adding these new functions
func SaveToFile(filename string, journal journal.Journal) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND, 0666)
	check(err)
	defer file.Close()
	_, err = file.WriteString(journal.String())
	check(err)
}

func LoadFileFromDisk(filename string, journal journal.Journal) {
	// implementation pending
}

func LoadFileFromWeb(url string, journal journal.Journal) {
	// implementation pending
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
```

Let's take another example. Say we have a module `Command` in a command-driven system. The `Command` module decode, validate and finally execute the incoming commands.

```go
type Command struct {
	commandType string
	args        []string
}

func (c Command) Decode(data []byte) {
	// decodes and initialise
}

func (c Command) ValidateCommandType() bool {
	// validates command type
}

func (c Command) ValidateArgs() bool {
	// validate provided args as if input
}

func (c Command) Execute() {
	// Executes separate types of commands
}

```

In this case, changes on how a `Command` gets decoded and validated and how a command get executed will directly affect the `Command` module. Hence the module performas multiple responsibilties and highly coupled. As per SRP the `Decode()` and `Validate()` is a separate concern than `Execute()`, and should be handle in separate modules.

We can introduce the `CommandFactory` module that parses, validates and initializes a command, where the `CommandExecutor` module executes the command. Now `CommandFactory` and `CommandExecutor` are loosely coupled via `Command` module.

```go
package command

type Command struct {
	commandType string
	args        []string
}

type CommandFactory struct {
	//
}

// Create decode and validate the command
func (cf CommandFactory) Create(data []byte) (*Command, error) {
	// decode command
	command, err := cf.Decode(data)
	if err != nil {
		return nil, err
	}
	// validate type
	switch cf.Type {
	case Foo:
	case Bar:
	default:
		return nil, InvalidCommandType
	}
	return command, nil
}

type CommandExecutor struct {
}

// Execute executes the command
func (ce CommandExecutor) Execute(command *Command) ([]byte, error) {
	// validate input and execute
	switch command.Type {
	case Foo:
		if len(args) == 0 || len(args[0]) == 0 {
			return nil, InvalidInput
		}
		return ExecuteFoo(command)
	case Bar: // Bar doesn't take any input
		return ExecuteBar(command)
	}
}
```



 

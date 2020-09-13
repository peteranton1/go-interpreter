---
# Writing an Interpreter in Go
---
# README

(view [CHANGELOG](./CHANGELOG.md) for notable changes)


This project contains example code
for an implementation of an interpreter 
for a target language. 

The target language is a made up language called Monkey. 
Monkey is quite interesting but by no means a production 
quality language, and this implementation is also not 
production quality. Libraries were not used to code generate 
e.g. the lexers and parsers. Instead these were hand-coded.
This ensures that there is minimal magic, i.e. the coder 
knows how it all works.

The idea is that the journey is more important than 
the destination. Writing this software allowed me to 
understand the architecture and mechanisms better
through the act of implementation. 

This was also a non-trivial implementation for practice 
in the Go language. 

## How to compile and run

### Run Tests

The Tests are run like any go program. 

From the folder containing a (file).go and (file)_test.go, issue 

    go test (file).go
    e.g.
    go test lexer.go
    go test parser.go

### Run REPL

The REPL is a command line program that presents the interactive interpreter. 

From the folder containing main.go

    go run main.go

This command compiles and runs the main program, which starts an interactive shell, 
known as the REPL. This accepts statements in the Monkey language and evaluates 
each line and prints the output to the screen. 

### Build executable and run

To build an executable and then run

    go build -o monkey . && ./monkey

### Example Monkey code

Here are some example statements in Monkey:

    20 + 100;
    if true {
        100
    } else {
        200
    }
    let five = 5;
	let ten = 10;
	let add = fn(x, y) {
		x + y;
	};
	let result = add(five,ten);
    result;

### Benchmarks

The code implements both an interpretter and also a compiler with VM. 

To create the benchmark program, issue:

    go build -o fibonacci ./benchmark

To run in interpretted mode, the command is:

    $ ./fibonacci -engine=eval
    
    engine=eval, result=9227465, duration=22.557914554s
   
To run in compiled mode, the command is:

    $ ./fibonacci -engine=vm
    
    engine=vm, result=9227465, duration=5.517457919s
   
The results shown were created on an ordinary i5 laptop. 
Interpretted mode took 23 seconds, Compiled mode took 6 seconds.
This indicates that the fibonnaci program ran much faster in compiled mode.

### Ideas for improvements

* Add more documentation and functionality to the monkey language
* Provide more example programs
* Additional run modes apart from REPL
* Implement a module system
* Implementations in languages other than Go

### Credits

This implementation was largely based on the excellent books called:

*    "Writing an Interpreter in Go" by Thorsten Ball
*    "Writing a Compiler in Go" by Thorsten Ball


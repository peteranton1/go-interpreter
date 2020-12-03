# Benchmark Results
## for interpreter vs compiler/VM

Results:

```
    $ cd go/proj/src/monkey
    $ go build -o fibonacci ./benchmark
    
    $ ./fibonacci -engine=eval
    engine=eval, result=9227465, duration=21.388178439s
    
    $ ./fibonacci -engine=vm
    engine=vm, result=9227465, duration=5.966806974s
```
# pipelinePatternInGo
Sample Pipeline Pattern in Go

This pattern is useful when you want to have a series of processing stages and each stage performs a specific operation and can work independently.
Example: You want to perform a series of operations on a stream of data and you want them to run concurrently.
In this pipeline pattern, each stage takes input from the previous stage, performs the operation, and passes the result to the next stage.

In this program, we have four stages.
1) Generator
2) Square
3) Double
4) Print

Each stage is implemented as a function which accepts an input channel and provides an output channel. The previous stage's output channel becomes the next stage's input.
Each stage runs independently of each other. As soon as the generator generates a number, it can get back to generating the next number without waiting for any other process to complete.

> Interface values - regardless of their type and whether they are empty or not - hold two things:
> 1. the concrete type of some value (or no type); and
> 2. the value of that concrete type (or no value)

Go more or less requires you to check a type before getting a value. (Type Assertions?)

The first step, figure out what the type is. Sometimes there is a step before a step.
```go
if i == nil {
    // if i has a typed value in it, and can hold nil pointers
    // the value part can be nil, and the i==nil is false. i does have a type
}
```
## Empty Interfaces

Are they just adhoc interfaces?
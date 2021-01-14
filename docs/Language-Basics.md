Go basics/syntax stuff.

## Packages

The differences between libraries and packages.

Library | Executable
--------|---------------
Created for Reusability | Created for Running Program
Non-Executables | Executable
Importable | Non-Importable
Name Agnostic | Only be Named Main
No Main Function | func main

## Scope

## Control Flow

### Conditionals
In computer science, control flow (flow of control) is the order in which individual statements, instructions or function calls of imperative program are executed or evaluated.

```go
if condition {
    // do something
} else if condition {
    // do something
} else {
    // do something
}

a := 3
b := 4
```
An expression is evaluated, and if the result is **true** the code in the conditional block body is executed. If it is **false** it is ignored. Go supports nesting it's conditional statements.

Conditionals rely on a Boolean expression (one that evaluates to true or false) to decide whether the code they contain should be executed. If you need to only execute code if a condition is **false** using the **Boolean negation operator** `!` will allow you to make a true statement false or a false statement true.

Using `&&` will run code if both expressions are evaluated true, and using `||` will run code if **either** expression is evaluated true.

### Loops

Code that runs repeatedly is known as a loop.



## Methods

Methods are functions that are associated with values of a particular type. Essentially, the dot indicates that the thing on the right belongs to the thing it's left. Methods belong to an individual value, where functions belong to a package.

Blank Identifiers

Error Handling

Block and Variable Scope

## Pointers

Values that represent the address of a variable are known as **pointers**, they point to the location where the variable is stored and can be accessed.

* https://www.golang-book.com/books/intro/8
* https://flaviocopes.com/golang-methods-receivers/
* https://medium.com/@vCabbage/go-are-pointers-a-performance-optimization-a95840d3ef85
* https://www.ardanlabs.com/blog/2017/05/language-mechanics-on-stacks-and-pointers.html
* https://medium.com/a-journey-with-go/go-should-i-use-a-pointer-instead-of-a-copy-of-my-struct-44b43b104963
* https://dev.to/chen/gos-method-receiver-pointer-vs-value-1kl8
* https://qvault.io/2019/09/25/the-proper-use-of-pointers-in-go-golang/
* https://www.digitalocean.com/community/conceptual_articles/understanding-pointers-in-go
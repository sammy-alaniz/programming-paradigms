# Question

2. We have a struct in C that describes a `Shape`:

typedef struct _shape {
    /* todo (as many members as you wish) */
} Shape;

As a user, I want to be able to use it like this:

Shape *a = createCircle(55/*radius*/);
Shape *b = createSquare(10/*side*/);
a->circumference(a);
b->circumference(b)

Show what would be behind `createCircle` and `createSquare` and update
the definition of the struct accordingly.

# Answer

Note: this answer is in Go instead of C. This semseter we didn't really go over C, but we did go over Go.
```
package main

import (
    "fmt"
    "math"
)

// Shape interface
type Shape interface {
    Circumference() float64
}

// Circle struct
type Circle struct {
    Radius float64
}

// Circle's Circumference method
func (c Circle) Circumference() float64 {
    return 2 * math.Pi * c.Radius
}

// Create a new Circle
func NewCircle(radius float64) *Circle {
    return &Circle{Radius: radius}
}

// Square struct
type Square struct {
    Side float64
}

// Square's Circumference method
func (s Square) Circumference() float64 {
    return 4 * s.Side
}

// Create a new Square
func NewSquare(side float64) *Square {
    return &Square{Side: side}
}

func main() {
    var a Shape = NewCircle(55) // radius
    var b Shape = NewSquare(10) // side

    fmt.Printf("Circle Circumference: %f\n", a.Circumference())
    fmt.Printf("Square Circumference: %f\n", b.Circumference())
}

```
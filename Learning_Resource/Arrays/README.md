# Arrays

An array is a contiguous collection of elements of the same type. It is an ordered sequence of elements stored contiguously in memory.

Basic Syntax
- created an array named `example`
- `<array-size>` : No. of elements in the array (i.e. 2 , 5, 10).
- `<type>` type of element in the array.
- `X1, X2, ...Xn ` are the elements it the array.

```go
example := [<array-size>]{<type>}{X1, X2, X3.......Xn}
```
> Remeber: Array is the data type whose length is immutable.

Another want array declaration is if we do not know the no. of elements in the array.
```go
[...]int{2, 4, 5, 7, 8, 9}
```
```go
[...]int{}
```

### Multi-Dimentional Array
```go
example := [x][y]{type}{{X11, X12 .. X1y},
                       {X21, X22 .. X2y},
                       {..           ..},
                       {Xx1, Xx2 .. Xxy}} 
```
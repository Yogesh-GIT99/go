# Pointers

* What are pointers ? 
- variables that stores value addresses instead of values.

* Why pointers ? 
- Avoid unecessary value copies: 
    By default, go creates a copy when passing values to function. 
    For very large and complex values this may take up too much memory space unnecessarily.
    With pointers only one value is stored in memory ( and the address is passed around)

- Directly mutate values: 
    Pass a pointer( address) instead of a value to a function.
    The function can then directly edit the underlying value - no return value is required.
    Can lead to less code ( but also to less understandable code or unexpected behaviors) 

* How to use a pointer ( declaration, retreiving and use in function) ? 
- Check code

* Pointer's Null value is nil. 

* In most of the cases since the values we are dealing with are generally small in mem size, we should not use pointers, unless required.

* Using pointers for data mutation
- Check code

* Scan function ( e.g )
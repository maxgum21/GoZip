# First steps 

This is the first part of my Go learning journey.

I will write my first impressions and useful info here

## Go as a language

Go ( if i understand correctly ) can be both built into binaries by the go compiler and be run immediately as if it were an interpreted language.

This makes it pretty easy for me to run first easy one-file programs to try out the syntax and stuff.

## Now for the details

Looks like you don't need semicolons at the ends of code blocks ( like the ones ending in '}' ).

Go has a "package" keyword. I'm going to have to investigate on them further.

Imports look pretty normal: just an "import" keyword followed by the package name string.

Functions have pretty familiar syntax, however I still have yet to investigate how to assign the return type.

Variables in Go are interesting in that Go can infer types by itself. You can also assign types by writing the type after the variable's name.
To declare variables you either use the "var" keyword, or the ":=" operator.
You can also assign multiple variables at once with a single "=" operator which is pretty neat.

For loops can be similar to C for loops ( assignment/condition/iteration ) but are also the while loop, as you can simply omit everthing except the condition.
So for loops are also while loops. You can make an infinite loop by not putting anything.

Conditionals are standard, however you can also put an expression before the condition. If you declare a variable it will be in the scope the condition
as well as every branch. One gripe i have with them is also that you can't put non-boolean values into conditions e.g. 1 or a non-empty string.

Switch-case statements are pretty good in Go. You can use them to substitute if/else logic by writing them without the checking value,
you can also check for multiple values with one case statement by separating them with a comma. 

Arrays seem pretty simple as well. Example: `"var a [5]int"` - initializes an array of 5 integers, all zeroed. you can initialize them in curly braces.
You can also add ':' colons to initialize only from one index on. Use `"[...]"` to make the array use the length of the initialization.

Go slices are basically C++ vectors with Python like slicing and are declared using `[]`. Use functinon "make" to create slices of predetermined length.
Also "slices" package.

Maps are kinda similar to Python dicts. Declare them with `map[key-type]val-type` syntax. Use "delete(map, keyword)" function to delete a key/value pair.
A very neat detail that Go maps have is that when you try to get the value from a map object it returns 2 values, the value itself and a boolean
value indicating if the **key** is present in the map. Also "map" package contains additinal functionality.

Ranges in Go are basically iterators that give you both the index and the value from the array.

Go functions are pretty intuitive for me. To add the return type just write the type between arguments and the opening curly brace.
They also can return multiple values at once with a very simpe syntax `"func function (...) (type1, type2, ...) { ... }"`
Variadic functions are very simple and interesting. Example: `func foo (vars...var-type)` - "vars" will be a slice of var-type elements

Go also has lamda functions which are called "closures". You can also curry functions. Closures can also be recursive, however they have to be declared first.

Go supports pointer use. They seem very C-like.

Go has "runes" instead of "chars". Runes can be multiple bytes long and are represented by UTF-8 characters. "utf8" is a package with utilities to work with utf-8 characters.

Go Structs are very C like. The biggest difference is that the '.' operator allows to get the specified fields from the struct pointer as well.
Structs support methods. Syntax is pretty easy: `func (s (*)struct-type) (...) {}`

Error handling is very nice in Go because of the way that functions allow returns and the way that if/else statements are handled. "error" package for utils.

Goroutines are a basic threading implementation. Syntax: `"go foo()"`
Channels allow for between-thread passing of values. This also allows to sync routines together.
You can also specify the direction of channels in the function arguments.
"select" keyword alllows you to wait for multiple channel outputs.
"<- time.After" allows you to set timeouts for goroutine management.
You can close channels and check if they are closed, using a second variable.
You can also range over the channel, which will continously iterate over the channel until it gets closed.
You can use it in a single thread after closing the channel and still recieve the values

"sync" and "waitGroups" packages allow for easier goroutine management. "sync" gives you mutexes and atomic values.

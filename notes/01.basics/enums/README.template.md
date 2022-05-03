{
"CreatedAt": "2022-04-28",
"UpdatedAt": "2022-04-29",
"Title": "Enums - Go notes"
}
---
# Enums
There are no enumns types in Go, but we can use constants and variables to get something similar. A good way to insure the type safety of enums is to define them as their own types.
This way if you have Month = 12 and Bananas = 12, then Month == Bananas will be false.
It's also useful to create a String() method for that type so packages like fmt and log
can print out the textual values of the enum.


<<enum.go>>

<<enum_test.go>>

Output:
```
december
```

## Using variables and structs
In this examples we create a type called Color that is a Struct and another type called NamedColor based on Color. 
We then create our enumns based on NamedColor.

<<enum2.go>>

Here we can also see that we can type cast NamedColor back to Color to check equality.

<<enum2_test.go>>

output:
```
magenta
```
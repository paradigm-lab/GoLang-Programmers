
Data Types:
-> In any Programming language you have multiple data types
-> Difference: which data types do they support exactly?

String -> For textual data, defined with double quotes, eg: "This is a String"

Integers -> Representing whole number, positive and negative, eg: 5, 120, -20
    -> There are many more numeric data types!
	Different Integer Types and Different Floating Types

-> Each data type can do different things and behaves differently

Go is a statically typed language
-> You need to tell Go Compiler, the data type when declaring the variable
	-> More Robust, Reduce the likelihood of errors
	-> Helps developers to catch type mismatches sooner (at compile time)

-> Type Inference: BUT, Go can infer the type when you assign a value using the assignment operator	   	


Variables: 

Variables are used to store values
Like containers for values 
Give the variable a name and reference everywhere in the app
Variable names should be descriptive
 
"Syntactic Sugar" in Programming
	-> A term to describe a feature in a language that let's you do samething more easily
	-> Makes the language "sweeter for human use
	-> But doesn't add any new functionality that it didn't already have

We can not use := with constants


Data structure to store collection of elements in a single variable

Using Arrays and slices

Arrays:
Arrays in Go have got the Fixed size (size = how many elements the array can hold) and data type or (type)

Slices: 
A list that is more dynamic in size 
Slice is an abstraction of an Array
Slice are more flexible and powerfull:
Variable-length or get an sub-array of its own
Slices are also index-based and have a size, but is resized when needed


Functions: 
-> Encapsulate code into own container (= function). Which logically belong together!
-> Like variable name, you should give a function a descriptive name
-> Call the function by its name, whenever you want to execute this block of code
-> Every program has at least one function, which is the main() function 

-> Function is only executed, when "called"!
-> You can call a function as many times you want
-> So a function is also used to reduce code duplication

Function Parameters
-> Information can be passed into functions as parameters

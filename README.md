-> Programming language developed at Google in 2007 by: Robert Griesemar, Rob Pike, Ken Thompson(UNIX)
-> Open-sourced in 2009

Why Go? Go use Cases
Evolution of Infrastructure:
	
Infrastructure changed a lot 
-> Scalable & Distribute
-> Dynamic
-> More Capacity

Existing Programming Languages did not fully take advantage of it
-> 1 task at a time vs doing multiple tasks at once

-> Multi-Threading Do multiple things at once.

-> Concurrency is about dealing with lots of things at once
	-> Developers need to write code to prevent conflicts

-> Go was designed to run on mutiple cores and built to support concurrency

-> Concurrency in Go is cheap and easy	


Main Use case of Go:
-> For Performant Applications
-> Running on scaled, distribute systems


Characterstics of Go:
-> Attempt to combine both:	
	-> Simple and readable syntax of a dynamically typed language like python
	-> Efficiency and Safety of a lower-level, statically typed language like C++

-> Server-Side or Backend Language
	-> Microservices
	-> Web Applications
	-> Database Services

-> Simple Systax: Easy to learn, read and write code

-> Fast build time, start up and run

-> Requires fewer resources

-> Compiled Language
	-> Compiles into single binary (Machine code)
	-> Consistent across different OS

-> Strong and Statically typed

-> Excellent Community

-> key features
	* Simplicity
	* Fast compile times
	* Garbage collected
	* Built-in concurrency
	* Compile to standalone binaries

For more Info visit: golang.org

----------------------------------------------------------------------------------------------------------------------------------------------------------

### Learning Wheels:
go mod init<module path>
-> Creates a new module
-> Module path can correspond to a repository you plan to publish your module to (eg: github.com/paradigm-lab/booking-app)
	-> Initialized a go.mod file
	-> Describes the module: with name/module path and go version used in the program


Packages in Go 
-> Go programs are organized into packages
-> A package is a collection of Go files


go run
go build 
go install


Variables:
- Variable declaration and Initialization
	* var foo int
	* var foo int = 42
	* foo := 32 (Syntactic Sugar)
	
	- Package Level Variables:
		* Defined at the top outside all functions
		* They can be accessed inside any of the functions
		* And in all files, which are in the same package

	- Local Variables:
		* Defined inside a function or a block function
		* They can be accessed only inside that function or block of code

- Redeclaration and shadowing
	* Can't redeclare variables, but can shadow then 
	* All variables must be used

- Visibility
	* Lower case first letter for package scope
	* Upper case first letter export globally 
	* no private scope (Scoping the variable into a block but declaring it in the scope)

- Naming conventions
	* Pascal or camelCase
	* Capitalize acronyms(HTTP, URL)
	* As short as reasonable
	* Longer names for longer lives

- Type convertions
	* destinationType(variable)
	* No implicity type conversion (Due to the possibility of losing the information eg: float to int)
	* Use strconv package for strings

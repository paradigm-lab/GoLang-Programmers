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


Primitives:
- Boolean Type
	* Values are true or false
	* Not an alias for other types(e.g int)
	* Zero values is false
	* Example: var n bool = false	
		   var n bool (By default it will be zero(0) which means false)

- Numeric Type
	* Integers
		* int() function
		
		* Signed Integers (int8, int16, int32, int64) -> int
		* int type has varying size, but min 32 bits
		* 8 bit(int8) through 64bit(int64)
		
		* Unsigned Inegers (uint8, uint16, uint32, NO uint64 but we have a byte aliase to represent uint8) -> uint (Positive numbers)
		* 8bit(byte and uint8) through 32 bit(uint32)

		* Arithmetic operations (
		* Bitwise Operations (
		* Zero value is 0 (Default in memory)
		* Can't mix types in same family!(uint16 + uint32 = error) 


	* Floating Point 
		* Follow IEEE-754 Standard (32 and 64 bits floating point numbers)
		* float32 +-1.18E-38 -> +-3.4E38
		* float64 +-2.23E-308 -> +-1.8E308 (Default)
		* float() function
		* Zero value is 0 
		* 32 and 64 bit versions
		* Literal styles
			* Decimal (3.14)
			* Exponential (13e18 or 2E10)
			* Mixed (13.7e12)
		* Arithmetic operations
			* Addition, Subtraction, multiplication and division
		* No bitwise operators, No remainder operators, No bit Shifting operators 
	
	* Complex numbers
		* Zero value is (0+0i)
		* 64 and 128 bit versions
		* complex64 and complex128
		* Built-in functions	
			* complex - make complex number from two floats
			* real - get real part as float
			* imag - get imaginary part as float
		* Arithmetic operations
			* Addition, Subtraction, multiplication and division

- Text types
	* Strings
		* UTF-8
		* Strings in Go are aliases for the bits
		* Strings are immutable
		* Can be concatenated with plus (+) oparator
		* Can be converted to []byte
	
	* Rune
		* UTF-32
		* Alias for int32
		* Special methods normally required to process
			* eg: strings.Reader#ReadRune


- Constants
	- Immutable, but can be shadowed
	- Replaced by the compiler at compile time
		* Value must be calculatable at compile time

	* Naming convention	
		* PascalCase for exported constants
		* camelCase for internal constants  

	* Typed constants
		* Work like immutable variables
		* Can interoperate only with same type

	* untyped constants
		* Work like literals
		* Can interoperate with similar types

	* Enumerated constants
		* Special symbol iota allows related constants to be created easily
		* Iota starts at 0 in each const block and increments by one
		* Watch out of constant values that match zero values for variables

	* Enumerated expressions
		* Operations that can be determined at compile time are allowed
			* Arithmetic
			* Bitwise operations
			* Bitshifting

- Arrays and Slices
	* Arrays
	        - Collection of items with same type
		- Fixed Size	
		
		- Creation
			- Declaration Style
				* a := [3]int{1, 2, 3}
				* a := [...]int{1, 2, 3}
				* var a [3]int
		
		- Access via zero-based index
			* a := [3]int {1, 3, 5} // a[1] == 3
		
		- Built-in Functions
			* len function returns size of array

		- Working with arrays
			* Copies refer to different underlying data 

	
	* Slices 
		- Creation 
		- Built-in Functions
		- Working with slices



# Go Interview Questions

Grisemer, Pike, Thompson

**What is a string literal in Go programming ?**

> A string literal specifies a string constant that is obtained from concatenating a sequence of characters.
>
> There are two types of string literals:
>
> - **Raw string literals**: The value of raw string literals are character sequence between back quotes, as in ``. The value of a raw string literal is the string composed of the uninterpreted characters between the quotes; in particular, backslashes have no special meaning and string may contain new lines. Carriage return character (\r) inside raw literals are discarded from the raw string value.
> - **Interpreted string literals**: The string literals are character sequences between double quotes, as in "bar". Within the quotes, any character, with backslash escapes interpreted as they are in rune literals(except that `\'` is illegal and `\"` is legal), with the same restrictions

```go
string_lit             = raw_string_lit | interpreted_string_lit .
raw_string_lit         = "`" { unicode_char | newline } "`" .
interpreted_string_lit = `"` { unicode_value | byte_value } `"` .
`abc`                // same as "abc"
`\n
\n`                  // same as "\\n\n\\n"
"\n"
"\""                 // same as `"`
"Hello, world!\n"
"日本語"
"\u65e5本\U00008a9e"
"\xff\u00FF"
"\uD800"             // illegal: surrogate half
"\U00110000"         // illegal: invalid Unicode code point
```

These examples all represent the same string:

```go
"日本語"                                 // UTF-8 input text
`日本語`                                 // UTF-8 input text as a raw literal
"\u65e5\u672c\u8a9e"                    // the explicit Unicode code points
"\U000065e5\U0000672c\U00008a9e"        // the explicit Unicode code points
"\xe6\x97\xa5\xe6\x9c\xac\xe8\xaa\x9e"  // the explicit UTF-8 bytes
```

**Rune literals ?**

> A rune literal represents a rune constant, an integer value identifying a Unicode code point. A rune literal is expressed as one or more characters enclosed in single quotes, as in 'x' and '\n'. Within the quotes, any characters may appear except newline and unescaped single quotes.
>
> After a backslash, certain single-character escapes represent special values:
>
> ```go
> \a   U+0007 alert or bell
> \b   U+0008 backspace
> \f   U+000C form feed
> \n   U+000A line feed or newline
> \r   U+000D carriage return
> \t   U+0009 horizontal tab
> \v   U+000b vertical tab
> \\   U+005c backslash
> \'   U+0027 single quote  (valid escape only within rune literals)
> \"   U+0022 double quote  (valid escape only within string literals)
> ```
>
> All other sequences starting with a backslash are illegal inside rune literals.
>
> ```go
> rune_lit         = "'" ( unicode_value | byte_value ) "'" .
> unicode_value    = unicode_char | little_u_value | big_u_value | escaped_char .
> byte_value       = octal_byte_value | hex_byte_value .
> octal_byte_value = `\` octal_digit octal_digit octal_digit .
> hex_byte_value   = `\` "x" hex_digit hex_digit .
> little_u_value   = `\` "u" hex_digit hex_digit hex_digit hex_digit .
> big_u_value      = `\` "U" hex_digit hex_digit hex_digit hex_digit
>                            hex_digit hex_digit hex_digit hex_digit .
> escaped_char     = `\` ( "a" | "b" | "f" | "n" | "r" | "t" | "v" | `\` | "'" | `"` ) .
> 'a'
> 'ä'
> '本'
> '\t'
> '\000'
> '\007'
> '\377'
> '\x07'
> '\xff'
> '\u12e4'
> '\U00101234'
> '\''         // rune literal containing single quote character
> 'aa'         // illegal: too many characters
> '\xa'        // illegal: too few hexadecimal digits
> '\0'         // illegal: too few octal digits
> '\uDFFF'     // illegal: surrogate half
> '\U00110000' // illegal: invalid Unicode code point
> ```

**Features of Golang**:

> - Provides concurrency at language level
> - Has Garbage collection.
> - Support CSP-style concurrent programming features (communicating sequential processes).
> - String and maps are built into the language.
> - Functions are first class objects in the language.

**Explain static type declaration in Go.**

> A static type variable declaration provides assurance to the compiler that there is one variable available with the given type and name so that the compiler can proceed for further compilation without requiring the complete detail of the variable. A variable declaration has its meaning at the time of compilation only, the compiler needs the actual variable declaration at the time of linking of the program.
>
> ```go
> package main
> 
> import "fmt"
> 
> func main() {
>    var x float64
>    x = 20.0
>    fmt.Println(x)
>    fmt.Printf("x is of type %T\n", x)
> }
> ```

**Explain dynamic type declaration/ Type inference in Go.**

> A dynamic type variable declaration requires the compiler to interpret the type of the variable based on the value passed to it. The compiler does not require a variable to have type statically as a necessary requirement.

**Is it true that short variable declaration `:=` cab  be used only inside a function ?**

> Yes, a short variable declaration `:=` can be used only inside a function.

**Does Go support operator overloading, method overloading, pointer arithmetic ?**

> Nope.

**Polymorphism with Interfaces**

> 1. Interfaces define the functional contract.
>
>    ```go
>    type Employee interface {
>      GetDetails() string
>      GetEmployeeSalary() int
>    }
>    
>    type Manager struct {
>      Name string
>      Age int
>      Designation string 
>      Salary int
>    }
>    
>    func (mgr Manager) GetDetails() string {
>      return mgr.Name + " " + mgr.Age;
>    }
>    
>    func (mgr Manager) GetEmployeeSalary int {
>      return mgr.Salary
>    }
>    ```
>
> 2. Interfaces are Types in Golang
>
>    ```go
>    newManager := Manager{Name: "Mayank", Age: 30, Designation: "Developer" Salary: 10}
>    
>    var employeeInterface Employee
>    
>    employeeInterface = newManager
>    
>    employeeInterface.GetDetails()
>    ```
>
> ```go
> type Employee interface {
>   GetDetails() string
>   GetEmployeeSalary() int
> }
> 
> // Creating a new type "Manager" containing all functions required by "Employee" Interface
> type Manager struct {
>   Name string
>   Age int
>   Designation string 
>   Salary int
> }
> 
> func (mgr Manager) GetDetails() string {
>   return mgr.Name + " " + mgr.Age;
> }
> 
> func (mgr Manager) GetEmployeeSalary int {
>   return mgr.Salary
> }
> 
> // Creating a new type "Lead" containing all functions required by "Employee" Interface
> type Lead struct {
>   Name string
>   Age int
>   TeamSize string 
>   Salary int
> }
> 
> func (ld Lead) GetDetails() string {
>   return ld.Name + " " + ld.Age;
> }
> 
> func (ld Lead) GetEmployeeSalary int {
>   return ld.Salary
> }
> 
> // Creating new Object of Type Manager
> newLead := Lead{Name: "Mayank", Age: 30, TeamSize: "30" Salary: 10}
> 
> // New variable of type "Employee" created
> var employeeInterface Employee
> 
> //Manager Object assigned to Interface type since Interface Contract is satisfied
> employeeInterface = newManager
> 
> // Invoke Functions that belong to Interface "Employee"
> employeeInterface.GetDetails()
> 
> //Same Interface can store value of Lead Object
> employeeInterface = newLead
> 
> // Invoke Functions of Lead Object that is assigned to Interface "Employee"
> employeeInterface.GetDetails()
> ```
>
> Interfaces as Function Parameters
>
> ```go
> // Function taking Interface as Input Parameter...
> func GetUserDetails(emp Employee) {
>   emp.GetDetails()
> }
> 
> type Employee interface {
>   GetDetails() string
> }
> 
> // Creating a new type "Manager" containing all functions required by "Employee" Interface
> type Manager struct {
>   Name string
>   Age int
>   Designation string 
>   Salary int
> }
> 
> func (mgr Manager) GetDetails() string {
>   return mgr.Name + " " + mgr.Age;
> }
> 
> // Creating a new type "Lead" containing all functions required by "Employee" Interface
> type Lead struct {
>   Name string
>   Age int
>   TeamSize string 
>   Salary int
> }
> 
> func (ld Lead) GetDetails() string {
>   return ld.Name + " " + ld.Age;
> }
> 
> // Creating new Object of Type Manager
> newLead := Lead{Name: "Mayank", Age: 30, TeamSize: "30" Salary: 10}
> 
> // New variable of type "Employee" created
> var employeeInterface Employee
> 
> //Manager Object assigned to Interface type since Interface Contract is satisfied
> GetUserDetails(newManager)
> 
> // Interface can be used to invoke function of either Lead or Manager...
> GetUserDetails(newLead)
> ```
>
> Using interfaces as types in structs
>
> ```go
> type Person struct {
>   name string
>   age string 
>   user Employee
> }
> ```
>
> ```go
> type Person struct {
>   name string
>   age string 
>   user Employee
> }
> 
> // Assigning Object of type "Manager" to "user" Property
> newPerson := Person{name: "Mayank", age: "32", Manager{Name: "Mayank", Age: 30, Designation: "Developer" Salary: 10}}
> 
> // Assigning Object of type "Lead" to "user" Property
> newPerson := Person{name: "Mayank", age: "32", Lead{Name: "Mayank", Age: 30, TeamSize: "30" Salary: 10}}
> ```
### Go Language Notes

**Package Declaration:**
- `package main`: Declares the package name. `main` is used for the entry point of an executable program.

**Imports:**
- `import "fmt"`: Imports the `fmt` package for formatted input and output operations.

**Constants:**
- Constants are declared with `const`.
- Example: `const LoginToken string = "Iamlovely"`
- Visibility: 
  - Start with an uppercase letter for public visibility (exported).
  - Start with a lowercase letter for private visibility (within the same package).

**Variables:**
- Declared with the `var` keyword.
- Must be used if declared.

**Data Types and Declarations:**
1. **String**
   - Example: `var username string = "venu"`
   - Prints the value and type.

2. **Boolean**
   - Example: `var isLoggedIn bool = true`
   - Prints the value and type.

3. **Integer**
   - Example: `var number int = 8`
   - Prints the value and type.

4. **Float**
   - Example: `var decimal float64 = 2.2351345345`
   - Prints the value and type.

5. **Default Values**
   - Uninitialized variables get default values (e.g., `0` for `int`).
   - Example: `var other int` (default value is `0`)
   - Prints the default value and type.

6. **Implicit Type**
   - Type is inferred from the value assigned.
   - Example: `var website = "venukumarmd.web.app"`
   - Prints the value and type.

7. **Short Variable Declaration**
   - Declared with `:=` for concise syntax.
   - Example: `userNo := 10`
   - Prints the value and type.

**Printing and Type Checking:**
- `fmt.Println()`: Prints the value.
- `fmt.Printf("variable is of type: %T \n", var)`: Prints the type of the variable.

**Code Example Breakdown:**
- String Type
- Boolean Type
- Integer Type
- Float Type
- Default Values
- Implicit Type
- No Var Style (Short Variable Declaration)
- Constant Example

**Summary:**
- **Package and Imports**: Essential for program structure and utilizing libraries.
- **Constants and Variables**: Proper declaration and usage are crucial.
- **Data Types**: Know the default values and type inference.
- **Printing**: Use `fmt` package for output and type checking.
- **Visibility**: Understand public vs private based on naming conventions.

By adhering to these guidelines, you can write clear, efficient, and idiomatic Go code.

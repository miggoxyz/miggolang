
# Miggolang 

### Key Features

- **Variable Bindings**:  
  ```miggolang
  let x = 42;
  let y = x + 8;
  ```
- **First-Class Functions**:  
  ```miggolang
  let multiply = fn(a, b) { a * b; };
  multiply(6, 7); // 42
  ```
- **Closures**:  
  ```miggolang
  let makeIncrementer = fn(x) { fn(y) { x + y; }};
  let addTwo = makeIncrementer(2);
  addTwo(3); // 5
  ```
- **Data Structures**:  
  - Arrays: `[1, 2, 3]`
  - Hash Maps: `{ "key": "value" }`
- **Conditional Expressions**:  
  ```miggolang
  if (5 > 3) { "Yes" } else { "No" }
  ```

## Running Miggolang

1. Clone the repository:
   ```bash
   git clone https://github.com/miggoxyz/miggolang.git
   cd miggolang
   ```

2. Build the interpreter:
   ```bash
   go build -o miggolang
   ```

3. Launch the REPL (Read-Eval-Print Loop):
   ```bash
   ./miggolang
   ```

4. Start writing Miggolang code:
   ```text
   >> let x = 10 * 2;
   >> x
   20
   ```

## Technical Highlights

- **Lexer**: Converts raw source code into tokens for processing.
- **Parser**: Constructs an Abstract Syntax Tree (AST) from tokens.
- **Evaluator**: Interprets the AST and executes operations dynamically.
- **Built-in Functions**: Includes functions like `len`, `first`, `last`, and custom user-defined functions.
- **Error Handling**: Gracefully identifies and reports syntax or runtime errors.

## Contibuting

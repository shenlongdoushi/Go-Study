# Variable Declaration

- var foo int
- var foo int = 42
- foo := 42

## Can't redeclare variables, but can shadow them

## All variables must be used if declared.

## Visibility

- lower case first letter for package scope
- upper case first letter for global use (export)
- no private scope

## Naming Conventions

### Pascal or camelCase - Capitalize acronyms (HTTP, URL)

### As short as reasonable - longer names for longer lives

## Type Conventions

- destinationType(variable)
- use **strconv** package for strings

# Primitives

## Boolean Type

- Values are true or false
- Not an alias for other types (e.g. int)
- Zero value is false

## Numeric Types

### Integers

#### Signed Integers

- int type has varying size but min 32 bits
- 8 bit (int8) through 64 bit (int64)
- first bit is used for `+` or `-`
- +0 -> 0, -0 still count negative, so int8 range (-2^7 ~ 2^7-1) is -128 ~127

#### Unsigned Integers

- 8 bit (byte and uint8) through 32 bit (uint32)

#### Arithmetic Operations

- Addition +, subtraction -, multiplication \*, Division /, Remainder %

#### Bitwise Operations

- And &, Or |, XOR ^, And Not &^

#### Zero value is 0

#### Can not mix types in same family! (uint16+uint32=Error)

### Floating Point Numbers

#### Follow IEEE-754 standard

#### Zero value is 0

#### 32 and 64 bit versions

#### Literal styles

- Decimal 3.14
- Exponential 13e18, 2E10
- Mixed 13.7e12

#### Arithmetic Operations

- Addition +, subtraction -, multiplication \*, Division /

### Complex Numbers

#### Zero value is (0+0i)

#### 64 and 128 bit versions (float32+float32, float64+float64)

#### Build-in Functions

- complex - make complex number from two floats
- real - get real part as float
- imag - get imaginary part as float

#### Arithmetic Operations

- Addition +, subtraction -, multiplication \*, Division /

## Text Types

### Strings

- UTF-8
- Immutable
- Can be concatenated with plus(+) operator
- Can be converted to []byte

### Rune

- UTF-32
- Alias for int32
- Special methods normally required to process (e.g. strings.Reader)

# Constants

### Immutable, but can be shadowed

### Replaced by the compiler at compile time

- Value must be calculable at compile time

### Named like variables

- PascalCase for exported constants
- camelCase for internal constants

### Typed constants work like immutable variables

- can interoperable only with same type

### Untyped constants work like literals

- Can interoperable with similar types

### Enumerated Constants

- Special symbol iota allows related constants to be created easily
- Iota starts at 0 in each const block and increments by one
- Watch out of constant values that match zero values for variables

### Enumerated Expressions

- Operations that can be determined at compile time are allowed. (e.g., Arithmetic, Bitwise Operations, Bitshifting.)

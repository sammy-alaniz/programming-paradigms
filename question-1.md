# Question 1

1. Write regular expressions for the following cases:

   a) keyword `abstract`.

   b) identifier that has to start with `_` followed by any number of
   digits and lower case letters.

   c) identifier that either (a) starts and ends with `_` and has any
   number of letters in between, or (b) contains only letters followed by
   digits only.

# Answers

a) keyword `abstract`.
```
\babstract\b
```
Note:
- `\b` This asserts a word boundary, ensuring that `abstract` is recognized as a standalone word and not part of a larger word.

b) identifier that has to start with `_` followed by any number of
   digits and lower case letters.
```
_[a-z0-9]*
```
Note: 
- `_` Matches the underscore character.
- `[a-z0-9]` This character class matches any lower case letter (a-z) or digit (0-9).
- `*` This quantifier matches zero or more of the preceding character class.

c) identifier that either (a) starts and ends with `_` and has any
   number of letters in between, or (b) contains only letters followed by
   digits only.
```
(_[a-zA-Z]+_)|([a-zA-Z]+\d+)
```
Note:
- `(_[a-zA-Z]+_)`
  - `_` Matches the starting underscore.
  - `[a-zA-Z]+` Matches one or more letters (both upper and lower case).
  - `_` Matches the ending underscore.
- `|`  This is the OR operator, allowing for either of the two patterns to be matched.
- `([a-zA-Z]+\d+)`
  - `[a-zA-Z]+` Matches one or more letters.
  - `\d+` Matches one or more digits.
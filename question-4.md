# Question

4. Design a declarative language to manipulate txt files.

A user can perform at least the following operations on a given file
(keep in mind that the file is your input):

  * get number of characters between C1 and C2
  * get number of lines
  * count number of a given character C (e.g., how many 'a' are in a document)

One function that is available to you is `count(c)` which counts the
number of `c` characters in the file.  (You cannot define any other
function, but your language can have an arbitrary number of operators
and other constructs.)

hint: be inspired by SQL and Cypher.

(a) Write a CFG for your language.

(b) Give at least one example for each of the required tasks using
your language.

(c)* Design classes (a sketch only needed) in Smalltalk-like syntax for
classes of your AST.

(d)* Using AST traversal show a sketch in a Smalltalk-like language to
implement "count number of a given character C"

# Answer

To design a declarative language for manipulating text files inspired by SQL and Cypher, we need to define a Context-Free Grammar (CFG), provide examples for each of the required tasks, sketch classes in Smalltalk-like syntax for the AST (Abstract Syntax Tree) classes, and show an AST traversal for implementing a specific function. Here's a breakdown of each requirement:

### (a) Write a CFG for your Language
The CFG for our language, which we'll call TextQuery, might look like this:
```
TextQuery ::= Statement ;
Statement ::= GetCharCount | GetLineCount | CountChar ;
GetCharCount ::= "GET CHAR COUNT BETWEEN" Char "AND" Char "IN" FileName ;
GetLineCount ::= "GET LINE COUNT IN" FileName ;
CountChar ::= "COUNT CHAR" Char "IN" FileName ;
Char ::= "'" AnyCharacter "'" ;
FileName ::= "'" AnyText "'" ;
AnyCharacter ::= /* any single character */ ;
AnyText ::= /* any sequence of characters */ ;
```

### (b) Examples for Each Required Task
1] Get number of characters between C1 and C2:
```
GET CHAR COUNT BETWEEN 'C1' AND 'C2' IN 'example.txt'
```

2] Get number of lines:
```
GET LINE COUNT IN 'example.txt'
```

3] Count number of a given character C:
```
COUNT CHAR 'a' IN 'example.txt'
```

### (c) Sketch Classes in Smalltalk-like Syntax for AST Classes
Here's a simplified sketch of the classes in Smalltalk-like syntax:
```
Object subclass: #TextQuery
    instanceVariableNames: 'statement'.

Object subclass: #Statement
    instanceVariableNames: 'type content fileName'.

Object subclass: #GetCharCount
    instanceVariableNames: 'char1 char2 fileName'.

Object subclass: #GetLineCount
    instanceVariableNames: 'fileName'.

Object subclass: #CountChar
    instanceVariableNames: 'char fileName'.
```

### (d) AST Traversal in Smalltalk-like Language for "Count Number of a Given Character C"
For the AST traversal, you would define a method in your AST class that traverses the tree and performs the necessary operation. Here's a sketch:
```
CountChar >> countInFile: file
    | fileContent charCount |
    fileContent := file readContents.
    charCount := fileContent count: self char.
    ^charCount
```
This sketch represents a method in the CountChar class that reads the contents of a file and uses the count: method to count the occurrences of the specified character. This is a high-level representation and would need more detailed implementation in a real-world scenario.


# Notes

### What is meant by "You cannot define any other function" in the question text?

You cannot introduce new functions beyond count(c) to perform operations. Instead, you have to rely on combining this function with other language constructs and operators to fulfill the required tasks (e.g., counting lines, finding the number of characters between two specific characters).

### What is meant by "operators or constructs"?

 Operators:

Operators are symbols or keywords used to perform operations on data.
They can include arithmetic operators (like +, -, *, / for addition, subtraction, multiplication, division), comparison operators (like =, !=, >, < for equality, inequality, greater than, less than), and logical operators (like AND, OR, NOT). In the context of your text manipulation language, operators could be designed to specify the range of characters (BETWEEN, AND), perform count operations (COUNT), or determine positions in the text.

 Constructs:

Constructs are the building blocks of a language that define its syntax and structure. They determine how statements are formed and how different elements of the language can be combined. This includes loops, conditionals, function calls, variable declarations, and query structures. In your language, constructs might include the format for specifying a file (IN 'filename.txt'), the structure of a query/command (like GET LINE COUNT), and how to specify characters or character ranges.
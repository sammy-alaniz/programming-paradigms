# CSP Notes

```
X:: ⋆[
    c: character;
    west?c → [
        c <> '⋆' → east!c; 
        []
        c = '⋆' → west?c; [
            c <> '⋆' → east!'⋆'; east!c; 
            []
            c = '⋆' → east!'@'
        ]
    ]
]
```

The CSP (Communicating Sequential Processes) program you've written for process X is designed to forward characters from a process named west to a process named east, with a specific rule for handling consecutive asterisk characters ('*'). The rule is that when two consecutive asterisks are encountered, they are replaced with a single '@' character. Let's break down how this program works:


- ``X:: ⋆[ ... ]`` : This defines process X with a repetitive command (indicated by ⋆[ ... ]). This means that the enclosed operations will be repeated indefinitely.

- ``c: character; west?c →`` : This is a declaration and a receive operation. c is a variable that holds a character, and west?c means that X is waiting to receive a character from west and store it in c.

- ``The First Layer of Decision Making [ ... ]`` : After receiving a character, the program makes a decision based on the value of c:

   - ``c <> '⋆' → east!c;`` : If c is not an asterisk ('⋆'), then it is immediately sent to east.
   - ``c = '⋆' → west?c;`` : If c is an asterisk, the program waits to receive the next character from west. This is to check if the next character is also an asterisk.

- ``The Second Layer of Decision Making [ ... ]`` : This part is executed if the first character received was an asterisk. Again, there are two paths:
   - ``c <> '⋆' → east!'⋆'; east!c;`` : If the second character is not an asterisk, X sends the first asterisk and then the second character to east.
   - ``c = '⋆' → east!'@'`` : If the second character is also an asterisk, X sends '@' to east, effectively replacing the two asterisks.


 In summary, this CSP program for process X handles a stream of characters from west and forwards them to east, replacing any sequence of two consecutive asterisks with a single '@' character. The program achieves this through a nested decision-making process that checks each received character and behaves accordingly.
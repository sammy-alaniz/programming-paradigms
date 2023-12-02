# Question

5.* Write a CSP program to simulate a FIFO (first-in-first-out) queue.
A queue can have no more than N integer elements.  A user can `add` a
value in the queue and `remove` a value from the queue.  You can
decide how to deal with an empty queue (when removing) and full queue
(when adding).  Please, document your decision.

hint: queue has to be a process in CSP.

hint: this is similar to the factorial example from the class.

# Answer

```
QUEUE(N: integer):: [
    queue: array[1..N] of integer;
    size: integer;
    size := 0;

    *[
        (size < N); user?add(value) → [queue[size] := value; size := size + 1; user!added;]
        []
        (size = N); user?add(value) → user!full;
        []
        (size > 0); user?remove → 
            [user!queue[1]; 
            -- Shift all elements to the left
            for i := 1 to size-1 do
                queue[i] := queue[i+1];
            size := size - 1]
        []
        (size = 0); user?remove → user!empty;
    ]
]
||
user:: USER
```


# Notes:
Factorial example from class.
```
[ fac(i: 1..LIMIT):
  *[
    n: integer;
    fac(i-1)?n -> [
      n = 0 -> fac(i - 1)!1
      []
      n > 0 -> fac(i + 1)!n-1;
      r: integer;
      fac(i + 1)?r;
      fac(i - 1)!(n*r)
    ]
  ]
  ||
  fac(0):: USER
]
```

### Breakdown of factorial

-  ```<process>!<message> - sending a message to the named process```
-  ```<process>?<message> - receiving a message from the named process```

```
[ fac(i: 1..LIMIT):               # This defines the process 'frac' that can go from i to LIMIT
  *[                              # repetitive command denoted by the *[] notation
    n: integer;                   # indicating that n will be used to store integer values
    fac(i-1)?n -> [               # fac(i-1) <process> ? <message> n ~ waiting for a message of n (int)
      n = 0 -> fac(i - 1)!1
      []
      n > 0 -> fac(i + 1)!n-1;
      r: integer;
      fac(i + 1)?r;
      fac(i - 1)!(n*r)
    ]
  ]
  ||                               # denotes that you are creating another parallel process
  fac(0):: USER                    # just another fac() process
]
```
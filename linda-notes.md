# Linda

## Operations:

- out(t)
  
  – puts a new (passive) tuple into tuplespace

  – the call is non-blocking

- in(p)

  – look up into the tuplespace to find passive tuple

  – matches a template tuple (contains concrete values and variables) 
  
  – blocking call

  – when matched, values are assigned to variables

  – deletes the matched tuple

  – if matched multiple, randomly choose one

- rd(t)

  – same as in(p), but does not delete the tuple

- eval(t)

  – puts an active tuple into the tuplespace

  – similar behavior to out(t), but computing fields happens inside the tuplespace – one process per function

  – non-blocking call

  – once all values are computed the tuple becomes passive

- inp(t)

  – checks if any tuple is matching

  – if none can be matches, return False

  – if one can be found, assign values to variables and return True – if there are several matching, choose one non-deterministically – delete the matching tuple

- rdp(t)

  – similar to inp(t), but this one does not delete the tuple



# Examples


## Semaphore
Sending out ```out("sem")``` means you're done and don't need the resource. Waiting ```in("sem")``` for sem to be available means that you'd like the shared resource. Some process needs to initalize the space with the amount of users that are allowed to access the shared resource by calling ```out("sem")``` N number of times. Those calls are blocking.
```
signal: out("sem")

wait: in("sem")

initialize semaphores to n by execute out("sem") n times
```

## Client - Server
In one distributed system there are many clients and one server. Number of clients is arbitrary large. Clients are sending requests to servers that servers requests and sends responses. Requests sent earlier have priorities.

```
void server() {
  int index = 1;
  ...
while (1) {
    in("request", index, ?req);
    ...
    out("response", index, response);
    index++;
  } 
}
```
- ```?``` means place the returned value in req

```
void client() {
  int index;
  ...
  in("server index", ?index);
  out("server index", index+1);
  ...
  out("request", index, request);
  in("response", index, ?response);
  ...
}
```

```
void init() {
  out("server index", 1);
  eval(server());
  for (int i=0; i<10; i++)
    eval(client());
}
```
- Note: use of eval(), so is this in tuplespace?

## Cigarette Smokers

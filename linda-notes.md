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
Assume a cigarette requires three ingredients to make and smoke: tobacco, paper, and matches. There are three smokers around a table, each of whom has an infinite supply of one of the three ingredients–one smoker has an infinite supply of tobacco, another has paper, and the third has matches. There is also a non-smoking agent who enables the smokers to make their cigarettes by arbitrarily (non-deterministically) selecting two of the supplies to place on the table. The smoker who has the third supply should remove the two items from the table, using them (along with their own supply) to make a cigarette, which they smoke for a while. Once the smoker has finished his cigarette, the agent places two new random items on the table. This process continues forever.

```
void agent(){
  int n;
  while(1) {
    n = (int)((rand()*3) / RAND_MAX);
    switch(n) {
        case 0:
          out("Paper");
          out("Tobacco");
          break;
        case 1:
          out("Tobacco");
          out("Matches");
          break;
        case 2: 
          out("Matches");
          out("Paper");
          break;
    }
  in("OK"); 
  }
}
```

```
void smoker_with_matches() {
  while(1) {
    in("Watch");
    if(rdp("Paper") && rdp("Tobacco")) {
     in("Paper");
     in("Tobacco");
     smoke();
     out("OK");
    }
    out("Watch");
  }
}
```

```
initialize () {
  eval(agent());
  eval(smoker_with_matches());
  eval(smoker_with_paper());
  eval(smoker_with_tobacco());
  out("Watch");
}
```

## Readers - Writers
Some threads may read and some may write, with the constraint that no thread may access the shared resource for either reading or writing while another thread is in the act of writing to it. (In particular, we want to prevent more than one thread modifying the shared resource simultaneously and allow for two or more readers to access the shared resource at the same time.)

```
void reader(){
  int id, num;
  while(1) {
    in("id", ?id);
    out("id", id + 1);
    in("ok to work", id);
    in("readers num", ?num);
    out("readers num", num + 1);
    out("ok to work", id + 1);
    reading();
    in("readers num", ?num);
    out("readers num", num - 1);
  } 
}
```

```
void writer() {
  int id;
  while(1) {
    in("id", ?id);
    out("id", id + 1);
    in("ok to work", id);
    rd("readers num", 0);
    writing();
    out("ok to work", id + 1);
  } 
}
```

```
void init () {
  int i;
  out("id", 0);
  out("ok to work", 0);
  out("readers num", 0);
  for (i=0; i<10; i++) {
    eval(reader());
    eval(writer());
  }
}
```
# Project

In this project, you will database as a service.

We cover graph databases (and talk about Neo4j, which is one
implementation), but if you need more resources check this wiki
article: https://en.wikipedia.org/wiki/Graph_database and this Neo4j
page: https://neo4j.com/docs/getting-started. You can also find free
books on the topic.


## You

Your UT EID plays a role when deciding what language to use and what
features to implement. It is very important to do this computation
properly.

```
hash(eid) => your group
```

hash function is the following:

```
hash(eid) = ascii(eid[0]) + ascii(eid[1]) + ... ascii(eid[n-1])
```

For example:
```
hash(sb43278) = 477
```

In the following section, we will compute problem that you should
solve by the following formula:

```
your group % N => your item
```

N will be defined in each section below, so pay attention.


## Components

This section describes components of the system. The design is modular
and you can easily replace any part of the system with a more robust
implementation.

```
 ---------        -------------
| client  |----->| visualizer  |
 ---------        -------------
   ^
   |
   v
 ---------
| server  |
 ---------
   ^
   |
   v
 ----------
| database |
 ----------
```

* client - client application that has API for sending
  queries/requests to server

* server - web server serving requests

* database - actual graph database implementation

* visualizer - (in a way part of the client) to visualize the result

In the following subsection we describe each of the components in more
detail.

### Client

Client application is responsible for:
* accepting user input
* checking correctness of the query (lexing + parsing)
* sending the request to server (REST)
* accepting the response (json)
* caching the results
* visualizing the results

Client application will be written in several programming languages.

Input by the user will be a Cypher query for the database (which might
be valid or invalid). You can decide in what way to accept the input.
You can decide in what language to write this code (it can even be
Python).

Once you have the input, you should check if the input can be properly
parsed. You should use ANTLR to obtain lexer and parser; grammar file
is already available for Cypher
(https://s3.amazonaws.com/artifacts.opencypher.org/M23/Cypher.g4), so
you do not need to write your own. (We define N=4 in this case: 0-Java
lexer+parser, 1-C/C++ lexer+parser, 2-Go lexer+parser, 3-Python
lexer+parser.) If the input cannot be parsed, give a nice error (your
decision what the error should say) and ask for the next input. If the
input is correct, move to the next step.

Once the input is successfully parsed, you should prepare a request
for the server and send it (json). You can write request code in any
language you wish.

Upon receiving a response (json), you should store the response into a
local relational database. We define N=2 in this case: 0-h2 database,
1-sqlite database. The result should always be a single table.

Every response should be cached (you already stored it in a local db),
so that any future input/request that uses the same input do not need
to be send to the server. This can be written in any language. (Do not
forget that you might need to invalidate cache; feel free to be more
coarse-grained and invalidate everything if there is any modification
to the database in any query.)

### Visualizer

Visualize the response in Smalltalk (proper GUI with nodes and edges)
by communicating between your client and the visualizer via a json
file.

Another version of a visualizer (text-based) should be implemented in
OCaml; communication from your client to OCaml should be the same file
as above. The result should be shown in a tabular
format. Additionally, OCaml should include stats about average values
for each int column (computed in OCaml not as part of a query).


### Server

Your server should be written in Go. The server will accept requests
and serve them. If the request has any error, appropriate message will
be sent to the client. If the request is valid, it will be sent to the
database, results will be accepted, packed, and sent to the user.

We leave to you to design communication protocol between the server
and database, e.g., log files, inter-process messages.

Client and server should communicate using REST. You have freedom to
define end points and arguments.

### Database

You should set up and use Neo4j as your actual database on the server.


## Testing

You should have tests for each part of your code; without tests, code
will be considered non-existent.


## Benchmarking

Graduate version only.

Write a bash script(s) that will collect benchmarking data for 100
queries. Each query should be run 100 times and averages should be
computed.


## Software

Your implementation should work under the following configuration:
* Linux (any recent distribution)
* N=2: 0 - Oracle Java 17 (https://www.oracle.com/java/technologies/downloads); 1 - Oracle Java 11
* N=2: 0 - Neo4j v5.x (cloud Graph Database Self-Managed community edition https://neo4j.com/deployment-center); 1 - Neo4j v4.x (USE4!!)
* Smalltalk Pharo 11 (https://pharo.org/download)
* Go 1.18+
* Python 3.8+
* gcc 9.4.0
* h2 2.2.222
* OCaml 4.08.1
* If you pick a language not in the list, please contact us for the version number

If you create a Docker image with required software and demo
everything using it, you will receive extra points.

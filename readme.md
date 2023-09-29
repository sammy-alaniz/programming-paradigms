# Programming Pradigms Project

hash(eid) = hash(sya369) = 495

1] Need to write an application that accepts user input. Any language.

2] Need to produce ANTLR code with already available grammer file. 
   - N = 4
   - Group = 3
   - Language = Python

3] From user input application, pass the cypher code into the ANTLR lexer and parser.

4] Client to Server send the Cypger if valid. Client to Server should be JSON. Code generating request can be in any language.
   - Every resposnse should be cached
   - Question: What should the local database be? ANSWER: sqlite
   - Don't forget that we may need to invalidate cache.
   - ! feel free to be more coarse grained !
   - ! invalidate everything if there is any modification to the database !
   - Question: does this invalidation require statusing before any query to support more than one client?
   - N = 2
   - Group = 1
   - Server = sqlite

5] Server recieves the JSON request, sends query to relational database 
   - Question: Store the request or the response in a relational database?
   - Question: What language should the server be in?


6] Visualize the response in Smalltalk GUI
   - GUI with nodes and edges
   - Communication between Client and Visualizer is via JSON file

7] Visualize the response in OCaml Text-Based.
   - Language=OCaml
   - Communication between Client and Text Visualizer is via JSON file, same as Smalltalks JSON file
   - Results should be shown in a tabular format
   - OCaml should include stats about average values for each int column
   - Computed in OCaml not as part of a query
   - Question: What does this mean to average values for each int column?

8] Server should be written in GO
   - Language = Go
   - Accept requests and serve them
   - If request has an error, appropriate message will be sent to the client
   - Question: what is the appropriate message?
   - If the request is valid, it will be sent to the database, results will be accepted, packed, and sent to the user.
   - Up to me to design protocol between server and database (log files, inter-process messages)
   - Client <-> Server should communicate using REST. We have freedom to define endpoints and arguments

8] Databasae
   - Database = Neo4j

9] Testing
   - Should have some test for each section

10] Benchmarking
   - Write Bash script that will collect benchmark data for 100 queries.
   - Each query should be ran 100 times and averages should be computed
   - Question: What metrics are we tracking on this test?
   - Question: Amongst the queries, how should they be divided? CRUD?

11] Software Versions
   - Linux (any recent distro)
   - N = 2 : Group = 1
   - Java version = Oracle Java 11
   - Neo4j version = Neo4j v4.x
   - Smalltalk Pharo 11 (I could only get this on MAC)
   - Question: I couldn not get an arm linux version of Smalltalk, only arm mac os. I have everything else linux. Is that okay?
   - Go version = Go 1.18+
   - Python version = Python 3.8+
   - GCC version = GCC 9.4.0
   - Question: I have GCC version 11.3.1, is that okay?
   - H2 version = h2 2.2.222
   - Question: Where is h2 supposed to be used again?
   - TODO: Figure out how to install H2
   - OCaml version = OCaml 4.08.1
   - Question: What needs to be 4.08.1? opam or the ocaml compiler? I have 5.1.0, is this okay?

12] Extra points
   - put all software (minus smalltalk) in docker container
   - Question: Can it be more than one container? Or is it better if it's only one?



### Other
- Docker install link: https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-on-rocky-linux-8

Question: v4 o4 v5 database
Question: how do we test gui code
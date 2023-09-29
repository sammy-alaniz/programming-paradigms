# Starting Neo4j

```
docker pull neo4j
```


```
docker volume create neo4jdata
```


```
docker run \
     --name testneo4j \
     -p7474:7474 -p7687:7687 \
     -d \
     -v neo4jdata:/data \
     -v neo4jdata:/logs \
     -e NEO4J_AUTH=neo4j/test2023test \
     neo4j
```

```
http://localhost:7474/browser/
```
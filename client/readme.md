# Lexer + Parser + Client

## Command to generate

```
java -jar {path}/antlr-4.13.0-complete.jar -Dlanguage=Python3 -visitor -o generated/ Cypher.g4
```

## Commands to start python virtual environment

```
python3 -m venv run-virt

source run-virt/bin/activate

pip install -r requirements.txt
```
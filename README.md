# Evaluator

# Setup:
    go mod download
    cd images/python3
    docker build --tag python3-eval .
    cd ../pypy3
    docker build --tag pypy3-eval .
    cd images/java
    docker build --tag java-eval .
    cd images/c
    docker build --tag c-eval .
    
# Run:
    cd ../../src
    go run .

# Test it with some http requests
**Python3**
    localhost:7070/submit/eval?id=korakora&lang=python3
    localhost:7070/submit/eval?id=korakora&lang=python3&timelimit=2s
    localhost:7070/submit/eval?id=korakora&lang=python3&timelimit=2s&memorylimit=64mb

**Pypy3**
    localhost:7070/submit/eval?id=korakora&lang=pypy3
    localhost:7070/submit/eval?id=dimbo&lang=pypy3

**Java**
    localhost:7070/submit/eval?id=sample&lang=java

**C** 
    localhost:7070/submit/eval?id=nutmeg&lang=c

    **Default time limit = 1s**
    **Default memory limit = 64MB**

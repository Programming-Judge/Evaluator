# Evaluator

# Setup:
    go mod download
    cd images/python3
    docker build --tag python3-eval
    cd ../pypy3
    docker build --tag pypy3-eval
    
# Run:
    cd ../../src
    go run .

# Test it with some http requests
    localhost:7070/submit/eval?id=korakora&lang=python3
    localhost:7070/submit/eval?id=korakora&lang=pypy3
    localhost:7070/submit/eval?id=dimbo&lang=pypy3
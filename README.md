# Evaluator

# Setup:
    go mod download
    cd images/python3
    docker build --tag python3-eval
    
# Run:
    cd ../../src
    go run .

# Test it with a http request
    http://localhost:7070/submit/eval?id=korakora&lang=python3&timelimit=2s
    **Default time limit = 1s**
#!/bin/bash


# used as ./evaluate.sh main.cpp input.txt output.txt time-limit

#compile code
g++ $1 -o main

if [ $? != 0 ]; then

echo "compile failed"
exit

fi

#file to store output
touch code_output.txt

timeout $4 ./main < $2 > code_output.txt

res=$?

#failure in execution
if [ $res -eq 124 ]; then

echo "time limit exceeded - given limit is"
exit

elif [ $res != 0 ]; then

echo "runtime error"
exit

fi

#check for difference in code output and expected output, added --strip-trailing to remove windows' \r insertion in file
diff --strip-trailing-cr code_output.txt $3 

if [ $? != 0 ]; then

echo "wrong output"
exit

fi

echo "sucessfully executed"
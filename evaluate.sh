#!/bin/bash


# used as ./evaluate.sh main.cpp input.txt output.txt

#compile code
g++ $1 -o main

if [ $? != 0 ]; then

echo "compile failed"
exit

fi

#file to store output
touch code_output.txt

./main < $2 > code_output.txt

#failure in execution
if [ $? != 0 ]; then

echo "run failed"
exit

fi

#check for difference in code output and expected output
diff code_output.txt $3 

if [ $? != 0 ]; then

echo "wrong output"
exit

fi

echo "sucessfully executed"
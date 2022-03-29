#!/bin/bash

function cleanup() {
    if [ -f "$1-code-output.txt" ]; then
        rm $1-code-output.txt
    fi
    if [ -f "$1-diff-messages.txt" ]; then
        rm $1-diff-messages.txt
    fi
    # additional cleanup for class files in java
    if [ -f "$2/$1_main.class" ];then
        rm $2/$1_main.class
    fi
}

# Input format
# $1 -> id of submission
# $2 -> extension of code file
# $3 -> bind mounted directory
# $4 -> time limit (to be added)
# $5 -> memory limit (to be added)

# (relative to user home)
#
# The corresponding source, input and output
# should be placed in the "bind_mnt_dir"
# directory with the following naming convention:
#
# source file = {id} + "-main." + {extension}
#
# input file = {id} + "-input.txt"
#
# output file = {id} + "-output.txt"

# compile the java file
a = 1
flag = 0
while [-e "$3/$1-input/input-$a.txt" ]
do
    touch $1-code-output.txt
    javac $3/$1_main.$2

    res = $?

    if [ $? != 0 ]; then
        echo "compile failed"
        flag = 1
        cleanup $1 $3
        exit
    fi

    cd $3
    # run the code and trap the output
    java $1_main < $1-input-$a.txt > ../$1-code-output.txt

    if [ $? != 0 ]; then
        echo "run failed"
        flag = 1
        cleanup $1 $3
        exit
    fi

    cd ..

    # Check if output matches
    diff --strip-trailing-cr $1-code-output.txt $3/$1-output-$a.txt > $1-diff-messages.txt
    if [ $? != 0 ]; then
        echo "wrong output"
        flag = 1
        cleanup $1 $3
        exit
    fi

    cleanup $1 $3
    a=$((a+1))
done



if [ $flag -eq 0 ]; then
    echo "successfully executed"
fi
#!/bin/bash

function cleanup() {
    if [ -f "$1-code-output.txt" ]; then
        rm $1-code-output.txt
    fi
    if [ -f "$1-diff-messages.txt" ]; then
        rm $1-diff-messages.txt
    fi
    # additional cleanup for class files in java
    if [ -f "Solution.class" ];then
        rm Solution.class
    fi
}

# Input format
# $1 -> id of submission
# $2 -> extension of code file
# $3 -> bind mounted directory
# $4 -> time limit 
# $5 -> filename 

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
javac $3/uploads/$5

if [ $? != 0 ]; then
    echo "compile failed"
    cleanup $1 $3
    exit
fi

a=1
flag=0
cd $3/uploads
while [ -e "../tests/$1-input/input/input-$a.txt" ]
do
    touch $1-code-output.txt

    # Execute and trap output
    java Solution < ../tests/$1-input/input/input-$a.txt > $1-code-output.txt 

    res=$?

    if [ $res -eq 124 ]; then 
        echo "TLE on test $a"
        cleanup $1
        flag=1
        exit
    elif [ $res -eq 137 ]; then
        echo "MLE on test $a"
        cleanup $1
        flag=1
        exit
    elif [ $res != 0 ]; then
        echo "Run time error on test $a", $res
        cleanup $1
        flag=1
        exit
    fi

    # Check if output matches
    diff --strip-trailing-cr $1-code-output.txt ../tests/$1-output/output/output-$a.txt > $1-diff-messages.txt
    if [ $? != 0 ]; then
        echo "WA on test case $a"
        cleanup $1
        flag=1
        exit
    fi

    rm $1-code-output.txt
    rm $1-diff-messages.txt

    a=$((a+1))
done

cleanup $1

if [ $flag -eq 0 ]; then
    echo "AC"
fi
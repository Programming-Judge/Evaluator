#!/bin/sh

function cleanup(){
    if [ -f "$1-code-output.txt" ]; then
        rm $1-code-output.txt
    fi
    if [ -f "$1-diff-messages.txt" ]; then
        rm $1-diff-messages.txt
    fi
    if [ -f "$1-main.out" ]; then
        rm $1-main.out
    fi
}
 
# $1 - id
# $2 - extension
# $3 - bind_mnt_dir
# $4 - timelimit

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

touch $1-code-output.txt

# compile the c code 
gcc $3/$1-main.$2 -o $1-main.out

if [ $? != 0 ]; then
    echo "compile failed"
    cleanup $1
    exit
fi

./$1-main.out < $3/$1-input.txt > $1-code-output.txt

if [ $? != 0 ]; then
    echo "run falied"
    cleanup $1
    exit
fi

diff -w $1-code-output.txt $3/$1-output.txt > $1-diff-messages.txt

if [ $? != 0 ]; then
    echo "wrong output"
    cleanup $1
    exit
fi

cleanup $1
echo "sucessfully executed"

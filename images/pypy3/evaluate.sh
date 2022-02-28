#!/bin/bash

function cleanup() {
    if [-f "$1-code-output.txt"]; then
        rm $1-code-output.txt
    fi
    if [-f "$1-diff-messages.txt"]; then
        rm $1-diff-messages.txt
    fi
}

# Input format
# $1 -> id of submission
# $2 -> extension of code file
# $3 -> bind mounted directory
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

ls -lR
# Execute and trap output
pypy3 $3/$1-main.$2 < $3/$1-input.txt &> $1-code-output.txt

if [ $? != 0 ]; then
    echo "run failed"
    cleanup $1
    exit
fi

# Check if output matches
diff $1-code-output.txt $3/$1-output.txt > $1-diff-messages.txt
if [ $? != 0 ]; then
    echo "wrong output"
    cleanup $1
    exit
fi

cleanup $1
echo "successfully executed"
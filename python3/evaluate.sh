#!/bin/bash

touch code_output.txt

# Execute and trap output
python3 $1 < $2 > code_output.txt

if [ $? != 0 ]; then
    echo "run failed"
    exit
fi

# Check if output matches
diff code_output.txt $3 > diff_messages.txt
if [ $? != 0 ]; then
    echo "wrong output"
    exit
fi

echo "successfully executed"
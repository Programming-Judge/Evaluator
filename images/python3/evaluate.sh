a=1
flag=0
while [ -e "$5/$7/$3/$a-input.txt" ]
do
    touch "$1-code-output.txt"
    timeout $4 python3 $5/$6/$1.$2 < "$5/$7/$3/$a-input.txt" > "$1-code-output.txt"
    #cat "$1-code-output.txt"
    res=$?
    if [ $res -eq 124 ]; then
        echo "Time limit exceeded on test $a"
        flag=1
        exit
    elif [ $res -eq 137 ]; then
        echo "Memory limit exceeded on test $a"
        flag=1
        exit
    elif [ $res != 0 ]; then
        echo "Runtime error on test case $a", $res
        flag=1
        exit
    fi

    diff "$1-code-output.txt" "$5/$7/$3/$a-output.txt" > "$1-diff-messages.txt"
    if [ $? != 0 ]; then
        echo "Wrong answer on test case $a"
        flag=1
        exit
    fi
    a=$((a+1))
done

if [ $flag -eq 0 ]; then
    echo "Accepted"
fi


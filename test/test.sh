#!/bin/bash

go run token_num.go > go_output.txt
python token_num.py > python_output.txt

diff -q go_output.txt python_output.txt > /dev/null
result=$?

if [ $result -eq 0 ]; then
    echo "Success"
else
    echo "Failure"
fi

rm go_output.txt python_output.txt
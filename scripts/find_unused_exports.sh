#!/usr/bin/env bash

if [ "$#" -gt 1 ]; then
    echo "Too many arugments passed in, expected one"
		exit 1
fi

echo "Used Exports: "; 
echo ""; 
rg -e "$1\.[A-Z]" | sed "s/.*$1.//;s/(.*//" | uniq | sort; 
echo ""
echo "All functions:"; 
cat $1/* 2>/dev/null | rg "func [A-Z]" | sed "s/.*func //;s/(.*//;s/Test.*//" | uniq | sort; 
echo ""

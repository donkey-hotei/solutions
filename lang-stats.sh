#!/bin/bash
git ls-files | 
    awk -F . '{ print $NF }' | 
    sort | 
    uniq -c | 
    sort -n -r | 
    awk '{ print $2, $1 }'
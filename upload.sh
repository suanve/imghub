#!/bin/bash
data=""
for arg in $*; do
    data=$data' -F "file[]=@'$arg'"'
done
echo curl $data http://127.0.0.1:8081/upload|tee ~/log_com
echo curl $data http://127.0.0.1:8081/upload|bash 2>&1|tee ~/log
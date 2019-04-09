#!/bin/bash
count=`grep ".go" goloomreport | wc -l`
echo "Number of errors $count"
if [ $count -le 100 ]
then
  echo "Errors within threshold"
  exit 0
else
  echo "Errors have exceeded threshold limit" >&2
  exit 1
fi

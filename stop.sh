#!/bin/bash

pids=$(ps -ef | grep '[d]experteventlistener' | awk '{print $2}')
for pid in $pids; do
    kill $pid
done


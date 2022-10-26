#!/bin/bash

processID=$(pidof main)
kill -SIGINT $processID; echo "process $processID was killed"

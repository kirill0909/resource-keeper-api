#!/bin/bash

processID=$(pidof main)
kill $processID; echo "process $processID was killed"

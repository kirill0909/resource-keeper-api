#!/bin/bash


./.bin/main&

processID=$(pidof main)
echo "process $processID was successfully launched"

#!/bin/bash
cd source
//-- git pull --rebase
docker build -t localhost:5000/wacron:latest .
docker push localhost:5000/wacron:latest

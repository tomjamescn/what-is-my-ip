#!/bin/bash

docker run \
    --name=what-is-my-ip \
    -p 9999:9999 \
    --restart unless-stopped \
    -d tomjamescn/what-is-my-ip:0.0.4

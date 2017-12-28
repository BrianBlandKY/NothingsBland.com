#!/bin/sh

# Simple script to wipe all running\stopped docker processes and images
DOC_PS=$(docker ps -aq)

l=${#DOC_PS}

if [ $l != "0" ]
    then
        # if there are processes, kill them
        docker stop ${DOC_PS} && \
        docker rm ${DOC_PS}
fi

echo "Docker Wiped!"
#!/bin/sh
set -ue
echo "Waiting jenkins to launch on 3306..."

while ! timeout 1 bash -c "echo > /dev/tcp/db/3306"; do   
  sleep 1
done

air
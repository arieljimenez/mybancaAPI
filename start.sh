#!/bin/bash
echo '+------------------+'
echo '+  Runing the API  +'
echo '+------------------+'
dev_appserver.py ./appengine/app.yaml \
    --port=9000 \
    --host=0.0.0.0 \
    --admin_port=9001 \
    --admin_host=0.0.0.0 \
    --api_port 9002 \
    --skip_sdk_update_check \
    --enable_watching_go_path=False

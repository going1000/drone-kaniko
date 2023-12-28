#!/usr/bin/env bash

docker build -t xhj-prod-registry.cn-hangzhou.cr.aliyuncs.com/image-common/kaniko:1.19.1 .
docker push xhj-prod-registry.cn-hangzhou.cr.aliyuncs.com/image-common/kaniko:1.19.1

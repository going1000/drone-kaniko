FROM bitnami/kaniko:1.19.1
ENV KANIKO_VERSION=1.19.1
ADD release/linux/amd64/kaniko-docker /kaniko/
ENTRYPOINT ["/kaniko/kaniko-docker"]

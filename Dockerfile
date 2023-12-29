FROM gcr.io/kaniko-project/executor:v1.9.1
ENV KANIKO_VERSION=1.9.1
ADD release/linux/amd64/kaniko-docker /kaniko/
ENTRYPOINT ["/kaniko/kaniko-docker"]

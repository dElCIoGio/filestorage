FROM ubuntu:latest
LABEL authors="dagos"

ENTRYPOINT ["top", "-b"]
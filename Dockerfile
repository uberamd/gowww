FROM scratch
LABEL maintainer "uberamd@gmail.com"

COPY ./bin/gowww /
EXPOSE 8080

USER 1001

ENTRYPOINT ["/gowww"]
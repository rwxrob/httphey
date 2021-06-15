FROM ubuntu
COPY httphey /usr/local/bin/
USER nobody
ENTRYPOINT ["httphey"]

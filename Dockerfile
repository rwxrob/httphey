FROM ubuntu
ADD https://raw.githubusercontent.com/rwxrob/httphey/main/httphey /usr/local/bin/
CMD ["httphey","9001","9002"]

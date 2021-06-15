# "Hey There" Minimal HTTP Responder (Server)

Useful little server that only sends "Hey there" responses tagged with
the requested ports passed as optional arguments (default: 8000, 8080)
created originally to reliably debug cloud-native plumbing (Istio
Service and Virtual Service specifically). 

## Install

```
go install github.com/rwxrob/httphey@latest
```

Or pull down the binary if you trust your network:

```
curl -L https://raw.githubusercontent.com/rwxrob/httphey/main/httphey
```

## Usage

```
$ httphey
Listening on 8000
Listening on 8080
```

Or

```
$ httphey 9001 9002
Listening on 9001
Listening on 9002
```

Or 

```dockerfile
FROM ubuntu
ADD https://raw.githubusercontent.com/rwxrob/httphey/main/httphey /usr/local/bin/
CMD ["httphey","9001","9002"]
```

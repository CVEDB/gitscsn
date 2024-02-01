FROM alpine:3.19.1
RUN apk --no-cache add ca-certificates git
COPY gitscan /usr/local/bin/gitscan
COPY contrib/*.tpl contrib/
ENTRYPOINT ["gitscan"]

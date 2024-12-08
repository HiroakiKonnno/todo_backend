FROM alpine:3.20.3 AS base

RUN apk add --no-cache ca-certificates curl
RUN curl -kLSO https://github.com/k0kubun/sqldef/releases/download/v0.17.19/mysqldef_linux_amd64.tar.gz && \
    tar xzf mysqldef_linux_amd64.tar.gz

FROM busybox

WORKDIR /sqldef
ADD schema.sql /sqldef/schema.sql
ADD run.sh /sqldef/run.sh
COPY --from=base /mysqldef /usr/bin/mysqldef

RUN chmod +x /sqldef/run.sh

CMD [ "sh", "/sqldef/run.sh" ]

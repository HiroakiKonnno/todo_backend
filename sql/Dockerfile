FROM alpine:3.20.3 AS base

RUN apk add --no-cache ca-certificates curl
RUN curl -LO https://github.com/sqldef/sqldef/releases/download/v0.16.13/psqldef_linux_amd64.tar.gz && \
    tar xf psqldef_linux_amd64.tar.gz

FROM busybox

WORKDIR /sqldef
ADD schema.sql /sqldef/schema.sql
ADD run.sh /sqldef/run.sh
COPY --from=base /psqldef /usr/bin/postgresdef

RUN chmod +x /sqldef/run.sh

CMD [ "sh", "/sqldef/run.sh" ]

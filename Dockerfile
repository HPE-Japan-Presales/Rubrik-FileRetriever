# Global Values
ARG ProjectDir=Rubrik-FileRetriever

#####
# Create build container
#####

FROM golang:1.14.0-buster as builder
MAINTAINER Tak

ARG GitBranch=master
ARG GitURL=https://github.com/fideltak/Rubrik-FileRetriever.git
ARG ProjectDir

WORKDIR /go/src/github.com
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

RUN git clone -b ${GitBranch} ${GitURL}
RUN cd ${ProjectDir} && go build -o main

#####
# Create target container
#####
FROM alpine

ARG ProjectDir

WORKDIR /${ProjectDir}
ENV CHILINFO_APP=$BuildTarget
RUN apk --update add tzdata && \
    ls /usr/share/zoneinfo && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    echo "Asia/Tokyo" >  /etc/timezone &&\
    date && \
#    apk del tzdata && \
    rm -rf /var/cache/apk/*
COPY --from=builder /go/src/github.com/${ProjectDir}/main .
CMD ["/Rubrik-FileRetriever/main"]

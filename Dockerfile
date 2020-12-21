FROM golang:1.14.3-alpine AS build
RUN apk add git; go get github.com/DowerX/covid-graph;cd /go/src/github.com/DowerX/covid-graph;CGO_ENABLED=0 go build; mkdir /static

FROM scratch AS bin
COPY --from=build /go/src/github.com/DowerX/covid-graph/covid-graph /
COPY --from=build /static /static
ENTRYPOINT [ "/covid-graph" ]
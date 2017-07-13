FROM golang:1.8


RUN mkdir -p /go/src/go-grafana-monitoringAPI/
WORKDIR /go/src/go-grafana-monitoringAPI/

COPY . /go/src/go-grafana-monitoringAPI/
RUN go-wrapper download && go-wrapper install

CMD ["go-wrapper", "run"]

EXPOSE 8082
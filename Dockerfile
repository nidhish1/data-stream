FROM golang as builder
ARG MODULE

RUN git clone https://github.com/edenhill/librdkafka.git
WORKDIR librdkafka
RUN ./configure --prefix /usr
RUN make
RUN make install


WORKDIR .
RUN go get gopkg.in/confluentinc/confluent-kafka-go.v1/kafka

COPY main.go .
CMD ["go", "run", "main.go"]


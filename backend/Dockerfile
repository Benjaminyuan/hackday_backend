FROM golang:1.13.1

COPY . /app

WORKDIR  /app

RUN mkdir out

RUN go build -o out/hackday

CMD [ "./out/hackday" ]
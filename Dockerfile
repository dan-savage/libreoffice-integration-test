FROM golang:1.15 as builder

WORKDIR /go/src/app
COPY . .

ENV CGO_ENABLED=0
RUN apt-get clean
RUN apt-get update
RUN apt-get install -y libssl-dev openssl libreoffice libreofficekit-dev

RUN go build -o /github_action

EXPOSE 8080

CMD [ "/github_action" ]
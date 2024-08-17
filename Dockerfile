FROM golang:1.22.5-bullseye as development

WORKDIR /app
COPY go.sum go.mod /app/

RUN go mod download
COPY . /app

ENV TZ=Europe/Istanbul

#RUN CGO_ENABLED=0 go build -o /bin/app ./cmd/
CMD ["go", "run", "./cmd/main.go"]
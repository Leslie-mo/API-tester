FROM golang:1.16.7 as builder
RUN mkdir /creation_omnial_simulator
WORKDIR /creation_omnial_simulator
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go generate
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o OmniSimulatorApp

FROM alpine:3.16.2
RUN apk add --no-cache tzdata zip unzip sudo curl
ENV TZ=Asia/Tokyo
WORKDIR /usr/local/bin
COPY env/all.env ./env/all.env
COPY --from=builder /creation_omnial_simulator/OmniSimulatorApp /usr/local/bin/OmniSimulatorApp

CMD ["/usr/local/bin/OmniSimulatorApp"]

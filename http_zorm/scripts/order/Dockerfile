FROM golang:1.18 as builder
ARG GOPROXY="https://goproxy.cn"
WORKDIR /app
ADD . /app
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN GOPROXY=$GOPROXY go mod download
RUN make build-order-svc

FROM alpine:latest
WORKDIR /
COPY --from=builder /app/order_svc/OrderSvc /OrderSvc
CMD ["/OrderSvc"]

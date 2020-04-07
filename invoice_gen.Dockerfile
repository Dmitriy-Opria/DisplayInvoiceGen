# Build executable.
# docker build --build-arg GITHUB_RAKOPS={your_token} -t invoice_gen -f invoice_gen.Dockerfile .
# docker run -it -d -p 9065:9065 --name invoice_gen invoice_gen:latest
# docker run -it -d -p 9065:9065 --name invoice_gen repo.rmgops.com/docker/display_invoice/invoice_gen:develop
FROM golang:1.14 AS build-env

ARG GITHUB_RAKOPS
ARG PG_USER
ARG PG_PASS
ARG PG_ADDR
ARG PG_NAME
ARG AUTH_TOKEN
ENV GO111MODULE=on

RUN git config --global url."https://rpx:${GITHUB_RAKOPS}@github.rakops.com/rpx/rules.git".insteadOf "https://github.rakops.com/rpx/rules.git"  && \
git config --global url."https://BNP:${GITHUB_RAKOPS}@github.rakops.com/BNP/DisplayInvoiceGen.git".insteadOf "https://github.rakops.com/BNP/DisplayInvoiceGen.git"

WORKDIR /src
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
WORKDIR /src

RUN GOGC=200 go build \
-ldflags "\
-X main.user=$PG_USER \
-X main.pass=$PG_PASS \
-X main.addr=$PG_ADDR \
-X main.name=$PG_NAME \
-X main.token=$AUTH_TOKEN" \
-o invoice_gen

# final stage
FROM golang:1.14
WORKDIR /src
COPY --from=build-env /src/config.yml .
COPY --from=build-env /src/invoice_gen .
ENTRYPOINT ["./invoice_gen"]
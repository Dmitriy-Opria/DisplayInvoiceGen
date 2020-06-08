# Build executable.
# docker build --build-arg PROD=true --build-arg GITHUB_RAKOPS={your_token} --build-arg ARTIFACTORY_APIKEY={your artifactory key} -t invoice_gen -f invoice_gen.Dockerfile .
# docker run -it -d -p 9065:9065 --name invoice_gen invoice_gen:latest
# docker run -it -d -p 9065:9065 --name invoice_gen repo.rmgops.com/docker/display_invoice/invoice_gen:develop
FROM golang:1.14 AS build-env
ARG GITHUB_RAKOPS
ARG ARTIFACTORY_APIKEY
# default argument when not provided in the --build-arg
ARG PROD

ENV GO111MODULE=on
RUN git config --global url."https://rpx:${GITHUB_RAKOPS}@github.rakops.com/rpx/rules.git".insteadOf "https://github.rakops.com/rpx/rules.git"  && \
git config --global url."https://BNP:${GITHUB_RAKOPS}@github.rakops.com/BNP/DisplayInvoiceGen.git".insteadOf "https://github.rakops.com/BNP/DisplayInvoiceGen.git"
WORKDIR /src

# download config
RUN apt-get update && apt-get -y install curl
RUN if [ "$PROD" = "true" ]; \
then curl -H X-JFrog-Art-Api:$ARTIFACTORY_APIKEY -o config.yaml -O "https://repo.rmgops.com:443/artifactory/vagrant-local/invoice_gen/invoice_gen.prod.yaml";\
else curl -H X-JFrog-Art-Api:$ARTIFACTORY_APIKEY -o config.yaml -O "https://repo.rmgops.com:443/artifactory/vagrant-local/invoice_gen/invoice_gen.qa.yaml";\
fi

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
WORKDIR /src

RUN GOGC=200 go build -o invoice_gen

# final stage
FROM golang:1.14
WORKDIR /src
COPY --from=build-env /src/config.yaml .
COPY --from=build-env /src/invoice_gen .
ENTRYPOINT ["./invoice_gen"]
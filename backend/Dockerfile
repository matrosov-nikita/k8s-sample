# build stage
FROM golang:alpine AS build-env
ADD . /src/
WORKDIR /src
RUN go build -o test-backend .

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/test-backend /app/
ENTRYPOINT ./test-backend
# Start from the official Golang base image
FROM golang:1.23-alpine
#ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# create workplace
WORKDIR /app
# copy go.mod file
COPY go.mod .
#copy go.sum file
COPY go.sum .
# download module
RUN go mod download
# copy source code
COPY . .
# run go build command 
RUN go build -o app .
EXPOSE 8080
# run app file
CMD [ "./app" ]

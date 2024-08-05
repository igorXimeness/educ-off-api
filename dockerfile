FROM golang:1.22-alpine 

WORKDIR /app 

COPY go.mod ./

COPY go.sum .

RUN go mod download 

COPY . .

RUN go build -o main ./cmd/educ-off-api

EXPOSE 8080

CMD ["./app/cmd/main"]




# [ alpine(linux) = [golang, pasta(app)   ]  ]
# cd ..  



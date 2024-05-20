FROM golang 
COPY main.go main.go
COPY vendor vendor
COPY go.sum go.sum

CMD [ "go", "run", "." ]
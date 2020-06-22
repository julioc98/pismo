FROM golang
WORKDIR /pismo
COPY . .
CMD [ "make","run/api"]
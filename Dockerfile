FROM golang
WORKDIR /pismo
COPY . .
RUN go mod vendor
ENV PORT 5001
EXPOSE 5001
CMD [ "make","run/api"]
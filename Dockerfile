FROM golang:1.8
WORKDIR /go/src/app
COPY . .
RUN echo "[url \"https://64e21058a79b7c3e97aad945314a28502af64dae:x-oauth-basic@github.com/\"]\n\tinsteadOf = https://github.com/" >> /root/.gitconfig
RUN go-wrapper download
RUN go-wrapper install
EXPOSE 3000
CMD ["go-wrapper", "run"]

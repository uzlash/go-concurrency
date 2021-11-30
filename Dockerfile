FROM alpine
WORKDIR app
COPY ./squareApp .
EXPOSE 3000
CMD ./squareApp
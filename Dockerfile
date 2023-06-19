FROM node:18.4.0 as ui
ARG NAME
ARG VERSION
WORKDIR /ui_build
ADD ui .
RUN npm install && npm run build

FROM golang:1.18 as build
ARG NAME
ARG VERSION
ARG COMMIT
ARG BUILD_TIME
ARG MAIN_PATH
ARG BASE="github.com/lishimeng/app-starter/version"
ENV GOPROXY=https://goproxy.cn,direct
ARG LDFLAGS=" \
    -X ${BASE}.AppName=${NAME} \
    -X ${BASE}.Version=${VERSION} \
    -X ${BASE}.Commit=${COMMIT} \
    -X ${BASE}.Build=${BUILD_TIME} \
    "
WORKDIR /release
ADD . .
COPY --from=ui /ui_build/dist/ static/

RUN go mod download && go mod verify
RUN go build -v --ldflags "${LDFLAGS} -X ${BASE}.Compiler=$(go version | sed 's/[ ][ ]*/_/g')" -o ${NAME} ${MAIN_PATH}

FROM lishimeng/alpine:3.17 as prod
ARG NAME
EXPOSE 80/tcp
WORKDIR /
COPY --from=build /release/${NAME} /
RUN ln -s /${NAME} /app
CMD [ "/app"]

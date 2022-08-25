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
ARG MAIN_PATH
ENV GOPROXY=https://goproxy.cn,direct
ENV LDFLAGS=" \
    -X 'github.com/lishimeng/app-starter/version.AppName=${NAME}' \
    -X 'github.com/lishimeng/app-starter/version.Version=${VERSION}' \
    -X 'github.com/lishimeng/app-starter/version.Commit=${COMMIT}' \
    -X 'github.com/lishimeng/app-starter/version.Build=`date +%FT%T%z`' \
    -X 'github.com/lishimeng/app-starter/version.Compiler=`go version`' \
    "
WORKDIR /release
ADD . .
COPY --from=ui /ui_build/dist/ static/
RUN go mod download && go mod verify
RUN go build -v --ldflags "${LDFLAGS}" -o ${NAME} ${MAIN_PATH}

FROM ubuntu:22.04 as prod
ARG NAME
EXPOSE 80/tcp
WORKDIR /
COPY --from=build /release/${NAME} /
RUN ln -s /${NAME} /app
CMD [ "/app"]
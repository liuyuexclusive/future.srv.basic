FROM alpine
ADD future.srv.basic /future.srv.basic
ADD appconfig.json /appconfig.json
ENTRYPOINT [ "/future.srv.basic"]

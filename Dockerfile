FROM alpine
ADD future.srv.basic /future.srv.basic
ADD appconfig.yml /appconfig.yml
ENTRYPOINT [ "/future.srv.basic"]

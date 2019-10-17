FROM alpine
ADD srv-basic /srv-basic
ADD appconfig.json /appconfig.json
ENTRYPOINT [ "/srv-basic"]

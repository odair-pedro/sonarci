FROM alpine:3.18.3

RUN wget https://github.com/odair-pedro/sonarci/releases/latest/download/sonarci-linux-x64.tar.gz &&\
tar -xf sonarci-linux-x64.tar.gz sonarci &&\
rm sonarci-linux-x64.tar.gz &&\
mv ./sonarci /usr/local/bin/sonarci &&\
chmod +x /usr/local/bin/sonarci

ENTRYPOINT ["/bin/sh", "-c", "sonarci $@"]

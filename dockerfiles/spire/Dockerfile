FROM debian

RUN useradd workload

RUN apt-get update && apt-get install -y zip curl

RUN curl -L https://github.com/spiffe/spire/releases/download/0.4/spire-0.4-linux-x86_64-glibc.tar.gz -o spire.tar.gz; \
    tar zvxf spire.tar.gz; \
    cp -r spire-0.4/. /opt/spire/; \
    chmod -R 777 /opt/spire/; \
    ln -s /opt/spire/spire-server /usr/bin/spire-server; \
    ln -s /opt/spire/spire-agent /usr/bin/spire-agent

RUN mkdir /.data

COPY spire_server.conf /opt/spire/conf/server/server.conf
COPY spire_agent.conf /opt/spire/conf/agent/agent.conf
COPY startup.sh /startup.sh

CMD /startup.sh 

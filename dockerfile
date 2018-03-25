FROM myregistry.com/public/centos

ADD ./logmaker /usr/local/bin
ENV MY_NAME=bruce
ENV HOST_NAME LOGMAKER

# comments
#RUN yum install golang -y

CMD "logmaker"

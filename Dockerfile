FROM scratch
MAINTAINER Sope Shen "shenshouer51@gmail.com"
EXPOSE 30001

COPY simple_request /

CMD ["/simple_request"]
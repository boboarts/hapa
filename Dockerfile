# for Qgen
# Version 0.0.1
FROM centos
Maintainer Gary SHEN "boboarts@gamil.com"
RUN mkdir -p /etc/qgen
WORKDIR /etc/qgen
RUN curl -O https://raw.githubusercontent.com/boboarts/qgen/master/bin/centos/qgen
RUN curl -O https://raw.githubusercontent.com/boboarts/qgen/master/bin/centos/template.html
RUN chmod 755 qgen
CMD ["/etc/qgen/qgen"]
EXPOSE 9923
# docker build -t "boboarts/qgen" .
# docker run -d -p 9923:9923 --name qgen boboarts/qgen
# docker run -d -p 9923:9923 --name qgen boboarts/qgen /etc/qgen/qgen
# docker run -it --name qgen boboarts/qgen /bin/bash
# docker exec -it qgen /bin/bash
# docker commit qgen boboarts/qgen
# curl localhost:9923





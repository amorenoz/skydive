ARG  BASE=fedora:30
FROM $BASE
ARG  ARCH=amd64
RUN dnf -y install  openvswitch libpcap libxml2 libvirt
COPY skydive.$ARCH /usr/bin/skydive
COPY skydive.yml /etc/skydive.yml
ENTRYPOINT ["/usr/bin/skydive", "--conf", "/etc/skydive.yml"]

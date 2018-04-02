FROM base/archlinux
ADD ./test-client /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/test-client"]

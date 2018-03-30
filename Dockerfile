FROM base/archlinux
ADD ./postgresql-test /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/postgresql-test"]

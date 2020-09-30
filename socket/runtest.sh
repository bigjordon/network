#!/bin/sh

#set -v on
GOTESTFILE=socket_test.go
BINFILE=socket.test
if test -e ${BINFILE}; then
	rm -f ${BINFILE}
fi

go test -c ${GOTESTFILE} -o ${BINFILE}

sudo ./${BINFILE} -test.v

#!/bin/bash
#
# Run the specified command, saving its output. Output is echoed when
# the command is done iff the command exited non-zero.

PROGNAME=$(basename $0)

if [[ $# -eq 0 ]] ; then
    echo "Usage: $PROGNAME cmd [args...]" >&2
    exit 2
fi

TMPFILE=$(mktemp)
function finish {
    rm -fr ${TMPFILE}
}
trap finish EXIT

EXECNAME=$1
shift

${EXECNAME} "$@" >${TMPFILE} 2>&1
if [[ $? -ne 0 ]] ; then
    cat ${TMPFILE}
fi

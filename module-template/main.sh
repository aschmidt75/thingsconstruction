#!/bin/ash

OUT="/out"
INJSON="/tmp/indata"
cat - >${INJSON}

# if input from STDIN is empty, display a default json container
# with metadata about this container
if [[ ! -s ${INJSON} ]]; then
    echo '{ "status": "info", "msg": "Sample THNGS:CONSTR generator module"}' | jq -M .
    exit 0
fi

cat ${INJSON} | jq -e . >/dev/null 2>&1
if [[ $? -eq 0 ]]; then
    # valid JSON as STDIN input, process. Here: default, dummy processing
    # place two static files in output folder, return metadata about
    # these files.
    touch ${OUT}/readme.md
    touch ${OUT}/sample.ino
    cat ${INJSON} | jq . | sed -e 's/^/\/\/ /g' >${OUT}/sample.ino

    echo '{
        "status": "ok",
        "msg": "Generated sample output",
        "files": [
            {
                "filename": "readme.md",
                "desc": "What to do with generated files",
                "ct": "text/markdown",
                "type": "doc"
            },
            {
                "filename": "sample.ino",
                "desc": "Arduino sketch",
                "ct": "text/plain",
                "type": "source",
                "language": "c"
            }
        ]
    }' | jq -M .
else
    # input from STDIN is not JSON, give back error message.
    echo '{ "status": "error", "msg": "Unable to parse input, must conform to JSON"}' | jq -M .
fi


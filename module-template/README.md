## Module Template

Things:Construction modules are docker containers. The Main app
communicates with containers over STDIN/STDOUT and files. STDIN/STDOUT is the
default way to exchange meta data, i.e. 

* what files have been produced
* generation status or error codes
* module meta information

Input files are read from `/in`, result files are written to `/out`.
Calling apps need to mount volumes to these points to be able to 
exchange data with modules.

### Sample walkthrough

This sample Dockerfile builds upon an alpine base image, adding
`jq` and `main.sh` as the main entrypoint.

```bash
$ docker build -t samplemod:0.1 .
```

Running this without any arguments returns meta data about the
module.

```bash
$ docker run samplemod:0.1
{
  "status": "info",
  "msg": "Sample THNGS:CONSTR generator module"
}
```

When supplying valid JSON meta data via STDIN, make sure to mount both `/in` and `/out` dirs:

```bash

$ mkdir /tmp/tc-in /tmp/tc-out

$ echo '{ "j": "son" }' | docker run -i -v /tmp/tc-in:/in -v /tmp/tc-out:/out samplemod:0.1
{
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
}

$ ls /tmp/tc-out
readme.md  sample.ino
```

Result metadata is written to STDOUT, stating that 2 files have been writte, `readme.md` and `sample.ino`.

In case of errors, no files are produces (or parts may be missing), and meta data containing
error details is written to STDOUT:

```bash
$ echo 'this-no-json' | docker run -i samplemod:0.1      
{
"status": "error",
"msg": "Unable to parse input, must conform to JSON"
}
```
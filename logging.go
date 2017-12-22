//    ThingsConstruction, a code generator for WoT-based models
//    Copyright (C) 2017  @aschmidt75
//
//    This program is free software: you can redistribute it and/or modify
//    it under the terms of the GNU Affero General Public License as published
//    by the Free Software Foundation, either version 3 of the License, or
//    (at your option) any later version.
//
//    This program is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//    GNU Affero General Public License for more details.
//
//    You should have received a copy of the GNU Affero General Public License
//    along with this program.  If not, see <http://www.gnu.org/licenses/>.
//
package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"syscall"
)

func writerFromLoggingTargetString(in string) io.Writer {
	if in == "" {
		return ioutil.Discard
	}
	if in == "stdout" {
		return os.NewFile(uintptr(syscall.Stdout), "/dev/stdout")
	}
	if in == "stderr" {
		return os.NewFile(uintptr(syscall.Stderr), "/dev/stderr")
	}
	f, err := os.Open(in)
	if err != nil {
		panic(err)
	}
	return bufio.NewWriter(f)
}

func InitializeBasicLogging() (debug *log.Logger, verbose *log.Logger, error *log.Logger) {
	debug = log.New(ioutil.Discard, "", 0)
	verbose = log.New(ioutil.Discard, "", 0)
	error = log.New(os.Stderr, "ERROR ", log.Lshortfile)
	return debug, verbose, error
}

func InitializeLogging(c *Config) (debug *log.Logger, verbose *log.Logger, error *log.Logger) {
	debug = log.New(writerFromLoggingTargetString(c.Logging.Debug), "DEBUG ", log.Lshortfile)
	verbose = log.New(writerFromLoggingTargetString(c.Logging.Verbose), "INFO ", 0)
	error = log.New(writerFromLoggingTargetString(c.Logging.Error), "ERROR ", 0)
	return debug, verbose, error
}

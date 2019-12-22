#!/bin/bash

CURDIR=`pwd`

export GOPATH=$GOPATH:${CURDIR}/v1:${CURDIR}/v2
echo ${GOPATH}
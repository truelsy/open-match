#!/bin/bash

make REGISTRY=172.16.1.13:5000/gv-ns-studio ALTERNATE_TAG=1.1.0 build-nbanow-image
make REGISTRY=172.16.1.13:5000/gv-ns-studio ALTERNATE_TAG=1.1.0 push-nbanow-image

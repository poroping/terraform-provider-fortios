#!/bin/bash
set -e

mv -f ./output/*.go ../fortios/
mv -f ./output/d/* ../website/docs/d/
mv -f ./output/r/* ../website/docs/r/
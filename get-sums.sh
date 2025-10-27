#!/bin/bash

# Check if debsums is installed
if ! dpkg -s debsums &> /dev/null; then
        # Install debsums from Debian site
        apt install wget
        wget http://ftp.us.debian.org/debian/pool/main/d/debsums/debsums_3.0.2.3_all.deb
        dpkg -i debsums_3.0.2.3_all.deb
fi

# Check sums of binaries
debsums | awk '$NF != "OK"'

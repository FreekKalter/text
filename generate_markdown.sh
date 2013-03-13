#!/bin/bash
go install github.com/FreekKalter/text/columnswriter
godoc-md github.com/FreekKalter/text/columnswriter > README.md

# Add license info
echo "*Copyright (c) 2013 Freek Kalter.  All rights reserved.
See the LICENSE file.*" >> README.md

#!/bin/bash
go install github.com/FreekKalter/text/columnswriter
godoc -ex=true -templates=/home/fkalter/godoc github.com/FreekKalter/text/columnswriter > README.md

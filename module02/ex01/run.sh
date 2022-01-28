#!/bin/bash
go build
./download https://github.com/42School/norminette/raw/master/pdf/en.norm.pdf
open norm.pdf

#!/bin/bash

# color
BLACK="\033[30m"
RED="\033[31m"
GREEN="\033[32m"
YELLOW="\033[33m"
BLUE="\033[34m"
MAGENTA="\033[35m"
CYAN="\033[36m"
WHITE="\033[37m"
BOLD="\033[1m"
UNDERLINE="\033[4m"
BOLD_UNDERLINE="\033[1;4m"
RESET="\033[0m"

# func for test
check () {
	if [ $1 -eq 0 ]; then
		printf ${GREEN}
		echo ok
		printf ${RESET}
	else
		printf ${RED}
		echo ko
		printf ${RESET}
	fi
}

test () {
	diff $1 $2
	check $?
}

# test 1
ORIROOT=images
ANSROOT=answer
SUBDIR=subdir
ORISUB=$ORIROOT/$SUBDIR
ANSSUB=$ANSROOT/$SUBDIR
FILE1=Icon.png
FILE2=IconNoExtension.png

./convert images
test $ORIROOT/$FILE1 $ANSROOT/$FILE1
test $ORIROOT/$FILE2 $ANSROOT/$FILE2
test $ORISUB/$FILE1 $ANSSUB/$FILE1
test $ORISUB/$FILE2 $ANSSUB/$FILE2

# test 2
ORIROOT=images_contain_no_permission

./convert images_contain_no_permission
test $ORIROOT/$FILE1 $ANSROOT/$FILE1
test $ORIROOT/$FILE2 $ANSROOT/$FILE2
test $ORISUB/$FILE1 $ANSSUB/$FILE1
test $ORISUB/$FILE2 $ANSSUB/$FILE2

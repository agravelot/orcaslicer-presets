#!/bin/bash

set -e

# exit 1 if syumbolic link exist
if [[ -L ~/.config/OrcaSlicer/user ]]; then
	echo "user symbolic link already exist"
	exit 1
fi

if [[ -d ~/.config/OrcaSlicer/user ]]; then
	echo "user profile already exist"
	echo "backup old user profile"
	mv -f ~/.config/OrcaSlicer/user ~/.config/OrcaSlicer/user.bak
fi

ln -s $(pwd)/user ~/.config/OrcaSlicer/user

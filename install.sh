#!/bin/bash

COL_GREEN="\033[1;32m"
COL_RED="\033[1;31m"
COL_YELLOW="\033[1;33m"
COL_NORMAL="\033[0;00m"
CMD_FOUND=0
CMD_NOT_FOUND=127

show_banner(){
	echo -e ""
    echo -e "$COL_GREEN  ____ $COL_NORMAL   _   ____ _____ _____ ____        "
    echo -e "$COL_GREEN |  _ \ $COL_NORMAL/ _ \|  _ \_   _| ____|  _ \       "
    echo -e "$COL_GREEN | |_) |$COL_NORMAL | | | |_) || | |  _| | |_) |      "
    echo -e "$COL_GREEN |  __/|$COL_NORMAL |_| |  _ < | | | |___|  _ <       "
    echo -e "$COL_GREEN |_|   $COL_NORMAL \___/|_| \_\|_| |_____|_| \_\      "
}

check_user(){
    user=$(whoami)
    if [ ! $user == "root" ]
    then
        echo -e $COL_RED"permission denied to install..."
        echo -e $COL_NORMAL"run again the script as root to install"
        exit 1
    fi
}

start_installer() {
    porter_path="/usr/share/porter"
    ports_json="ports.json"
    porter_build="go.mod"
    porter_exe="porter"
    sys_bin_path="/bin"
    usr_bin_path="/usr/bin"
    echo -e $COL_GREEN"checking$COL_NORMAL $porter_path"
    sleep 1
    if [ ! -d $porter_path ]
    then
        echo -e $COL_GREEN"creating$COL_NORMAL $porter_path"
        mkdir $porter_path
    fi

    echo -e $COL_GREEN"checking$COL_NORMAL $ports_json..."
    sleep 1
    if [ ! -f $(pwd)"/"$ports_json ]
    then
        echo -e "$COL_RED""unable to find $ports_json "$COL_NORMAL
        echo -e "$COL_RED""terminating installation...$COL_NORMAL"
        exit 1
    else
        echo -e $COL_GREEN"found$COL_NORMAL $ports_json..."
        if [ ! -f $porter_path"/"$ports_json ]
        then
            echo -e $COL_GREEN"copying$COL_NORMAL ""$ports_json -> $porter_path"
            cp $(pwd)/$ports_json $porter_path
        fi
    fi

    echo -e $COL_GREEN"checking golang..."$COL_NORMAL
    sleep 1
    res=$(go env GOROOT)
    if [ $? -eq $CMD_NOT_FOUND ]
    then
        echo -e $COL_RED"go compiler not found..."$COL_NORMAL
        echo -e $COL_RED"install golang and run the install script again..."$COL_NORMAL
        exit 1
    fi

    echo -e $COL_GREEN"checking$COL_NORMAL build file..."
    sleep 1
    if [ ! -f "$(pwd)"/$porter_build ]
    then
        echo -e $COL_RED"build file not found..."$COL_NORMAL
        echo -e $COL_RED"terminating installation..."$COL_NORMAL
        exit 1
    fi

    echo -e $COL_GREEN"building script$COL_NORMAL"
    go build
    if [ $? -ne 0 ]
    then
        echo -e $COL_RED"build not finished..."$COL_NORMAL
        echo -e $COL_NORMAL"terminating installation..."
        exit 1
    fi

    if [ ! -f $(pwd)"/"$porter_exe ]
    then
        echo -e $COL_RED"porter executable not found..."
        exit 1
    else
        cp $(pwd)/$porter_exe $sys_bin_path"/"$porter_exe
        if [ $? -ne 0 ]
        then
            echo -e $COL_RED"failed copying porter to $sys_bin_path"$COL_NORMAL
            echo -e $COL_YELLOW"copy manually or add $porter_exe to path"$COL_NORMAL
            exit 1
        fi

    fi
    echo -e $COL_GREEN"done installing..."$COL_NORMAL
    exit 0
}

show_banner
check_user
start_installer

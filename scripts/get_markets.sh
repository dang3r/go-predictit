#!/bin/bash

curl -s https://www.predictit.org/Browse/Featured | pup 'ul[class="dropdown-menu"]' | grep 'Group'

#!/bin/bash

curl -X GET https://www.predictit.org/api/marketdata/ticker/RNOM16 | grep 'Group'

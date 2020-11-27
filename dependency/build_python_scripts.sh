#!/bin/bash

SCRIPT_PATH=$(readlink -f $(dirname $0))
cd $SCRIPT_PATH
pipenv run pyinstaller --onefile create_retention_policy.py
pipenv run pyinstaller --onefile nginx_form_conf.py
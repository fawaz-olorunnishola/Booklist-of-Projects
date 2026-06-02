#!/bin/bash
grep "^170:" superhero | awk -F: '{print $2"\n"$3"\n"$4}'
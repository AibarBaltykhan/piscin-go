find -name '*.sh' | cut -d '.' -f2 | sed 's/\///g' | sed 's/sh//g' | sed 's/test//g'


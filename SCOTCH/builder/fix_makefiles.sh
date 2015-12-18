#!/bin/bash
# The makefiles provided by scotch require some replacements in order to work under
# Windows. The biggest of these is that $(LIB) needs to be replaced with $(LIB_EXT) as
# the LIB environment variable is used to indiate the link-time library path.
# The replacement rules are defined in the file: makefile_replacement_rules

target=$1

rule_file=makefile_replacement_rules


# Replace $(LIB) with $(LIB_EXT)
list="`ls $target/*/Makefile` $target/Makefile"
for f in $list
do
  echo "sed -i -f $rule_file $f"
  sed -i -f $rule_file $f
done


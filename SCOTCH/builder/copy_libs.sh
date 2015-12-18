#!/bin/bash
# Copies the generated libraries to the required target, taking the COMMLIB into account.

src_dir=$1
target_dir=$2
commlib=$3
lib_ext=$4

# We need to append the MPI flavour to the parallel libraries.
# The fit the pattern: libpt*.${lib_ext}
parallel_libs=`ls ${src_dir}/libpt*${lib_ext}`
sequential_libs=`ls ${src_dir}/lib*${lib_ext} | grep -v -E "libpt.*${lib_ext}"`

# echo "parallel"
# echo ${parallel_libs}
# echo ""
# echo "seqential"
# echo ${sequential_libs}


for f in `ls ${src_dir}/libpt*${lib_ext}`
do
  # If a renaming of the library itself is required, then it can be performed here
  # target=${f/$lib_ext/\.$commlib$lib_ext}
  target=$target_dir/${commlib}${f/$src_dir/}
  echo "cp $f $target"
  cp $f $target
done

for f in `ls ${src_dir}/lib*${lib_ext} | grep -v -E "libpt.*${lib_ext}"`
do
  # If a renaming of the library itself is required, then it can be performed here
  # target=${f/\$lib_ext/\.$commlib\$lib_ext}
  target=$target_dir${f/$src_dir/}
  echo "cp $f $target"
  cp $f $target
done

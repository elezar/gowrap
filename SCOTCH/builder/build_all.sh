#!/bin/bash
# Usage:
#   bash ./build_all.sh TYPE [VERSION]
#
# Where TYPE is one of win64, win32, linux64, linux32 and VERSION is an optional
# version string.
#
# If a version is not specified, the script looks for an archive scotch_VERSION.tar.gz
# to determine the version.
#
# If the version is specified, then the archive scotch_VERSION.tar.gz is used, and
# must exist.

what=$1

if [ $# -eq 2 ]
then
# The version has been specifed.
  version=$2
  archive=scotch_$version.tar.gz
  echo "Using version: $version"
  echo "Using archive: $archive"
else
  archive=`ls scotch_*.tar.gz`

  if [[ x"$archive" == x"" ]]
  then
    echo "Source archive not found."
    exit 1
  fi

  echo "Found source archive: $archive"
  archive_base=${archive/.tar.gz/}
  version=${archive_base/scotch_/}
  echo "Detected version: $version"

fi

if [ ! -f $archive ]
then
  echo "Source archive ($archive) not found."
  exit 1
fi

if [[ x"$what" == x"win64" ]]
then
platform="win64_x86_64"
commlibs="MPICH MSMPI IMPI"
elif [[ x"$what" == x"win32" ]]
then
platform="win32_ia32"
commlibs="MPICH MSMPI IMPI"
elif [[ x"$what" == x"linux64" ]]
then
platform="linux64_x86_64"
commlibs="MPICH MPT IMPI"
elif [[ x"$what" == x"linux32" ]]
then
platform="linux32_ia32"
commlibs="MPICH IMPI"
else
echo "Undefined platform $what"
exit 1
fi


echo "Running build on $platform for $commlibs"
for c in $commlibs
do
make install PLAT_NAME=$platform COMMLIB=$c version=$version
done




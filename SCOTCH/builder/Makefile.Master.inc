
# This make include file specifies the required MAKE configuration for actually
# bulding SCOTCH. It is COPIED to the scotch_<VERSION>_<PLAT_NAME>_<COMMLIB>_source/src
# folder, with the definitions of PLAT_NAME and COMMLIB prepended.


IS_WINDOWS=0
ifeq (,$(findstring linux,$(PLAT_NAME)))
IS_WINDOWS=1
EXE		= .exe
LIB_EXT	= .lib
OBJ		= .obj
else
EXE		=
LIB_EXT	= .a
OBJ		= .obj
endif

# Specify the executables used.
MAKE		= make
CAT		= cat

CP		 = cp
LEX		 = flex -Pscotchyy -olex.yy.c
LN		 = ln
MKDIR	 = mkdir -p
MV		 = mv
YACC	 = bison -pscotchyy -y -b y

# Platform-specific executables.
ifeq ($(IS_WINDOWS),1)
# On windows we use the compiler driver for both lib and icl
AR		= ../winwrap lib
ARFLAGS = /nologo
CC    = ../winwrap icl

RANLIB = echo
else
AR		= ar
ARFLAGS = -ruv
CC    = icc

RANLIB = ranlib
endif

CCS		= $(CC)
CCP		= $(CC)
CCD		= $(CC)

# Specify the compiler and linker flags.
CFLAGS	= -O3
CFLAGS += -DHAVE_STDINT_H -DHAVE_UINT_T
CFLAGS += -DCOMMON_FILE_COMPRESS_GZ
CFLAGS += -DSCOTCH_DETERMINISTIC
CFLAGS += -DSCOTCH_RENAME -DSCOTCH_RENAME_PARSER

ifeq ($(IS_WINDOWS),1)
# Setup CFLAGS.
CFLAGS += -DCOMMON_STUB_FORK
CFLAGS += /Qrestrict -DIDXSIZE64

CFLAGS += -DCOMMON_WINDOWS
# CFLAGS += -DCOMMON_RANDOM_FIXED_SEED
CFLAGS += -DYY_NO_UNISTD_H

# Include some dummy header files for the strings and sys/time.h
CFLAGS += -I../../../win_includes
CFLAGS += /nologo /Qvc12

CLIBFLAGS	= /MD
LDFLAGS = /link /fixed:no /incremental:no
LDFLAGS += /NODEFAULTLIB:libcmt /NODEFAULTLIB:libcmtd
else
CFLAGS += -DCOMMON_PTHREAD
CFLAGS += -restrict -DIDXSIZE64
# LDFLAGS		= -lz -lm -lrt -pthread
endif

# Thirdpary.
CFLAGS += -I$(ZLIB_INC)
CFLAGS += -I$(MPI_INC)

LDFLAGS += $(ZLIB_LIBRARY)
LDFLAGS += $(MPI_LIBRARY)


# Include thirdparty for the correct MPI and ZLIB paths.
includedir ?= .
HOME_DIR = $(includedir)/../../../../..
MAKEINC_DIR = $(HOME_DIR)/src/BUILD_tools/make_includes/
THIRDPARTY_ROOT = $(HOME_DIR)/thirdparty

ifeq ($(COMMLIB),IMPI)
BUILD_IMPI=1
MPI_INC=$(IMPI_INC)
MPI_LIBRARY=$(IMPI_LIBRARY)
endif
ifeq ($(COMMLIB),MSMPI)
BUILD_MSMPI=1
MPI_INC=$(MSMPI_INC)
MPI_LIBRARY=$(MSMPI_LIBRARY)
endif
ifeq ($(COMMLIB),MPICH)
BUILD_MPICH=1
MPI_INC=$(MPICH_INC)
MPI_LIBRARY=$(MPICH_LIBRARY)

# There is a problem building for MPICH on LINUX32. The Intel Fortran libs are required.
ifneq (,$(findstring linux32_ia32,$(PLAT_NAME)))
CFLAGS += $(IFORT90_LIBRARY)
endif

endif
ifeq ($(COMMLIB),MPT)
BUILD_MPT=1
MPI_INC=$(MPT_INC)
MPI_LIBRARY=$(MPT_LIBRARY)
endif


# include $(MAKEINC_DIR)/platforms.inc
include $(MAKEINC_DIR)/thirdparty_versions.inc

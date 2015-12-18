# Specify the version that we are building and the archive.
version ?= 6.0.5rc3
archive ?= scotch_$(version).tar.gz

# For windows PLAT_NAME=win64_x86_64 COMMLIB=IMPI
PLAT_NAME ?= $(error The platform name must be defined using PLAT_NAME=)
COMMLIB ?= $(error The MPI library must be defined using COMMLIB=)


IS_WINDOWS = 0
ifeq (,$(findstring linux,$(PLAT_NAME)))
IS_WINDOWS = 1
LIB_EXT = .lib
else
LIB_EXT = .a
endif



source_dir = $(version)_$(PLAT_NAME)_$(COMMLIB)_source
target_src = $(source_dir)/src


target_dist = ../$(version)/$(PLAT_NAME)
commlib_makefile = Makefile.$(PLAT_NAME).$(COMMLIB).inc

$(info $(target_src))


distclean:
	rm -rf $(target_dist)

clean:
	rm -rf $(source_dir)

ifeq ($(IS_WINDOWS),1)
# For windows, a compiler wrapper is required to minimise the changes to the SCOTCH
# makefiles.
# This requires the go tool (http://golang.org/)
compiler_wrapper: winwrap.go
	go build winwrap.go
	cp -p winwrap.exe $(target_src)
else
compiler_wrapper:
	echo No compiler wrapper needed.
endif


source_dir:
	mkdir -p $(source_dir)


target_dist:
	mkdir -p $(target_dist)
	mkdir -p $(target_dist)/include
	mkdir -p $(target_dist)/lib
	mkdir -p $(target_dist)/lib/$(COMMLIB)

source: source_dir $(archive)
	tar -xzvf $(archive) -C $(source_dir) --strip-components=1
	bash ./fix_makefiles.sh $(target_src)


copy_input: source copy_makefile_inc compiler_wrapper


copy_makefile_inc: Makefile.Master.inc
	echo "# Generated from "$^ > $(target_src)/Makefile.inc
	echo PLAT_NAME = $(PLAT_NAME) >> $(target_src)/Makefile.inc
	echo COMMLIB = $(COMMLIB) >> $(target_src)/Makefile.inc
	cat $^ >> $(target_src)/Makefile.inc


sequential: copy_input
	make -C $(target_src) scotch

parallel: copy_input sequential
	make -C $(target_src) ptscotch


esmumps: parallel sequential
	make -C $(target_src)/esmumps install
	make -C $(target_src)/esmumps ptinstall


install: target_dist sequential parallel esmumps
	cp -p $(source_dir)/include/* $(target_dist)/include/
	bash ./copy_libs.sh $(source_dir)/lib $(target_dist)/lib $(COMMLIB) $(LIB_EXT)




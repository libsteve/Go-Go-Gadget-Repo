include $(GOROOT)/src/MAKE.inc

all: 	prepro.go pp.go
	6g prepro.go
	6g pp.go
	6l -o pp pp.6

doc: 	; ( godoc -http=:6060 -path=./ & \
  sleep 5; open -a Safari http://localhost:6060/pkg/prepro; \
  sleep 15; kill $$! )

include $(GOROOT)/src/Make.cmd
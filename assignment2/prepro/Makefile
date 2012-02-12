include $(GOROOT)/src/MAKE.inc

all: 	pp.go main.go
	6g pp.go
	6g main.go
	6l -o pp main.6

doc: 	; ( godoc -http=:6060 -path=./ & \
  sleep 5; open -a Safari http://localhost:6060/pkg/prepro; \
  sleep 15; kill $$! )

include $(GOROOT)/src/Make.cmd
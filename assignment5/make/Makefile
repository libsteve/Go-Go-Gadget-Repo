include $(GOROOT)/src/MAKE.inc

all:	main.go dag.6 edge_implementation.go
	6g edge_implementation.go
	6g parser.go
	6g main.go
	6l -o make main.6 

doc:	; ( godoc -http=:6060 -path=./ & \
  sleep 5; open -a Safari http://localhost:6060/pkg/dag; \
  sleep 15; kill $$! )

include $(GOROOT)/src/Make.cmd

include $(GOROOT)/src/MAKE.inc

all:	main.go dag.go dag_implementaton.go edge_implementation.go
	6g dag.go edge_implementation.go dag_implementaton.go
	6g main.go
	6l -o dag main.6

doc:	; ( godoc -http=:6060 -path=./ & \
  sleep 5; open -a Safari http://localhost:6060/pkg/dag; \
  sleep 15; kill $$! )

include $(GOROOT)/src/Make.cmd

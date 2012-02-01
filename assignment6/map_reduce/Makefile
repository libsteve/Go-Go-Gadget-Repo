include $(GOROOT)/src/MAKE.inc

all: matrixServer matrixClient regserver regclient

parser.6: parser.go
	6g parser.go

########
### This deals with the matrix

matrixServer: imatrix.6 matrixserver.go
	6g matrixserver.go
	6l -o matrixServer matrixserver.6

matrixClient: imatrix.6 matrixclient.go parser.6
	6g matrixclient.go
	6l -o matrixClient matrixclient.6

imatrix.6: imatrix.go matrixcommon.go
	6g imatrix.go matrixcommon.go
###
########


########
### This deals with the registry

regserver: regserver.6
	6l -o regserver regserver.6

regclient: regclient-run.6 regclient.6 parser.6
	6l -o regclient regclient-run.6

regserver.6: regserver.go registry.6
	6g regserver.go

regclient-run.6: regclient-run.go regclient.6 parser.6
	6g regclient-run.go

regclient.6: regclient.go registry.6
	6g regclient.go

registry.6: registry.go iregistry.go
	6g registry.go iregistry.go
###
########

doc:	; ( godoc -http=:6060 -path=./ & \
  sleep 5; open -a Safari http://localhost:6060/pkg/map_reduce; \
  sleep 15; kill $$! )

include $(GOROOT)/src/Make.cmd
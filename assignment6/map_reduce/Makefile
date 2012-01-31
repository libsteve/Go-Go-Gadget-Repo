include $(GOROOT)/src/MAKE.inc

all: matrixServer matrixClient

matrixServer: imatrix.6 matrixserver.go
	6g matrixserver.go
	6l -o matrixServer matrixserver.6

matrixClient: imatrix.6 matrixcommon.go
	6g matrixclient.go
	6l -o matrixClient matrixclient.6

imatrix.6: imatrix.go matrixcommon.go
	6g imatrix.go matrixcommon.go

doc:	; ( godoc -http=:6060 -path=./ & \
  sleep 5; open -a Safari http://localhost:6060/pkg/map_reduce; \
  sleep 15; kill $$! )

include $(GOROOT)/src/Make.cmd
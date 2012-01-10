include $(GOROOT)/src/MAKE.inc

all:	main.go ls.go
	6g ls.go
	6g main.go
	6l -o ls main.6

doc: 	; ( godoc -http=:6060 -path=./ & \
  sleep 5; open -a Safari http://localhost:6060/pkg/ls; \
  sleep 15; kill $$! )

include $(GOROOT)/src/Make.cmd

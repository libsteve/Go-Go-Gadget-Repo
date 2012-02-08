include $(GOROOT)/src/MAKE.inc

all: wag

wag: main.go wagon.6 screen.6
	6g main.go
	6l -o wag main.6

wagon.6: wagon.go
	6g wagon.go

screen.6: screen.go
	6g screen.go

doc:	; ( godoc -http=:6060 -path=./ & \
  sleep 5; open -a Safari http://localhost:6060/pkg/wagon; \
  sleep 15; kill $$! )

include $(GOROOT)/src/Make.cmd
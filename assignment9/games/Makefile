include $(GOROOT)/src/MAKE.inc

all:	games

games:	main.go games.6 rps.6 ttt.6
	6g main.go
	6l -o games main.6

games.6:	games.go igame.go referee.go view.go
	6g games.go igame.go referee.go view.go

rps.6:	rps.go games.6
	6g rps.go

ttt.6:	ttt.go games.6
	6g ttt.go

doc:	; ( godoc -http=:6060 -path=./ & \
	sleep 5; open -a Safari http://localhost:6060/pkg/games; \
	sleep 15; kill $$! )

include $(GOROOT)/src/MAKE.cmd
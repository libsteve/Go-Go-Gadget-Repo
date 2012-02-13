include $(GOROOT)/src/MAKE.inc

all:	games

games:	main.go view.6 referee.6 rps.6 ttt.6
	6g main.go
	6l -o games main.go

view.6:	view.go games.6
	6g view.go

referee.6:	referee.go games.6
	6g referee.go

games.6:	games.go igame.go
	6g igame.go games.go

rps.6:	rps.go
	6g rps.go

ttt.6:	ttt.go
	6g ttt.go

doc:	; ( godoc -http=:6060 -path=./ & \
	sleep 5; open -a Safari http://localhost:6060/pkg/games; \
	sleep 15; kill $$! )

include $(GOROOT)/src/MAKE.cmd
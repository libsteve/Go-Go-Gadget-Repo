include $(GOROOT)/src/MAKE.inc

all:	local remote server

server: server.go
	6g server.go
	6l -o server server.6

remote: remote.go games.6 rps.6 ttt.6
	6g remote.go
	6l -o remote remote.6

local:	local.go games.6 rps.6 ttt.6
	6g local.go
	6l -o local local.6

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
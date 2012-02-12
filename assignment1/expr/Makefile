include $(GOROOT)/src/MAKE.inc

all:	main.go expr.go
	6g expr.go
	6g main.go
	6l -o expr main.6

include $(GOROOT)/src/Make.cmd

doc:	 	 ; ( godoc -http=:6060 -path=./ & \
  sleep 5; open -a Safari http://localhost:6060/pkg/expr; \
  sleep 15; kill $$! )

test:
	./expr 1 + 2
	expr 1 + 2
	@echo 
	./expr 2 \* \( 3 + 4 \) / \( 2 - 4 \)
	expr 2 \* \( 3 + 4 \) / \( 2 - 4 \)
	@echo 
	./expr 3 + \( \( 2 + 5 \) \*  \( 3 + 4 \) \) / \( 2 - 4 \)
	expr 3 + \( \( 2 + 5 \) \*  \( 3 + 4 \) \) / \( 2 - 4 \)
	@echo
	./expr 3 + 5 \* 4
	expr 3 + 5 \* 4
	@echo 
	./expr 3 + 7 + 7 + 7 + 7
	expr 3 + 7 + 7 + 7 + 7
	@echo 
	./expr \( 5 \* 6 \) + 7
	expr \( 5 \* 6 \) + 7
	@echo 
	- ./expr \(
	- expr \(
	@echo 
	- ./expr
	- expr
	@echo 
	- ./expr \( 5 + 6
	- expr \( 5 + 6
	@echo 
	- ./expr 5+6
	- expr 5+6
	@echo 
	- ./expr 1 + 5+6
	- expr 1 + 5+6
	@echo 
	- ./expr *
	- expr *
	@echo 
	- ./expr +
	- expr +

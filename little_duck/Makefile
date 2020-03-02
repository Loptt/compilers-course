duck: lex.yy.c duck.tab.c duck.tab.h
	g++ duck.tab.c lex.yy.c -lfl -o duck

duck.tab.c duck.tab.h: duck.y
	bison -d duck.y

lex.yy.c: duck.l duck.tab.h
	flex duck.l
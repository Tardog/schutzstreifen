PROJECT=schutzstreifen

install:
	buffalo db create -a

start:
	buffalo dev

test:
	buffalo test

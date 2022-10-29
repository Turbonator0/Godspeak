GC=go
OUT=godspeak

godspeak: server.go speak.go; $(GC) build -o $(OUT) server.go speak.go ; chmod +x $(OUT)
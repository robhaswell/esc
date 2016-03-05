bin/rsrc -manifest esc.manifest -o rsrc.syso
go build -ldflags="-H windowsgui"

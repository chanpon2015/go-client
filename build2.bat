xcopy vue\dist resources\app /s/e/i/y
move bind.go work\bind.go
astilectron-bundler -v
del /Q bind*.go
move work\bind.go bind.go
pause
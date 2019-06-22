cd ./vue
yarn build
cp -r -f ./dist/* ../resources/app
cd ../
mv ./bind.go ./work/bind.go
astilectron-bundler -w
mv ./work/bind.go ./bind.go
rm bind_darwin_amd64.go
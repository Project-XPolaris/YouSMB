go build main.go
rm -rf ./pack-output
mkdir "./pack-output"
cp ./config.json ./pack-output/config.json
cp ./main ./pack-output/yousmb
cp ./install.sh ./pack-output/install.sh
#example parameter 0.15.0
#--consensus.create_empty_blocks=false
#--p2p.seeds="bftx0.blockfreight.net:888,bftx1.blockfreight.net:888,bftx2.blockfreight.net:888,bftx3.blockfreight.net:888"
#--p2p.seeds=104.42.43.66:888
#--consensus.create_empty_blocks=false
$1=0.15.0
VERSION=0.15.0
rm -rf ${VERSION}
#rm ${VERSION}/tendermint_${VERSION}_darwin_386.zip
#rm ${VERSION}/tendermint
#wget https://github.com/tendermint/tendermint/releases/download/v${VERSION}/tendermint_${VERSION}_darwin_386.zip -P ${VERSION}
#wget https://github.com/blockfreight/go-bftx/releases/download/v0.5.15/bftnode -P ${VERSION}
#wget https://github.com/blockfreight/go-bftx/releases/download/v0.5.15/tendermint -P ${VERSION}
wget https://github.com/blockfreight/go-bftx/releases/download/v0.5.15/blockfreight.zip -P ${VERSION}
unzip ${VERSION}/blockfreight.zip -d ${VERSION}
chmod +x ${VERSION}/tendermint
chmod +x ${VERSION}/bftnode
#unzip ${VERSION}/tendermint_${VERSION}_darwin_386.zip -d ${VERSION}
${VERSION}/tendermint version
${VERSION}/tendermint unsafe_reset_all --home ${VERSION}
${VERSION}/tendermint init --home ./${VERSION}
bftnode & ${VERSION}/tendermint node --home ./${VERSION} --consensus.create_empty_blocks=false --p2p.seeds="bftx0.blockfreight.net:888,bftx1.blockfreight.net:888,bftx2.blockfreight.net:888,bftx3.blockfreight.net:888"

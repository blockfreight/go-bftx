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
wget https://github.com/tendermint/tendermint/releases/download/v${VERSION}/tendermint_${VERSION}_darwin_386.zip -P ${VERSION}
unzip ${VERSION}/tendermint_${VERSION}_darwin_386.zip -d ${VERSION}
${VERSION}/tendermint version
${VERSION}/tendermint unsafe_reset_all --home ${VERSION}
${VERSION}/tendermint init --home ./${VERSION}
${VERSION}/tendermint node --home ./${VERSION} --p2p.seeds="bftx0.blockfreight.net:888,bftx1.blockfreight.net:888,bftx2.blockfreight.net:888,bftx3.blockfreight.net:888"

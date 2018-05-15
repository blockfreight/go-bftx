#example parameter 0.15.0
#rm ${VERSION}/tendermint_${VERSION}_darwin_386.zip
#rm ${VERSION}/tendermint
#wget https://github.com/tendermint/tendermint/releases/download/v${VERSION}/tendermint_${VERSION}_darwin_386.zip -P ${INSTALL_DIR}
#wget https://github.com/blockfreight/go-bftx/releases/download/v0.5.15/bftnode -P ${INSTALL_DIR}
#wget https://github.com/blockfreight/go-bftx/releases/download/v0.5.15/tendermint -P ${INSTALL_DIR}
#--consensus.create_empty_blocks=false
#--p2p.seeds="bftx0.blockfreight.net:888,bftx1.blockfreight.net:888,bftx2.blockfreight.net:888,bftx3.blockfreight.net:888"
#--p2p.seeds=104.42.43.66:888
#--consensus.create_empty_blocks=false
#unzip ${VERSION}/tendermint_${VERSION}_darwin_386.zip -d ${INSTALL_DIR}
$1=0.15.0
VERSION=0.15.0
INSTALL_DIR=~/.blockfreight
rm -rf ${INSTALL_DIR}
#wget https://github.com/blockfreight/go-bftx/releases/download/v0.5.15/blockfreight.zip -P ${INSTALL_DIR}
curl -o ${INSTALL_DIR}/blockfreight.zip -L https://github.com/blockfreight/go-bftx/releases/download/v0.5.15/blockfreight.zip
unzip ${INSTALL_DIR}/blockfreight.zip -d ${INSTALL_DIR}
chmod +x ${INSTALL_DIR}/tendermint
chmod +x ${INSTALL_DIR}/bftnode
${INSTALL_DIR}/tendermint version
${INSTALL_DIR}/tendermint unsafe_reset_all --home ${INSTALL_DIR}
${INSTALL_DIR}/tendermint init --home ./${INSTALL_DIR}
${INSTALL_DIR}/bftnode & ${INSTALL_DIR}/tendermint node --home ./${INSTALL_DIR} --consensus.create_empty_blocks=false --p2p.seeds="bftx0.blockfreight.net:888,bftx1.blockfreight.net:888,bftx2.blockfreight.net:888,bftx3.blockfreight.net:888"

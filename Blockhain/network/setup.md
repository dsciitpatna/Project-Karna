 cryptogen generate --config=./crypto-config.yaml


 mkdir channel-artifacts


 configtxgen -profile Genesis -outputBlock ./channel-artifacts/genesis.block -channelID genesischannel 
 configtxgen -profile KarnaChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID karnachannel
 configtxgen -profile KarnaChannel --outputAnchorPeersUpdate ./channel-artifacts/NGOOrgAnchorUpdate.tx -asOrg NGO -channelID karnachannel
 configtxgen -profile KarnaChannel --outputAnchorPeersUpdate ./channel-artifacts/VolOrgAnchorUpdate.tx -asOrg Volunteers -channelID karnachannel


 
 change file name under ca folder : ca-certp=.pem to cert.pem and private_sk to PRIVATE_KEY in bot ngo.com and vol.com


cli commands
for devpeer
export CORE_PEER_ADDRESS=devpeer:7051
export CORE_PEER_LOCALMSPID=NGOMSP
export CORE_PEER_MSPCONFIGPATH=/crypto-config/peerOrganizations/ngo.com/users/Admin@ngo.com/msp

for firstngo
export CORE_PEER_ADDRESS=firstngo:7051
export CORE_PEER_LOCALMSPID=NGOMSP
export CORE_PEER_MSPCONFIGPATH=/crypto-config/peerOrganizations/ngo.com/users/Admin@ngo.com/msp

for devvol
export CORE_PEER_ADDRESS=devvol:7051
export CORE_PEER_LOCALMSPID=VolunteersMSP
export CORE_PEER_MSPCONFIGPATH=/crypto-config/peerOrganizations/vol.com/users/Admin@vol.com/msp/

installing chaincode 

peer chaincode install -n karna -v 0 -p KarnaChanincode
peer chaincode instantiate -n karna -v 0 -C karnachannel -c '{"args":[]}'
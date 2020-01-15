 cryptogen generate --config=./crypto-config.yaml


 mkdir channel-artifacts


 configtxgen -profile Genesis -outputBlock ./channel-artifacts/genesis.block -channelID genesischannel 
 configtxgen -profile KarnaChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID karnachannel
 configtxgen -profile KarnaChannel --outputAnchorPeersUpdate ./channel-artifacts/NGOOrgAnchorUpdate.tx -asOrg NGO -channelID karnachannel
 configtxgen -profile KarnaChannel --outputAnchorPeersUpdate ./channel-artifacts/VolOrgAnchorUpdate.tx -asOrg Volunteers -channelID karnachannel


 
 change file name under ca folder : ca-certp=.pem to cert.pem and private_sk to PRIVATE_KEY in bot ngo.com and vol.com

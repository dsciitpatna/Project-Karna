const fs = require('fs')
const yaml = require('js-yaml')
const {FileSystemWallet,Gateway} = require('fabric-network')
const CONNECTION_PROFILE_PATH= "./connection.yaml"
const WALLET_PATH="./wallet"
const IDENTITY_NAME = "AppClient"
const CHANNEL_NAME = "karnachannel"
const CONTRACT_NAME="karna"

async function connectNetwork(){
       const ccp = yaml.safeLoad(fs.readFileSync(CONNECTION_PROFILE_PATH))
       const wallet = new FileSystemWallet(WALLET_PATH)
       const is = await wallet.exists(IDENTITY_NAME) 
       if (!is){
           return
       }
       const gateway = new Gateway()
       await gateway.connect(ccp,{wallet:wallet,identity:IDENTITY_NAME,discovery:{enabled:false,asLocalhost:true}})
       const network = await gateway.getNetwork(CHANNEL_NAME)
       const contract = network.getContract(CONTRACT_NAME)
       return contract
}
module.exports = {connectNetwork}
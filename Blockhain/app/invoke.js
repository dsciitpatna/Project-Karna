const fs = require('fs')
const yaml = require('js-yaml')
const {FileSystemWallet,Gateway} = require('fabric-network')
const CONNECTION_PROFILE_PATH= "./connection.yaml"
const WALLET_PATH="./wallet"
const IDENTITY_NAME = "AppClient"
const CHANNEL_NAME = "karnachannel"
const CONTRACT_NAME="karna"

async function invoke(){
    try {
        const wallet = new  FileSystemWallet(WALLET_PATH)
        const is = await wallet.exists(IDENTITY_NAME)
        if (!is){
            console.log(`${IDENTITY_NAME} doesn't exists`)
            return
        }
        const ccp = yaml.safeLoad(fs.readFileSync(CONNECTION_PROFILE_PATH))
        const gateway = new Gateway()
        await gateway.connect(ccp,{wallet:wallet,identity:IDENTITY_NAME,discovery:{enabled:false,asLocalhost:true}})
        const newtork = await gateway.getNetwork(CHANNEL_NAME)
        const contract = newtork.getContract(CONTRACT_NAME)
        response = await contract.evaluateTransaction("getUser","Zzocker")
        console.log(response.Message)
    } catch (error) {
        console.log(error)
        process.exit(1)
    }
}
invoke()
const fs = require('fs')
const yaml  = require('js-yaml')
const {FileSystemWallet,Gateway,X509WalletMixin} = require("fabric-network")
const CONNECTION_PROFILE_PATH= "./connection.yaml"
const WALLET_PATH="./wallet"
const IDENTITY_NAME = "AppClient"
const MSPID = "NGOMSP"

async function main(){
    try {
        const wallet = new FileSystemWallet(WALLET_PATH)
        let is = await wallet.exists(IDENTITY_NAME)
        if (is){
            console.log(`${IDENTITY_NAME} already exists`)
            return
        }
        is = await wallet.exists("admin")
        if (!is){
            console.log("Admin doesn't exists")
            return
        }
        ccp = yaml.safeLoad(fs.readFileSync(CONNECTION_PROFILE_PATH))
        const gateway = new Gateway()
        await gateway.connect(ccp,{wallet:wallet,identity:"admin",discovery:{ enabled: true, asLocalhost: true}})
        const ca = gateway.getClient().getCertificateAuthority()
        const admin = gateway.getCurrentIdentity()
        await ca.register({enrollmentID:IDENTITY_NAME,enrollmentSecret:"pw",role:'client',affiliation:""},admin)
        const enrollment = await ca.enroll({enrollmentID:IDENTITY_NAME,enrollmentSecret:"pw",}) 
        const identity = X509WalletMixin.createIdentity(MSPID,enrollment.certificate,enrollment.key.toBytes())
        await wallet.import(IDENTITY_NAME,identity)
        console.log(`${IDENTITY_NAME} successfully enrolled`)
    } catch (error) {
        console.log(error)
        process.exit(1)
    }
}
main()
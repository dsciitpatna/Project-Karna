const fs = require('fs')
const yaml = require('js-yaml')
const {X509WalletMixin,FileSystemWallet,Gateway} = require('fabric-network')
const FabricCa = require('fabric-ca-client')
const CONNECTION_PROFILE_PATH= "./connection.yaml"
const WALLET_PATH="./wallet"
const MSPID = "NGOMSP"

async function main(){
    try {
        const ccp = yaml.safeLoad(fs.readFileSync(CONNECTION_PROFILE_PATH))
        const wallet = new FileSystemWallet(WALLET_PATH)
        const adminExists = await wallet.exists('admin')
        console.log(ccp.certificateAuthorities.ngoca.url)
        if (adminExists){
            console.log("Admin already exists")
            process.exit(1)
        }
        const ca = new FabricCa(ccp.certificateAuthorities.ngoca.url)
        const enrollment = await ca.enroll({enrollmentID:"admin",enrollmentSecret:"adminpw"})
        const identity = X509WalletMixin.createIdentity(MSPID,enrollment.certificate,enrollment.key.toBytes())
        await wallet.import('admin',identity)
        console.log("Successfully enrolled admin of ngo")
    } catch (error) {
        console.log(error)
        process.exit(1)
    }
}
main()
const fs = require('fs')
const yaml = require('js-yaml')
const { Gateway,FileSystemWallet, X509WalletMixin } = require('fabric-network')
const CONNECTION_PROFILE_PATH= './connection.yaml'
const WALLET_PATH='./wallet/ngo'
const cpp = yaml.safeLoad(fs.readFileSync(CONNECTION_PROFILE_PATH))
console.log(cpp.certificateAuthorities)
async function main(){
    try {
        const wallet = new FileSystemWallet(WALLET_PATH)
        const userExists = await wallet.exists('ngoAdmin')
        if (userExists){
            console.log("org admin alredy exists")
        }
        const admin = await wallet.exists("admin")
        if (!admin){
            console.log('admin doesn exists')
            return 
        }
        const gateway = new Gateway()
        await gateway.connect(cpp,{wallet:wallet,identity:'admin',discovery:{ enabled: true, asLocalhost: true}})
        const ca = gateway.getClient().getCertificateAuthority()
        const adminIdentity = gateway.getCurrentIdentity()
        // console.log(adminIdentity)
        await ca.register({enrollmentID:'ngoAdmin',role:'admin',affiliation:'',role:"admin",enrollmentSecret:"pw"},adminIdentity)
        const enrollment = await ca.enroll({enrollmentID:'ngoAdmin',enrollmentSecret:"pw"})
        const identity = X509WalletMixin.createIdentity("NGOMSP",enrollment.certificate,enrollment.key.toBytes())
        await wallet.import("ngoAdmin",identity)
        console.log('Successfully registered and enrol ngoAdmin to the network')
    } catch (error) {
        console.log(error)
        process.exit(1)
    }
}
main()
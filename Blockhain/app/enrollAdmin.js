const fs = require('fs')
const {X509WalletMixin,FileSystemWallet} = require('fabric-network')
const yaml = require('js-yaml')
const FabricCAServices = require('fabric-ca-client')
const CONNECTION_PROFILE_PATH = 'connection.yaml'
const ccp = yaml.safeLoad(fs.readFileSync(CONNECTION_PROFILE_PATH))
const ngowallet = 'wallet/ngo/'
async function main(){
    try {
        const caInfo = ccp.certificateAuthorities['ngoca']
        console.log(caInfo.url)
        //console.log(fs.readFileSync(pem).toString())
        const ca = new FabricCAServices(caInfo.url)
        const wallet = new FileSystemWallet(ngowallet)
        const admin = await wallet.exists('admin')
        if (admin){
            console.log("An identity for admin already ecists in walled ")
        }
        const enrollment = await ca.enroll({enrollmentID:'admin',enrollmentSecret:'adminpw'})
        const identity = X509WalletMixin.createIdentity('NGOMSP',enrollment.certificate,enrollment.key.toBytes())
        await wallet.import('admin',identity)
        console.log('Successfully enrolled admin user "admin" and imported it into the wallet')
    } catch (error) {
        console.error(`Failed to enroll admin user "admin": ${error}`)
        process.exit(1)
    }
}
main()
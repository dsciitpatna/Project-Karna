const express = require('express')
const api =  express()
const PORT = "3000"
const logger= (req,res,next)=>{
    console.log(`${req.protocol}://${req.get('host')}${req.originalUrl}`)
    next()
}
api.use(logger)
api.use(express.json())
const network = require('./contract')

api.get('/api/user/get/:username',async (req,res)=>{
    try {
        const contract = await network.connectNetwork()
        const response = await contract.evaluateTransaction('getUser',req.params.username)
        let result  = JSON.parse(response)
        res.status(200).json({result:result})
    } catch (error) {
        console.error(`Failed to evaluate transaction: ${error}`);
        res.status(500).json({
      error: error
    })
    }
})
api.post('/api/user/register/',async (req,res)=>{
    try {
    const contrct = await network.connectNetwork()
    await contrct.submitTransaction('userRegistration',req.body.username,req.body.name,req.body.password)
    res.status(200).json({msg:"Successfully created user"})
    } catch (error) {
        console.error(`Failed to evaluate transaction: ${error}`);
        res.status(500).json({
        error: error
    })
    }
})
api.get('/api/user/login',async (req,res)=>{
    try {
        const contract = await network.connectNetwork()
        const response = await contract.evaluateTransaction('userGateway','userLogin',req.body.username,req.body.password)
        res.status(200).json({
            msg:"Successfully loged in",
            token : response.toString()
        })
    } catch (error) {
        res.status(500).json({
            msg:"uable to login, re-try"
        })
    }
})
api.put('/api/user/donate',async (req,res)=>{
    try {
        const contract = await network.connectNetwork()
        const rbody = req.body
        const response =  await contract.evaluateTransaction('userGateway',rbody.token,"donate",rbody.ngo_id,rbody.mission_id,rbody.donation)
        res.status(200).json(JSON.parse(response))
    } catch (error) {
        res.status(500)
    }
})
api.get('/api/user/getDonatedMission/:token',async (req,res)=>{
    try {
        const contract = network.connectNetwork()
        const response = await contract.evaluateTransaction('userGateway',req.params.token,"getDonatedMission")
        res.status(200).json(JSON.parse(response))
    } catch (error) {
        res.status(500)
    }
})
api.post('/api/ngo/register',async (req,res)=>{
    try {
        const contract = await network.connectNetwork()
        await contract.submitTransaction('NGORegistration',req.body.username,req.body.name,req.body.address,req.body.description,req.body.password)
        res.status(200).json({
            msg:"Successfully created NGO"
        })
    } catch (error) {
        res.status(500).json({
            msg:"unable to register NGO"
        })
    }
})
api.get('/api/ngo/get/:username',async (req,res)=>{
    try {
        const contract = await network.connectNetwork()
        const response = await contract.evaluateTransaction('getNgo',req.params.username) 
        res.status(200).json(JSON.parse(response))
    } catch (error) {
        res.status(500).json({
            msg:"unable to get ngo, re-try"
        })
    }
})
api.get('/api/ngo/login',async (req,res)=>{
    try {
        const contract = await network.connectNetwork()
        const response= await contract.evaluateTransaction('userGateway','userLogin',req.body.username,req.body.password)
        res.status(200).json({
            msg:"Successfully loged in",
            token : response.toString()
        })
    } catch (error) {
        res.status(500).json({
            msg:"Unable to login NGO"
        })
    }
})
api.post('/api/ngo/createMission',async (req,res)=>{
    try {
        const contract = await network.connectNetwork()
        const rbody = req.body
        const response = await contract.submitTransaction('ngoGateway',rbody.token,"createMission",rbody.mission_id,rbody.name,rbody.description,rbody.target)
        res.status(200).json(
            JSON.parse(response)
        )
    } catch (error) {
        res.status(500)
    }
})
api.get('/api/ngo/getNgoMission/:token',async (req,res)=>{
    try {
        const contract = await network.connectNetwork()
        const response = await contract.evaluateTransaction('ngoGateway',req.params.token,"getNgoMission")
        res.status(200).json(
            JSON.parse(response)
        )
    } catch (error) {
        res.status(500)
    }
})
api.listen(PORT,()=>{
    console.log(`listening on port ${PORT}`)
})

/*
 * SPDX-License-Identifier: Apache-2.0
 */

'use strict';

const config = require("../config/hlf_network.json")
const connectionProfilePath = config.connectionProfilePath;
const walletDirPath = config.walletPath;
const appUser=config.appUser;
const channelName=config.channel;
const chaincodeName = config.chaincode;

// const { FileSystemWallet, Gateway } = require('fabric-network');
const { Gateway, Wallets } = require('fabric-network');
const fs = require('fs');

const path = require('path');
// const fs = require('fs');
const ccpPath = path.resolve(connectionProfilePath);

let ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));

// console.log("controller/invoke ccp",ccp);
// console.log("controller/invoke ccppath",ccpPath);

// // ** create user **//
var createuser = async function createuser(req, res) {
    try {

        // Create a new file system based wallet for managing identities.
        const walletPath = path.resolve(walletDirPath);
        const wallet = await Wallets.newFileSystemWallet(walletPath);

        // console.log(`Wallet path: ${walletPath}`);

        // Check to see if we've already enrolled the user.
        const identity = await wallet.get(appUser);
        if (!identity) {
            console.log('An identity for the user "user2" does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return res.status(400).send({ 'result': `user not registered` });
        }

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();

        // console.log("gateway",gateway);
        // console.log("wallet",wallet);
        // console.log("ccpPath",ccpPath);

        await gateway.connect(ccp, { wallet, identity: appUser, discovery: { enabled: true, asLocalhost: true } });

        // console.log("next to geteway.connect",conn);

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork('mychannel');

        //  console.log("network",network);

        // Get the contract from the network.
        const contract = network.getContract('scott');

        // console.log("contract",contract);

        var obj = {
            "fName": req.body.fName,
            "lName": req.body.lName,
            "EmailId": req.body.EmailId,
            "mobile": req.body.mobile,
            "password": req.body.password,
            "country": req.body.country,
            "vinnumber": req.body.vinnumber,
        }

        await contract.submitTransaction('createUser', JSON.stringify(obj));

        // Disconnect from the gateway.
        console.log('Transaction has been submitted');
        await gateway.disconnect();
        return res.status(200).send({ 'result': `Transaction has been submitted` });

    } catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        return res.status(400).send({ 'result': `Failed to submit transaction: ${error}` });

        //process.exit(1);
    }
}

// // ** update user **//
var updateUser = async function updateUser(req, res) {
    try {

        // Create a new file system based wallet for managing identities.
        const walletPath = path.resolve(walletDirPath);
        const wallet = await Wallets.newFileSystemWallet(walletPath);

        // Check to see if we've already enrolled the user.
        const identity = await wallet.get(appUser);
        if (!identity) {
            console.log('An identity for the user "user2" does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return res.status(400).send({ 'result': `user not registered` });
        }

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();

        await gateway.connect(ccp, { wallet, identity: appUser, discovery: { enabled: true, asLocalhost: true } });

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork('mychannel');

        // Get the contract from the network.
        const contract = network.getContract('scott');
        var obj = {
            "fName": req.body.fName,
            "lName": req.body.lName,
            "EmailId": req.body.EmailId,
            "mobile": req.body.mobile,
            "password": req.body.password,
            "country": req.body.country,

        }

        await contract.submitTransaction('updateUser', JSON.stringify(obj));

        // Disconnect from the gateway.
        console.log('Transaction has been submitted');
        await gateway.disconnect();
        return res.status(200).send({ 'result': `Transaction has been submitted` });

    } catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        return res.status(400).send({ 'result': `Failed to submit transaction: ${error}` });

        //process.exit(1);
    }
}

// // ** create Admin **//
var createAdministration = async function createAdmin(req, res) {
    try {

        // Create a new file system based wallet for managing identities.
        const walletPath = path.resolve(walletDirPath);
        const wallet = await Wallets.newFileSystemWallet(walletPath);

        // Check to see if we've already enrolled the user.
        const identity = await wallet.get(appUser);
        if (!identity) {
            console.log('An identity for the user "user2" does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return res.status(400).send({ 'result': `user not registered` });
        }

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();

        await gateway.connect(ccp, { wallet, identity: appUser, discovery: { enabled: true, asLocalhost: true } });

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork('mychannel');

        // Get the contract from the network.
        const contract = network.getContract('scott');
        var obj = {
            "adminId": req.body.adminId,
            "adminFirstName": req.body.adminFirstName,
            "adminLastName": req.body.adminLastName,
            "adminPhone": req.body.adminPhone,
            "adminEmail": req.body.adminEmail,
            "cts": req.body.cts,
            "dob": req.body.dob,
            "location": req.body.location,
            "password": req.body.password,
        }

        await contract.submitTransaction('createAdministration', JSON.stringify(obj));

        // Disconnect from the gateway.
        console.log('Transaction has been submitted');
        await gateway.disconnect();
        return res.status(200).send({ 'result': `Transaction has been submitted` });

    } catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        return res.status(400).send({ 'result': `Failed to submit transaction: ${error}` });

        //process.exit(1);
    }
}
// // ** update Admin **//
var updateAdministration = async function updateAdministration(req, res) {
    try {

        // Create a new file system based wallet for managing identities.
        const walletPath = path.resolve(walletDirPath);
        const wallet = await Wallets.newFileSystemWallet(walletPath);

        // Check to see if we've already enrolled the user.
        const identity = await wallet.get(appUser);
        if (!identity) {
            console.log('An identity for the user "user2" does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return res.status(400).send({ 'result': `user not registered` });
        }

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();

        await gateway.connect(ccp, { wallet, identity: appUser, discovery: { enabled: true, asLocalhost: true } });

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork('mychannel');

        // Get the contract from the network.
        const contract = network.getContract('scott');
        var obj = {
            "adminId": req.body.adminId,
            "adminFirstName": req.body.adminFirstName,
            "adminLastName": req.body.adminLastName,
            "adminPhone": req.body.adminPhone,
            "adminEmail": req.body.adminEmail,
            "cts": req.body.cts,
            "dob": req.body.dob,
            "location": req.body.location,
            "password": req.body.password,
        }

        await contract.submitTransaction('updateAdministration', JSON.stringify(obj));

        // Disconnect from the gateway.
        console.log('Transaction has been submitted');
        await gateway.disconnect();
        return res.status(200).send({ 'result': `Transaction has been submitted` });

    } catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        return res.status(400).send({ 'result': `Failed to submit transaction: ${error}` });

        //process.exit(1);
    }
}
// // ** create claim **//
var createClaim = async function createClaim(req, res) {
    try {

        // Create a new file system based wallet for managing identities.
        const walletPath = path.resolve(walletDirPath);
        const wallet = await Wallets.newFileSystemWallet(walletPath);

        // Check to see if we've already enrolled the user.
        const identity = await wallet.get(appUser);
        if (!identity) {
            console.log('An identity for the user "user2" does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return res.status(400).send({ 'result': `user not registered` });
        }

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();

        await gateway.connect(ccp, { wallet, identity: appUser, discovery: { enabled: true, asLocalhost: true } });

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork('mychannel');

        // Get the contract from the network.
        const contract = network.getContract('scott');
        var obj = {
            "contractno": req.body.contractNo,
            "fname": req.body.fname,
            "lname": req.body.lname,
            "mobile": req.body.mobile,
            "dealer": req.body.dealer,
            "address": req.body.address,
            "emailid": req.body.emailid,
            "dateofloss": req.body.dateofloss,
            "mileage": req.body.mileage,
            "insurance": req.body.insurance,
        }

        await contract.submitTransaction('createClaim', JSON.stringify(obj));

        // Disconnect from the gateway.
        console.log('Transaction has been submitted');
        await gateway.disconnect();
        return res.status(200).send({ 'result': `Transaction has been submitted` });

    } catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        return res.status(400).send({ 'result': `Failed to submit transaction: ${error}` });

        //process.exit(1);
    }
}

// // ** create Pre paid Agreement**//
var createPrePaidAgreement = async function createPrePaidAgreement(req, res) {
    try {

        // Create a new file system based wallet for managing identities.
        const walletPath = path.resolve(walletDirPath);
        const wallet = await Wallets.newFileSystemWallet(walletPath);

        // Check to see if we've already enrolled the user.
        const identity = await wallet.get(appUser);
        if (!identity) {
            console.log('An identity for the user '+appUser+' does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return res.status(400).send({ 'result': `user not registered` });
        }

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();

        await gateway.connect(ccp, { wallet, identity: appUser, discovery: { enabled: true, asLocalhost: true } });

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork('mychannel');

        // Get the contract from the network.
        const contract = network.getContract('scott');
        var obj = {
            "ppid": req.body.ppid,
            "customername": req.body.customerName,
            "customeraddress": req.body.customerAddress,
            "vehicleyear": req.body.vehicleYear,
            "vehiclename": req.body.vehicleName,
            "vehiclemodel": req.body.vehicleModel,
            "odometer": req.body.odoMeter,
            "vinnumber": req.body.vinNumber,
            "customerphone": req.body.customerPhone,
            "dealersname": req.body.dealersName,
            "dealersaddress": req.body.dealersAddress,
            "dealersphone": req.body.dealersPhone,
        }

        await contract.submitTransaction('createPrePaidAgreement', JSON.stringify(obj));

        // Disconnect from the gateway.
        console.log('Transaction has been submitted');
        await gateway.disconnect();
        return res.status(200).send({ 'result': `Transaction has been submitted` });

    } catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        return res.status(400).send({ 'result': `Failed to submit transaction: ${error}` });

        //process.exit(1);
    }
}

// // ** create PaintlessRepair **//
var createPaintlessRepair = async function createPaintlessRepair(req, res) {
    try {

        // Create a new file system based wallet for managing identities.
        const walletPath = path.resolve(walletDirPath);
        const wallet = await Wallets.newFileSystemWallet(walletPath);

        // Check to see if we've already enrolled the user.
        const identity = await wallet.get(appUser);
        if (!identity) {
            console.log('An identity for the user "user2" does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return res.status(400).send({ 'result': `user not registered` });
        }

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();

        await gateway.connect(ccp, { wallet, identity: appUser, discovery: { enabled: true, asLocalhost: true } });

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork('mychannel');

        // Get the contract from the network.
        const contract = network.getContract('scott');
        var obj = {
            "prid": req.body.prid,
            "customername": req.body.customerName,
            "customeraddress": req.body.customerAddress,
            "vehicleyear": req.body.vehicleYear,
            "vehiclename": req.body.vehicleName,
            "vehiclemodel": req.body.vehicleModel,
            "mileage": req.body.mileage,
            "vinnumber": req.body.vinNumber,
            "customerphone": req.body.customerPhone,
            "dealersname": req.body.dealersName,
            "dealersaddress": req.body.dealersAddress,
            "dealersphone": req.body.dealersPhone,
            "vehicleprice": req.body.vehiclePrice,
            "term": req.body.term,
        }

        await contract.submitTransaction('createPaintlessRepair', JSON.stringify(obj));

        // Disconnect from the gateway.
        console.log('Transaction has been submitted');
        await gateway.disconnect();
        return res.status(200).send({ 'result': `Transaction has been submitted` });

    } catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        return res.status(400).send({ 'result': `Failed to submit transaction: ${error}` });

        //process.exit(1);
    }
}

// // ** Battery Agreement **//
var createBatteryAgreement = async function createBatteryAgreement(req, res) {
    try {

        // Create a new file system based wallet for managing identities.
        const walletPath = path.resolve(walletDirPath);
        const wallet = await Wallets.newFileSystemWallet(walletPath);

        // Check to see if we've already enrolled the user.
        const identity = await wallet.get(appUser);
        if (!identity) {
            console.log('An identity for the user "user2" does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return res.status(400).send({ 'result': `user not registered` });
        }

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();

        await gateway.connect(ccp, { wallet, identity: appUser, discovery: { enabled: true, asLocalhost: true } });

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork('mychannel');

        // Get the contract from the network.
        const contract = network.getContract('scott');
        var obj = {
            "bid": req.body.bid,
            "customername": req.body.customerName,
            "customeraddress": req.body.customerAddress,
            "vehicleyear": req.body.vehicleYear,
            "vehiclename": req.body.vehicleName,
            "vehiclemodel": req.body.vehicleModel,
            "mileage": req.body.mileage,
            "vinnumber": req.body.vinNumber,
            "customerphone": req.body.customerPhoneNo,
            "issuingloc": req.body.issuingLoc,
            "issuingid": req.body.issuingId,
            "iaddress": req.body.iAddress,
            "iphone": req.body.iPhone,
            "sassociate": req.body.sAssociate,
            "lienholder": req.body.lienHolder,
            "laddress": req.body.lAddress,
            "deductible": req.body.deductible,
            "term": req.body.term,
        }

        await contract.submitTransaction('createBatteryAgreement', JSON.stringify(obj));

        // Disconnect from the gateway.
        console.log('Transaction has been submitted');
        await gateway.disconnect();
        return res.status(200).send({ 'result': `Transaction has been submitted` });

    } catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        return res.status(400).send({ 'result': `Failed to submit transaction: ${error}` });

        //process.exit(1);
    }
}
// // ** CREATE SERVICE AGREEMENT **//
var createServiceAgreement = async function createServiceAgreement(req, res) {
    try {

        // Create a new file system based wallet for managing identities.
        const walletPath = path.resolve(walletDirPath);
        const wallet = await Wallets.newFileSystemWallet(walletPath);

        // Check to see if we've already enrolled the user.
        const identity = await wallet.get(appUser);
        if (!identity) {
            console.log('An identity for the user "user2" does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return res.status(400).send({ 'result': `user not registered` });
        }

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();

        await gateway.connect(ccp, { wallet, identity: appUser, discovery: { enabled: true, asLocalhost: true } });

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork('mychannel');

        // Get the contract from the network.
        const contract = network.getContract('scott');
        var obj = {
            "sid": req.body.sid,
            "vehicleyear": req.body.vehicleYear,
            "vehiclename": req.body.vehicleName,
            "vehiclemodel": req.body.vehicleModel,
            "vehiclemileage": req.body.vehicleMileage,
            "vinnumber": req.body.vinNumber,
            "vehicleprice": req.body.vehiclePrice,
            "customername": req.body.customerName,
            "customeraddress": req.body.customerAddress,
            "customerphone": req.body.customerPhoneNo,
            "dealersname": req.body.dealersName,
            "dealersaddress": req.body.dealersAddress,
            "dealersphone": req.body.dealersPhone,
            "lienholder": req.body.lienHolder,
            "laddress": req.body.lAddress,
            "agreementprice": req.body.agreementPrice,
            "agreementdate": req.body.agreementDate,
            "coverage": req.body.coverage,
            "deductible": req.body.deductible,
            "term": req.body.term,
        }

        await contract.submitTransaction('createServiceAgreement', JSON.stringify(obj));

        // Disconnect from the gateway.
        console.log('Transaction has been submitted');
        await gateway.disconnect();
        return res.status(200).send({ 'result': `Transaction has been submitted` });

    } catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        return res.status(400).send({ 'result': `Failed to submit transaction: ${error}` });

        //process.exit(1);
    }
}

var createRefundRequest = async function createRefundRequest(req, res) {
    try {

        // Create a new file system based wallet for managing identities.
        const walletPath = path.resolve(walletDirPath);
        const wallet = await Wallets.newFileSystemWallet(walletPath);

        // Check to see if we've already enrolled the user.
        const identity = await wallet.get(appUser);
        if (!identity) {
            console.log('An identity for the user "user2" does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return res.status(400).send({ 'result': `user not registered` });
        }

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();

        await gateway.connect(ccp, { wallet, identity: appUser, discovery: { enabled: true, asLocalhost: true } });

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork('mychannel');

        // Get the contract from the network.
        const contract = network.getContract('scott');
        var obj = {
            "rid": req.body.rid,
            "contractno": req.body.contractNo,
            "cname": req.body.customerName,
            "effectivedate": req.body.effectiveDate,
            "starttime": req.body.startTime,
            "endtime": req.body.endTime,
            "purchaseprice": req.body.agreementPurchasePrice,
            "canceldate": req.body.cancelDate,
            "approvalfromAdmin": req.body.approvalFromAdmin,
            "id": req.body.id,
            "years": req.body.years,
            "presentmonth": req.body.presentMonth,
            
        }
        console.log(obj)

        var result =await contract.submitTransaction('createRefundRequest', JSON.stringify(obj));

        // Disconnect from the gateway.
        console.log('Transaction has been submitted');
        await gateway.disconnect();
        return res.status(200).send({ 'result': `Transaction has been submitted`, msg:JSON.parse(result) });

    } catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        return res.status(400).send({ 'result': `Failed to submit transaction: ${error}` });

        //process.exit(1);
    }
}

module.exports = {
    createuser,
    updateUser,
    createAdministration,
    updateAdministration,
    createPrePaidAgreement,
    createPaintlessRepair,
    createBatteryAgreement,
    createServiceAgreement,
    createClaim,
    createRefundRequest
}













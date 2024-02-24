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

const { Gateway, Wallets } = require('fabric-network');
const path = require('path');
const fs = require('fs');

const ccpPath = path.resolve(connectionProfilePath);
const ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));


// ** query brand **//
var queryUser = async function queryUser(req,res) {
    try {

// Create a new file system based wallet for managing identities.
  const walletPath = path.resolve(walletDirPath);
  const wallet = await Wallets.newFileSystemWallet(walletPath);
  console.log(`Wallet path: ${walletPath}`);

  // Check to see if we've already enrolled the user.
  const identity = await wallet.get(appUser);
  console.log("identity:", identity);

  if (!identity) {
      console.log('An identity for the '+appUser+'  does not exist in the wallet');
      console.log('Run the registerUser.js application before retrying');
      return res.status(400).send({'result':`user not registered`});
  }
 
  // Create a new gateway for connecting to our peer node.
  const gateway = new Gateway();
  console.log(`Wallet path: `,gateway,wallet);

  await gateway.connect(ccp, { wallet, identity: appUser, discovery: { enabled: true, asLocalhost: true } });
  console.log("ccp*******",ccp);
  // Get the network (channel) our contract is deployed to.
  const network = await gateway.getNetwork(channelName);

        // Get the contract from the network.
        const contract = network.getContract(chaincodeName);


        // Evaluate the specified transaction.
        // queryCar transaction - requires 1 argument, ex: ('queryCar', 'CAR4')
        // queryAllCars transaction - requires no arguments, ex: ('queryAllCars')
        const result = await contract.evaluateTransaction("queryUser",req.body.Id);
        console.log(`Transaction has been evaluated, result is: ${result.toString()}`);
        return res.status(200).send({'result': `Transaction has been evaluated, result is: ${result.toString()}`});

    }catch(error) {
        console.error(`Failed to evaluate transaction: ${error}`);
        return res.status(400).send({'result':`Failed to submit transaction: ${error}`});
    }
}

// ** query cluster **//
var queryClaim=async function queryClaim(req,res) {
    try {

  // Create a new file system based wallet for managing identities.
  const walletPath = path.resolve(walletDirPath);
  const wallet = await Wallets.newFileSystemWallet(walletPath);
  console.log(`Wallet path: ${walletPath}`);

  // Check to see if we've already enrolled the user.
  const identity = await wallet.get(appUser);
  console.log("identity:", identity);

  if (!identity) {
      console.log('An identity for the '+appUser+'  does not exist in the wallet');
      console.log('Run the registerUser.js application before retrying');
      return res.status(400).send({'result':`user not registered`});
  }
 
  // Create a new gateway for connecting to our peer node.
  const gateway = new Gateway();
  console.log(`Wallet path: `,gateway,wallet);

  await gateway.connect(ccp, { wallet, identity: appUser, discovery: { enabled: true, asLocalhost: true } });
  console.log("ccp*******",ccp);
  // Get the network (channel) our contract is deployed to.
  const network = await gateway.getNetwork(channelName);

        // Get the contract from the network.
        const contract = network.getContract(chaincodeName);

        // Evaluate the specified transaction.
        // queryCar transaction - requires 1 argument, ex: ('queryCar', 'CAR4')
        // queryAllCars transaction - requires no arguments, ex: ('queryAllCars')
        const result = await contract.evaluateTransaction("queryClaim",req.body.Id);
        console.log(`Transaction has been evaluated, result is: ${result.toString()}`);
        return res.status(200).send({'result': `Transaction has been evaluated, result is: ${result.toString()}`});

    } catch (error) {
        console.error(`Failed to evaluate transaction: ${error}`);
        return res.status(400).send({'result':`Failed to submit transaction: ${error}`});
    }
}

//** query Admin */
var queryAdministration=async function queryAdmin(req,res) {
    try {

         // Create a new file system based wallet for managing identities.
         const walletPath = path.resolve(walletDirPath);
         const wallet = await Wallets.newFileSystemWallet(walletPath);
         console.log(`Wallet path: ${walletPath}`);
       
         // Check to see if we've already enrolled the user.
         const identity = await wallet.get(appUser);
         console.log("identity:", identity);
       
         if (!identity) {
             console.log('An identity for the '+appUser+'  does not exist in the wallet');
             console.log('Run the registerUser.js application before retrying');
             return res.status(400).send({'result':`user not registered`});
         }
        
         // Create a new gateway for connecting to our peer node.
         const gateway = new Gateway();
         console.log(`Wallet path: `,gateway,wallet);
       
         await gateway.connect(ccp, { wallet, identity: appUser, discovery: { enabled: true, asLocalhost: true } });
         console.log("ccp*******",ccp);
         // Get the network (channel) our contract is deployed to.
         const network = await gateway.getNetwork(channelName);
       
               // Get the contract from the network.
               const contract = network.getContract(chaincodeName);
       

        // Evaluate the specified transaction.
        // queryCar transaction - requires 1 argument, ex: ('queryCar', 'CAR4')
        // queryAllCars transaction - requires no arguments, ex: ('queryAllCars',A00001)
        const result = await contract.evaluateTransaction("queryAdministration",req.body.Id);
        console.log(`Transaction has been evaluated, result is: ${result.toString()}`);
        return res.status(200).send({'result': `Transaction has been evaluated, result is: ${result.toString()}`});
    } catch (error) {
        console.error(`Failed to evaluate transaction: ${error}`)
        return res.status(400).send({'result':`Failed to submit transaction: ${error}`});
    }
}

//** User Login */
var userLogin=async function userLogin(req,res) {
    try {

         // Create a new file system based wallet for managing identities.
         const walletPath = path.resolve(walletDirPath);
         const wallet = await Wallets.newFileSystemWallet(walletPath);
         console.log(`Wallet path: ${walletPath}`);
       
         // Check to see if we've already enrolled the user.
         const identity = await wallet.get(appUser);
         console.log("identity:", identity);
       
         if (!identity) {
             console.log('An identity for the '+appUser+'  does not exist in the wallet');
             console.log('Run the registerUser.js application before retrying');
             return res.status(400).send({'result':`user not registered`});
         }
        
         // Create a new gateway for connecting to our peer node.
         const gateway = new Gateway();
         console.log(`Wallet path: `,gateway,wallet);
       
         await gateway.connect(ccp, { wallet, identity: appUser, discovery: { enabled: true, asLocalhost: true } });
         console.log("ccp*******",ccp);
         // Get the network (channel) our contract is deployed to.
         const network = await gateway.getNetwork(channelName);
       
               // Get the contract from the network.
               const contract = network.getContract(chaincodeName);
               var obj = {
               
                "mobile": req.body.mobile,
                "password": req.body.password,
                "vinnumber": req.body.vinnumber,           
            }

        // Evaluate the specified transaction.
        // queryCar transaction - requires 1 argument, ex: ('queryCar', 'CAR4')
        // queryAllCars transaction - requires no arguments, ex: ('queryAllCars',A00001)
        const result = await contract.evaluateTransaction("userLogin",JSON.stringify(obj));
        console.log(`Transaction has been evaluated, result is: ${result.toString()}`);
        return res.status(200).send({'result': `Transaction has been evaluated, result is: ${result.toString()}`});
    } catch (error) {
        console.error(`Failed to evaluate transaction: ${error}`)
        return res.status(400).send({'result':`Failed to submit transaction: ${error}`});
    }
}
//** findByNumberLastName */
var findByNumberLastName=async function findByNumberLastName(req,res) {
    try {

         // Create a new file system based wallet for managing identities.
         const walletPath = path.resolve(walletDirPath);
         const wallet = await Wallets.newFileSystemWallet(walletPath);
         console.log(`Wallet path: ${walletPath}`);
       
         // Check to see if we've already enrolled the user.
         const identity = await wallet.get(appUser);
         console.log("identity:", identity);
       
         if (!identity) {
             console.log('An identity for the '+appUser+'  does not exist in the wallet');
             console.log('Run the registerUser.js application before retrying');
             return res.status(400).send({'result':`user not registered`});
         }
        
         // Create a new gateway for connecting to our peer node.
         const gateway = new Gateway();
         console.log(`Wallet path: `,gateway,wallet);
       
         await gateway.connect(ccp, { wallet, identity: appUser, discovery: { enabled: true, asLocalhost: true } });
         console.log("ccp*******",ccp);
         // Get the network (channel) our contract is deployed to.
         const network = await gateway.getNetwork(channelName);
       
               // Get the contract from the network.
               const contract = network.getContract(chaincodeName);
              
        // Evaluate the specified transaction.
        console.log("body values",req.body.mnumber,req.body.lname)
        const result = await contract.evaluateTransaction("findByNumberLastName",req.body.mnumber,req.body.lname);
        console.log(`Transaction has been evaluated, result is: ${result.toString()}`);
        return res.status(200).send({'result': `Transaction has been evaluated, result is: ${result.toString()}`});
    } catch (error) {
        console.error(`Failed to evaluate transaction: ${error}`)
        return res.status(400).send({'result':`Failed to submit transaction: ${error}`});
        // process.exit(1);
    }
}
//** findByContractLastName */
var findByContractLastName=async function findByContractLastName(req,res) {
    try {

         // Create a new file system based wallet for managing identities.
         const walletPath = path.resolve(walletDirPath);
         const wallet = await Wallets.newFileSystemWallet(walletPath);
         console.log(`Wallet path: ${walletPath}`);
       
         // Check to see if we've already enrolled the user.
         const identity = await wallet.get(appUser);
         console.log("identity:", identity);
       
         if (!identity) {
             console.log('An identity for the '+appUser+'  does not exist in the wallet');
             console.log('Run the registerUser.js application before retrying');
             return res.status(400).send({'result':`user not registered`});
         }
        
         // Create a new gateway for connecting to our peer node.
         const gateway = new Gateway();
         console.log(`Wallet path: `,gateway,wallet);
       
         await gateway.connect(ccp, { wallet, identity: appUser, discovery: { enabled: true, asLocalhost: true } });
         console.log("ccp*******",ccp);
         // Get the network (channel) our contract is deployed to.
         const network = await gateway.getNetwork(channelName);
       
               // Get the contract from the network.
               const contract = network.getContract(chaincodeName);
              
        // Evaluate the specified transaction.

        const result = await contract.evaluateTransaction("findByContractLastName",req.body.cnumber,req.body.lname);
        console.log(`Transaction has been evaluated, result is: ${result.toString()}`);
        return res.status(200).send({'result': `Transaction has been evaluated, result is: ${result.toString()}`});
    } catch (error) {
        console.error(`Failed to evaluate transaction: ${error}`)
        return res.status(400).send({'result':`Failed to submit transaction: ${error}`});
    }
}
module.exports ={    
    queryUser,
    queryClaim,
    queryAdministration,
    userLogin,
    findByNumberLastName,
    findByContractLastName,

}

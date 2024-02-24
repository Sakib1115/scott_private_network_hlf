/*
 * SPDX-License-Identifier: Apache-2.0
 */

'use strict';

const config = require("../config/hlf_network.json")

// const { FileSystemWallet, Gateway, X509WalletMixin } = require('fabric-network');
const { Wallets } = require('fabric-network');
const FabricCAServices = require('fabric-ca-client');
const fs = require('fs');
const path = require('path');

const connectionProfilePath = config.connectionProfilePath;
const walletDirPath = config.walletPath;

const ccpPath = path.resolve(connectionProfilePath);
const ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));

var registerUserApi=async function registerUserApi(req,res) {
    try {
         console.log("req.body.UserId***************", req.body.UserId) 
        // Create a new file system based wallet for managing identities.
        
        const caURL = ccp.certificateAuthorities['ca.org1.example.com'].url;
        const ca = new FabricCAServices(caURL);
        
        // Create a new file system based wallet for managing identities.
		const walletPath = path.resolve(walletDirPath);
        const wallet = await Wallets.newFileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        // Check to see if we've already enrolled the user.
        const userIdentity = await wallet.get(req.body.UserId);
        if (userIdentity) {
            console.log('An identity for the user  already exists in the wallet');
            // return;
          return res.status(200).send({'result':`An identity for the user  already exists in the wallet`});
        }

        // Check to see if we've already enrolled the admin user.
        const adminIdentity = await wallet.get('admin');
        if (!adminIdentity) {
            console.log('An identity for the admin user "admin" does not exist in the wallet');
            console.log('Run the enrollAdmin.js application before retrying');
            // return;
          return res.status(200).send({'result':`An identity for the admin user "admin" does not exist in the wallet run enrolladmin`});
        }

        // build a user object for authenticating with the CA
        const provider = wallet.getProviderRegistry().getProvider(adminIdentity.type);
        const adminUser = await provider.getUserContext(adminIdentity, 'admin');

        // Register the user, enroll the user, and import the new identity into the wallet.
        const secret = await ca.register({
            affiliation: 'org1.department1',
            enrollmentID: req.body.UserId,
            role: 'client'
        }, adminUser);
        const enrollment = await ca.enroll({
            enrollmentID: req.body.UserId,
            enrollmentSecret: secret
        });
        const x509Identity = {
            credentials: {
                certificate: enrollment.certificate,
                privateKey: enrollment.key.toBytes(),
            },
            mspId: 'Org1MSP',
            type: 'X.509',
        };
        await wallet.put(req.body.UserId, x509Identity);
        console.log(`Successfully registered and enrolled admin user ${req.body.UserId} and imported it into the wallet`);
     return res.status(200).send({'result':`Successfully register user`});
    } catch (error) {
        console.error(`Failed to register user ${req.body.UserId}: ${error}`);
        // process.exit(1);
      return res.status(400).send({'result':`Failed to register user:${error}`});
    }
}

module.exports = {
    registerUserApi
}

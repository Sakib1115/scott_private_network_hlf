var express = require('express');
var router = express.Router();
var api = require('../controller/query');
console.log("inside routes/query.js");
/* GET home page. */
router.post('/queryClaim', api.queryClaim);  // t

router.post('/queryUser',api.queryUser);

router.post('/queryAdmin',api.queryAdministration);
router.post('/userLogin',api.userLogin);
router.post('/findByNumberLastName',api.findByNumberLastName);
router.post('/findByContractLastName',api.findByContractLastName);

module.exports = router;
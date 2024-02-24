var express = require('express');
var router = express.Router();
var api = require('../controller/invoke');
/* GET home page. */

router.post('/createuser', api.createuser);
router.post('/updateuser', api.updateUser);
router.post('/createAdministration', api.createAdministration);
router.post('/updateAdministration', api.updateAdministration);
router.post('/createPrePaidAgreement', api.createPrePaidAgreement);
router.post('/createPaintlessRepair', api.createPaintlessRepair);
router.post('/createBatteryAgreement', api.createBatteryAgreement);
router.post('/createServiceAgreement', api.createServiceAgreement);
router.post('/createClaim', api.createClaim);
router.post('/createRefundRequest', api.createRefundRequest);




module.exports = router;








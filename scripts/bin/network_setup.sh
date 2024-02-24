# ! /bin/bash
##########################################
# Script performs docker based setup of 
# Hyperledger fabric network with couchdb,
# fabric ca ,ordere and peer service
############################################

function clearContainers() {
  CONTAINER_IDS=$(docker ps -a | awk '($2 ~ /dev-peer.*/) {print $1}')
  if [ -z "$CONTAINER_IDS" -o "$CONTAINER_IDS" == " " ]; then
    echo "No containers available for deletion"
  else
    docker rm -f $CONTAINER_IDS
  fi

  CONTAINER_IDS=$(docker ps -a | awk '($2 ~ /logspout.*/) {print $1}')
  if [ -z "$CONTAINER_IDS" -o "$CONTAINER_IDS" == " " ]; then
    echo "No containers available for deletion"
  else
    docker rm -f $CONTAINER_IDS
  fi

}
#Bring down docker containers
function networkDown() {
	echo "Bring down docker containers"
	
	export IMAGETAG=latest
	docker-compose -f scott-network-ca-docker.yaml down --volumes --remove-orphans
	rc=$?
	
	if [[ $rc -ne 0 ]];then
		echo "Terminating process"
		exit 1
	fi
	clearContainers
	docker ps -a
}

function couchDbDown() {
	echo "Bring down docker containers"
	
	export IMAGETAG=latest
	docker-compose -f docker-compose-couch.yaml down --volumes --remove-orphans
	rc=$?
	
	if [[ $rc -ne 0 ]];then
		echo "Terminating process"
		exit 1
	fi
	clearContainers
	docker ps -a
}

function bringDownFabricCA() {

                CA_IMAGETAG="latest"
                COMPOSE_FILE_CA=docker-compose-couch.yaml
                IMAGE_TAG=${CA_IMAGETAG} docker-compose -f $COMPOSE_FILE_CA down --volumes --remove-orphans 2>&1

                CONTAINER_IDS=$(docker ps -a | awk '($2 ~ /fabric-ca.*/) {print $1}')
                if [ -z "$CONTAINER_IDS" -o "$CONTAINER_IDS" == " " ]; then
                       echo "No containers available for deletion"
                else
                       docker rm -f $CONTAINER_IDS
                fi
}

function networkSetUp(){
    
	
       
      # docker-compose -f scott-network-ca-docker.yaml down --volume
      # docker-compose -f scott-network-docker.yaml down --volume
      # docker-compose -f docker-compose-couch.yaml down --volume

       echo "Bringing up docker container of couch db"
       docker-compose -f docker-compose-couch.yaml up -d 
       echo "Bringing up of docker container of fabric ca"
       docker-compose -f scott-network-ca-docker.yaml up -d

       cd $BASH_DIR/scripts/bin

#-------------------------------------------------------------------------------------------------------------------------------------------------------------------------
#Generate Config
       echo "Generating NodeOU Config file for orderer"
       ./mspUtils.sh createNodeOUConfig -n ca-orderer -f $BASH_DIR/scripts/conf/test-net-msp-orderer01.cfg -u admin -p adminpw  -o $BASH_DIR/docker/network/ordererOrganizations/scott.com/msp
       echo "Generating NodeOU Config file for peer"
       ./mspUtils.sh createNodeOUConfig -n ca-scott -f $BASH_DIR/scripts/conf/test-net-msp-peer0.cfg -o $BASH_DIR/docker/network/peerOrganizations/hlf.scott.com/msp ; rc=$? ; echo $rc

#------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
#orderer01 command
       echo "Generating MSP and TLS Profile of orderer"
       ./mspUtils.sh generateMsp -n ca-orderer -i orderer1 -s orderer1pw -t orderer -f $BASH_DIR/scripts/conf/test-net-msp-orderer01.cfg -u admin -p adminpw  -o $BASH_DIR/docker/network/ordererOrganizations/scott.com/orderers/orderer01.scott.com/msp
 
       ./mspUtils.sh generateTlsCerts -n ca-orderer -i orderer1 -s orderer1pw -t orderer -f $BASH_DIR/scripts/conf/test-net-msp-orderer01.cfg -u admin -p adminpw -o $BASH_DIR/docker/network/ordererOrganizations/scott.com/orderers/orderer01.scott.com/tls
 
# -----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
#peer0 command
       echo "Generating MSP and TLS Profile of Peer0"
       ./mspUtils.sh generateMsp -n ca-scott -i peer0 -s peer0pw -t peer -f $BASH_DIR/scripts/conf/test-net-msp-peer0.cfg -u admin -p adminpw  -o $BASH_DIR/docker/network/peerOrganizations/hlf.scott.com/peers/peer0.hlf.scott.com/msp
   
       ./mspUtils.sh generateTlsCerts -n ca-scott -i peer0 -s peer0pw -t peer -f $BASH_DIR/scripts/conf/test-net-msp-peer0.cfg -u admin -p adminpw -o $BASH_DIR/docker/network/peerOrganizations/hlf.scott.com/peers/peer0.hlf.scott.com/tls
    
#------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
#peer1 command
       echo "Generating MSP and TLS Profile of Peer1"
       ./mspUtils.sh generateMsp -n ca-scott -i peer1 -s peer1pw -t peer -f $BASH_DIR/scripts/conf/test-net-msp-peer1.cfg -u admin -p adminpw  -o $BASH_DIR/docker/network/peerOrganizations/hlf.scott.com/peers/peer1.hlf.scott.com/msp
   
       ./mspUtils.sh generateTlsCerts -n ca-scott -i peer1 -s peer1pw -t peer -f $BASH_DIR/scripts/conf/test-net-msp-peer1.cfg -u admin -p adminpw -o $BASH_DIR/docker/network/peerOrganizations/hlf.scott.com/peers/peer1.hlf.scott.com/tls


#--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
#orderer
       echo "Copy TLS Certs and rename tls keystore keys for Orderer"
       cp $BASH_DIR/docker/network/ordererOrganizations/scott.com/msp/config.yaml $BASH_DIR/docker/network/ordererOrganizations/scott.com/orderers/orderer01.scott.com/msp

       cp -r $BASH_DIR/docker/network/ordererOrganizations/scott.com/orderers/orderer01.scott.com/tls/tlscacerts $BASH_DIR/docker/network/ordererOrganizations/scott.com/msp

       mv $BASH_DIR/docker/network/ordererOrganizations/scott.com/orderers/orderer01.scott.com/tls/keystore/*_sk $BASH_DIR/docker/network/ordererOrganizations/scott.com/orderers/orderer01.scott.com/tls/keystore/priv_key

#---------------------------------------------------------------------------------------------------------------------------------------------------------------------------
#peer0
       echo "Copy TLS Certs and rename tls keystore keys for Peer0"
         cp $BASH_DIR/docker/network/peerOrganizations/hlf.scott.com/msp/config.yaml $BASH_DIR/docker/network/peerOrganizations/hlf.scott.com/peers/peer0.hlf.scott.com/msp

         cp -r $BASH_DIR/docker/network/peerOrganizations/hlf.scott.com/peers/peer0.hlf.scott.com/tls/tlscacerts $BASH_DIR/docker/network/peerOrganizations/hlf.scott.com/msp

         mv $BASH_DIR/docker/network/peerOrganizations/hlf.scott.com/peers/peer0.hlf.scott.com/tls/keystore/*_sk $BASH_DIR/docker/network/peerOrganizations/hlf.scott.com/peers/peer0.hlf.scott.com/tls/keystore/priv_key

#-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
#peer1
       echo "Copy TLS Certs and rename tls keystore keys for Peer1"
        cp $BASH_DIR/docker/network/peerOrganizations/hlf.scott.com/msp/config.yaml $BASH_DIR/docker/network/peerOrganizations/hlf.scott.com/peers/peer1.hlf.scott.com/msp

        mv $BASH_DIR/docker/network/peerOrganizations/hlf.scott.com/peers/peer1.hlf.scott.com/tls/keystore/*_sk $BASH_DIR/docker/network/peerOrganizations/hlf.scott.com/peers/peer1.hlf.scott.com/tls/keystore/priv_key

#-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

#genesis block command
       echo "Generating Genesis Block"
       ./genesisUtils.sh generateGenesis -n ScottOrdererGenesis -c system-channel -o $BASH_DIR/docker/network/ordererOrganizations/scott.com/system-genesis-block/genesis.block  -f $BASH_DIR/scripts/conf/test-net-genesis.cfg


#------------------------------------------------------------------------------------------------------------------------------------------------------------------    
       echo "Starting Docker container for Orderer and Peers"
        cd $BASH_DIR/docker/compose-files

         docker-compose -f scott-network-docker.yaml up -d

}

echo "Setting up network"
CHANNEL_CFG=$1

if [ -e  $CHANNEL_CFG ]
then
	echo "Loading config file $CHANNEL_CFG"
else
	echo "config file $CHANNEL_CFG not exist"
fi
. $CHANNEL_CFG
cd $BASH_DIR/docker/compose-files
echo "Bringing down docker containers"
bringDownFabricCA
couchDbDown
clearContainers

echo "Bringing up network docker containers"
networkSetUp

echo "Docker Container setups done"


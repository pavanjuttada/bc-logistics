##############################################################
################# Tree Structure #########################
##############################################################
Directory Structure for Node App(UI,Servies,NodeSDK Client)::

logistics
├── app.js
├── enrollAdmin.js
├── hfc-key-store
│   ├── 44077c8015ddc7277ab8f2c86cf10974221dca492edd4417e9ae37701281b92a-priv
│   ├── 44077c8015ddc7277ab8f2c86cf10974221dca492edd4417e9ae37701281b92a-pub
│   ├── 5a0b5420ee2dfa727721578373ccd5778667dbe53b54eeffce0eb9be4091e634-priv
│   ├── 5a0b5420ee2dfa727721578373ccd5778667dbe53b54eeffce0eb9be4091e634-pub
│   ├── admin
│   └── user1
├── invoke.js
├── node_modules
├── package.json
├── package-lock.json
├── public
│   ├── css
│   │   └── bootstrap.min.css
│   ├── favicon.ico
│   ├── html
│   │   ├── index.html
│   │   └── login.html
│   └── js
│       ├── angular.js
│       ├── angular.min.js
│       ├── core.js
│       └── jquery-2.2.4.min.js
├── query.js
├── registerUser.js
├── routes
│   └── logistics.js
└── startFabric.sh

Directory Structure for Chaincode::

chaincode
    └── logistics
           └── go
               ├── logistics.go
               └── ReadMe.md

##############################################################
## Setup Network Hyperledger Fabric Logistics
##############################################################

Go to Basic network
Step1: Clean the network & docker contatiners

    root@integra-Aspire-E5-573:~/workspace/pavan-fabric/basic-network#
        #./stop.sh
        #docker network prune
        #./teardown.sh

    check any container are running:  docker ps -aq

Step2:Start the network from the logistics directory

    root@integra-Aspire-E5-573:~/workspace/pavan-fabric/logistics# ./startFabric.sh 

Step3: Install node_modules, Enroll admin & Register User
    
    root@integra-Aspire-E5-573:~/workspace/pavan-fabric/logistics# npm install
    
    root@integra-Aspire-E5-573:~/workspace/pavan-fabric/logistics#  node enrollAdmin.js
    root@integra-Aspire-E5-573:~/workspace/pavan-fabric/logistics#  node registerUser.js    

Step4: Start node application
    node app.js


##############################################################
## UI & Services
##############################################################

## UI ::

URL : http://localhost:4000/home

Assumptions : 
            Assumptions created 3 diffrent users & with id's (Seller-111,Buyer-222,LogisticProvider-333)


Step1: Select Login As Seller

Step2: Create Shipment
            With BuyerId: 222 & Logistic Provider: 333
Step3: Check the shipment status of recently added record & move it for shipping
            Temprature will be recorder from now onwords into the blocchain till it delevers.
            NOTE: This feature is not enabled as of now.
Step4: Now switch login user to Logistic Provider
Step5: Check the shipment status of recently added record & delever it.
        
Step6: Now switch login user to Buyer.
        Buyer can take a action Accept or Reject. (This decision taken based on temprature checking w.r.t this shipment, but this not enabled as of now).

Step7: Once Buyer is Accepted/Rejected, the Seller & Logistic Provder will get to now  about it. we check it by login in with corresponding accounts.
########################################UI END ###################################################

Node Webservices:
    
     #1 AddShipment Service:
        url==> localhost:4000/api/newShipment
        request ==> {
                        "operationType": "addShipment",
                        "Shipment_Name":"Shipment12",
                        "Status": "store",
                        "SellerId": "111",
                        "SellerName": "Pavan",
                        "SellerLocation": "Banglore",
                        "BuyerId": "222",
                        "BuyerName": "Vijay",
                        "BuyerLocation": "Chennai",
                        "Logisticprovider": "333"
                    }

     #2 Service to get Shipments of Seller:
        url==> localhost:4000/api/getShipments
        request ==> {
                     "queryName": "queryShipmentsForSeller",
                     "queryParam": "111"
                    }

        Other similar services:
            queryName: queryShipmentsForBuyer/queryShipmentsForLogisticprovider
            queryParam: 222/333


    #3 Seller sends the shipment to Shipment provider:
        url==> localhost:4000/api/updateShipment                    
        request ==> {
                        "operationType": "changeShipmentStatus",
                        "existing_tx_id": "TRANSACTION_ID_OF_PREVIOUSLY_CREATED_SHIPMENT",
                        "new_status": "intransit",
                        "user_id": "111"
                    }

        Other similar services:
            new_status: delivered/accepted/rejected
            user_id: 333/222/222

######################################## Node Services END ##############################################


########################################
Network Status ::
########################################

CONTAINER ID        IMAGE                                                                                                       COMMAND                  CREATED             STATUS              PORTS                                            NAMES
4595759e5f2e        dev-peer0.org1.example.com-logistics-1.0-ab4d04d4cb94b4c693fa18c29d361fd4fc4f16d0e71cb84d84489dc30f1e88b6   "chaincode -peer.add…"   3 hours ago         Up 3 hours                                                           dev-peer0.org1.example.com-logistics-1.0
c45ecaf579d3        hyperledger/fabric-tools                                                                                    "/bin/bash"              3 hours ago         Up 3 hours                                                           cli
54143807e532        hyperledger/fabric-peer                                                                                     "peer node start"        3 hours ago         Up 3 hours          0.0.0.0:7051->7051/tcp, 0.0.0.0:7053->7053/tcp   peer0.org1.example.com
8c0498cecc1c        hyperledger/fabric-couchdb                                                                                  "tini -- /docker-ent…"   3 hours ago         Up 3 hours          4369/tcp, 9100/tcp, 0.0.0.0:5984->5984/tcp       couchdb
20c523508c71        hyperledger/fabric-orderer                                                                                  "orderer"                3 hours ago         Up 3 hours          0.0.0.0:7050->7050/tcp                           orderer.example.com
c11159e9aa74        hyperledger/fabric-ca                                                                                       "sh -c 'fabric-ca-se…"   3 hours ago         Up 3 hours          0.0.0.0:7054->7054/tcp                           ca.example.com


NOTE: This is basic test network is only for development.
    In the production it will be having 3 Organizations & the 3 users(Seller,Buyer,Shipment Provider) will be part of seprate organization. And each Organizaion contains 2 peers & 1 MSP. will use kafaka as Ordering service node. A Smart device is integrated for temprature readings.

######################################## END ###################################################
Contact::
        Author: Pavan Kumar J
        MailId: pavanjuttada@gmail.com
        Mobile: 91-6363117217
######################################## END ###################################################
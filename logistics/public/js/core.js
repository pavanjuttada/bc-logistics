// public/core.js
var logisticsApp = angular.module('logisticsApp', []);

function mainController($scope, $http) {

    // seting the global configurations, with 3 default users.
    $scope.users = [{id: 111, name: "Seller"},
                    {id: 222, name: "Buyer"},
                    {id: 333, name: "Logisticprovider"}];
    $scope.reqQuery = {
                        "queryName": "queryShipmentsForSeller",
                        "queryParam": "111"
                        };

    $scope.formData = {"operationType": "addShipment",
                        "Status": "store",
                        "SellerId": "111",
                        "SellerName": "Pavan",
                        "SellerLocation": "Banglore",
                        };

    $scope.changeStatusReq = {
                        "operationType": "changeShipmentStatus",
                        "existing_tx_id": "",
                        "new_status": "intransit",
                        "user_id": "111"
                        };

    // when landing on the page, get intial configurations.
    $scope.init = function () {
        console.log("Login User is :: "+$scope.user_id);
        if($scope.user_id==111){
            $scope.reqQuery.queryName="queryShipmentsForSeller";
            $scope.reqQuery.queryParam=$scope.user_id;

            $scope.formData.SellerId=$scope.user_id;
            $scope.formData.SellerName="Pavan";
            $scope.formData.SellerLocation="Banglore";
            $scope.formData.Status="store";

            $scope.changeStatusReq.user_id=$scope.user_id;
        }else if($scope.user_id==222){
            $scope.reqQuery.queryName="queryShipmentsForBuyer";
            $scope.reqQuery.queryParam=$scope.user_id;

            $scope.formData.SellerId=$scope.user_id;
            $scope.formData.SellerName="Pavan";
            $scope.formData.SellerLocation="Banglore";
            $scope.formData.Status=$scope.user_id;

            $scope.changeStatusReq.user_id=$scope.user_id;
        }else if($scope.user_id==333){
            $scope.reqQuery.queryName="queryShipmentsForLogisticprovider";
            $scope.reqQuery.queryParam=$scope.user_id;

            $scope.changeStatusReq.user_id=$scope.user_id;
        }

        $http.post('/api/getShipments', $scope.reqQuery)
            .success(function(data) {
                console.log(data);
                $scope.todos = data.result;
            })
            .error(function(data) {
                console.log('Error: ' + data);
            });
    };
    
    // when submitting create shipment form calling the node API service
    $scope.createTodo = function() {
        console.log($scope.formData);
        $http.post('/api/newShipment', $scope.formData)
            .success(function(data) {
                console.log(data);
                $scope.todos.push({Key: data.tid, Record: {shipment_name: $scope.formData.Shipment_Name,
                                    logistic_provider_id: $scope.formData.Logisticprovider,
                                    id: data.tid,
                                    seller_name: $scope.formData.SellerName,
                                    seller_location: $scope.formData.SellerLocation,
                                    buyer_name: $scope.formData.BuyerName,
                                    buyer_location: $scope.formData.BuyerLocation,
                                    status: $scope.formData.Status}
                                });

                $scope.formData.Shipment_Name = "";
                $scope.formData.BuyerId = "";
                $scope.formData.BuyerName = "";
                $scope.formData.BuyerLocation = "";
                $scope.formData.Logisticprovider = ""
            })
            .error(function(data) {
                console.log('Error: ' + data);
            });
    };

    $scope.changeStatus = function(id,status) {
        $scope.changeStatusReq.existing_tx_id = id;
        console.log("id :: "+id);
        console.log("status :: "+status);
        $scope.changeStatusReq.new_status=status;
        $http.post('/api/updateShipment', $scope.changeStatusReq)
            .success(function(data) {
                angular.forEach($scope.todos, function(record) {
                        if(record.Key == id){
                            $scope.todos[$scope.todos.indexOf(record)].Record.status = status;
                        }
                })
            })
            .error(function(data) {
                console.log('Error: ' + data);
            });
    };
}
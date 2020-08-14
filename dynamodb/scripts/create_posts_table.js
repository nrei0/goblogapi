const AWS = require('aws-sdk');

AWS.config.update({region: 'eu-central-1'});

var dynamodb = new AWS.DynamoDB({apiVersion: '2012-08-10', endpoint: 'http://localhost:8000'});

var params = {
    TableName: 'Posts',
    KeySchema: [
        {
            AttributeName: 'ID',
            KeyType: 'HASH',
        }
    ],
    AttributeDefinitions: [
        {
            AttributeName: 'ID',
            AttributeType: 'S',
        }
    ],
    ProvisionedThroughput: {
        ReadCapacityUnits: 1, 
        WriteCapacityUnits: 1, 
    },
};
dynamodb.createTable(params, function(err, data) {
    if (err) console.error(err);
    else console.log('Table created', data);
});
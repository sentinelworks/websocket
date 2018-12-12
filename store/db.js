/*
var clients = [
    { date: '12/1/2011', reading: 3, id: 20055 },
    { date: '13/1/2011', reading: 5, id: 20053 },
    { date: '14/1/2011', reading: 6, id: 45652 }
];
*/

var clients = {};
var conn;
var infosight;

var addClient = function (conn) {
    clients.push(conn);
}

function getClient()
{
    return clients[0];
}

module.exports = clients;
module.exports = infosight;
module.exports = conn;
module.exports = addClient;
module.exports = getClient;


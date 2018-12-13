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


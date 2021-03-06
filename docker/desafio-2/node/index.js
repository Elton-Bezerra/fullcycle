const express = require('express')
const app = express()
const port = 3000

const config = {
    host: 'db',
    user: 'root',
    password: 'root',
    database: 'nodedb'
};

const mysql = require('mysql')
const connection = mysql.createConnection(config)

const create_table_statement = `CREATE TABLE IF NOT EXISTS people(
    id int not null auto_increment,
    name varchar(255),
    primary key (id)
 )`;
const sql = `INSERT INTO people(name) values('Wesley')`
connection.query(create_table_statement);
connection.query(sql);


app.get('/', async (req, res) => {
    //connection.
    connection.query("SELECT name FROM people", function (err, result, fields) {
        if (err) throw err;

        let response = '<h1>Full Cycle</h1> <ul>';

        result.forEach(i => {
            response += `<li>
                ${i.name}
            </li>`
        })
        response += "</ul>"
        res.send(response);
    });
})

app.listen(port, () => {
    console.log('Rodando na porta ' + port)
})



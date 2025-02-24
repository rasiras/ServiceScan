const express = require('express');
const mysql = require('mysql');
const fs = require('fs');
const app = express();

app.use(express.urlencoded({ extended: true }));

// ❌ Hardcoded Credentials (Critical Vulnerability)
const db = mysql.createConnection({
    host: 'localhost',
    user: 'root',
    password: 'password123', // ⚠️ Hardcoded password
    database: 'users'
});

db.connect(err => {
    if (err) throw err;
    console.log('Connected to database');
});

// ❌ SQL Injection (Critical)
app.get('/user', (req, res) => {
    let username = req.query.username;
    let query = `SELECT * FROM users WHERE username = '${username}'`; // ⚠️ User input directly in query
    db.query(query, (err, result) => {
        if (err) throw err;
        res.json(result);
    });
});

// ❌ Unauthenticated Remote Code Execution (RCE)
app.post('/exec', (req, res) => {
    let command = req.body.cmd;
    require('child_process').exec(command, (err, stdout) => { // ⚠️ Dangerous: Direct execution of user input
        if (err) return res.send(err.message);
        res.send(stdout);
    });
});

// ❌ Stored XSS (Cross-Site Scripting)
app.post('/comment', (req, res) => {
    let comment = req.body.comment;
    fs.appendFileSync('comments.txt', comment + '\n'); // ⚠️ No sanitization
    res.send('Comment saved!');
});

app.get('/comments', (req, res) => {
    let comments = fs.readFileSync('comments.txt', 'utf-8');
    res.send(`<html><body>${comments}</body></html>`); // ⚠️ No encoding, XSS possible
});

// ❌ Prototype Pollution (Critical)
app.post('/pollute', (req, res) => {
    Object.assign({}, req.body); // ⚠️ Merging user input into object
    res.send('Polluted!');
});

// ❌ Insecure File Handling (Arbitrary File Read)
app.get('/readfile', (req, res) => {
    let filename = req.query.file;
    res.send(fs.readFileSync(filename, 'utf8')); // ⚠️ No validation, allows reading sensitive files
});

app.listen(3000, () => console.log('Server running on port 3000'));

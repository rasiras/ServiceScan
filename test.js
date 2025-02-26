const express = require('express');
const fs = require('fs');
const { exec } = require('child_process');
const mysql = require('mysql');
const app = express();

app.use(express.json());

// 🚨 1. Remote Code Execution (RCE) - Untrusted user input in exec()
app.get('/exec', (req, res) => {
    const cmd = req.query.cmd; // 🚨 User-controlled input
    exec(cmd, (error, stdout, stderr) => {
        res.send(stdout || stderr || error);
    });
});

// 🚨 2. SQL Injection - Directly concatenating user input
const db = mysql.createConnection({ host: "localhost", user: "root", password: "", database: "test" });
app.get('/user', (req, res) => {
    const userId = req.query.id;
    db.query(`SELECT * FROM users WHERE id = '${userId}'`, (err, result) => { // 🚨 SQL Injection
        if (err) return res.send("DB error");
        res.json(result);
    });
});

// 🚨 3. XSS - Unescaped user input in HTML response
app.get('/xss', (req, res) => {
    const name = req.query.name;
    res.send(`<h1>Welcome ${name}</h1>`); // 🚨 XSS: Name is directly injected into HTML
});

// 🚨 4. Insecure Deserialization - Allowing eval-based JSON execution
app.post('/deserialize', (req, res) => {
    const data = req.body.data;
    const obj = eval('(' + data + ')'); // 🚨 Insecure eval()
    res.send(obj);
});

// 🚨 5. Arbitrary File Read - Reading files from user input
app.get('/read', (req, res) => {
    const filePath = req.query.file;
    fs.readFile(filePath, 'utf8', (err, data) => { // 🚨 Path traversal possible
        if (err) return res.send("Error reading file");
        res.send(data);
    });
});

// 🚨 6. Hardcoded Secrets - Exposing sensitive credentials
const API_KEY = "123456-SECRET-API-KEY"; // 🚨 Hardcoded secret

app.listen(3000, () => {
    console.log('Vulnerable app running on port 3000');
});

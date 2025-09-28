const express = require('express');

const app = express();
const port = 8080;

app.use(express.static('./'))

app.get('/', (req, res)=>
{
    res.sendFile(__dirname + '/index.html')
})

app.get('/redirection_error', (req, res) => {
    res.sendFile(__dirname + '/html/error_page.html')
})

// 5. Start the server
app.listen(port, () => {
    console.log(`Server is running on http://localhost:${port}`);
});
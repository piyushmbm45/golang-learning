const express = require('express');

const app = express();

app.use(express.json());
app.use(
  express.urlencoded({
    extended: true,
  })
);

app.get('/api/get', (req, res) => {
  console.log(req.body);
  res.status(200).json({ message: 'hello from get api' });
});

app.get('/api/home', (req, res) => {
  res.status(200).send('<h1>Welcome To Go lang</h1>');
});

app.post('/api/post', (req, res) => {
  console.log(req.body);
  res.status(200).send(req.body);
});

app.post('/api/post-form/', (req, res) => {
  console.log(req.body);
  const data = JSON.stringify(req.body);
  res.status(200).send(data);
});

app.listen(8000, () => {
  console.log('listening on port 8000');
});

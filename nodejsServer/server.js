// server.js
const express = require("express");
const app = express();
const PORT = 4000;

// Middleware to parse JSON bodies
app.use(express.json());

// GET request handler at /api/hello
app.get("/api/hello", (req, res) => {
  res.json({
    message: "Hello, World!",
  });
});

// POST request handler at /api/data
app.post("/api/data", (req, res) => {
  const { name, age } = req.body; // Destructure name and age from the request body

  if (!name || !age) {
    return res.status(400).json({ error: "Name and age are required" });
  }

  // Construct a JSON response with application/json Content-Type
  res.setHeader("Content-Type", "application/json");
  res.status(200).json({
    message: `Hello, ${name}! You are ${age} years old.`,
  });
});

// Start the server
app.listen(PORT, () => {
  console.log(`Server is running on http://localhost:${PORT}`);
});

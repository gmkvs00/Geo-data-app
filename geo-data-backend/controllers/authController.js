const express = require('express');
const router = express.Router();

// Mock user data (replace this with your actual user database logic)
const users = [
    { email: 'test@example.com', password: 'password123' } // Example user
];

// Handle login request
exports.login = (req, res) => {
    const { email, password } = req.body;

    // Validate input
    if (!email || !password) {
        return res.status(400).json({ error: 'Email and password are required' });
    }

    // Check if user exists
    const user = users.find(user => user.email === email && user.password === password);
    
    if (user) {
        // Successful login
        return res.status(200).json({ message: 'Login successful', user });
    } else {
        // Invalid credentials
        return res.status(401).json({ error: 'Invalid email or password' });
    }
};

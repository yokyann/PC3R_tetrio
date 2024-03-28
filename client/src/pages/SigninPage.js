import React, { useState } from "react";

function SigninPage() {
    // State variables to store the values of username/email and password
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    // Function to handle form submission
    const handleSubmit = (e) => {
        e.preventDefault();
        // Here you can perform authentication or any other action with the entered username and password
        console.log("Username:", username);
        console.log("Password:", password);
    };

    return (
        <div>
            <h1>Signin Page</h1>
            <form onSubmit={handleSubmit}>
                <div>
                    <label htmlFor="username">Username:</label>
                    <input
                        type="text"
                        id="username"
                        value={username}
                        onChange={(e) => setUsername(e.target.value)}
                    />
                </div>
                <div>
                    <label htmlFor="password">Password:</label>
                    <input
                        type="password"
                        id="password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                    />
                </div>
                <button type="submit">Sign In</button>
            </form>
        </div>
    );
};

export default SigninPage;

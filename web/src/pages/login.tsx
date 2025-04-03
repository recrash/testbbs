import React, { useState } from "react";
import { TextField, Button, Typography, Box } from "@mui/material";
import axios from "axios";

function Login() {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault()
        try {
            const response = await axios.post("http://localhost:8081/login", {
                email,
                password
            });
            console.log("로그인 성공:", response.data);
        } catch (error) {
            console.error("로그인 실패:", error);
        }
    }

    return (
        <Box
        sx={{
          width: 300,
          margin: "100px auto",
          padding: 3,
          border: "1px solid #ccc",
          borderRadius: 2,
          boxShadow: 3,
        }}
        >
            <Typography variant="h5" align="center" gutterBottom>
            로그인
            </Typography>

            <form onSubmit={handleSubmit}>
                <TextField
                    fullWidth
                    label="이메일"
                    variant="outlined"
                    type="email"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    margin="normal"
                />

                <TextField
                    fullWidth
                    label="비밀번호"
                    variant="outlined"
                    type="password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    margin="normal"
                />

                <Button
                    type="submit"
                    fullWidth
                    variant="contained"
                    color="primary"
                    sx={{ mt:2 }}
                >
                로그인
                </Button>
            </form>
        </Box>

    )
}

export default Login;
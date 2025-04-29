import React, { useState } from "react";
import { TextField, Button, Typography, Box } from "@mui/material";
import axios from "axios";
import { Link } from "react-router-dom";

function Login() {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault()
        try {
            const response = await axios.post("http://localhost:8081/login", {
                email,
                password,
            }, {
                withCredentials: true,
            });
            console.log("로그인 성공:", response.data);
        } catch (error) {
            if(axios.isAxiosError(error) && error.response) {
                alert(error.response.data.error);
            } else {
                alert("알 수 없는 에러가 발생하였습니다.");
            }
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
                <Box sx={{ textAlign: "right", mt: 1 }}>
                    <Link to="/register">회원가입</Link>
                </Box>
                         
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
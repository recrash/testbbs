import React, { useState } from "react";
import { TextField, Button, Typography, Box } from "@mui/material";
import { useNavigate } from "react-router-dom"
import { api, isAxiosError } from "../api/api";


function Register() {
    const [email, setEmail] = useState("");
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    const navigate = useNavigate();

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault()
        let response = null

        try {
            response = await api.post("/register", { username, email, password });
            console.log("회원가입 성공!", response.data);
            alert("정상적으로 회원가입이 되었습니다.");
            navigate("/login");
        } catch (error) {
            if(isAxiosError(error) && error.response) {
                alert(error.response.data.error);
            } else {
                alert("알 수 없는 에러가 발생하였습니다.");
            }
            console.log("회원가입 실패:", error);
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
            회원가입
            </Typography>

            <form onSubmit={handleSubmit}>
                <TextField
                    fullWidth
                    label="이름 또는 별명"
                    variant="outlined"
                    type="text"
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                    margin="normal"
                />
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
                확인
                </Button>
            </form>
        </Box>        
    )
}

export default Register;
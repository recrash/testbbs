import React, { useContext, useState } from "react";
import { TextField, Button, Typography, Box } from "@mui/material";
import { Link } from "react-router-dom";
import { AuthContext } from "../contexts/AuthContext";
import { api,  isAxiosError } from "../api/api";

function Login() {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const { login } = useContext(AuthContext);

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault()
        try {
            const response = await api.post("/login", { email, password });
            console.log("로그인 성공:", response.data);

            // 사용자 정보를 로컬 스토리지에 저장하여 페이지 새로고침 후에도 로그인 상태 유지
            localStorage.setItem("user", JSON.stringify(response.data.user));
            login(response.data.user);
        } catch (error) {
            if(isAxiosError(error) && error.response) {
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
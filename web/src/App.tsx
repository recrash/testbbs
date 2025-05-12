import './App.css'
import Login from './pages/login'
import Register from './pages/register';
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { AuthProvider } from './contexts/AuthContext';



function App() {  
  return (
    <BrowserRouter>
      <AuthProvider>
        <Routes>
          <Route path ='/' element={<Login />} />
          <Route path='/login' element={<Login />} />
          <Route path='/register' element={<Register />} />
        </Routes>
      </AuthProvider>
    </BrowserRouter>
  );
}

export default App

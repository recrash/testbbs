import './App.css'
import Login from './pages/login'
import Register from './pages/register';
import { BrowserRouter, Routes, Route } from "react-router-dom";


function App() {  
  return (
    <BrowserRouter>
      <Routes>
        <Route path ='/' element={<Login />} />
        <Route path='/login' element={<Login />} />
        <Route path='/register' element={<Register />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App

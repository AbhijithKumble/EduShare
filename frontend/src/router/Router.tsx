import Login from "@/pages/login";
import Register from "@/pages/register";
import { Route, Routes } from "react-router-dom";

const Routing = () => {
    
    
   return (
    <Routes>
        <Route path="/login" element={<Login />} />
        <Route path="/signup" element={<Register />} />
    </Routes>
   ); 

};


export default Routing;

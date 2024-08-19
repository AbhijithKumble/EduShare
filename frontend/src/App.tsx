import { BrowserRouter, Navigate, Route, Routes } from "react-router-dom";
import Login from "./pages/Login";
import Register from "./pages/Register";
import { Home } from "./pages/Home";
import Explore from "./pages/Explore";
import NotFound from "./pages/NotFound";

const App = () => {
  return (
    <BrowserRouter>
      <Routes>
        {/*makes the default path as mycourses for / */}
        <Route path="/" element={<Navigate to="/mycourses" />} />

        {/*Auth layout*/}
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />

        {/*User Layout*/}
        <Route path="/mycourses" element={<Home />} />
        <Route path="/explore" element={<Explore />} />

        {/*Admin Layout*/}

        {/*Non existant route is redirected to not found page*/}
        <Route path="*" element={<NotFound />} />

      </Routes>
    </BrowserRouter>
  );
};

export default App;

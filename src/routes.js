import { Routes, Route, Navigate } from "react-router-dom";
import Homepage from "./pages/homepage";
import Login from "./pages/login";
import Cars from "./pages/cars";
import About from "./pages/about";
import Promotion from "./pages/promotion";
import Contact from "./pages/contact";
import Review from "./pages/review";

const AppRoutes = () => {
    return (
      <Routes>
        <Route path="/" element={<Navigate to="/home" />} />
        <Route path="/home" element={<Homepage />} />
        <Route path="/login" element={<Login />} />        
        <Route path="/cars" element={<Cars />} />     
        <Route path="/about" element={<About />} />     
        <Route path="/promotion" element={<Promotion />} />     
        <Route path="/contact" element={<Contact />} />     
        <Route path="/review" element={<Review />} />     

      </Routes>
    );
  };
  
  export default AppRoutes;
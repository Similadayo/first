import React from 'react';
import { BrowserRouter as Router, Routes, Route  } from 'react-router-dom';
import Header from './components/Header/Header';
import Footer from './components/Footer/Footer';
import Home from './pages/Home/Home';
import ProductDetail from './pages/ProductDetail/ProductDetail';
import ProductCatalog from './pages/ProductCatalog/ProductCatalog';
import ShoppingCart from './pages/ShoppingCart/ShoppingCart';
import Checkout from './pages/Checkout/Checkout';
import UserAccount from './pages/UserAccount/UserAccount';
import AdminDashboard from './pages/Admin/AdminDashboard/AdminDashboard';
import AdminLogin from './pages/Admin/AdminLogin/AdminLogin';
import './App.css';
import Register from './pages/auth/Register';
import Login from './pages/auth/Login';

function App() {
  return (
    <Router>
      <div className="App">
      {window.location.pathname !== '/register' && window.location.pathname !== '/login' && <Header />}
        <main className="main-content">
          <Routes>
            <Route path="/" exact element={<Home/>} />
            <Route path="/product/:productId" element={<ProductDetail/>} />
            <Route path="/catalog" element={<ProductCatalog/>} />
            <Route path="/cart" element={<ShoppingCart/>} />
            <Route path="/checkout" element={<Checkout/>} />
            <Route path="/register" element={<Register />} />
            <Route path="/login" element={<Login/>} />
            <Route path="/account" element={<UserAccount/>} />
            <Route path="/admin" exact element={<AdminLogin/>} />
            <Route path="/admin/dashboard" element={<AdminDashboard/>} />
          </Routes>
        </main>
        <Footer />
      </div>
    </Router>
  );
}

export default App;

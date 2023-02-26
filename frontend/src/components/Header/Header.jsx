import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import { FaShoppingCart, FaBars, FaTimes } from 'react-icons/fa';
import './Header.css';

function Header() {
  const [isOpen, setIsOpen] = useState(false);
  const [loggedIn, setLoggedIn] = useState(false); // set the initial state to false

  const toggleMenu = () => {
    setIsOpen(!isOpen);
  };

  const handleLogin = () => {
    // handle login logic
    setLoggedIn(true); // update the state to true when user logs in
  };

  const handleLogout = () => {
    // handle logout logic
    setLoggedIn(false); // update the state to false when user logs out
  };

  return (
    <header className="header">
      <div className="header__navbar">
        <Link to="/" className="header__logo">
          FIRST-CLOSET
        </Link>
        <nav className={`header__nav ${isOpen ? 'header__nav--open' : ''}`}>
          <Link to="/" className="header__nav-item">
            Home
          </Link>
          <Link to="/products" className="header__nav-item">
            Products
          </Link>
          <Link to="/about" className="header__nav-item">
            About Us
          </Link>
        </nav>
      </div>
      <div className="search-container">
        <input type="text" className="search-input" placeholder="Search for products..." />
        <button className="search-button">Search</button>
      </div>
      {loggedIn ? ( // conditionally render the avatar if user is logged in
        <div className="header__avatar">
          <img src="https://www.example.com/avatar.png" alt="User Avatar" />
        </div>
      ) : (
        <div className="header__auth">
          <button className="header__login" onClick={handleLogin}>Sign Up</button>
          <button className="header__logout" onClick={handleLogout}>Login</button>
        </div>
      )}
      <div className="cart-icon-container">
        <FaShoppingCart className="header__cart-icon cart-icon" />
        <span className="header__cart-count">0</span>
      </div>
      <div className={`header__nav-mobile ${isOpen ? 'header__nav-mobile--open' : ''}`} onClick={toggleMenu}>
        <div className="header__hamburger">
          {isOpen ? <FaTimes /> : <FaBars />}
        </div>
        <nav className="header__nav-mobile-menu">
          <Link to="/" className="header__nav-item">
            Home
          </Link>
          <Link to="/products" className="header__nav-item">
            Products
          </Link>
          <Link to="/about" className="header__nav-item">
            About Us
          </Link>
          {loggedIn ? ( // conditionally render the avatar in mobile view
            <div className="header__avatar">
              <img src="https://www.example.com/avatar.png" alt="User Avatar" />
            </div>
          ) : (
            <div className="header__auth-mobile">
              <button className="header__login" onClick={handleLogin}>Sign Up</button>
              <button className="header__logout" onClick={handleLogout}>Login</button>
            </div>
          )}
        </nav>
      </div>
    </header>
  );
}

export default Header;

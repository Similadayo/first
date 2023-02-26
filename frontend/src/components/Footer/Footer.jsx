import React from 'react';
import './Footer.css';
import { FaFacebook, FaTwitter, FaInstagram } from 'react-icons/fa';

const Footer = () => {
  return (
    <footer>
      <div className="footer-container">
        <div className="footer-section about">
          <h2>About Us</h2>
          <p>
            Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do
            eiusmod tempor incididunt ut labore et dolore magna aliqua.
          </p>
        </div>
        <div className="footer-section contact">
          <h2>Contact Us</h2>
          <ul>
            <li>
              <span>
                <i className="fas fa-map-marker-alt"></i>
              </span>
              <span>123 Main St, New York, NY 10001</span>
            </li>
            <li>
              <span>
                <i className="fas fa-envelope"></i>
              </span>
              <span>info@example.com</span>
            </li>
            <li>
              <span>
                <i className="fas fa-phone"></i>
              </span>
              <span>(123) 456-7890</span>
            </li>
          </ul>
        </div>
        <div className="footer-section social">
          <h2>Follow Us</h2>
          <ul>
            <li>
              <a href="/">
                <FaFacebook className="social-icon" />
              </a>
            </li>
            <li>
              <a href="/">
                <FaTwitter className="social-icon" />
              </a>
            </li>
            <li>
              <a href="/">
                <FaInstagram className="social-icon" />
              </a>
            </li>
          </ul>
        </div>
      </div>
      <div className="footer-bottom">
        <p>© 2023 First-Closet. All Rights Reserved</p>
      </div>
    </footer>
  );
};

export default Footer;

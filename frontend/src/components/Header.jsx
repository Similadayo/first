import React, { useState } from 'react';
import styled, { ThemeProvider } from 'styled-components';
import { AppBar, Toolbar, IconButton, Typography, InputBase } from '@material-ui/core';
import { Menu as MenuIcon, Search as SearchIcon } from '@material-ui/icons';
import { createTheme } from '@material-ui/core/styles';

const HeaderContainer = styled(AppBar)`
  background-color: #000;
  padding: 1rem;
  box-shadow: none;
  display: flex;
  justify-content: space-around;

  ${({ theme }) => theme.breakpoints.down('sm')} {
    padding: 0.5rem;
  }
`;

const Logo = styled.img`
  height: 50px;
  margin-right: 1rem;
`;

const SearchBar = styled(InputBase)`
  height: 25px;
  width: 250px;
  border: none;
  background-color: #e0e0e0;
  border-radius: 15px;
  padding: 0.5rem;

  ${({ theme }) => theme.breakpoints.down('sm')} {
    display: none;
  }
`;

const MenuButton = styled(IconButton)`
  margin-left: auto;

  ${({ theme }) => theme.breakpoints.up('md')} {
    display: none;
  }
`;

const NavItems = styled.div`
  display: flex;
  align-items: center;

  ${({ theme }) => theme.breakpoints.down('sm')} {
    display: none;
  }
`;

const NavItem = styled(Typography)`
  color: #000;
  font-size: 12px;
  margin-right: 11px;
  cursor: pointer;
  text-transform: uppercase;

  &:hover {
    text-decoration: underline;
  }
`;

const HamburgerMenu = styled.div`
  position: absolute;
  top: 60px;
  right: 0;
  width: 100%;
  padding: 1rem;
  background-color: #f8f8f8;
  display: none;

  ${({ theme }) => theme.breakpoints.down('sm')} {
    display: flex;
    flex-direction: column;
    align-items: center;
  }
`;

const HamburgerNavItem = styled(Typography)`
  color: #000;
  font-size: 16px;
  margin-bottom: 1rem;
  cursor: pointer;
  text-transform: uppercase;
`;

const theme = createTheme({
  breakpoints: {
    values: {
      xs: 0,
      sm: 600,
      md: 960,
      lg: 1280,
      xl: 1920,
    },
  },
});

const Header = () => {
  const [isHamburgerMenuOpen, setIsHamburgerMenuOpen] = useState(false);

  const handleHamburgerMenuOpen = () => {
    setIsHamburgerMenuOpen(true);
  };

  const handleHamburgerMenuClose = () => {
    setIsHamburgerMenuOpen(false);
  };

  return (
    <ThemeProvider theme={theme}>
      <HeaderContainer position="static">
        <Toolbar>
          <Logo src="logo.png" alt="My Logo" />
          <SearchBar placeholder="Search..." startAdornment={<SearchIcon />} />
          <MenuButton edge="end" color="inherit" aria-label="menu" onClick={handleHamburgerMenuOpen}>
          <MenuIcon />
          </MenuButton>
          <NavItems>
            <NavItem>Home</NavItem>
            <NavItem>About</NavItem>
            <NavItem>Contact</NavItem>
          </NavItems>
        </Toolbar>
        {isHamburgerMenuOpen && (
          <HamburgerMenu>
            <HamburgerNavItem>Home</HamburgerNavItem>
            <HamburgerNavItem>About</HamburgerNavItem>
            <HamburgerNavItem>Contact</HamburgerNavItem>
            <IconButton edge="end" color="inherit" aria-label="close" onClick={handleHamburgerMenuClose}>
              <SearchIcon />
            </IconButton>
          </HamburgerMenu>
        )}
      </HeaderContainer>
    </ThemeProvider>
  );
};
export default Header;

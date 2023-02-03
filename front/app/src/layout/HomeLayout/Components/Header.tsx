import {observer} from "mobx-react";
import {Box, Button, Container, IconButton, Typography} from "@mui/material";
import AdbIcon from '@mui/icons-material/Adb';
import SearchIcon from '@mui/icons-material/Search';
import React from "react";
import "@/assets/styles/scss/index.scss"
import {useNavigate} from "react-router-dom";

export const Header = observer(function () {
    const navigate = useNavigate();
    return (
        <header className="header">
            <Container fixed className="menu">
                <AdbIcon sx={{display: {xs: 'none', md: 'flex'}, mr: 1}}/>
                <Typography
                    variant="h6"
                    noWrap
                    component="a"
                    href="/"
                    sx={{
                        mr: 2,
                        display: {xs: 'none', md: 'flex'},
                        fontFamily: 'monospace',
                        fontWeight: 700,
                        letterSpacing: '.3rem',
                        color: 'inherit',
                        textDecoration: 'none',
                    }}
                >

                    LOGO
                </Typography>
                <Box>
                    <nav>
                        <ul className="flex flex-row">
                            <li>123</li>
                            <li>123</li>
                            <li>123</li>
                            <li>123</li>
                        </ul>
                    </nav>
                </Box>
                <Box className="ml-auto"/>
                <div>
                    <IconButton>
                        <SearchIcon />
                    </IconButton>
                    <Button onClick={()=>navigate('/login')} variant="outlined" size="small">
                        Sign up
                    </Button>
                </div>
            </Container>
        </header>
    );
})
export default Header
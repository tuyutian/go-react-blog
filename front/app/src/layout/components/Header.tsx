import {observer} from "mobx-react";
import {Box, Container, Typography} from "@mui/material";
import AdbIcon from '@mui/icons-material/Adb';
import React, {useState} from "react";


export const Header = observer(function () {
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
                        <ul>
                            <li>123</li>
                            <li>123</li>
                            <li>123</li>
                            <li>123</li>
                        </ul>
                    </nav>
                </Box>
                <Box className="ml-auto"/>
                <div>
                    1231231
                </div>
            </Container>
        </header>
    );
})
export default Header
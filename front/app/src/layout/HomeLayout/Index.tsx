import Footer from "~/layout/HomeLayout/Components/Footer";
import {FullScreenLoading} from "~/router/Router";
import {observer} from "mobx-react";
import React, {ReactNode, useEffect} from "react";
import { stores } from "@/store";
import Header from "@/layout/HomeLayout/Components/Header";
import {Container, createTheme, ThemeProvider} from "@mui/material";

export const HomeLayout = observer(function ({children}: { children: ReactNode }) {
    const {globalLoading} = stores.commonStore
    useEffect(function (){
        stores.commonStore.setGlobalLoading(true)
    },[])
    const theme = createTheme();
    return globalLoading ? (
                    <ThemeProvider theme={theme}>
                        <Header />
                        <Container>
                            {children}
                        </Container>
                        <Footer />
                    </ThemeProvider>
            ) : <FullScreenLoading />

});
export default HomeLayout
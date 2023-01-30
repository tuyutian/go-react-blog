import Footer from "~/layout/components/Footer";
import Main from "~/layout/components/Main";
import {FullScreenLoading} from "~/router/Router";
import {observer} from "mobx-react";
import React, {useEffect} from "react";
import { stores } from "@/store";

interface Props {
}

export const Index = observer(function (props: Props) {
    const {globalLoading} = stores.commonStore
    useEffect(function (){
        stores.commonStore.setGlobalLoading(true)
    },[])
    return globalLoading ? (
                <div style={{minHeight: '75vh'}}>
                    <Main />
                    <Footer />
                </div>
            ) : <FullScreenLoading />

});
export default Index
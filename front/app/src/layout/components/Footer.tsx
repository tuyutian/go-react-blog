import {stores} from "~/store";
import {observer} from "mobx-react";
import {useLocation} from "react-router-dom";

const Footer = observer(function () {
    const params = useLocation()
    const navigateName = params.pathname;
    const {footerShow} = stores.commonStore

    return <div className="">
        tomaxut
    </div>
})

export default Footer;
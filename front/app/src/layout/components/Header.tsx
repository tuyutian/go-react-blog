import {stores} from "~/store";
import {observer} from "mobx-react";

export const Header = observer(function () {
    const stateLang = stores.commonStore.lang
    return (
        <>
            Top
        </>

    );
})
export default Header
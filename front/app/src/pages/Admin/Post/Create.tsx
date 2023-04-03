import {Modal} from "@mui/material";
import {useState} from "react";
import Box from "@mui/material/Box";

type Props = {
    open:boolean
    onClose:Function
}
export default function CreatePost({open,onClose}:Props) {


    return <div>
        <Modal open={open} onClose={()=>onClose()}>
            <Box sx={{ backgroundColor:"white", width: 400 }}>
                <h2 id="parent-modal-title">Text in a modal</h2>
                <p id="parent-modal-description">
                    Duis mollis, est non commodo luctus, nisi erat porttitor ligula.
                </p>
            </Box>
        </Modal>
    </div>
}
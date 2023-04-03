import {Container} from "@mui/material";
import Button from "@mui/material/Button";
import {Outlet, useNavigate} from "react-router-dom";
import CreatePost from "@/pages/Admin/Post/Create";
import {useState} from "react";

export default function Post() {
    const navigate = useNavigate();

    const [open, setOpen] = useState(false);
    return (<div>
        <Container>
            <Button variant="contained" onClick={() => navigate("/admin/post/create")}>Add Post</Button>
            <CreatePost open={open} onClose={()=>{setOpen(false)}} />
        </Container>
    </div>);
}
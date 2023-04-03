import {lazy, Suspense} from 'react';
import {NonIndexRouteObject, RouteObject, useRoutes} from "react-router-dom";
import classNames from "classnames";
import HomeLayout from "@/layout/HomeLayout/Index"
import AdminLayout from "@/layout/AdminLayout/Index"
import {IndexRouteObject} from "react-router/dist/lib/context";

const Page404Screen = lazy(() => import('~/layout/404'));
const Home = lazy(() => import('@/pages/Home/Index'));
const Login = lazy(() => import('@/pages/Admin/Login'));
const Dashboard = lazy(() => import('@/pages/Admin/Dashboard/Index'));
const Post = lazy(() => import('@/pages/Admin/Post/Index'));
const CreatePost = lazy(() => import('@/pages/Admin/Post/Create'));


export function FullScreenLoading(props: { loading?: boolean }) {
    const _loading = props?.loading || true
    return _loading ? (
        <div
            className={classNames(['flex items-center justify-center absolute top-0 bg-white bg-opacity-90 w-full h-full'])}
            style={{
                zIndex: 100
            }}>
        </div>
    ) : <></>
}

export const Router = () => {
    return (
        <InnerRoutes/>
    );
};

const InnerRoutes = () => {
    const routes: RouteObject[] = [
        {
            path: '/',
            element: <HomeLayout><Home/></HomeLayout>,
            index:true
        },
        {
            path: '/login',
            element: <Login/>,
            index:true
        },
        {
            path:"/admin",
            element:<AdminLayout />,
            children:[
                {
                  path:"",
                  element:<Dashboard />,
                },
                {
                    path: 'post',
                    element: <Post/>,
                }
            ]
        },
        {
            path: '*',
            element: <Page404Screen/>,
        },
    ];

    const element = useRoutes(routes);
    return (
        <div>
            <Suspense fallback={<FullScreenLoading/>}>{element}</Suspense>
        </div>
    );
};

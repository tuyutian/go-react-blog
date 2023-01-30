import {lazy, Suspense} from 'react';
import {RouteObject, useRoutes} from 'react-router-dom';
import classNames from "classnames";

const Page404Screen = lazy(() => import('~/layout/404'));
const Dashboard = lazy(() => import('~/pages/Dashboard'));


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
        <InnerRoutes />
    );
};

const InnerRoutes = () => {
    const PageRouter:RouteObject[] = [];

    const routes: RouteObject[] =[
        {
            path:'/',
            element:<Dashboard />,
            index:true,
        },
        {
            path: '*',
            element: <Page404Screen />,
        },
    ];

    const element = useRoutes(routes);
    return (
        <div>
            <Suspense fallback={<FullScreenLoading />}>{element}</Suspense>
        </div>
    );
};

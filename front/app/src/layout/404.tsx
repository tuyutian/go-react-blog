import {Head} from "~/layout/shared/Head";
import {Link, useLocation} from "react-router-dom";

function Page404() {
  const location = useLocation();
  console.log(location.pathname)
  return (
    <>
      <Head title="The page is not found" />
      <div className="hero min-h-screen ">
        <div className="text-center hero-content text-3xl font-bold">
          <div className="pt-20">
            <h1 >
              The page is not found.
            </h1>
            <div className="mt-4">
              <Link to="/" className="link-primary">Top Page</Link>
            </div>
          </div>
        </div>
      </div>
    </>
  )
}

export default Page404

import LoginForm from "@/components/auth/loginform";

const Login = () => {

  return (
    <div className="bg-black h-screen p-10 flex items-center justify-center">
      <div className="bg-white h-full w-full rounded-xl grid md:grid-cols-2 p-2">
        
        {/* contains two grids with left side is the company logo and right side is auth*/}
        <div className="max-md:hidden bg-green-50 ">
          <div className="bg-blue-50 h-full pt-20">
            <img alt="logo" src="/logo.svg" className="h-[165px] w-[120px] m-auto " />
            <h1 className="text-center font-grotesque font-bold text-8xl">Edu Share</h1>
            <h3 className="font-grotesque font-medium text-4xl text-center pt-10">Unlock your potential with Shared Knowledge</h3>
          </div>
        </div>
        
        <div className="bg-lightpurple rounded-xl  h-full space-y-16">
          <h1 className="font-poppins font-medium text-4xl text-center pt-10">Login</h1>
          <LoginForm />
        </div>
      </div>
    </div>
  );
};

export default Login;

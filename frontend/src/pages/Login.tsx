import LoginForm from "@/components/auth/loginform";

const Login = () => {
  return (
    <div className="bg-black h-screen p-4 md:p-10 lg:p-20 flex items-center justify-center">
      <div className="bg-white h-full w-full max-w-6xl min-w-[300px] rounded-xl grid grid-cols-1 md:grid-cols-2">

        {/* Left Side: Company Logo */}
        <div className="hidden md:flex bg-green-50 rounded-l-xl">
          <div className="bg-blue-50 h-full flex flex-col items-center justify-center space-y-4 p-4 lg:p-10 rounded-l-xl">
            <img alt="logo" src="/logo.svg" className="h-[165px] w-[120px] lg:h-[200px] lg:w-[150px]" />
            <h1 className="text-center font-grotesque font-bold text-6xl lg:text-7xl">Edu Share</h1>
            <h3 className="font-grotesque font-medium text-2xl lg:text-3xl text-center px-4">Unlock your potential with Shared Knowledge</h3>
          </div>
        </div>

        {/* Right Side: Auth */}
        <div className="bg-lightpurple rounded-r-xl h-full flex flex-col items-center justify-center space-y-8 p-4 md:p-10 lg:p-20">
          <h1 className="font-poppins font-medium text-3xl md:text-4xl lg:text-5xl text-center">Login</h1>
          <div className="w-full max-w-lg min-w-[250px]">
            <LoginForm />
          </div>
        </div>
      </div>
    </div>
  );
};

export default Login;


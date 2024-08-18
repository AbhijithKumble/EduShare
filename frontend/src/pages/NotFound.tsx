import { useNavigate } from "react-router-dom";

const Error404 = () => {
  const navigate = useNavigate();

  const goBack = () => {
    navigate("/");
  };

  return (
    <div className="h-screen flex items-center justify-center bg-white md:bg-black">
      <div className="bg-white h-screen w-full min-w-[320px] md:min-w-[640px] rounded-xl grid md:grid-cols-2">
        {/* Left side with the logo */}
        <div className="max-md:hidden bg-green-50">
          <div className="bg-blue-50 h-full ">
            <img alt="logo" src="/logo.svg" className="h-[165px] w-[120px] m-auto" />
            <h1 className="text-center font-grotesque font-bold text-8xl">Edu Share</h1>
            <h3 className="font-grotesque font-medium text-4xl text-center pt-10">Unlock your potential with Shared Knowledge</h3>
          </div>
        </div>

        {/* Right side with the 404 message */}
        <div className="bg-lightpurple rounded-xl h-full flex flex-col items-center justify-center space-y-10">
          <h1 className="font-poppins font-medium text-9xl text-center">404</h1>
          <p className="font-poppins text-2xl text-center">Oops! The page you are looking for does not exist.</p>
          <button
            onClick={goBack}
            className="bg-blue-500 text-white font-poppins font-medium py-3 px-6 rounded-lg hover:bg-blue-700 transition duration-300"
          >
            Go Back Home
          </button>
        </div>
      </div>
    </div>
  );
};

export default Error404;


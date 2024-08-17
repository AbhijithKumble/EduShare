import LoginForm from "@/components/auth/loginform";

const Login = () => {

    return (
        <div className="bg-white h-full grid sm:rounded-xl  md:grid-cols-2 ">
            <div className="bg-lightpurple rounded-xl pr-10 pl-10 w-full h-full space-y-16 pb-10">
                <h1 className="font-poppins font-medium text-4xl text-center pt-10">Login</h1>
                <LoginForm />
            </div>
        </div>
    );
};

export default Login;

import RegisterForm from "@/components/auth/registerform";

const Register = () => {

    return (
        <div className="bg-white h-full grid sm:rounded-xl  md:grid-cols-2">
            <div className="bg-lightpurple rounded-xl pr-10 pl-10 w-full h-full space-y-4 pb-4">
                <h1 className="font-poppins font-medium text-4xl text-center pt-10">Register</h1>
                <RegisterForm />
            </div>
        </div>
    );
};

export default Register;

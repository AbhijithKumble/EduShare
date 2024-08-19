import React from 'react'
import {  useNavigate } from 'react-router-dom';

function Signin() {
  const navigate= useNavigate();
  return (
   
    <div className="bg-black h-screen px-12 py-8">
    <div className="bg-white h-full rounded-xl grid md:grid-cols-2 p-2">
        <div className="hidden md:block bg-green-50">
            <div className="bg-blue-50 h-full pt-20">
                <img alt="logo" src="/logo.svg" className="h-[165px] w-[120px] m-auto" />
                <h1 className="text-center font-grotesque font-bold text-8xl">Edu Share</h1>
                <h3 className="font-grotesque font-medium text-4xl text-center pt-10">
                    Unlock your potential with Shared Knowledge
                </h3>
            </div>
        </div>
        <div className="bg-50 p-20  flex justify-center">
            <div className="bg-green-200 min-h-full     rounded-lg p-2 w-[400px]">
                <h1 className="font-poppins font-medium text-4xl text-center mt-4">Login</h1>
                <input type="email" className="rounded-full h-10 min-w-72 mt-8 mx-12  text-center text-black  text-opacity-55" placeholder="Email"  />
                <input type="password" className="rounded-full h-10 min-w-72 mt-8 mx-12  text-center text-black  text-opacity-55" placeholder="password"  />
                <button className=" mx-12 mt-8 bg-green-400   rounded-full h-10 min-w-72">LOGIN</button>
                <h1 className="mx-16 mt-2 text-sm font-light text-center"> or Continue with </h1>
                <button  className=" mx-12 mt-2   text-gray-950 bg-white rounded-full h-10 min-w-72 flex items-center justify-center"><img src="googleimage.png" alt="Google logo" className="h-10 w-10" />
                     <span>Sign in with Google</span>
                </button>
                <button onClick={()=>navigate('/Signup')}  className="ml-32  mt-2 text-sm font-light text-center"> Don't have an Account </button>
  
            
           
            </div>
        </div>
    </div>
  </div>
  )
}

export default Signin;

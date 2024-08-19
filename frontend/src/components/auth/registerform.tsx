import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import { z } from "zod"

import { RegisterSchema } from "@/schema/auth"

import { Button } from "@/components/ui/button"
import {
    Form,
    FormControl,
    FormField,
    FormItem,
    FormMessage,
} from "@/components/ui/form"
import { Input } from "@/components/ui/input"

const RegisterForm = () => {

    const onSubmit = () => {

    };


    const form = useForm<z.infer<typeof RegisterSchema>>({
        resolver: zodResolver(RegisterSchema),
        defaultValues: {
            email: "",
            password: "",
            name: "",
            confirmPassword: ""
        },
    })

    return (
        <div className="space-y-4">
            <Form {...form}>
                <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4 font-poppins justify-center">

                    <FormField
                        control={form.control}
                        name="name"
                        render={({ field }) => (
                            <FormItem>
                                <FormControl>
                                    <Input placeholder="Name" {...field} className="rounded-3xl" />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="email"
                        render={({ field }) => (
                            <FormItem>
                                <FormControl>
                                    <Input placeholder="Email" {...field} className="rounded-3xl" />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="password"
                        render={({ field }) => (
                            <FormItem>
                                <FormControl>
                                    <Input placeholder="Password" {...field} className="rounded-3xl" />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="confirmPassword"
                        render={({ field }) => (
                            <FormItem>
                                <FormControl>
                                    <Input placeholder="Confirm Password" {...field} className="rounded-3xl" />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        )}
                    />
                    <Button type="submit" className="w-full rounded-3xl bg-btnpink text-black border font-semibold border-black" >Create Account</Button>
                </form>
            </Form>
            <div className="flex flex-col  space-y-4 justify-center items-center font-poppins ">
                <h3 className="mt-2 text-center">or Continue with </h3>
                <Button type="submit" className="w-full rounded-3xl bg-white text-black font-semibold " >Sign in with Google</Button>
                <p className="text-xs">Already have an account ?</p>
                <Button type="submit" className="w-full rounded-3xl bg-btnpink text-black border font-semibold border-black" >Go to Login</Button>
            </div>
        </div>
    );
};
/*
 *
<div className="bg-50 p-20  flex justify-center">
//            <div className="bg-green-200 min-h-full     rounded-lg p-2 w-[400px]">
//                <h1 className="font-poppins font-medium text-4xl text-center mt-4">Login</h1>
//                <input type="email" className="rounded-full h-10 min-w-72 mt-8 mx-12  text-center text-black  text-opacity-55" placeholder="Email"  />
//                <input type="password" className="rounded-full h-10 min-w-72 mt-8 mx-12  text-center text-black  text-opacity-55" placeholder="password"  />
//                <button className=" mx-12 mt-8 bg-green-400   rounded-full h-10 min-w-72">LOGIN</button>
//                <h1 className="mx-16 mt-2 text-sm font-light text-center"> or Continue with </h1>
//                <button  className=" mx-12 mt-2   text-gray-950 bg-white rounded-full h-10 min-w-72 flex items-center justify-center"><img src="googleimage.png" alt="Google logo" className="h-10 w-10" />
//                     <span>Sign in with Google</span>
//                </button>
//                <button onClick={()=>navigate('/Signup')}  className="ml-32  mt-2 text-sm font-light text-center"> Don't have an Account </button>
//
//
//
//            </div>
//        </div>
 *
 * */
export default RegisterForm;

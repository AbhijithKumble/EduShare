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

export default RegisterForm;

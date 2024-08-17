import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import { z } from "zod"

import { LoginSchema } from "@/schema/auth"

import { Button } from "@/components/ui/button"
import {
    Form,
    FormControl,
    FormField,
    FormItem,
    FormMessage,
} from "@/components/ui/form"
import { Input } from "@/components/ui/input"

const LoginForm = () => {

    const onSubmit = () => {

    };


    const form = useForm<z.infer<typeof LoginSchema>>({
        resolver: zodResolver(LoginSchema),
        defaultValues: {
            email: "",
            password: "",
        },
    })

    return (
        <div className="space-y-8">
            <Form {...form}>
                <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8 font-poppins justify-center">
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
                    <Button type="submit" className="w-full rounded-3xl bg-btnpink text-black border font-semibold border-black" >Login</Button>
                </form>
            </Form>
            <div className="flex flex-col  space-y-8 justify-center items-center font-poppins ">
                <h3 className="mt-2 text-center">or Continue with </h3>
                <Button type="button" onClick={signInWithGoogle} className="w-full rounded-3xl bg-white text-black font-semibold " >Sign in with Google</Button>
                <p className="text-xs">Don't have an account ?</p>
                <Button type="button" onClick={} className="w-full rounded-3xl bg-btnpink text-black border font-semibold border-black" >Create an account</Button>
            </div>
        </div>
    );
};

export default LoginForm;

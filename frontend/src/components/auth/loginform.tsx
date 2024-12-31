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
import { useNavigate } from "react-router-dom"
import axios from "axios"

import toast, { Toaster } from 'react-hot-toast';


const LoginForm = () => {

  const navigate = useNavigate();
  const form = useForm<z.infer<typeof LoginSchema>>({
    resolver: zodResolver(LoginSchema),
    defaultValues: {
      email: "",
      password: "",
    }
  })

  const onSubmit = () => {

  };

  const onEmailPasswordSubmit = async (body: z.infer<typeof LoginSchema>) => {
    try {
      const response: any = await axios.post('http://localhost:8080/api/v1/login', body);
      console.log(response);

      // Assuming you store the token here, if returned in the response
      const { token } = response.data;
      localStorage.setItem('edushare-token', token);

      navigate('/mycourses');
    } catch (error: any) {
      // Check if the error is due to a response from the server
      if (error.response) {
        // If there is a response (status code != 2xx)
        toast.error(error.response.data?.message || 'Login failed');
        console.error('Login failed', error.response);
      } else if (error.request) {
        // If there is no response (e.g., network error)
        toast.error('Network error, please try again later');
        console.error('Login failed - No response', error.request);
      } else {
        // If error occurred during setup or something else
        toast.error('Something went wrong');
        console.error('Error', error.message);
      }
    }
  };

  return (
    <div className="space-y-8">
      <Toaster />
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onEmailPasswordSubmit)} className="space-y-8 font-poppins justify-center">
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
        <h3 className="mt-2 text-center" >or Continue with </h3>
        <Button type="button" onClick={onSubmit} className="w-full rounded-3xl bg-white text-black font-semibold " >Sign in with Google</Button>
        <p className="text-xs">Don't have an account ?</p>
        <Button type="button" onClick={() => navigate('/signup')} className="w-full rounded-3xl bg-btnpink text-black border font-semibold border-black" >Create an account</Button>
      </div>
    </div>
  );
};

export default LoginForm;

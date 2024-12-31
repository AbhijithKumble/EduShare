import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import { z } from "zod"
import { useNavigate } from "react-router-dom"
import { SignUpSchema } from "@/schema/auth"
import { Button } from "@/components/ui/button"
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormMessage,
} from "@/components/ui/form"
import { Input } from "@/components/ui/input"
import axios from "axios"
import toast, { Toaster } from "react-hot-toast"

const SignUpForm = () => {
  const navigate = useNavigate();

  const onEmailPasswordSubmit = async (body: z.infer<typeof SignUpSchema>) => {
    const toastId = toast.loading('Loading...')
    try {
      const response: any = await axios.post('http://localhost:8080/api/v1/signup', body);
      console.log(response);
      // Assuming the response contains a success message or user info
      toast.success('User created successfully', {
        id: toastId
      });
      const userID = response.data.userID;
      navigate(`/verifyemail/${userID}`)
    } catch (error: any) {
      // Check if the error is due to a response from the server
      console.log(error);
      if (error.response) {
        // If there is a response (status code != 2xx)
        const errorMessage = error.response.data?.message || 'Signup failed';
        toast.error(errorMessage, {
          id: toastId
        });
        console.error('Signup failed', error.response);
      } else if (error.request) {
        // If there is no response (e.g., network error)
        toast.error('Network error, please try again later', {
          id: toastId
        });
        console.error('Signup failed - No response', error.request);
      } else {
        // If error occurred during setup or something else
        toast.error('Something went wrong', {
          id: toastId
        });
        console.error('Error', error.message);
      }
    }
  };
  const form = useForm<z.infer<typeof SignUpSchema>>({
    resolver: zodResolver(SignUpSchema),
    defaultValues: {
      email: "",
      password: "",
      name: "",
      confirmPassword: ""
    },
  })

  return (
    <div className="space-y-4">
      <Toaster />
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onEmailPasswordSubmit)} className="space-y-4 font-poppins justify-center">

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
        <Button type="submit" className="w-full rounded-3xl bg-btnpink text-black border font-semibold border-black" onClick={() => navigate('/login')} >Go to Login</Button>
      </div>
    </div>
  );
};

export default SignUpForm;

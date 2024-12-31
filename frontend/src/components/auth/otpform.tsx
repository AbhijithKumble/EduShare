import axios from "axios";
import { useState } from "react";
import toast, { Toaster } from "react-hot-toast";
import { useNavigate, useParams } from "react-router";

const OtpForm = () => {
  const [otp, setOtp] = useState<string>("");

  const navigate = useNavigate();

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setOtp(e.target.value);
  };

  const { userID } = useParams(); // Extract userID from the URL
  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    // Submit OTP logic here
    //console.log("OTP submitted:", otp);
    console.log(userID)
    try {
      // Send only userID and OTP to the backend
      const response = await axios.post(
        `http://localhost:8080/api/v1/verifyemail/${userID}`,
        {
          "userID": userID,
          "otp": otp,
        }
      );

      if (response.status === 200) {
        // OTP verified successfully, navigate to another page (e.g., dashboard)
        toast.success("User successfully Logged In!")
        navigate("/login");
      }
    } catch (error) {
      console.log(error);
      if (axios.isAxiosError(error)) {
        toast.error(error.response?.data?.message || "An error occurred");
      } else {
        toast.error("Unexpected error occurred");
      }
    }
  };

  return (
    <>
      <Toaster />
      <form onSubmit={handleSubmit} className="space-y-4">
        <div>
          <label htmlFor="otp" className="block text-xl font-medium">Enter OTP</label>
          <input
            type="text"
            id="otp"
            value={otp}
            onChange={handleChange}
            className="w-full p-3 mt-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
            maxLength={6}
            placeholder="Enter 6-digit OTP"
          />
        </div>
        <button
          type="submit"
          className="w-full p-3 bg-blue-500 text-white rounded-lg font-semibold hover:bg-blue-600 transition duration-300"
        >
          Verify OTP
        </button>
      </form>
    </>
  );
};

export default OtpForm;


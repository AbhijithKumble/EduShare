import { ReactNode } from "react";

interface LayoutProps  {
  children : ReactNode;
}

const Layout = ({children} : LayoutProps) => {

  return (
    <main className="bg-black h-screen">
      {children} 
    </main>
  );
};


export default Layout;

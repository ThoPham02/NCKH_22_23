import Header from "./Header";
import Footer from "./Footer";

const DefaultLayout = ({ children }) => {
  return (
    <>
      <Header />
      <div style={{minHeight: "calc(100vh - 275px)"}}>
        {children}
      </div>
      <Footer />
    </>
  );
};

export default DefaultLayout;

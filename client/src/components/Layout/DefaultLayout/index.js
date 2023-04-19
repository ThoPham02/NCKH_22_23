import Header from "./Header";
import Footer from "./Footer";
import NavBar from "./Navbar";

const DefaultLayout = ({ children }) => {
  return (
    <>
      <Header />
      <NavBar />
      <div style={{ minHeight: "calc(100vh - 212px)", paddingTop: "4px" }} className="container">
        {children}
      </div>
      <Footer />
    </>
  );
};

export default DefaultLayout;

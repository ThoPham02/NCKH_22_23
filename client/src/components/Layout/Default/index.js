import Header from "./Header";
import Footer from "./Footer";
import NavBar from "./Navbar";

const background = require("./background.png");

const Default = ({ children }) => {
  return (
    <div
      style={{
        backgroundImage: `url(${background})`,
        backgroundSize: "100% auto",
      }}
    >
      <Header />
      <NavBar />
      <div
        className="container"
        style={{
          minHeight: "calc(100vh - 215px)",
          marginTop: "4px",
        }}
      >
        {children}
      </div>
      <Footer />
    </div>
  );
};

export default Default;

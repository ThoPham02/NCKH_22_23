import { useState } from "react";
import Pagination1 from "../../components/Shares/Pagination";

const Home = () => {
  const [pagi, setPagi] = useState(1)
  return (
    <div className="home">
      <Pagination1 currentPage={pagi} total={10} limit={2} setPagi={setPagi}/>
    </div>
  );
}

export default Home;
